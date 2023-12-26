package user

import (
	"log"
	"regexp"

	"github.com/duongnln96/blog-realworld/internal/pkg/serror"
	"github.com/duongnln96/blog-realworld/internal/pkg/token"
	"github.com/duongnln96/blog-realworld/internal/user/core/domain"
	"github.com/duongnln96/blog-realworld/internal/user/core/port"
	"github.com/duongnln96/blog-realworld/pkg/config"
	"github.com/google/wire"
)

type service struct {
	config *config.Configs

	jwtMaker token.TokenMakerI
	userRepo port.UserRepoI
}

var _ port.UserServiceI = (*service)(nil)

var ServiceSet = wire.NewSet(NewService)

func NewService(
	config *config.Configs,

	userRepoInstance port.UserRepoI,
) port.UserServiceI {

	jwtSecretKey, ok := config.Other.Get("jwt_secret_key").(string)
	if !ok {
		log.Panicf("jwt_secret_key invalid")
	}
	jwtMaker, err := token.NewJWTTokenMaker(jwtSecretKey)
	if err != nil {
		log.Panicf("token.NewJWTTokenMaker %s", err.Error())
	}

	return &service{
		config: config,

		jwtMaker: jwtMaker,
		userRepo: userRepoInstance,
	}
}

func (s *service) validateEmail(email string) error {
	emailValidator := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	if !emailValidator.MatchString(email) {
		return serror.NewSError(domain.EmailInvalidErrUser, "email invalid format")
	}

	return nil
}
