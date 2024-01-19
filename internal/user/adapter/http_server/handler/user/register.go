package user

import (
	"log/slog"
	"net/http"

	"github.com/duongnln96/blog-realworld/internal/pkg/serror"
	"github.com/duongnln96/blog-realworld/internal/user/core/port"
	"github.com/labstack/echo/v4"
)

func (h *handler) Register(c echo.Context) error {
	slog.Info("Register user")

	httpCtx := c.Request().Context()

	var request = new(port.RegisterUserDTO)
	if err := c.Bind(request); err != nil {
		return serror.NewErrorResponse(http.StatusBadRequest, "", err.Error())
	}
	if err := c.Validate(request); err != nil {
		return serror.NewErrorResponse(http.StatusBadRequest, "", err.Error())
	}

	user, err := h.userService.Register(httpCtx, request)
	if err != nil {
		return serror.Service2EchoErr(err)
	}

	return c.JSON(serror.EchoSuccess(user))
}
