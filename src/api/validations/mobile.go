package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/krmsaeed/golang-web-api/common"
)

func IranianMobileNumberValidator(fld validator.FieldLevel) bool {

	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}

	return common.IranianMobileNumberValidate(value)
}
