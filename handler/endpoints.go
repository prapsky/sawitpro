package handler

import (
	goerrors "errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/prapsky/sawitpro/common/errors"
	"github.com/prapsky/sawitpro/common/response"
)

// (POST /register)
func (s *Server) Register(ctx echo.Context) error {
	req := RegisterRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.NewError(errors.ErrInvalidRequestPayload))
	}

	if errs := req.validate(); len(errs) != 0 {
		return ctx.JSON(http.StatusBadRequest, response.NewError(errs...))
	}

	registerInput, err := req.registerInput()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.NewError(err))
	}

	registerResponse, err := s.service.Register(ctx.Request().Context(), registerInput)
	if err != nil {
		res := response.NewError(err)
		if goerrors.Is(err, errors.ErrPhoneNumberAlreadyRegistered) {
			return ctx.JSON(http.StatusBadRequest, res)
		}

		return ctx.JSON(http.StatusInternalServerError, res)
	}

	return ctx.JSON(http.StatusOK, registerResponse)
}

// (POST /login)
func (s *Server) Login(ctx echo.Context) error {
	req := LoginRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.NewError(errors.ErrInvalidRequestPayload))
	}

	loginResponse, err := s.service.Login(ctx.Request().Context(), req.loginInput())
	if err != nil {
		res := response.NewError(err)
		if goerrors.Is(err, errors.ErrPhoneNumberNotRegistered) {
			return ctx.JSON(http.StatusBadRequest, res)
		}

		if goerrors.Is(err, errors.ErrIncorrectPassword) {
			return ctx.JSON(http.StatusBadRequest, res)
		}

		return ctx.JSON(http.StatusInternalServerError, response.NewError(err))
	}

	return ctx.JSON(http.StatusOK, loginResponse)
}

// (GET /profile)
func (s *Server) Profile(ctx echo.Context) error {
	token := s.getToken(ctx)
	if token == "" {
		return ctx.JSON(http.StatusForbidden, response.NewError(errors.ErrEmptyToken))
	}

	profileResponse, err := s.service.GetProfile(ctx.Request().Context(), token)
	if err != nil {
		res := response.NewError(err)
		if goerrors.Is(err, errors.ErrInvalidToken) {
			return ctx.JSON(http.StatusForbidden, res)
		}

		if goerrors.Is(err, errors.ErrPhoneNumberNotRegistered) {
			return ctx.JSON(http.StatusForbidden, res)
		}

		return ctx.JSON(http.StatusInternalServerError, res)
	}

	return ctx.JSON(http.StatusOK, profileResponse)
}

// (PATCH /profile)
func (s *Server) UpdateProfile(ctx echo.Context) error {
	token := s.getToken(ctx)
	if token == "" {
		return ctx.JSON(http.StatusForbidden, response.NewError(errors.ErrEmptyToken))
	}

	req := UpdateProfileRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.NewError(errors.ErrInvalidRequestPayload))
	}

	if req.PhoneNumber != "" {
		if errPhone := validatePhoneNumber(req.PhoneNumber); errPhone != nil {
			return ctx.JSON(http.StatusForbidden, response.NewError(errPhone))
		}
	}

	if req.FullName != "" {
		if errName := validateFullName(req.FullName); errName != nil {
			return ctx.JSON(http.StatusForbidden, response.NewError(errName))
		}
	}

	profileResponse, err := s.service.UpdateProfile(ctx.Request().Context(), token, req.updateProfileInput())
	if err != nil {
		res := response.NewError(err)
		if goerrors.Is(err, errors.ErrInvalidToken) {
			return ctx.JSON(http.StatusForbidden, res)
		}

		if goerrors.Is(err, errors.ErrPhoneNumberNotRegistered) {
			return ctx.JSON(http.StatusForbidden, res)
		}

		if goerrors.Is(err, errors.ErrPhoneNumberAlreadyRegistered) {
			return ctx.JSON(http.StatusConflict, res)
		}

		return ctx.JSON(http.StatusInternalServerError, res)
	}

	return ctx.JSON(http.StatusOK, profileResponse)
}

func (s *Server) getToken(ctx echo.Context) string {
	key := "Bearer "
	reqToken := ctx.Request().Header.Get("Authorization")
	if !strings.HasPrefix(reqToken, key) {
		return ""
	}

	token := strings.TrimPrefix(reqToken, key)
	return token
}
