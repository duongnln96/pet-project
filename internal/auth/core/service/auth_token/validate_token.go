package auth_token

import (
	"context"
	"fmt"

	"github.com/duongnln96/blog-realworld/internal/auth/core/port"
	"github.com/duongnln96/blog-realworld/internal/pkg/serror"
)

func (u *usecases) ValidateToken(ctx context.Context, req port.ValidateTokenRequest) (port.ValidateTokenResponse, error) {

	response := port.ValidateTokenResponse{
		IsValid: false,
	}

	tokenPayload, err := u.jwtMaker.VerifyToken(req.JwtToken)
	if err != nil {
		return response, nil
	}

	authToken, err := u.authTokenRepo.GetOneByPrimary(ctx, tokenPayload.ID)
	if err != nil {
		return response, serror.NewSystemSError(fmt.Sprintf("authTokenRepo.GetOneByPrimary %s", err.Error()))
	}
	if authToken.IsExpired() || authToken.IsDeleted() {
		return response, nil
	}

	return port.ValidateTokenResponse{
		IsValid: true,
	}, nil
}
