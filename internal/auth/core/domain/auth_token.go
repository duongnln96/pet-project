package domain

import (
	"time"

	"github.com/google/uuid"
)

const (
	DEFAULT_EXPIRATION_MIN_AUTH_TOKEN = 7 * 24 * 60
)

type JwtAuthToken string

func NewJwtAuthTokenFromStr(token string) JwtAuthToken {
	return JwtAuthToken(token)
}

func (m JwtAuthToken) ToString() string {
	return string(m)
}

type AuthTokenStatus string

func (m AuthTokenStatus) ToString() string {
	return string(m)
}

func NewAuthTokenStatusFromStr(tokenStatus string) AuthTokenStatus {
	return AuthTokenStatus(tokenStatus)
}

const (
	ActiveAuthTokenStatus  AuthTokenStatus = "active"
	ExpiredAuthTokenStatus AuthTokenStatus = "expired"
	DeletedAuthTokenStatus AuthTokenStatus = "deleted"
)

func (m AuthTokenStatus) IsSupportedStatus() bool {
	switch m {
	case ActiveAuthTokenStatus,
		ExpiredAuthTokenStatus,
		DeletedAuthTokenStatus:
		return true
	}

	return false
}

type AuthToken struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	DeviceID    uuid.UUID
	UserAgent   string
	JwtToken    JwtAuthToken
	RemoteIP    string
	ExpiredDate time.Time
	Status      AuthTokenStatus

	CreatedDate *time.Time
	UpdatedDate *time.Time
}

func (m *AuthToken) IsDeleted() bool {
	return m.Status == DeletedAuthTokenStatus
}

func (m *AuthToken) IsExpired() bool {
	return m.Status == ExpiredAuthTokenStatus
}
