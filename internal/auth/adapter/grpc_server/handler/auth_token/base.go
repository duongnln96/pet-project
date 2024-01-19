package auth_token

import (
	authTokenGen "github.com/duongnln96/blog-realworld/gen/go/auth/v1"
	"github.com/duongnln96/blog-realworld/internal/auth/core/port"
	"github.com/google/wire"
)

var HandlerSet = wire.NewSet(NewHandler)

var _ authTokenGen.AuthTokenServiceServer = (*handler)(nil)

type handler struct {
	authTokenGen.UnimplementedAuthTokenServiceServer

	authTokenUC port.AuthTokenUseCasesI
}

func NewHandler(
	authTokenUCInstance port.AuthTokenUseCasesI,
) authTokenGen.AuthTokenServiceServer {
	return &handler{
		authTokenUC: authTokenUCInstance,
	}
}
