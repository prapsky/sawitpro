package errors

import (
	pkgErrors "github.com/pkg/errors"
)

var (
	ErrInvalidRequestPayload                = pkgErrors.New("Invalid request payload")
	ErrPhoneNumberBelowMinimumCharacters    = pkgErrors.New("Phone number is below a minimum of 10 characters.")
	ErrPhoneNumberAboveMaximumCharacters    = pkgErrors.New("Phone number is above a maximum of 13 characters.")
	ErrPhoneNumberNotIndonesiaCountryCode   = pkgErrors.New("Phone number does not start with Indonesia country code +62.")
	ErrPasswordBelowMinimumCharacters       = pkgErrors.New("Password must be at least 6 characters")
	ErrPasswordAboveMaximumCharacters       = pkgErrors.New("Password must be at most 64 characters")
	ErrPasswordNotContainsSpecialCharacters = pkgErrors.New("Password must contain at least 1 capital letter, 1 number, and 1 special character")
)
