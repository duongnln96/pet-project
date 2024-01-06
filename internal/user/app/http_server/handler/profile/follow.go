package profile

import (
	"net/http"

	"github.com/duongnln96/blog-realworld/internal/pkg/serror"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *handler) Follow(c echo.Context) error {

	httpCtx := c.Request().Context()

	followingUserID, err := uuid.Parse(c.Param("follow_user_id"))
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

	follow, err := h.followService.Follow(httpCtx, userID, followingUserID)
	if err != nil {
		return serror.Service2EchoErr(err)
	}

	return c.JSON(serror.EchoSuccess(follow))
}
