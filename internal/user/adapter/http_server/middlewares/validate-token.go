package middlewares

import (
	"github.com/duongnln96/blog-realworld/pkg/config"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

type userAuth struct {
	config *config.Configs

	// TODO: grpc client
}

func NewUserAuthMiddleware(
	config *config.Configs,
) *userAuth {
	return &userAuth{
		config: config,
	}
}

var UserAuthMiddlewareSet = wire.NewSet(NewUserAuthMiddleware)

// Process is the middleware function.
func ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		auth := c.Request().Header.Get(echo.HeaderAuthorization)

		// TODO:
		_ = auth

		// token.

		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}
