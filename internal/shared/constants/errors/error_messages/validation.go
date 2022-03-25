package error_messages

import (
	"fmt"
	"time"
)

const (
	Validation_FieldNotEmpty      = "The field can't be empty"
	Validation_FieldSizeMax       = "The field is too long"
	Validation_FieldSizeMin       = "The field is too short"
	Validation_FieldInvalidLength = "The field length is not correct"
	Validation_FieldNotValidChars = "The field contains invalid characters"
	Validation_FieldEmail         = "Email isn't valid"
	Validation_FieldCardNumber    = "Card number isn't valid"
	Validation_FieldPhone         = "The phone number isn't valid"
	Validation_FieldDuplicate     = "The field value should be unique"
)

var Validation_FieldMax = func(maxNumber string) string { return fmt.Sprintf("The number can't be greater than %v", maxNumber) }
var Validation_FieldMin = func(minNumber string) string { return fmt.Sprintf("The number can't be less than %v", minNumber) }
var Validation_FieldFuture = func(maxDate time.Time) string { return fmt.Sprintf("The date should be later than %s", maxDate) }
var Validation_FieldPast = func(minDate time.Time) string { return fmt.Sprintf("The date should be early than %s", minDate) }
