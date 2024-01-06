package domain

import (
	"time"

	"github.com/google/uuid"
)

const (
	NotFoundErrFollow string = "not_found_err_follow"
	ExistedErrFollow  string = "existed_err_follow"
)

type FollowStatus string

func (m FollowStatus) ToString() string {
	return string(m)
}

func NewFollowStatusFromString(input string) FollowStatus {
	return FollowStatus(input)
}

const (
	ActiveFollowStatus   FollowStatus = "active"
	DeactiveFollowStatus FollowStatus = "deactive"
)

type Follow struct {
	ID              int64
	FollowedUserID  uuid.UUID
	FollowingUserID uuid.UUID
	Status          FollowStatus

	CreatedDate *time.Time
	UpdatedDate *time.Time
}

func NewEmptyFollow() Follow {
	return Follow{}
}

type Follows []Follow

func NewEmptyFollows(size int) Follows {
	return make(Follows, 0, size)
}

func (m *Follow) IsExist() bool {
	return m.ID != 0
}
