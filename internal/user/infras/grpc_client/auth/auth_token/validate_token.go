package auth_token

import (
	"context"

	authTokenGen "github.com/duongnln96/blog-realworld/gen/go/auth/v1"
	"github.com/duongnln96/blog-realworld/internal/user/core/port"
)

func (c *authTokenDomainClient) ValidateToken(ctx context.Context, req *port.ValidateTokenRequest) (*port.ValidateTokenResponse, error) {

	validatedToken, err := c.domain.ValidateAuthToken(ctx, &authTokenGen.ValidateAuthTokenRequest{
		JwtToken: req.JwtToken,
	})
	if err != nil {
		return nil, err
	}

	return &port.ValidateTokenResponse{
		IsValid: validatedToken.GetIsValid(),
	}, nil
}
