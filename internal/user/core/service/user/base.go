package user

import (
	"regexp"

	"github.com/duongnln96/blog-realworld/internal/pkg/serror"
	"github.com/duongnln96/blog-realworld/internal/user/core/domain"
	"github.com/duongnln96/blog-realworld/internal/user/core/port"
	"github.com/duongnln96/blog-realworld/pkg/config"
	"github.com/google/wire"
)

type service struct {
	config *config.Configs

	userRepo        port.UserRepoI
	authTokenDomain port.AuthTokenDomainI
}

var _ port.UserServiceI = (*service)(nil)

var ServiceSet = wire.NewSet(NewService)

func NewService(
	config *config.Configs,

	userRepoInstance port.UserRepoI,
	authTokenDomain port.AuthTokenDomainI,
) port.UserServiceI {

	return &service{
		config: config,

		userRepo:        userRepoInstance,
		authTokenDomain: authTokenDomain,
	}
}

func (s *service) validateEmail(email string) error {
	emailValidator := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	if !emailValidator.MatchString(email) {
		return serror.NewSError(domain.EmailInvalidErrUser, "email invalid format")
	}

	return nil
}
