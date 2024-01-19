package user

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/duongnln96/blog-realworld/internal/pkg/serror"
	"github.com/duongnln96/blog-realworld/internal/user/core/port"
	"github.com/labstack/echo/v4"
)

func (h *handler) Login(c echo.Context) error {
	defer slog.Info("Login user")

	var request = new(port.LoginUserDTO)
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

	user, err := h.userService.LogIn(ctx, request)
	if err != nil {
		serr, _ := err.(*serror.SError)
		if serr.IsSystem {
			return serror.NewErrorResponse(http.StatusInternalServerError, serr.Code, serr.Msg)
		} else {
			return serror.NewErrorResponse(http.StatusOK, serr.Code, serr.Msg)
		}
	}

	return c.JSON(serror.EchoSuccess(user))
}
