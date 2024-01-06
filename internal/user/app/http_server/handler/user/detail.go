package user

import (
	"log/slog"
	"net/http"

	"github.com/duongnln96/blog-realworld/internal/pkg/serror"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type detailReq struct {
	ID uuid.UUID `query:"id" validate:"required"`
}

func (h *handler) Detail(c echo.Context) error {

	slog.Info("Get detail user")

	httpCtx := c.Request().Context()

	var request = new(detailReq)
	if err := c.Bind(request); err != nil {
		return serror.NewErrorResponse(http.StatusBadRequest, "", err.Error())
	}
	if err := c.Validate(request); err != nil {
		return serror.NewErrorResponse(http.StatusBadRequest, "", err.Error())
	}

	user, err := h.userService.Detail(httpCtx, request.ID)
	if err != nil {
		serr, _ := err.(*serror.SError)
		if serr.IsSystem {
			return serror.NewErrorResponse(http.StatusBadGateway, serr.Code, serr.Msg)
		} else {
			return serror.NewErrorResponse(http.StatusOK, serr.Code, serr.Msg)
		}
	}

	return c.JSON(serror.EchoSuccess(user))
}
