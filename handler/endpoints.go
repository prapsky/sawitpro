package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// (POST /registration)
func (s *Server) Registration(ctx echo.Context) error {
	req := new(RegistrationRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	validate := validator.New()

	validate.RegisterValidation("phone", validatePhoneNumber)
	validate.RegisterValidation("password", validatePassword)

	if err := validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, resp)
}
