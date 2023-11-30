package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/haodam/Bank-Go/simplebank/util"
)

var valiCurrency validator.Func = func(FieldLevel validator.FieldLevel) bool {
	if currency, ok := FieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)

	}
	return false
}
