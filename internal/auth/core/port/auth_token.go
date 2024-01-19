package port

import (
	"context"

	"github.com/duongnln96/blog-realworld/internal/auth/core/domain"
	"github.com/google/uuid"
)

type AuthTokenRepoI interface {
	Create(context.Context, domain.AuthToken) (domain.AuthToken, error)
	GetOneByPrimary(ctx context.Context, tokenID uuid.UUID) (domain.AuthToken, error)
}

type AuthTokenUseCasesI interface {
	GenToken(context.Context, GenAuthTokenRequest) (GenAuthTokenResponse, error)
	ValidateToken(ctx context.Context, req ValidateTokenRequest) (ValidateTokenResponse, error)
}

type GenAuthTokenRequest struct {
	UserID    string `json:"user_id"`
	DeviceID  string `json:"device_id"`
	UserAgent string `json:"user_agent"`
	RemoteIP  string `json:"remote_ip"`
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