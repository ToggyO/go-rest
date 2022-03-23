package validation

import (
	"fmt"
	"time"
)

const (
	FieldNotEmpty      = "The field can't be empty"
	FieldSizeMax       = "The field is too long"
	FieldSizeMin       = "The field is too short"
	FieldInvalidLength = "The field length is not correct"
	FieldNotValidChars = "The field contains invalid characters"
	FieldEmail         = "Email isn't valid"
	FieldCardNumber    = "Card number isn't valid"
	FieldPhone         = "The phone number isn't valid"
	FieldDuplicate     = "The field value should be unique"
)

var FieldMax = func(maxNumber int) string { return fmt.Sprintf("The number can't be greater than %v", maxNumber) }
var FieldMin = func(minNumber int) string { return fmt.Sprintf("The number can't be less than %v", minNumber) }
var FieldFuture = func(maxDate time.Time) string { return fmt.Sprintf("The date should be later than %s", maxDate) }
var FieldPast = func(minDate time.Time) string { return fmt.Sprintf("The date should be early than %s", minDate) }
