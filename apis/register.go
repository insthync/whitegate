package apis

import (
	"fmt"

	"gopkg.in/kataras/iris.v6"
	"gopkg.in/mgo.v2/bson"
	"github.com/insthync/whitegate/configurations"
	"github.com/insthync/whitegate/models/user"
	"github.com/insthync/whitegate/utils/crypto"
)

type registerForm struct {
	Username string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
	Email    string `form:"email" validate:"required"`
}

// Register ... Register with receving post body
func Register(ctx *iris.Context) {
	var err error
	response := user.NewEmpty()
	appConfig := configurations.GetAppConfiguration()
	form := registerForm{}

	err = ValidateForm(ctx, &form)
	if ResponseSystemError(ctx, err) {
		return
	}

	fmt.Println("Registering with data: ", form)

	col := user.GetCollection()
	username := form.Username
	password := crypto.GetMD5Hash(appConfig.UserPasswordSalt + form.Password)
	email := form.Email

	var countResult int
	countResult, err = col.Find(bson.M{"$or": []bson.M{bson.M{"username": username}, bson.M{"email": email}}}).Count()
	if ResponseSystemError(ctx, err) {
		return
	}

	if countResult > 0 {
		ResponseErrorMessage(ctx, "Register fail, username or email already in use.")
		return
	}

	// Insert new user if not exists
	newUser := user.New()
	newUser.Username = username
	newUser.Password = password
	newUser.Email = email
	col.Insert(&newUser)

	// Set response body
	response.ID = newUser.ID
	response.LoginToken = newUser.LoginToken

	// Responding
	ctx.JSON(iris.StatusOK, &response)
}
