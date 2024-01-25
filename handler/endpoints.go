package handler

import (
	goerrors "errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/prapsky/sawitpro/common/errors"
	"github.com/prapsky/sawitpro/common/response"
	"github.com/prapsky/sawitpro/generated"
)

// (POST /register)
func (s *Server) Register(ctx echo.Context) error {
	req := RegisterRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.NewError(errors.ErrInvalidRequestPayload))
	}

	if errValidate := req.validate(); errValidate != nil {
		return ctx.JSON(http.StatusBadRequest, response.NewError(errValidate))
	}

	registerInput, err := req.registerInput()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.NewError(err))
	}

	userID, err := s.service.Register(ctx.Request().Context(), registerInput)
	if err != nil {
		res := response.NewError(err)
		if goerrors.Is(err, errors.ErrPhoneNumberAlreadyRegisterd) {
			return ctx.JSON(http.StatusBadRequest, res)
		}

		return ctx.JSON(http.StatusInternalServerError, res)
	}

	data := &generated.RegisterResponseData{
		UserID: &userID,
	}

	return ctx.JSON(http.StatusOK, generated.RegisterResponse{
		Data: data})
}

// (POST /login)
func (s *Server) Login(ctx echo.Context) error {
	req := LoginRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.NewError(errors.ErrInvalidRequestPayload))
	}

	loginOutput, err := s.service.Login(ctx.Request().Context(), req.loginInput())
	if err != nil {
		res := response.NewError(err)
		if goerrors.Is(err, errors.ErrPhoneNumberNotRegisterd) {
			return ctx.JSON(http.StatusBadRequest, res)
		}

		return ctx.JSON(http.StatusInternalServerError, response.NewError(err))
	}

	data := &generated.LoginResponseData{
		Token:  &loginOutput.Token,
		UserID: &loginOutput.UserID,
	}

	return ctx.JSON(http.StatusOK, generated.LoginResponse{
		Data: data})
}
