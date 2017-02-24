package apis

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"gopkg.in/kataras/iris.v6"
	"gopkg.in/mgo.v2/bson"
	"suriyun.com/suriyun/whitegate/models/user"
)

type bindFacebookAccountForm struct {
	ID            string `form:"id" validate:"required"`
	LoginToken    string `form:"loginToken" validate:"required"`
	FacebookToken string `form:"facebookToken" validate:"required"`
}

// BindFacebookAccount ... Binding Faceboock account with receving post body
func BindFacebookAccount(ctx *iris.Context) {
	var err error
	response := user.NewEmpty()
	form := bindFacebookAccountForm{}

	err = ValidateForm(ctx, &form)
	if ResponseSystemError(ctx, err) {
		return
	}

	col := user.GetCollection()
	userID := form.ID
	loginToken := form.LoginToken
	facebookToken := form.FacebookToken

	// Check access token with facebook api and get email
	var facebookID string
	var fbResp *http.Response
	fbResp, err = http.Get("https://graph.facebook.com/me?access_token=" + facebookToken + "&fields=id")
	if fbResp.StatusCode != iris.StatusOK {
		return
	}
	var fbRespBody []byte
	defer fbResp.Body.Close()
	fbRespBody, err = ioutil.ReadAll(fbResp.Body)
	fbRespResult := facebookLoginBody{}
	err = json.Unmarshal(fbRespBody, &fbRespResult)
	facebookID = fbRespResult.ID

	var result user.User
	findQuery := user.NewEmpty()
	findQuery.FacebookID = facebookID
	err = col.Find(findQuery).One(&result)
	if err == nil {
		ResponseErrorMessage(ctx, "Facebook ID already bind to another game ID.")
		return
	}
	updateSelectQuery := user.NewEmpty()
	updateSelectQuery.ID = bson.ObjectIdHex(userID)
	updateSelectQuery.LoginToken = loginToken
	updateQueryData := user.NewEmpty()
	updateQueryData.FacebookID = facebookID
	updateQueryData.FacebookToken = facebookToken
	err = col.Update(updateSelectQuery, bson.M{"$set": updateQueryData})
	if ResponseLoginTokenError(ctx, err) {
		return
	}

	// Set response body
	response.ID = bson.ObjectIdHex(userID)
	response.FacebookID = facebookID
	response.FacebookToken = facebookToken

	// Responding
	ctx.JSON(iris.StatusOK, &response)
}
