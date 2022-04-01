package responses

import (
	"go-rest/internal/shared/constants/errors/error_codes"
	"go-rest/internal/shared/constants/errors/error_messages"
	"net/http"
)

func NewNotFoundErrorResponse[T any]() *Response[T] {
	return NewErrorResponse[T](
		http.StatusNotFound,
		error_codes.Global_NotFound,
		error_messages.Global_NotFound,
		DefaultValidationError(),
	)
}
