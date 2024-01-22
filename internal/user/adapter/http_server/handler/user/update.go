package user

import (
	"log/slog"
	"net/http"

	"github.com/duongnln96/blog-realworld/internal/pkg/serror"
	"github.com/duongnln96/blog-realworld/internal/user/core/port"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type updateUserRequest struct {
	ID       uuid.UUID `json:"id"`
	Name     *string   `json:"name"`
	Bio      *string   `json:"bio"`
	Email    *string   `json:"email"`
	Password *string   `json:"password"`
}

func (h *handler) Update(c echo.Context) error {
	slog.Info("Update user")

	httpCtx := c.Request().Context()

	var request = new(updateUserRequest)
	if err := c.Bind(request); err != nil {
		return serror.NewErrorResponse(http.StatusBadRequest, "", err.Error())
	}
	if err := c.Validate(request); err != nil {
		return serror.NewErrorResponse(http.StatusBadRequest, "", err.Error())
	}

	user, err := h.userService.Update(httpCtx, &port.UpdateUserRequest{
		ID:       request.ID,
		Name:     request.Name,
		Bio:      request.Bio,
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return serror.Service2EchoErr(err)
	}

	return c.JSON(serror.EchoSuccess(user))
}
