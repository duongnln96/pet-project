package profile

import (
	"context"

	"github.com/duongnln96/blog-realworld/internal/pkg/serror"
	"github.com/google/uuid"
)

func (s *service) CountByFollowedUserID(ctx context.Context, followedUserID uuid.UUID) (int64, error) {

	count, err := s.followRepo.CountByFollowedUserID(ctx, followedUserID)
	if err != nil {
		return 0, serror.NewSystemSError(err.Error())
	}

	return count, nil
}
