package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/prapsky/sawitpro/common/errors"
	"github.com/prapsky/sawitpro/common/response"
	"github.com/prapsky/sawitpro/generated"
)

// (POST /register)
func (s *Server) Register(ctx echo.Context) error {
	req := RegistrationRequest{}
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
		return ctx.JSON(http.StatusInternalServerError, response.NewError(err))
	}

	data := &generated.RegisterResponseData{
		UserID: &userID,
	}

	return ctx.JSON(http.StatusOK, generated.RegisterResponse{
		Data: data})
}
