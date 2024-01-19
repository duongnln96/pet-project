package profile

import (
	"github.com/duongnln96/blog-realworld/internal/user/core/port"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

type HandlerI interface {
	Follow(echo.Context) error
	Unfollow(echo.Context) error
	Profile(echo.Context) error
}

type handler struct {
	followService port.FollowServiceI
	userService   port.UserServiceI
}

var HandlerSet = wire.NewSet(NewHandler)

func NewHandler(
	followServiceInstance port.FollowServiceI,
	userServiceInstance port.UserServiceI,

) HandlerI {
	return &handler{
		followService: followServiceInstance,
		userService:   userServiceInstance,
	}
}
