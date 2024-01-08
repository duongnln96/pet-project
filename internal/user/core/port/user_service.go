package port

import (
	"context"

	"github.com/duongnln96/blog-realworld/internal/user/core/domain"
	"github.com/google/uuid"
)

type UserServiceI interface {
	LogIn(context.Context, LoginUserDTO) (string, error)
	LogOut(context.Context) error
	Register(context.Context, RegisterUserDTO) (UserDTO, error)
	Update(context.Context, UpdateUserDTO) (UserDTO, error)
	Detail(context.Context, uuid.UUID) (UserDTO, error)
}

type UserDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Bio  string    `json:"bio"`

	TracingDTO `json:",inline"`
}

func NewEmptyUserDTO() UserDTO {
	return UserDTO{
		ID: uuid.Nil,
	}
}

func (m *UserDTO) IsExist() bool {
	return m.ID != uuid.Nil
}

func (m *UserDTO) Domain2Port(domain domain.User) {
	m.ID = domain.ID
	m.Name = domain.Name
	m.Bio = domain.Bio

	m.TracingDTO.CreatedDate = domain.CreatedDate
	m.TracingDTO.UpdatedDate = domain.UpdatedDate
}

type RegisterUserDTO struct {
	Name     string `json:"name" validate:"required,max=500"`
	Bio      string `json:"bio" validate:"omitempty,max=1000"`
	Email    string `json:"email" validate:"required,max=500"`
	Password string `json:"password" validate:"required"`
}

type LoginUserDTO struct {
	Email    string `json:"email" validate:"required,max=500"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserDTO struct {
	ID       uuid.UUID `json:"id"`
	Name     *string   `json:"name"`
	Bio      *string   `json:"bio"`
	Email    *string   `json:"email"`
	Password *string   `json:"password"`
}
