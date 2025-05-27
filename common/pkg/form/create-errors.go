package form

import (
	"errors"
	"strings"
	
	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
)

type Errors map[string][]string

func (e Errors) GetErrors(fieldName string) string {
	fieldErrs, exist := e[fieldName]
	if !exist {
		return ""
	}
	return strings.Join(fieldErrs, "\n")
}

func CreateErrors(err error) Errors {
	result := make(Errors)
	var validatorErrors validator.ValidationErrors
	if !errors.As(err, &validatorErrors) {
		return result
	}
	for _, fieldErr := range validatorErrors {
		fieldName := strcase.ToLowerCamel(fieldErr.Field())
		if _, ok := result[fieldName]; !ok {
			result[fieldName] = make([]string, 0)
		}
		tag := fieldErr.Tag()
		var msg string
		switch tag {
		case "required":
			msg = "Field is required"
		case "email":
			msg = "Invalid e-mail"
		case "gt":
			msg = "Field must be greater than " + fieldErr.Param()
		case "lt":
			msg = "Field must be less than " + fieldErr.Param()
		case "max":
			msg = "Field must be greater than " + fieldErr.Param()
		case "min":
			msg = "Field must be greater than " + fieldErr.Param()
		default:
			msg = "Field is invalid"
		}
		result[fieldName] = append(result[fieldName], msg)
	}
	return result
}
