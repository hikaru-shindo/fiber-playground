package handler

import "github.com/go-playground/validator/v10"

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (validator *Validator) Validate(value interface{}) error {
	return validator.validator.Struct(value)
}
