package user

import (
	"context"
	"strings"

	gUtils "github.com/duongnln96/blog-realworld/pkg/utils"

	"github.com/duongnln96/blog-realworld/internal/pkg/serror"
	"github.com/duongnln96/blog-realworld/internal/pkg/utils"
	"github.com/duongnln96/blog-realworld/internal/user/core/domain"
	"github.com/duongnln96/blog-realworld/internal/user/core/port"
	"github.com/google/uuid"
)

func (s *service) Register(ctx context.Context, req *port.RegisterUserRequest) (*port.UserDTO, error) {

	if err := s.validateEmail(req.Email); err != nil {
		return nil, err
	}

	secretKey, ok := s.config.Other.Get("password_secret_key").(string)
	if !ok {
		return nil, serror.NewSystemSError("cannot get password secret key")
	}
	hashPassword, err := utils.HashPassword(req.Password, secretKey)
	if err != nil {
		return nil, serror.NewSystemSError(err.Error())
	}

	user, err := s.userRepo.GetOneByEmail(ctx, req.Email)
	if err != nil {
		return nil, serror.NewSystemSError(err.Error())
	}
	if user.IsExist() {
		return nil, serror.NewSError(domain.NotFoundErrUser, "user with email is existed")
	}

	doaminUser, err := s.userRepo.Create(ctx, domain.User{
		ID:       uuid.New(),
		Name:     gUtils.UnicodeNorm(strings.TrimSpace(req.Name)),
		Email:    strings.TrimSpace(req.Email),
		Password: hashPassword,
		Bio:      gUtils.UnicodeNorm(strings.TrimSpace(req.Bio)),
		Status:   domain.ActiveUserStatus,
	})
	if err != nil {
		return nil, serror.NewSystemSError(err.Error())
	}

	var userDTORes = port.NewEmptyUserDTO()
	userDTORes.Domain2Port(doaminUser)

	return &userDTORes, nil
}
