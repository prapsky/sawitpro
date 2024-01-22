package response

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

func NewError(errors ...error) *Error {
	return &Error{
		Errors: errors,
		Meta:   nil,
	}
}
