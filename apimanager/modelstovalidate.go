package apimanager

import "github.com/go-playground/validator"

type (
	// UserRegister this struct help you to validate user email and password
	UserRegister struct {
		Email    string `json:"Email" validate:"required,email"`
		Password string `json:"assword" validate:"required"`
	}
	CustomValidator struct {
		validator *validator.Validate
	}
)

//Validate : use validator package to validate data that it receive from client
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
