package handler

import (
	"rename-service-name-here/internal/util"
	"strings"

	"github.com/labstack/echo/v4"
)

func authorizationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		errs := util.AppErrors{}
		bearerToken := ctx.Request().Header.Get("Authorization-Signin")
		token := strings.TrimPrefix(bearerToken, "Bearer ")

		userId, expiresAt, err := util.ValidateToken(token)
		if err != nil {
			errs.Add(util.NewUnauthorizedError(err.Error()))
			return errorHandler(ctx, errs)
		}
		// Call the next handler in the chain
		ctx.Set("userId", userId)
		ctx.Set("expiresAt", expiresAt)

		return next(ctx)
	}
}
