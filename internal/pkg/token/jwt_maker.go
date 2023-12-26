package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	__minSecretKeySize = 32
)

type JWTTokenMaker struct {
	secretKey string
}

var _ TokenMakerI = (*JWTTokenMaker)(nil)

func NewJWTTokenMaker(secretKey string) (TokenMakerI, error) {
	if len([]rune(secretKey)) < __minSecretKeySize {
		return nil, fmt.Errorf("Invalid key size - must be at least %d characters", __minSecretKeySize)
	}
	return &JWTTokenMaker{
		secretKey: secretKey,
	}, nil
}

func (s *JWTTokenMaker) CreateToken(userID string, duration time.Duration) (string, *TokenPayload, error) {

	payload, err := NewTokenPayload(userID, duration)
	if err != nil {
		return "", payload, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(s.secretKey))
	return token, payload, err
}

func (s *JWTTokenMaker) VerifyToken(token string) (*TokenPayload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(s.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &TokenPayload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*TokenPayload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
