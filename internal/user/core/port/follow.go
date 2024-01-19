package port

import (
	"context"

	"github.com/duongnln96/blog-realworld/internal/user/core/domain"
	"github.com/google/uuid"
)

type FollowRepoI interface {
	Create(ctx context.Context, followedUserID uuid.UUID, followingUserID uuid.UUID) (domain.Follow, error)
	Update(ctx context.Context, followedUserID uuid.UUID, followingUserID uuid.UUID, status domain.FollowStatus) (domain.Follow, error)
	AllByFollowedUserID(ctx context.Context, followedUserID uuid.UUID, offset int64, limit int32) (domain.Follows, error)
	GetOne(ctx context.Context, followedUserID uuid.UUID, followingUserID uuid.UUID) (domain.Follow, error)
	CountByFollowedUserID(ctx context.Context, followedUserID uuid.UUID) (int64, error)
}

type FollowServiceI interface {
	Unfollow(ctx context.Context, followedUserID uuid.UUID, followingUserID uuid.UUID) (FollowDTO, error)
	CountByFollowedUserID(ctx context.Context, followedUserID uuid.UUID) (int64, error)
	Follow(ctx context.Context, followedUserID uuid.UUID, followingUserID uuid.UUID) (FollowDTO, error)
	Profile(ctx context.Context, followingUserID, followedUserID uuid.UUID) (ProfileDTO, error)
}

type FollowDTO struct {
	ID              int64     `json:"id"`
	FollowedUserID  uuid.UUID `json:"followed_user_id"`
	FollowingUserID uuid.UUID `json:"following_user_id"`

	TracingDTO `json:",inline"`
}

func NewEmptyFollowDTO() FollowDTO {
	return FollowDTO{}
}

func (m *FollowDTO) IsExist() bool {
	return m.ID != 0
}

func (m *FollowDTO) Domain2Port(domain domain.Follow) {
	m.ID = domain.ID
	m.FollowedUserID = domain.FollowedUserID
	m.FollowingUserID = domain.FollowingUserID

	m.TracingDTO.CreatedDate = domain.CreatedDate
	m.TracingDTO.UpdatedDate = domain.UpdatedDate
}

type ProfileDTO struct {
	Name string `json:"name"`
	Bio  string `json:"bio"`

	Following bool `json:"following"`
}

func NewEmptyProfileDTO() ProfileDTO {
	return ProfileDTO{}
}
