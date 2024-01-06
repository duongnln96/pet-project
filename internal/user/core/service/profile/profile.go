package profile

import (
	"context"

	"github.com/duongnln96/blog-realworld/internal/pkg/serror"
	"github.com/duongnln96/blog-realworld/internal/user/core/domain"
	"github.com/duongnln96/blog-realworld/internal/user/core/port"
	"github.com/google/uuid"
)

func (s *service) Profile(ctx context.Context, followingUserID, followedUserID uuid.UUID) (port.ProfileDTO, error) {

	profileRes := port.NewEmptyProfileDTO()

	followingUser, err := s.userRepo.GetOneByID(ctx, followingUserID)
	if err != nil {
		return profileRes, serror.NewSystemSError(err.Error())
	}
	if !followingUser.IsExist() {
		return profileRes, serror.NewSError(domain.NotFoundErrUser, "user not found")
	}

	folowed, err := s.followRepo.GetOne(ctx, followedUserID, followingUserID)
	if err != nil {
		return profileRes, serror.NewSystemSError(err.Error())
	}
	if folowed.IsExist() {
		profileRes.Following = true
	}

	profileRes.Name = followingUser.Name
	profileRes.Bio = followingUser.Bio

	return profileRes, nil
}
