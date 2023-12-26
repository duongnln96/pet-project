package utils

import (
	"crypto/sha1"
	"errors"
	"fmt"
)

func HashPassword(raw string, secretKey string) (string, error) {
	if len([]rune(secretKey)) < 6 {
		return "", errors.New(fmt.Sprintf("invlaid secret_key, must inclue at least %d character", 6))
	}

	h := sha1.New()
	h.Write([]byte(fmt.Sprintf("%s%s", raw, secretKey)))
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
