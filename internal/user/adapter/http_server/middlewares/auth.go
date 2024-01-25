package middlewares

import (
	"fmt"
	"net/http"

	"github.com/duongnln96/blog-realworld/internal/pkg/serror"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"

	"github.com/duongnln96/blog-realworld/internal/user/core/port"
)

var AuthMiddlewareSet = wire.NewSet(NewAuthMiddleware)

type AuthMiddleware struct {
	authTokenDomain port.AuthTokenDomainI
}

func NewAuthMiddleware(
	authTokenDomain port.AuthTokenDomainI,
) *AuthMiddleware {
	return &AuthMiddleware{
		authTokenDomain: authTokenDomain,
	}
}

func (m *AuthMiddleware) ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		authToken := c.Request().Header.Get(echo.HeaderAuthorization)
		if authToken == "" {
			return serror.NewErrorResponse(http.StatusBadRequest, serror.ErrUserCommon, fmt.Sprintf("%s not found", echo.HeaderAuthorization))
		}

		validateToken, err := m.authTokenDomain.ValidateToken(c.Request().Context(), &port.ValidateTokenRequest{
			JwtToken: authToken,
		})
		if err != nil {
			return serror.NewErrorResponse(http.StatusInternalServerError, serror.ErrSystemInternal, fmt.Sprintf("authTokenDomain.ValidateToken %s", err.Error()))
		}
		if !validateToken.IsValid {
			return serror.NewErrorResponse(http.StatusUnauthorized, serror.ErrUnauthorized, fmt.Sprint("user auth_token is invalid"))
		}

		return next(c)
	}
}
