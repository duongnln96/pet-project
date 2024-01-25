package auth_token

import (
	"context"

	"github.com/duongnln96/blog-realworld/internal/user/core/port"

	authTokenGen "github.com/duongnln96/blog-realworld/gen/go/auth/v1"
)

func (c *authTokenDomainClient) GenAuthToken(ctx context.Context, req *port.GenAuthTokenRequest) (*port.GenAuthTokenResponse, error) {

	authenToken, err := c.domain.GenAuthToken(ctx, &authTokenGen.GenAuthTokenRequest{
		UserId:    req.UserID.String(),
		DeviceId:  req.DeviceID.String(),
		UserAgent: req.UserAgent,
		RemoteIp:  req.RemoteIP,
	})
	if err != nil {
		return nil, err
	}

	return &port.GenAuthTokenResponse{
		JwtToken: authenToken.GetJwtToken(),
	}, nil
}
