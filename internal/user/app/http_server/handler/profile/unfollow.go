package profile

import (
	"net/http"

	"github.com/duongnln96/blog-realworld/internal/pkg/serror"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *handler) Unfollow(c echo.Context) error {

	httpCtx := c.Request().Context()

	unfollowingUserID, err := uuid.Parse(c.Param("unfollow_user_id"))
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

	unfollow, err := h.followService.Unfollow(httpCtx, userID, unfollowingUserID)
	if err != nil {
		return serror.Service2EchoErr(err)
	}

	return c.JSON(serror.EchoSuccess(unfollow))
}
