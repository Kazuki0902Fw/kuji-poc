package authtoken

import (
	// "context"
	// "net/http"
	// "strings"

	// "github.com/labstack/echo/v4"
	// "github.com/myuoncorp/usednet-server/domain/service"
)

type echoUserIDContextKeyType string

const (
	EchoUserIDContextKey = echoUserIDContextKeyType("echo.user_id")
)

// type AddUserIDMiddleware struct {
// 	Service service.AuthTokenService
// }

// func (f *AddUserIDMiddleware) MiddlewareFunc() echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			ctx := c.Request().Context()

// 			authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
// 			if authHeader == "" {
// 				// このミドルウェア token があれば userID を埋め込むためのものなので、トークンがない場合はスルーする
// 				return next(c)
// 			}

// 			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
// 			userID, err := f.Service.VerifyAccessToken(ctx, tokenString)
// 			if err != nil {
// 				return c.JSON(http.StatusUnauthorized, "invalid access token format")
// 			}

// 			req := c.Request()
// 			ctx = context.WithValue(ctx, EchoUserIDContextKey, *userID)
// 			c.SetRequest(req.WithContext(ctx))

// 			return next(c)
// 		}
// 	}
// }
