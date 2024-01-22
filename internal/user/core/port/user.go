package port

import (
	"context"

	"github.com/duongnln96/blog-realworld/internal/user/core/domain"
	"github.com/google/uuid"
)

type UserRepoI interface {
	Create(context.Context, domain.User) (domain.User, error)
	Update(context.Context, domain.User) (domain.User, error)
	GetOneByID(context.Context, uuid.UUID) (domain.User, error)
	GetOneByEmail(context.Context, string) (domain.User, error)
}

type UserServiceI interface {
	LogIn(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	LogOut(context.Context) error
	Register(context.Context, *RegisterUserRequest) (*UserDTO, error)
	Update(context.Context, *UpdateUserRequest) (*UserDTO, error)
	Detail(context.Context, uuid.UUID) (*UserDTO, error)
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

type RegisterUserRequest struct {
	Name     string
	Bio      string
	Email    string
	Password string
}

type LoginUserRequest struct {
	Email    string
	Password string
}

type LoginUserResponse struct {
	JwtToken string
}

type UpdateUserRequest struct {
	ID       uuid.UUID `json:"id"`
	Name     *string   `json:"name"`
	Bio      *string   `json:"bio"`
	Email    *string   `json:"email"`
	Password *string   `json:"password"`
}
