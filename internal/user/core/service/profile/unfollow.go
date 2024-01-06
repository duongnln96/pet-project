package profile

import (
	"context"
	"fmt"

	"github.com/duongnln96/blog-realworld/internal/pkg/serror"
	"github.com/duongnln96/blog-realworld/internal/user/core/domain"
	"github.com/duongnln96/blog-realworld/internal/user/core/port"
	"github.com/google/uuid"
)

func (s *service) Unfollow(ctx context.Context, followedUserID uuid.UUID, followingUserID uuid.UUID) (port.FollowDTO, error) {

	response := port.NewEmptyFollowDTO()

	repoFollow, err := s.followRepo.GetOne(ctx, followedUserID, followingUserID)
	if err != nil {
		return response, serror.NewSystemSError(err.Error())
	}
	if !repoFollow.IsExist() {
		return response, serror.NewSError(domain.NotFoundErrFollow, fmt.Sprintf("%s still not follow %s", followedUserID.String(), followingUserID.String()))
	}

	updated, err := s.followRepo.Update(ctx, followedUserID, followingUserID, domain.DeactiveFollowStatus)
	if err != nil {
		return response, serror.NewSystemSError(err.Error())
	}

	response.Domain2Port(updated)

	return response, nil
}
