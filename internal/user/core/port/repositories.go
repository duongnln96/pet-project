package port

import (
	"context"

	"github.com/duongnln96/blog-realworld/internal/user/core/domain"
	"github.com/google/uuid"
)

type UserRepoI interface {
	Create(context.Context, domain.User) (domain.User, error)
	Update(context.Context, domain.User) (domain.User, error)
	GetOneByID(context.Context, uuid.UUID) (domain.User, error)
	GetOneByEmail(context.Context, string) (domain.User, error)
}

type FollowRepoI interface {
	Create(ctx context.Context, followedUserID uuid.UUID, followingUserID uuid.UUID) (domain.Follow, error)
	Update(ctx context.Context, followedUserID uuid.UUID, followingUserID uuid.UUID, status domain.FollowStatus) (domain.Follow, error)
	AllByFollowedUserID(ctx context.Context, followedUserID uuid.UUID, offset int64, limit int32) (domain.Follows, error)
	GetOne(ctx context.Context, followedUserID uuid.UUID, followingUserID uuid.UUID) (domain.Follow, error)
	CountByFollowedUserID(ctx context.Context, followedUserID uuid.UUID) (int64, error)
}
