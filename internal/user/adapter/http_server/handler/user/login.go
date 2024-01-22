package user

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/duongnln96/blog-realworld/internal/pkg/serror"
	"github.com/duongnln96/blog-realworld/internal/user/core/port"
	"github.com/labstack/echo/v4"
)

type loginRequest struct {
	Email    string `json:"email" validate:"required,max=500"`
	Password string `json:"password" validate:"required"`
}

type loginResponse struct {
	JwtToken string `json:"jwt_token"`
}

func (h *handler) Login(c echo.Context) error {
	defer slog.Info("Login user")

	var request = new(loginRequest)
	if err := c.Bind(request); err != nil {
		return serror.NewErrorResponse(http.StatusBadRequest, "", err.Error())
	}
	if err := c.Validate(request); err != nil {
		return serror.NewErrorResponse(http.StatusBadRequest, "", err.Error())
	}

	httpCtx := c.Request().Context()
	ctx := context.WithValue(httpCtx, "user-agent", c.Request().UserAgent())
	ctx = context.WithValue(ctx, "device-id", c.Request().Header.Get("Device-Id"))
	ctx = context.WithValue(ctx, "remote-ip", c.RealIP())

	token, err := h.userService.LogIn(ctx, &port.LoginUserRequest{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		serr, _ := err.(*serror.SError)
		if serr.IsSystem {
			return serror.NewErrorResponse(http.StatusInternalServerError, serr.Code, serr.Msg)
		} else {
			return serror.NewErrorResponse(http.StatusOK, serr.Code, serr.Msg)
		}
	}

	return c.JSON(serror.EchoSuccess(loginResponse{
		JwtToken: token.JwtToken,
	}))
}
