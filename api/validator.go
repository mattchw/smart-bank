package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/mattchw/smart-bank/util"
)

func validateCurrency(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		return util.IsValidCurrency(currency)
	}
	return false
}
