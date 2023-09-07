package valid

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

func NewValidator() *validator.Validate {
	valid := validator.New(validator.WithRequiredStructEnabled())
	valid.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	return valid
}
