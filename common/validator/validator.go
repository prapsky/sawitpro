package validator

import (
	"github.com/go-playground/validator/v10"
)

type FormValidator struct {
	validator *validator.Validate
}

func NewFormValidator() *FormValidator {
	return &FormValidator{
		validator: validator.New(),
	}
}

func (fv *FormValidator) Validate(i interface{}) error {
	return fv.validator.Struct(i)
}
