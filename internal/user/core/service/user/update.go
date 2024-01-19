package user

import (
	"context"
	"strings"

	"github.com/duongnln96/blog-realworld/internal/pkg/serror"
	"github.com/duongnln96/blog-realworld/internal/pkg/utils"
	"github.com/duongnln96/blog-realworld/internal/user/core/domain"
	"github.com/duongnln96/blog-realworld/internal/user/core/port"

	gUtils "github.com/duongnln96/blog-realworld/pkg/utils"
)

func (s *service) Update(ctx context.Context, req *port.UpdateUserDTO) (user *port.UserDTO, err error) {

	domainUser, err := s.userRepo.GetOneByID(ctx, req.ID)
	if err != nil {
		return nil, serror.NewSystemSError(err.Error())
	}
	if !domainUser.IsExist() {
		return nil, serror.NewSError(domain.NotFoundErrUser, "user not found")
	}

	if req.Name != nil {
		domainUser.Name = gUtils.UnicodeNorm(strings.TrimSpace(*req.Name))
	}

	if req.Bio != nil {
		domainUser.Name = gUtils.UnicodeNorm(strings.TrimSpace(*req.Bio))
	}

	if req.Email != nil {

		if err := s.validateEmail(*req.Email); err != nil {
			return nil, err
		}

		domainUser.Email = strings.TrimSpace(*req.Email)
	}

	if req.Password != nil {
		secretKey, ok := s.config.Other.Get("password_secret_key").(string)
		if !ok {
			return nil, serror.NewSystemSError("cannot get password secret key")
		}
		hashPassword, err := utils.HashPassword(*req.Password, secretKey)
		if err != nil {
			return nil, serror.NewSystemSError(err.Error())
		}

		domainUser.Password = hashPassword
	}

	updatedUser, err := s.userRepo.Update(ctx, domainUser)
	if err != nil {
		return nil, serror.NewSystemSError(err.Error())
	}

	var userDTORes = port.NewEmptyUserDTO()

	userDTORes.Domain2Port(updatedUser)

	return &userDTORes, nil
}
