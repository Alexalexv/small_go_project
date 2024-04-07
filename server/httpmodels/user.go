package httpmodels

import (
	"github.com/go-playground/validator/v10"
)

type CreateUserRequest struct {
	Name    string `json:"Name" validate:"required"`
	Surname string `json:"Surname" validate:"required"`
	Age     int    `json:"Age" validate:"gte=18"`
}

func (cus *CreateUserRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(cus)
	if err != nil {
		return err
	}

	return nil
}

type GetUserRequest struct {
	Id      int    `json:"Id"`
	Name    string `json:"Name"`
	Surname string `json:"Surname"`
	Age     int    `json:"Age"`
}
