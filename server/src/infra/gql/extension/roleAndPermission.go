package extension

import (
	"context"
	"errors"
	"strings"

	"github.com/99designs/gqlgen/graphql"

	"kujicole/domain/model"
	"kujicole/domain/repository"
	echoMiddleware "kujicole/infra/echo/middleware"
)

// RoleAndPermissionExtension は gqlgen のフィールドインターセプトを行う拡張
// "ディレクティブが付いているかどうか" を調べ、
// 付いていなければ認可チェックを行わない ( = ユーザ情報があるかどうか )
type RoleAndPermissionExtension struct {
	userRepository repository.UserRepository
}

// NewRoleAndPermissionExtension creates a new RoleAndPermissionExtension with the given UserRepository
func NewRoleAndPermissionExtension(userRepository repository.UserRepository) *RoleAndPermissionExtension {
	return &RoleAndPermissionExtension{
		userRepository: userRepository,
	}
}

func (r *RoleAndPermissionExtension) ExtensionName() string {
	return "RoleAndPermissionExtension"
}
func (r *RoleAndPermissionExtension) Validate(schema graphql.ExecutableSchema) error {
	return nil
}

// Operation 全体を通して何かしたい場合に使う (今回は何もしない)
func (r *RoleAndPermissionExtension) InterceptOperation(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	return next(ctx)
}

// フィールドがリゾルバ実行される直前のフック
func (r *RoleAndPermissionExtension) InterceptField(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		// 何かおかしい場合はそのまま
		return next(ctx)
	}

	// ディレクティブ @isAdmin または @isItemTransferPermission が付いていれば認可必須とする
	if !isAdminOrAdministrativePrivilege(fc) {
		return next(ctx)
	}

	// 付いている → 認可必須とする
	// ミドルウェアで埋め込んだユーザ情報を取り出す
	userIDAny := ctx.Value(echoMiddleware.EchoUserIDContextKey)
	if userIDAny == nil {
		return nil, errors.New("unauthorized")
	}
	userID := userIDAny.(model.ID)
	user, err := r.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return nil, errors.New("unauthorized")
	}

	if user.Role != model.UserRoleAdmin && !user.AdministrativePrivilege {
		return nil, errors.New("unauthorized")
	}

	return next(ctx)
}

// isAdminOrAdministrativePrivilege はフィールド定義に `@isAdmin` または `@isAdministrativePrivilege` が付いているかどうかを判定
func isAdminOrAdministrativePrivilege(fc *graphql.FieldContext) bool {
	// path が __ で始まる場合は introspection なのでスキップ
	if strings.HasPrefix(fc.Path().String(), "__") {
		return true
	}

	// ディレクティブ @isAdmin または @isAdministrativePrivilegeが付いているかどうかを判定
	for _, d := range fc.Field.Definition.Directives {
		if d.Name == "isAdmin" || d.Name == "isAdministrativePrivilege" {
			return true
		}
	}
	return false
}
