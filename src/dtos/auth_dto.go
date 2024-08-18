package dtos

import validation "github.com/go-ozzo/ozzo-validation"

type AuthDTO struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func (ad AuthDTO) Validate() error {
	return validation.ValidateStruct(&ad,
		validation.Field(&ad.UserName, validation.Required),
		validation.Field(&ad.Password, validation.Required),
	)
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
}
