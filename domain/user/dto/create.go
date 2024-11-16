package user

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type UserCreate struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (a *UserCreate) Validate() error {

	return validation.ValidateStruct(a,
		validation.Field(&a.Email,
			validation.Required.Error("email is required"),
			is.Email.Error("email must be a valid email address"),
		),
		validation.Field(&a.Name,
			validation.Required.Error("name is required"),
			validation.Length(2, 50).Error("name must be between 2 and 50 characters"),
		),
	)
}
