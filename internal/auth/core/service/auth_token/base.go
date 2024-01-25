package auth_token

import (
	"github.com/duongnln96/blog-realworld/internal/auth/core/port"
	"github.com/duongnln96/blog-realworld/internal/pkg/token"

	"github.com/google/wire"
)

var _ port.AuthTokenServiceI = (*service)(nil)

var ServiceSet = wire.NewSet(NewService)

type service struct {
	jwtMaker token.TokenMakerI

	authTokenRepo port.AuthTokenRepoI
}

func NewService(
	authTokenRepoInstance port.AuthTokenRepoI,
	jwtMakerInstance token.TokenMakerI,
) port.AuthTokenServiceI {
	return &service{
		jwtMaker:      jwtMakerInstance,
		authTokenRepo: authTokenRepoInstance,
	}
}
