package user

import (
	"log/slog"
	"net/http"

	"github.com/duongnln96/blog-realworld/internal/pkg/serror"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *handler) Detail(c echo.Context) error {

	slog.Info("Get detail user")

	httpCtx := c.Request().Context()

	userID, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		return serror.NewErrorResponse(http.StatusBadRequest, serror.ErrUserCommon, err.Error())
	}

	user, err := h.userService.Detail(httpCtx, userID)
	if err != nil {
		return serror.Service2EchoErr(err)
	}

	return c.JSON(serror.EchoSuccess(user))
}
