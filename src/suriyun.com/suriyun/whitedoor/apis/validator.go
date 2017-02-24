package apis

import (
	"fmt"

	"gopkg.in/bluesuncorp/validator.v9"
	"gopkg.in/kataras/iris.v6"
)

var validate *validator.Validate

// ValidateForm ... validating form by its definations
func ValidateForm(ctx *iris.Context, formObject interface{}) error {

	if validate == nil {
		fmt.Println("Validate instance is initialized.")
		validate = validator.New()
	}

	err := ctx.ReadForm(formObject)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = validate.Struct(formObject)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
