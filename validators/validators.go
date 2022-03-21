package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateOkayTitle(fl validator.FieldLevel) bool {
	return strings.Contains(fl.Field().String(), "okay")
}