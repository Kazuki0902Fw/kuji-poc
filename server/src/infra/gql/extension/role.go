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

// RoleExtension は gqlgen のフィールドインターセプトを行う拡張
// "ディレクティブが付いているかどうか" を調べ、
// 付いていなければ認可チェックを行わない ( = ユーザ情報があるかどうか )
type RoleExtension struct {
	userRepository repository.UserRepository
}

// NewRoleExtension creates a new RoleExtension with the given UserRepository
func NewRoleExtension(userRepository repository.UserRepository) *RoleExtension {
	return &RoleExtension{
		userRepository: userRepository,
	}
}

func (r *RoleExtension) ExtensionName() string {
	return "RoleExtension"
}
func (r *RoleExtension) Validate(schema graphql.ExecutableSchema) error {
	return nil
}

// Operation 全体を通して何かしたい場合に使う (今回は何もしない)
func (r *RoleExtension) InterceptOperation(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	return next(ctx)
}

// フィールドがリゾルバ実行される直前のフック
func (r *RoleExtension) InterceptField(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		// 何かおかしい場合はそのまま
		return next(ctx)
	}

	// ディレクティブ @isAdmin が付いていれば認可必須とする
	if !isAdmin(fc) {
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

	if user.Role != model.UserRoleAdmin {
		return nil, errors.New("unauthorized")
	}

	return next(ctx)
}

// isAdmin はフィールド定義に `@isAdmin` が付いているかどうかを判定
func isAdmin(fc *graphql.FieldContext) bool {
	// path が __ で始まる場合は introspection なのでスキップ
	if strings.HasPrefix(fc.Path().String(), "__") {
		return true
	}

	// ディレクティブ @isAdmin が付いているかどうかを判定
	for _, d := range fc.Field.Definition.Directives {
		if d.Name == "isAdmin" {
			return true
		}
	}
	return false
}
