package apis

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/mgo.v2/bson"
	"github.com/insthync/whitegate/models/user"
)

type validateLogin struct {
	ID         string `form:"id" validate:"required"`
	LoginToken string `form:"loginToken" validate:"required"`
}

// ValidateLogin ... Validating login with ID and LoginToken
func ValidateLogin(ctx *iris.Context) {
	var err error
	response := user.NewEmpty()
	form := validateLogin{}

	err = ValidateForm(ctx, &form)
	if ResponseSystemError(ctx, err) {
		return
	}

	col := user.GetCollection()
	userID := form.ID
	loginToken := form.LoginToken

	var result user.User
	selectQuery := user.NewEmpty()
	selectQuery.ID = bson.ObjectIdHex(userID)
	selectQuery.LoginToken = loginToken
	err = col.Find(selectQuery).One(&result)
	if ResponseLoginTokenError(ctx, err) {
		return
	}

	// Set response body
	response.ID = bson.ObjectIdHex(userID)
	response.LoginToken = loginToken

	// Responding
	ctx.JSON(iris.StatusOK, &response)
}
