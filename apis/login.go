package apis

import (
	"fmt"

	"gopkg.in/kataras/iris.v6"
	"gopkg.in/mgo.v2/bson"
	"github.com/insthync/whitegate/configurations"
	"github.com/insthync/whitegate/models/user"
	"github.com/insthync/whitegate/utils/crypto"
)

type loginForm struct {
	Username string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
}

// Login ... Login with receving post body
func Login(ctx *iris.Context) {
	var err error
	response := user.NewEmpty()
	appConfig := configurations.GetAppConfiguration()
	form := loginForm{}

	err = ValidateForm(ctx, &form)
	if ResponseSystemError(ctx, err) {
		return
	}

	col := user.GetCollection()
	username := form.Username
	password := crypto.GetMD5Hash(appConfig.UserPasswordSalt + form.Password)

	result := user.NewEmpty()
	findQuery := user.NewEmpty()
	findQuery.Username = username
	findQuery.Password = password
	err = col.Find(findQuery).One(&result)
	if ResponseSystemError(ctx, err) {
		return
	}

	fmt.Println("Found User: ", result)
	// Login success, so update login token
	query := user.NewEmpty()
	query.UpdateLogin()
	updateQuery := bson.M{"$set": query}
	err = col.UpdateId(result.ID, updateQuery)
	if ResponseSystemError(ctx, err) {
		return
	}

	// Set response body
	response.ID = result.ID
	response.LoginToken = query.LoginToken

	// Responding
	ctx.JSON(iris.StatusOK, &response)
}
