package response

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

type Response struct {
	Message string              `json:"message,omitempty"`
	Errors  map[string][]string `json:"errors,omitempty"`
}

// %f - field
// %ts - target size
var validationTags = map[string]string{
	"required":     "field %f is a required field",
	"email":        "field %f must be an valid email",
	"min":          "field %f can't be less than %ts",
	"max":          "field %f can't be more than %ts",
	"eqfield":      "must be equal field %ts",
	"defaultError": "field %f is not valid",
}

//func OK() Response {
//	return Response{
//		Status: StatusOK,
//	}
//}

func Error(msg string) Response {
	return Response{
		Message: msg,
	}
}

func ValidationErrors(msg string, errors map[string][]string) Response {
	return Response{
		Message: msg,
		Errors:  errors,
	}
}

func ValidationError(errs validator.ValidationErrors) Response {
	errors := make(map[string][]string)
	for _, err := range errs {
		var vErr string

		if msg, ok := validationTags[err.Tag()]; ok {
			vErr = strings.ReplaceAll(msg, "%f", err.Field())
			vErr = strings.ReplaceAll(vErr, "%ts", err.Param())
		} else {
			vErr = strings.ReplaceAll(validationTags["defaultError"], "%f", err.Field())
		}

		if errArray, ok := errors[err.Field()]; ok {
			errArray = append(errArray, vErr)
		} else {
			errors[err.Field()] = []string{vErr}
		}
	}

	return ValidationErrors("Some fields have errors", errors)
}
