package utils

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

func GetFieldErrorsArray(err error) map[string]interface{} {
	var ve validator.ValidationErrors
	fieldErrorsArray := make(map[string]interface{})
	if errors.As(err, &ve) {
		for _, fe := range ve {
			fieldErrorsArray[fe.Field()] = fe.Error()
		}
		return fieldErrorsArray
	}
	return nil
}
