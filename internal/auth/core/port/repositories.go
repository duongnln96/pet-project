package port

import (
	"context"

	"github.com/duongnln96/blog-realworld/internal/auth/core/domain"
	"github.com/google/uuid"
)

type AuthTokenRepoI interface {
	Create(context.Context, domain.AuthToken) (domain.AuthToken, error)
	GetOneByPrimary(ctx context.Context, tokenID uuid.UUID) (domain.AuthToken, error)
}
