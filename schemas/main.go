package schemas

import "github.com/go-playground/validator/v10"

var validate = validator.New();
func ValidateFields(fields interface{}) error {
	return validate.Struct(fields)
}