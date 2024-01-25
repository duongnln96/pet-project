package auth_token

import (
	"context"
	"fmt"

	authTokenGen "github.com/duongnln96/blog-realworld/gen/go/auth/v1"
	"github.com/duongnln96/blog-realworld/internal/auth/core/port"
	"github.com/duongnln96/blog-realworld/internal/pkg/serror"
	"github.com/google/uuid"
)

func (h *handler) DeleteAuthToken(ctx context.Context, request *authTokenGen.DeleteAuthTokenRequest) (*authTokenGen.DeleteAuthTokenResponse, error) {

	tokenID, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, serror.NewSError(port.TokenInvalidErrAuthToken, fmt.Sprintf("uuid.Parse %s", err.Error()))
	}

	err = h.authTokenUC.DeleteToken(ctx, tokenID)
	if err != nil {
		return nil, err
	}

	return &authTokenGen.DeleteAuthTokenResponse{}, nil
}
