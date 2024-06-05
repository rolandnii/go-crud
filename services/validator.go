package services

import (
	"errors"
	"fmt"
	"strings"
	"gopkg.in/go-playground/validator.v9"
)

var Validate *validator.Validate = validator.New()

type ValidationErrorMessage struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

func Validator(data interface{}) ([]ValidationErrorMessage,error) {
	validationErrors := []ValidationErrorMessage{}

	err := Validate.Struct(data)

	if _, ok := err.(*validator.InvalidValidationError); ok {

		return []ValidationErrorMessage{}, errors.New(err.Error())
	}

	if err != nil {

		for _, err := range err.(validator.ValidationErrors) {
			var elem ValidationErrorMessage

			elem.FailedField = err.Field() 
			elem.Tag = err.Tag()          
			elem.Value = err.Value()       
			elem.Error = true
			
			validationErrors = append(validationErrors, elem)

		}
	}

	return validationErrors, nil

}

func ValidationResponse(errors []ValidationErrorMessage) string {

	errMsgs := make([]string, 0)

	for _, err := range errors {
		errMsgs = append(errMsgs, fmt.Sprintf(
			"[%s]: '%v' | Needs to implement '%s'",
			err.FailedField,
			err.Value,
			err.Tag,
		))
	}

	return strings.Join(errMsgs, " and ")
}
