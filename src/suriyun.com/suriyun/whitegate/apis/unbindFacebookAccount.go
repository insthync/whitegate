package apis

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/mgo.v2/bson"
	"suriyun.com/suriyun/whitedoor/models/user"
)

type unbindFacebookAccountForm struct {
	ID         string `form:"id" validate:"required"`
	LoginToken string `form:"loginToken" validate:"required"`
}

// UnbindFacebookAccount ... Unbinding Faceboock account with receving post body
func UnbindFacebookAccount(ctx *iris.Context) {
	var err error
	response := user.NewEmpty()
	form := unbindFacebookAccountForm{}

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
	if ResponseSystemError(ctx, err) {
		return
	}

	if len(result.Username) == 0 {
		ResponseErrorMessage(ctx, "Can not unbind, Game ID username is empty")
		return
	}

	err = col.UpdateId(result.ID, bson.M{"$unset": bson.M{"facebookId": "", "facebookToken": ""}})
	if ResponseLoginTokenError(ctx, err) {
		return
	}

	// Set response body
	response.ID = bson.ObjectIdHex(userID)

	// Responding
	ctx.JSON(iris.StatusOK, &response)
}
