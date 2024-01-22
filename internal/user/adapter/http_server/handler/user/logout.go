package user

import (
	"log/slog"

	"github.com/duongnln96/blog-realworld/internal/pkg/serror"
	"github.com/labstack/echo/v4"
)

func (h *handler) Logout(c echo.Context) error {
	defer slog.Info("Logout user")

	// httpCtx := c.Request().Context()

	// token, err := h.userService.LogIn(httpCtx, &port.LoginUserRequest{
	// 	Email:    request.Email,
	// 	Password: request.Password,
	// })
	// if err != nil {
	// 	serr, _ := err.(*serror.SError)
	// 	if serr.IsSystem {
	// 		return serror.NewErrorResponse(http.StatusInternalServerError, serr.Code, serr.Msg)
	// 	} else {
	// 		return serror.NewErrorResponse(http.StatusOK, serr.Code, serr.Msg)
	// 	}
	// }

	return c.JSON(serror.EchoSuccess(nil))
}
