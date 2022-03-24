package responses

import (
	"go-rest/internal/shared/constants/errors/error_codes"
	"go-rest/internal/shared/constants/errors/error_messages"
	"net/http"
)

func NewNotFoundErrorResponse[T any]() *Response[T] {
	r := &Response[T]{
		Success:         false,
		HttpStatusCode:  http.StatusNotFound,
		Code:            error_codes.Global_NotFound,
		SuccessResponse: nil,
		ErrorResponse: &errorResponse[T]{
			Message: error_messages.Global_NotFound,
		},
	}
	r.ErrorResponse.Errors = r.DefaultValidationError()
	return r
}
