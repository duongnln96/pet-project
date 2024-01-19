package auth_token

import (
	"context"
	"log/slog"

	authTokenGen "github.com/duongnln96/blog-realworld/gen/go/auth/v1"
	"github.com/duongnln96/blog-realworld/internal/auth/core/port"
	"github.com/duongnln96/blog-realworld/internal/pkg/serror"
)

func (h *handler) ValidateAuthToken(ctx context.Context, request *authTokenGen.ValidateAuthTokenRequest) (*authTokenGen.ValidateAuthTokenResponse, error) {
	// defer slog.InfoContext(ctx, "ValidateAuthToken", "request", request.String())

	response := &authTokenGen.ValidateAuthTokenResponse{}

	validationInfo, err := h.authTokenUC.ValidateToken(ctx, port.ValidateTokenRequest{
		JwtToken: request.GetJwtToken(),
	})
	if err != nil {
		serr, ok := err.(*serror.SError)
		if !ok {
		}

		if serr.IsSystem {
			slog.ErrorContext(ctx, "system error", "code", serr.Code, "message", serr.Msg)
		} else {
			slog.ErrorContext(ctx, "user error", "code", serr.Code, "message", serr.Msg)
		}

		return response, serr
	}

	response.IsValid = validationInfo.IsValid
	return response, nil
}
