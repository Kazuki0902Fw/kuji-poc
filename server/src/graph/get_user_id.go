package graph

import (
	"context"

	"kujicole/domain/model"
	customEchoMiddleware "kujicole/infra/echo/middleware"
)

func getUserID(ctx context.Context) (model.ID, error) {
	userID := ctx.Value(customEchoMiddleware.EchoUserIDContextKey).(model.ID)
	return userID, nil
}
