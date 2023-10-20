package helpers

import "github.com/go-playground/validator/v10"

func ValidateRequest(request interface{}) error {
	validate := validator.New()

	err := validate.Struct(request)
	if err != nil {
		return err
	}

	return nil
}
