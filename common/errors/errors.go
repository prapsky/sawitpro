package errors

import (
	"errors"
)

var (
	ErrInvalidRequestPayload                = errors.New("Invalid request payload")
	ErrPhoneNumberBelowMinimumCharacters    = errors.New("Phone number is below a minimum of 10 characters.")
	ErrPhoneNumberAboveMaximumCharacters    = errors.New("Phone number is above a maximum of 13 characters.")
	ErrPhoneNumberNotIndonesiaCountryCode   = errors.New("Phone number does not start with Indonesia country code +62.")
	ErrFullNameBelowMinimumCharacters       = errors.New("Full name is below a minimum of 3 characters.")
	ErrFullNameAboveMaximumCharacters       = errors.New("Full name is above a maximum of 60 characters.")
	ErrPasswordBelowMinimumCharacters       = errors.New("Password must be at least 6 characters")
	ErrPasswordAboveMaximumCharacters       = errors.New("Password must be at most 64 characters")
	ErrPasswordNotContainsSpecialCharacters = errors.New("Password must contain at least 1 capital letter, 1 number, and 1 special character")
)
