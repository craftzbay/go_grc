package helpers

import (
	"net/mail"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// validate struct
func Validate(data interface{}) error {
	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return validationErrors
	}
	return nil
}

// validate email
func ValidEmail(str string) bool {
	_, err := mail.ParseAddress(str)
	return err == nil
}

// validate phone
func ValidPhone(str string) bool {
	ok, _ := regexp.MatchString(`^\d{8}$`, str)
	return ok
}

// validate regNo
func ValidRegNo(str string) bool {
	ok, _ := regexp.MatchString("^[а-яА-ЯөӨүҮ]{2}[0-9]{8}$", str)
	return ok
}
