package profile

import (
	"context"
	"fmt"

	"github.com/duongnln96/blog-realworld/internal/pkg/serror"
	"github.com/duongnln96/blog-realworld/internal/user/core/domain"
	"github.com/duongnln96/blog-realworld/internal/user/core/port"
	"github.com/google/uuid"
)

func (s *service) Follow(ctx context.Context, followedUserID uuid.UUID, followingUserID uuid.UUID) (port.FollowDTO, error) {

	response := port.NewEmptyFollowDTO()

	repoFollow, err := s.followRepo.GetOne(ctx, followedUserID, followingUserID)
	if err != nil {
		return response, serror.NewSystemSError(err.Error())
	}
	if repoFollow.IsExist() {
		return response, serror.NewSError(domain.ExistedErrFollow, fmt.Sprintf("%s is following %s", followedUserID.String(), followingUserID.String()))
	}

	followed, err := s.followRepo.Create(ctx, followedUserID, followingUserID)
	if err != nil {
		return response, serror.NewSystemSError(err.Error())
	}

	response.Domain2Port(followed)

	return response, nil
}
