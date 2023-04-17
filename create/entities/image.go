package entities

import (
	"github.com/go-playground/validator/v10"
)

type Image struct {
	Data string `json:"data,omitempty" bson:"data,omitempty" validate:"required"`
	Alt  string `json:"alt,omitempty" bson:"alt,omitempty" validate:"required"`
}

type ValidateErr struct {
	Param   string `json:"param"`
	Message string `json:"message"`
}

func mapErrorMsg(fe validator.FieldError) string {
	msgs := map[string]string{
		"required": "is required",
	}

	if val, ok := msgs[fe.Tag()]; ok {
		return val
	}

	return "unknown validation error "
}

func (s *Image) Validate() []ValidateErr {
	validate := validator.New()

	if err := validate.Struct(s); err != nil {

		validateErrs := make([]ValidateErr, len(err.(validator.ValidationErrors)))

		for i, fe := range err.(validator.ValidationErrors) {
			validateErrs[i] = ValidateErr{Param: fe.Field(), Message: mapErrorMsg(fe)}
		}

		return validateErrs
	}

	return nil
}
