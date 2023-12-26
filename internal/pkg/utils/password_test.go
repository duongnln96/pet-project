package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	tc := []struct {
		Input  string
		Expect string
	}{
		{
			"123456",
			"",
		},
	}

	_ = tc

	_ = assert.New(t)
}
