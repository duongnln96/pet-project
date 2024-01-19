package domain

import (
	"time"

	"github.com/google/uuid"
)

const (
	EmailInvalidErrUser     string = "email_invalid_err_user"
	PasswordInvalidErrUser  string = "password_invalid_err_user"
	NotFoundErrUser         string = "not_found_err_user"
	ExistedErrUser          string = "existed_err_user"
	LoginInfoInvalidErrUser string = "login_info_invalid_err_user"
)

type UserStatus string

func (m UserStatus) ToString() string {
	return string(m)
}

func NewUserStatusFromString(input string) UserStatus {
	return UserStatus(input)
}

const (
	ActiveUserStatus   UserStatus = "active"
	DeactiveUserStatus UserStatus = "deactive"
	DeletedUserStatus  UserStatus = "deleted"
)

type User struct {
	ID       uuid.UUID
	Name     string
	Email    string
	Password string
	Bio      string
	Status   UserStatus

	CreatedDate *time.Time
	UpdatedDate *time.Time
}

func (m *User) IsExist() bool {
	return m.ID != uuid.Nil
}

func (m *User) IsActive() bool {
	return m.Status == ActiveUserStatus
}
