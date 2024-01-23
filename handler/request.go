package handler

import (
	"regexp"
	"time"

	"github.com/prapsky/sawitpro/common/errors"
	"github.com/prapsky/sawitpro/entity"
	"golang.org/x/crypto/bcrypt"
)

type RegistrationRequest struct {
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber" validate:"required,min=10,max=13,phone"`
	FullName    string `json:"fullName" form:"fullName" validate:"required,min=3,max=60"`
	Password    string `json:"password" form:"password" validate:"required,min=6,max=64,password"`
}

func (r RegistrationRequest) validatePhoneNumber() error {
	const (
		minLength   = 10
		maxLength   = 13
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

func (r RegistrationRequest) validateFullName() error {
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

func (r RegistrationRequest) validatePassword() error {
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

func (r RegistrationRequest) validate() []error {
	errors := []error{}

	if errPhone := r.validatePhoneNumber(); errPhone != nil {
		errors = append(errors, errPhone)
	}

	if errName := r.validateFullName(); errName != nil {
		errors = append(errors, errName)
	}

	if errPass := r.validatePassword(); errPass != nil {
		errors = append(errors, errPass)
	}

	return errors
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func comparePasswords(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func buildUserEntity(request RegistrationRequest) (entity.User, error) {
	currentTime := time.Now()

	passwordHash, err := hashPassword(request.Password)
	if err != nil {
		return entity.User{}, err
	}

	return entity.User{
		PhoneNumber:  request.PhoneNumber,
		FullName:     request.FullName,
		PasswordHash: passwordHash,
		CreatedAt:    currentTime,
	}, nil
}
