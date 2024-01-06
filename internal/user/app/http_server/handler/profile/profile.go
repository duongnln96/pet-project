package profile

import (
	"net/http"

	"github.com/duongnln96/blog-realworld/internal/pkg/serror"
	"github.com/duongnln96/blog-realworld/internal/user/core/port"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type profileResponse struct {
	port.UserDTO
	Following bool `json:"following"`
}

func (h *handler) Profile(c echo.Context) error {

	httpCtx := c.Request().Context()

	profileUserID, err := uuid.Parse(c.Param("profile_user_id"))
	if err != nil {
		return serror.NewErrorResponse(http.StatusBadRequest, serror.ErrUserCommon, err.Error())
	}

	userIDStr, ok := c.Get("user_id").(string)
	if !ok {
		return serror.NewErrorResponse(http.StatusBadRequest, serror.ErrUserCommon, "Cant get user_id from token")
	}
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return serror.NewErrorResponse(http.StatusBadRequest, serror.ErrUserCommon, err.Error())
	}

	profile, err := h.followService.Profile(httpCtx, userID, profileUserID)
	if err != nil {
		return serror.Service2EchoErr(err)
	}

	return c.JSON(serror.EchoSuccess(profile))
}
