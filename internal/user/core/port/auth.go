package port

import (
	"context"

	"github.com/google/uuid"
)

type (
	AuthTokenDomainI interface {
		GenAuthToken(ctx context.Context, req *GenAuthTokenRequest) (*GenAuthTokenResponse, error)
		ValidateToken(ctx context.Context, req *ValidateTokenRequest) (*ValidateTokenResponse, error)
	}
)

type GenAuthTokenRequest struct {
	UserID    uuid.UUID `json:"user_id"`
	DeviceID  uuid.UUID `json:"device_id"`
	UserAgent string    `json:"user_agent"`
	RemoteIP  string    `json:"remote_ip"`
}

type GenAuthTokenResponse struct {
	JwtToken string `json:"jwt_token"`
}

type ValidateTokenRequest struct {
	JwtToken string `json:"jwt_token"`
}

type ValidateTokenResponse struct {
	IsValid bool `json:"is_valid"`
}
