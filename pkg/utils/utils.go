package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Utils struct {
}

func NewUtils() *Utils {
	return &Utils{}
}

type ValidateError struct {
	Param   string `json:"param"`
	Message string `json:"message"`
}

func (u *Utils) ValidatorGetCustomMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "this field is required"
	case "email":
		return "invalid email"
	case "min":
		if fe.Field() == "Password" {
			return fmt.Sprintf("%s length password should be %s", fe.Tag(), fe.Param())
		}
		return "this field doesn't meet the minimum length"
	case "max":
		if fe.Field() == "Password" {
			return fmt.Sprintf("%s length password should be %s", fe.Tag(), fe.Param())
		}
		return "this field doesn't meet the maximum length"
	default:
		return fe.Error()
	}
}
