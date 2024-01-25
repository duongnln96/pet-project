package auth_token

import (
	"context"
	"time"

	"github.com/duongnln96/blog-realworld/internal/auth/core/domain"
	"github.com/google/uuid"
)

func (r *repoManager) UpdateDelete(ctx context.Context, tokenID uuid.UUID, expiredDate time.Time) error {

	if err := authenTokenTable.UpdateQueryContext(ctx, r.db.GetSession(), "status").
		Bind(domain.DeletedAuthTokenStatus, GoogleUUIDMarshaler{tokenID}, expiredDate).
		ExecRelease(); err != nil {
		return err
	}

	return nil
}
