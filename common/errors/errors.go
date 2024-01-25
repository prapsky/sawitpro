package errors

import (
	"errors"
)

var (
	ErrUnexpectedSigning                    = errors.New("Unexpected signing.")
	ErrInvalidRequestPayload                = errors.New("Invalid request payload.")
	ErrPhoneNumberBelowMinimumCharacters    = errors.New("Phone number must be at least 10 characters.")
	ErrPhoneNumberAboveMaximumCharacters    = errors.New("Phone number must be at most 13 characters.")
	ErrPhoneNumberNotIndonesiaCountryCode   = errors.New("Phone number does not start with Indonesia country code +62.")
	ErrPhoneNumberAlreadyRegisterd          = errors.New("Phone number is already registered.")
	ErrPhoneNumberNotRegisterd              = errors.New("Phone number is not registered.")
	ErrFullNameBelowMinimumCharacters       = errors.New("Full name must be at least 3 characters.")
	ErrFullNameAboveMaximumCharacters       = errors.New("Full name must be at most 60 characters.")
	ErrPasswordBelowMinimumCharacters       = errors.New("Password must be at least 6 characters.")
	ErrPasswordAboveMaximumCharacters       = errors.New("Password must be at most 64 characters.")
	ErrPasswordNotContainsSpecialCharacters = errors.New("Password must contain at least 1 capital letter, 1 number, and 1 special character.")
	ErrIncorrectPassword                    = errors.New("Password is incorrect.")
	ErrEmptyToken                           = errors.New("Token is empty.")
	ErrInvalidToken                         = errors.New("Token is invalid.")
)
