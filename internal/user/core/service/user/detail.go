package user

import (
	"context"

	"github.com/duongnln96/blog-realworld/internal/pkg/serror"
	"github.com/duongnln96/blog-realworld/internal/user/core/domain"
	"github.com/duongnln96/blog-realworld/internal/user/core/port"
	"github.com/google/uuid"
)

func (s *service) Detail(ctx context.Context, req uuid.UUID) (*port.UserDTO, error) {

	userRes := port.NewEmptyUserDTO()

	user, err := s.userRepo.GetOneByID(ctx, req)
	if err != nil {
		return nil, serror.NewSystemSError(err.Error())
	}
	if !user.IsExist() {
		return nil, serror.NewSError(domain.NotFoundErrUser, "user not found")
	}

	userRes.Domain2Port(user)

	return &userRes, nil
}
