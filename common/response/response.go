package response

import (
	"github.com/prapsky/sawitpro/generated"
)

type Success struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta"`
}

type Error struct {
	Errors []error     `json:"errors"`
	Meta   interface{} `json:"meta"`
}

type EmptyMeta struct{}

func NewSuccess(data, meta interface{}) *Success {
	return &Success{
		Data: data,
		Meta: meta,
	}
}

func NewError(errs ...error) generated.ErrorResponse {
	var errorResponse generated.ErrorResponse

	for _, err := range errs {
		msg := struct {
			Message string "json:\"message\""
		}{Message: err.Error()}
		errorResponse.Errors = append(errorResponse.Errors, msg)
	}

	return errorResponse
}
