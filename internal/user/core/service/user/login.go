package user

import (
	"context"
	"fmt"
	"time"

	"github.com/duongnln96/blog-realworld/internal/pkg/serror"
	"github.com/duongnln96/blog-realworld/internal/pkg/utils"
	"github.com/duongnln96/blog-realworld/internal/user/core/domain"
	"github.com/duongnln96/blog-realworld/internal/user/core/port"
)

func (s *service) LogIn(ctx context.Context, req port.LoginUserDTO) (string, error) {

	var loginToken string

	err := s.validateEmail(req.Email)
	if err != nil {
		return loginToken, err
	}

	secretKey, ok := s.config.Other.Get("secret_key").(string)
	if !ok {
		return loginToken, serror.NewSystemSError("cannot get password secret key")
	}
	reqPasswordHashed, err := utils.HashPassword(req.Password, secretKey)
	if err != nil {
		return loginToken, serror.NewSystemSError(err.Error())
	}

	user, err := s.userRepo.GetOneByEmail(ctx, req.Email)
	if err != nil {
		return loginToken, serror.NewSystemSError(err.Error())
	}
	if !user.IsExist() {
		return loginToken, serror.NewSError(domain.NotFoundErrUser, "user not found")
	}
	if user.Password != reqPasswordHashed {
		return loginToken, serror.NewSError(domain.PasswordInvalidErrUser, "email or password is invalid")
	}

	// TODO: move to authen service/ constant
	loginToken, _, err = s.jwtMaker.CreateToken(user.ID.String(), 7*24*60*time.Minute)
	if err != nil {
		return loginToken, serror.NewSystemSError(fmt.Sprintf("jwtMaker.CreateToken %s", err.Error()))
	}

	return loginToken, nil
}
