package extension

import (
	"context"
	"errors"
	"strings"

	"github.com/99designs/gqlgen/graphql"

	echoMiddleware "kujicole/infra/echo/middleware"
)

// AuthExtension は gqlgen のフィールドインターセプトを行う拡張
// "ディレクティブが付いているかどうか" を調べ、
// 付いていなければ認証チェックを行う ( = ユーザ情報があるかどうか )
type AuthExtension struct{}

func (a *AuthExtension) ExtensionName() string {
	return "AuthExtension"
}
func (a *AuthExtension) Validate(schema graphql.ExecutableSchema) error {
	return nil
}

// Operation 全体を通して何かしたい場合に使う (今回は何もしない)
func (a *AuthExtension) InterceptOperation(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	return next(ctx)
}

// フィールドがリゾルバ実行される直前のフック
func (a *AuthExtension) InterceptField(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		// 何かおかしい場合はそのまま
		return next(ctx)
	}

	// ディレクティブ @withoutAuth が付いていればスキップ (認証チェックしない)
	if isWithoutAuth(fc) {
		return next(ctx)
	}

	// 付いていない → 認証必須とする
	// ミドルウェアで埋め込んだユーザ情報を取り出す
	user := ctx.Value(echoMiddleware.EchoUserIDContextKey)
	if user == nil {
		// ユーザ情報が無ければエラー (401)
		return nil, errors.New("unauthorized")
	}

	return next(ctx)
}

// isWithoutAuth はフィールド定義に `@withoutAuth` が付いているかどうかを判定
func isWithoutAuth(fc *graphql.FieldContext) bool {
	// path が __ で始まる場合は introspection なのでスキップ
	if strings.HasPrefix(fc.Path().String(), "__") {
		return true
	}

	// ディレクティブ @withoutAuth が付いているかどうかを判定
	for _, d := range fc.Field.Definition.Directives {
		if d.Name == "withoutAuth" {
			return true
		}
	}
	return false
}
