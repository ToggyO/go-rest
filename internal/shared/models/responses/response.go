package responses

import "net/http"

type successResponse[T any] struct {
	Data T `json:"data"`
}

type errorResponse[T any] struct {
	Message string             `json:"message"`
	Errors  []*ValidationError `json:"errors"`
}

type ValidationError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Field   string `json:"field"`
}

type Response[T any] struct {
	HttpStatusCode int    `json:"-"`
	Success        bool   `json:"success"`
	Code           string `json:"code"`

	SuccessResponse *successResponse[T] `json:"successResponse"`
	ErrorResponse   *errorResponse[T]   `json:"errorResponse"`
}

func NewResponse[T interface{}]() *Response[T] {
	return &Response[T]{
		HttpStatusCode: http.StatusOK,
		Success:        true,
		Code:           "success",

		SuccessResponse: new(successResponse[T]),
		ErrorResponse:   nil,
	}
}

func NewErrorResponse[T interface{}](
	httpStatus int,
	code string,
	message string,
	errors []*ValidationError,
) *Response[T] {
	return &Response[T]{
		Success:         false,
		HttpStatusCode:  httpStatus,
		Code:            code,
		SuccessResponse: nil,
		ErrorResponse: &errorResponse[T]{
			Message: message,
			Errors:  errors,
		},
	}
}

func (r *Response[T]) ToErrorResponse(
	httpStatus int,
	code string,
	message string,
	errors []*ValidationError,
) {
	r.Success = false
	r.HttpStatusCode = httpStatus
	r.Code = code
	r.SuccessResponse = nil
	r.ErrorResponse = &errorResponse[T]{
		Message: message,
		Errors:  errors,
	}
}

func DefaultValidationError() []*ValidationError {
	return make([]*ValidationError, 0)
}
