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

type UpdateProfileRequest struct {
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber" validate:"required"`
	FullName    string `json:"fullName" form:"fullName" validate:"required"`
}

func validatePhoneNumber(phoneNumber string) error {
	const (
		minLength   = 10
		maxLength   = 16
		idCodeRegex = `^\+62`
	)

	if len(phoneNumber) < minLength {
		return errors.ErrPhoneNumberBelowMinimumCharacters
	}

	if len(phoneNumber) > maxLength {
		return errors.ErrPhoneNumberAboveMaximumCharacters
	}

	if !regexp.MustCompile(idCodeRegex).MatchString(phoneNumber) {
		return errors.ErrPhoneNumberNotIndonesiaCountryCode
	}

	return nil
}

func validateFullName(fullName string) error {
	const (
		minLength = 3
		maxLength = 60
	)

	if len(fullName) < minLength {
		return errors.ErrFullNameBelowMinimumCharacters
	}

	if len(fullName) > maxLength {
		return errors.ErrFullNameAboveMaximumCharacters
	}

	return nil
}

func validatePassword(password string) error {
	const (
		minLength = 6
		maxLength = 64
	)

	if len(password) < minLength {
		return errors.ErrPasswordBelowMinimumCharacters
	}

	if len(password) > maxLength {
		return errors.ErrPasswordAboveMaximumCharacters
	}

	if !regexp.MustCompile(`[A-Z]`).MatchString(password) ||
		!regexp.MustCompile(`[0-9]`).MatchString(password) ||
		!regexp.MustCompile(`[^a-zA-Z0-9]`).MatchString(password) {
		return errors.ErrPasswordNotContainsSpecialCharacters
	}

	return nil
}

func (r RegisterRequest) validate() []error {
	var errs []error
	if errPhone := validatePhoneNumber(r.PhoneNumber); errPhone != nil {
		errs = append(errs, errPhone)
	}

	if errName := validateFullName(r.FullName); errName != nil {
		errs = append(errs, errName)
	}

	if errPass := validatePassword(r.Password); errPass != nil {
		errs = append(errs, errPass)
	}

	return errs
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

func (u UpdateProfileRequest) updateProfileInput() service.UpdateProfileInput {
	return service.UpdateProfileInput{
		PhoneNumber: u.PhoneNumber,
		FullName:    u.FullName,
	}
}
