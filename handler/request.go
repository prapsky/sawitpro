package handler

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

type RegistrationRequest struct {
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber" validate:"required,min=10,max=13,phone"`
	FullName    string `json:"fullName" form:"fullName" validate:"required,min=3,max=60"`
	Password    string `json:"password" form:"password" validate:"required,min=6,max=64,password"`
}

func (r RegistrationRequest) validatePhoneNumber(fl validator.FieldLevel) bool {
	const (
		minLength = 10
		maxLength = 13
		idCodeRegex = `^\+62`
	)

	phoneNumber := fl.Field().String()
	return len(phoneNumber) >= minLength && len(phoneNumber) <=  && regexp.MustCompile(idCodeRegex).MatchString(phoneNumber)
}

func (r RegistrationRequest) validatePassword(fl validator.FieldLevel) bool {
	const (
		minLength = 6
		maxLength = 64
		idCodeRegex = `^\+62`
	)

	password := fl.Field().String()
	return len(password) >= minLength && len(password) <= maxLength &&
		regexp.MustCompile(`[A-Z]`).MatchString(password) &&
		regexp.MustCompile(`[0-9]`).MatchString(password) &&
		regexp.MustCompile(`[^a-zA-Z0-9]`).MatchString(password)
}
