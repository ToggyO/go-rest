package validation

import (
	"github.com/go-playground/validator/v10"
	"go-rest/internal/shared/constants/errors/error_codes"
	"go-rest/internal/shared/constants/errors/error_messages"
	"go-rest/internal/shared/models/responses"
	"net/http"
	"strconv"
	"strings"
)

func ValidateModel[TInputModel interface{}, TReturnModel interface{}](model TInputModel) *responses.Response[TReturnModel] {
	r := responses.NewResponse[TReturnModel]()

	v := validator.New()
	err := v.Struct(model)

	if err != nil {
		var errors []*responses.ValidationError

		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, buildValidationError(err))
		}

		r.ToErrorResponse(
			http.StatusBadRequest,
			error_codes.Global_BadParameters,
			error_messages.Global_BadParameters,
			errors)
	}

	return r
}

func buildValidationError(fe validator.FieldError) *responses.ValidationError {
	ve := &responses.ValidationError{Field: strings.ToLower(fe.Field())}

	// TODO: дополнить
	tag := fe.Tag()
	switch tag {
	case "required":
		ve.Code = error_codes.Validation_FieldNotEmpty
		ve.Message = error_messages.Validation_FieldNotEmpty

	case "email":
		ve.Code = error_codes.Validation_FieldEmail
		ve.Message = error_messages.Validation_FieldEmail

	case "lt", "lte":
		ve.Code = error_codes.Validation_FieldMax
		if tag == "lte" {
			i, _ := strconv.Atoi(fe.Param())
			ve.Message = error_messages.Validation_FieldMax(strconv.Itoa(i + 1))
		} else {
			ve.Message = error_messages.Validation_FieldMax(fe.Param())
		}

	case "gt", "gte":
		ve.Code = error_codes.Validation_FieldMin
		if tag == "gte" {
			i, _ := strconv.Atoi(fe.Param())
			ve.Message = error_messages.Validation_FieldMin(strconv.Itoa(i - 1))
		} else {
			ve.Message = error_messages.Validation_FieldMin(fe.Param())
		}

	}

	return ve
}
