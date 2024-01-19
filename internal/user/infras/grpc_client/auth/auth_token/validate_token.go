package auth_token

import (
	"context"

	authTokenGen "github.com/duongnln96/blog-realworld/gen/go/auth/v1"
	"github.com/duongnln96/blog-realworld/internal/user/core/port"
)

func (m *authTokenGRPCClient) ValidateToken(ctx context.Context, req *port.ValidateTokenRequest) (*port.ValidateTokenResponse, error) {

	client := authTokenGen.NewAuthTokenServiceClient(m.conn)

	validatedToken, err := client.ValidateAuthToken(ctx, &authTokenGen.ValidateAuthTokenRequest{
		JwtToken: req.JwtToken,
	})
	if err != nil {
		return nil, err
	}

	return &port.ValidateTokenResponse{
		IsValid: validatedToken.GetIsValid(),
	}, nil
}
