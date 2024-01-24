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

func NewError(err error) generated.ErrorResponse {
	msg := err.Error()
	return generated.ErrorResponse{
		Message: &msg,
	}
}
