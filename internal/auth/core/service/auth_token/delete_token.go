package auth_token

import (
	"context"
	"fmt"

	"github.com/duongnln96/blog-realworld/internal/auth/core/port"
	"github.com/duongnln96/blog-realworld/internal/pkg/serror"
	"github.com/google/uuid"
)

func (u *service) DeleteToken(ctx context.Context, tokenID uuid.UUID) error {

	token, err := u.authTokenRepo.GetOneByPrimary(ctx, tokenID)
	if err != nil {
		return serror.NewSystemSError(fmt.Sprintf("authTokenRepo.GetOneByPrimary %s", err.Error()))
	}
	if !token.IsExist() || token.IsDeleted() || token.IsExpired() {
		return serror.NewSError(port.TokenInvalidErrAuthToken, "auth_token is invalid")
	}

	if err := u.authTokenRepo.UpdateDelete(ctx, tokenID, token.ExpiredDate); err != nil {
		return serror.NewSystemSError(fmt.Sprintf("authTokenRepo.UpdateDelete %s", err.Error()))
	}
	return nil
}
