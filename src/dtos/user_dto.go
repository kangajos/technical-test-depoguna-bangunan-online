package dtos

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type UserCreateDTO struct {
	Name     string
	UserName string
	Password string
}

func (ucd UserCreateDTO) Validate() error {
	return validation.ValidateStruct(&ucd,
		validation.Field(&ucd.Name, validation.Required),
		validation.Field(&ucd.UserName, validation.Required),
		validation.Field(&ucd.Password, validation.Required),
	)
}
func (ucd UserUpdateDTO) Validate() error {
	return validation.ValidateStruct(&ucd,
		validation.Field(&ucd.Name, validation.Required),
		validation.Field(&ucd.UserName, validation.Required),
		validation.Field(&ucd.Password),
	)
}

type UserUpdateDTO struct {
	Name     string
	UserName string
	Password string
}
