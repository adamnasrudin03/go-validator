package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) []string {
	if fieldErrors, ok := err.(validator.ValidationErrors); ok {
		resp := make([]string, len(fieldErrors))

		for i, e := range fieldErrors {
			switch e.Tag() {
			case "required":
				resp[i] = fmt.Sprintf("%s is a required field", e.Field())
			case "email":
				resp[i] = fmt.Sprintf("%v must be type %v", e.Field(), e.Tag())
			case "e164":
				resp[i] = fmt.Sprintf("%v must be type phone number", e.Field())
			case "gte":
				resp[i] = fmt.Sprintf("%v must be greater than or equal to %v", e.Field(), e.Param())
			case "lte":
				resp[i] = fmt.Sprintf("%v must be less than or equal to %v", e.Field(), e.Param())
			default:
				resp[i] = fmt.Sprintf("%v is %v %v", e.Field(), e.Tag(), e.Param())
			}
		}

		return resp
	}

	return []string{}
}
