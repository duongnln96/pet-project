package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// TokenMakerI is an interface
type TokenMakerI interface {
	CreateToken(userID string, duration time.Duration) (string, *TokenPayload, error)
	VerifyToken(token string) (*TokenPayload, error)
}

// Different types of error returned by the VerifyToken function
var (
	ErrInvalidToken = errors.New("Token is invalid")
	ErrExpiredToken = errors.New("Token has expired")
)

type TokenPayload struct {
	ID        uuid.UUID `json:"id"`
	UserID    string    `json:"user_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// NewTokenPayload creates a new token payload with a specific username and duration
func NewTokenPayload(userID string, duration time.Duration) (*TokenPayload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	now := time.Now()

	payload := TokenPayload{
		ID:        tokenID,
		UserID:    userID,
		IssuedAt:  now,
		ExpiredAt: now.Add(duration),
	}

	return &payload, nil
}

func (m *TokenPayload) Valid() error {
	if time.Now().After(m.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
