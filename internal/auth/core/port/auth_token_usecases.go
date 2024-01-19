package port

import (
	"context"
)

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
