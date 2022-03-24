package responses

type Response[T any] struct {
	HttpStatusCode int    `json:"-"`
	Success        bool   `json:"success"`
	Code           string `json:"code"`

	SuccessResponse *successResponse[T] `json:"successResponse"`
	ErrorResponse   *errorResponse[T]   `json:"errorResponse"`
}

func NewResponse[T any]() *Response[T] {
	return &Response[T]{
		// TODO: into constants
		HttpStatusCode: 200,
		Success:        true,
		Code:           "success",

		SuccessResponse: new(successResponse[T]),
		ErrorResponse:   nil,
	}
}

func (h Response[T]) ToErrorResponse(
	httpStatus int,
	code string,
	message string,
	errors []ValidationError,
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

func (h Response[T]) DefaultValidationError() []ValidationError {
	return make([]ValidationError, 0)
}

type successResponse[T any] struct {
	Data T `json:"data"`
}

type errorResponse[T any] struct {
	Message string            `json:"message"`
	Errors  []ValidationError `json:"errors"`
}

type ValidationError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Field   string `json:"field"`
}
