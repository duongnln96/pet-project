package user

import (
	"log/slog"
	"net/http"

	"github.com/duongnln96/blog-realworld/internal/pkg/serror"
	"github.com/duongnln96/blog-realworld/internal/user/core/port"
	"github.com/labstack/echo/v4"
)

type registerUserRequest struct {
	Name     string `json:"name" validate:"required,max=500"`
	Bio      string `json:"bio" validate:"omitempty,max=1000"`
	Email    string `json:"email" validate:"required,max=500"`
	Password string `json:"password" validate:"required"`
}

func (h *handler) Register(c echo.Context) error {
	slog.Info("Register user")

	httpCtx := c.Request().Context()

	var request = new(registerUserRequest)
	if err := c.Bind(request); err != nil {
		return serror.NewErrorResponse(http.StatusBadRequest, "", err.Error())
	}
	if err := c.Validate(request); err != nil {
		return serror.NewErrorResponse(http.StatusBadRequest, "", err.Error())
	}

	user, err := h.userService.Register(httpCtx, &port.RegisterUserRequest{
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
