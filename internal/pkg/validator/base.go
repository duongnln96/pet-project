package validator

import (
	"github.com/go-playground/validator/v10"
)

type sValidator struct {
	v *validator.Validate
}

func NewSValidator() *sValidator {
	return &sValidator{
		v: validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (sv *sValidator) Validate(i interface{}) error {
	return sv.v.Struct(i)
}
