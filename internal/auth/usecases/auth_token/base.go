package auth_token

import (
	"github.com/duongnln96/blog-realworld/internal/auth/port"
	"github.com/duongnln96/blog-realworld/internal/pkg/token"

	"github.com/google/wire"
)

var _ port.AuthTokenUseCasesI = (*usecases)(nil)

var UsecasesSet = wire.NewSet(NewUsecases)

type usecases struct {
	jwtMaker token.TokenMakerI

	authTokenRepo port.AuthTokenRepoI
}

func NewUsecases(
	authTokenRepoInstance port.AuthTokenRepoI,
	jwtMakerInstance token.TokenMakerI,
) port.AuthTokenUseCasesI {
	return &usecases{
		jwtMaker:      jwtMakerInstance,
		authTokenRepo: authTokenRepoInstance,
	}
}
