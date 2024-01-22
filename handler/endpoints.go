package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/prapsky/sawitpro/common/errors"
	"github.com/prapsky/sawitpro/common/response"
)

// (POST /registration)
func (s *Server) Registration(ctx echo.Context) error {
	req := RegistrationRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.NewError(errors.ErrInvalidRequestPayload))
	}

	if errs := req.validate(); len(errs) != 0 {
		return ctx.JSON(http.StatusBadRequest, response.NewError(errs...))
	}

	_, err := buildUserEntity(req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.NewError(err))
	}

	return ctx.JSON(http.StatusOK, nil)
}
