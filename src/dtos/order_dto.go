package dtos

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type OrderCreateDTO struct {
	Name   string
	Qty    int32
	Price  int64
	UserID int64
}

func (ucd OrderCreateDTO) Validate() error {
	return validation.ValidateStruct(&ucd,
		validation.Field(&ucd.Name, validation.Required),
		validation.Field(&ucd.Qty, validation.Required, validation.Min(1)),
		validation.Field(&ucd.Price, validation.Required, validation.Min(1)),
		validation.Field(&ucd.UserID, validation.Required),
	)
}

func (ucd OrderUpdateDTO) Validate() error {
	return validation.ValidateStruct(&ucd,
		validation.Field(&ucd.Name, validation.Required),
		validation.Field(&ucd.Qty, validation.Required, validation.Min(1)),
		validation.Field(&ucd.Price, validation.Required, validation.Min(1)),
		validation.Field(&ucd.UserID, validation.Required),
	)
}

type OrderUpdateDTO struct {
	Name   string
	Qty    int32
	Price  int64
	UserID int64
}
