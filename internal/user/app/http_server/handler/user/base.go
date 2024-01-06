package user

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"

	"github.com/duongnln96/blog-realworld/internal/user/core/port"
)

type HandlerI interface {
	Detail(echo.Context) error
	Register(echo.Context) error
	Update(echo.Context) error
}

type handler struct {
	userService port.UserServiceI
}

var HandlerSet = wire.NewSet(NewHandler)

func NewHandler(
	userServiceInstance port.UserServiceI,
) HandlerI {
	return &handler{
		userService: userServiceInstance,
	}
}
