package auth_token

import (
	"context"
	"log/slog"

	authTokenGen "github.com/duongnln96/blog-realworld/gen/go/auth/v1"
	"github.com/duongnln96/blog-realworld/internal/auth/core/port"
	"github.com/duongnln96/blog-realworld/internal/pkg/serror"
)

func (h *handler) GenAuthToken(ctx context.Context, request *authTokenGen.GenAuthTokenRequest) (*authTokenGen.GenAuthTokenResponse, error) {
	response := authTokenGen.GenAuthTokenResponse{}

	generatedToken, sErr := h.authTokenUC.GenToken(ctx, port.GenAuthTokenRequest{
		UserID:    request.GetUserId(),
		DeviceID:  request.GetDeviceId(),
		UserAgent: request.GetUserAgent(),
		RemoteIP:  request.GetRemoteIp(),
	})
	if sErr != nil {
		serr, ok := sErr.(*serror.SError)
		if !ok {
		}

		if serr.IsSystem {
			slog.ErrorContext(ctx, "system error", "code", serr.Code, "message", serr.Msg)
		} else {
			slog.ErrorContext(ctx, "user error", "code", serr.Code, "message", serr.Msg)
		}

		return &response, serr
	}

	response.JwtToken = generatedToken.JwtToken

	return &response, nil
}
