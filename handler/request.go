package handler

import (
	"regexp"

	"github.com/prapsky/sawitpro/common/errors"
	"github.com/prapsky/sawitpro/service"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber" validate:"required"`
	FullName    string `json:"fullName" form:"fullName" validate:"required"`
	Password    string `json:"password" form:"password" validate:"required"`
}

type LoginRequest struct {
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber" validate:"required"`
	Password    string `json:"password" form:"password" validate:"required"`
}

func (r RegisterRequest) validatePhoneNumber() error {
	const (
		minLength   = 10
		maxLength   = 16
		idCodeRegex = `^\+62`
	)

	if len(r.PhoneNumber) < minLength {
		return errors.ErrPhoneNumberBelowMinimumCharacters
	}

	if len(r.PhoneNumber) > maxLength {
		return errors.ErrPhoneNumberAboveMaximumCharacters
	}

	if !regexp.MustCompile(idCodeRegex).MatchString(r.PhoneNumber) {
		return errors.ErrPhoneNumberNotIndonesiaCountryCode
	}

	return nil
}

func (r RegisterRequest) validateFullName() error {
	const (
		minLength = 3
		maxLength = 60
	)

	if len(r.FullName) < minLength {
		return errors.ErrFullNameBelowMinimumCharacters
	}

	if len(r.FullName) > maxLength {
		return errors.ErrFullNameAboveMaximumCharacters
	}

	return nil
}

func (r RegisterRequest) validatePassword() error {
	const (
		minLength = 6
		maxLength = 64
	)

	if len(r.Password) < minLength {
		return errors.ErrPasswordBelowMinimumCharacters
	}

	if len(r.Password) > maxLength {
		return errors.ErrPasswordAboveMaximumCharacters
	}

	if !regexp.MustCompile(`[A-Z]`).MatchString(r.Password) ||
		!regexp.MustCompile(`[0-9]`).MatchString(r.Password) ||
		!regexp.MustCompile(`[^a-zA-Z0-9]`).MatchString(r.Password) {
		return errors.ErrPasswordNotContainsSpecialCharacters
	}

	return nil
}

func (r RegisterRequest) validate() error {
	if errPhone := r.validatePhoneNumber(); errPhone != nil {
		return errPhone
	}

	if errName := r.validateFullName(); errName != nil {
		return errName
	}

	if errPass := r.validatePassword(); errPass != nil {
		return errPass
	}

	return nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (r RegisterRequest) registerInput() (service.RegisterInput, error) {
	passwordHash, err := hashPassword(r.Password)
	if err != nil {
		return service.RegisterInput{}, err
	}

	return service.RegisterInput{
		PhoneNumber:  r.PhoneNumber,
		FullName:     r.FullName,
		PasswordHash: passwordHash,
	}, nil
}

func (l LoginRequest) loginInput() service.LoginInput {
	return service.LoginInput{
		PhoneNumber: l.PhoneNumber,
		Password:    l.Password,
	}
}
