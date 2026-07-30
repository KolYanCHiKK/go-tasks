package validations

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func GmailEmailValidate(f1 validator.FieldLevel) bool {
	params := f1.Param()
	exp, err := regexp.Compile(fmt.Sprintf(`^.*%s\..*$`, params))
	if err != nil {
		return false
	}
	match := exp.MatchString(f1.Field().String())
	if !match {
		return false
	}
	return true
}
