package api

import (
	"github.com/Golang/bank_management_system/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, isOk := fieldLevel.Field().Interface().(string); isOk {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
