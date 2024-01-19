package user

import (
	"log/slog"
	"net/http"

	"github.com/duongnln96/blog-realworld/internal/pkg/serror"
	"github.com/duongnln96/blog-realworld/internal/user/core/port"
	"github.com/labstack/echo/v4"
)

func (h *handler) Update(c echo.Context) error {
	slog.Info("Update user")

	httpCtx := c.Request().Context()

	var request = new(port.UpdateUserDTO)
	if err := c.Bind(request); err != nil {
		return serror.NewErrorResponse(http.StatusBadRequest, "", err.Error())
	}
	if err := c.Validate(request); err != nil {
		return serror.NewErrorResponse(http.StatusBadRequest, "", err.Error())
	}

	user, err := h.userService.Update(httpCtx, *request)
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
