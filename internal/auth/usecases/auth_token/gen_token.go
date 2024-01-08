package auth_token

import (
	"context"
	"fmt"
	"time"

	"github.com/duongnln96/blog-realworld/internal/auth/domain"
	"github.com/duongnln96/blog-realworld/internal/auth/port"
	"github.com/duongnln96/blog-realworld/internal/pkg/serror"
	"github.com/google/uuid"
)

func (u *usecases) GenToken(ctx context.Context, req port.GenAuthTokenRequest) (port.GenAuthTokenResponse, error) {

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		return port.GenAuthTokenResponse{}, serror.NewSystemSError(fmt.Sprintf("uuid.Parse %s", err.Error()))
	}

	deviceID, err := uuid.Parse(req.DeviceID)
	if err != nil {
		return port.GenAuthTokenResponse{}, serror.NewSystemSError(fmt.Sprintf("uuid.Parse %s", err.Error()))
	}

	token, tokenPayload, err := u.jwtMaker.CreateToken(req.UserID, domain.DEFAULT_EXPIRATION_MIN_AUTH_TOKEN*time.Minute)
	if err != nil {
		return port.GenAuthTokenResponse{}, serror.NewSystemSError(err.Error())
	}

	tokenID, err := uuid.NewRandom()
	if err != nil {
		return port.GenAuthTokenResponse{}, serror.NewSystemSError(err.Error())
	}

	authToken, err := u.authTokenRepo.Create(ctx, domain.AuthToken{
		ID:          tokenID,
		UserID:      userID,
		DeviceID:    deviceID,
		UserAgent:   req.UserAgent,
		JwtToken:    domain.NewJwtAuthTokenFromStr(token),
		RemoteIP:    req.RemoteIP,
		ExpiredDate: tokenPayload.ExpiredAt,
	})
	if err != nil {
		return port.GenAuthTokenResponse{}, serror.NewSystemSError(fmt.Sprintf("authTokenRepo.Create %s", err.Error()))
	}

	return port.GenAuthTokenResponse{
		JwtToken: authToken.JwtToken.ToString(),
	}, nil
}
