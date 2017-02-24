package apis

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"gopkg.in/kataras/iris.v6"
	"gopkg.in/mgo.v2/bson"
	"suriyun.com/suriyun/whitedoor/models/user"
)

type loginWithFacebookForm struct {
	FacebookToken string `form:"facebookToken" validate:"required"`
}

type facebookLoginBody struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

// LoginWithFacebook ... Login with receving post body
func LoginWithFacebook(ctx *iris.Context) {
	var err error
	response := user.NewEmpty()
	form := loginWithFacebookForm{}

	err = ValidateForm(ctx, &form)
	if ResponseSystemError(ctx, err) {
		return
	}

	col := user.GetCollection()
	facebookToken := form.FacebookToken

	// Check access token with facebook api and get email
	var facebookID string
	var email string
	var fbResp *http.Response
	fbResp, err = http.Get("https://graph.facebook.com/me?access_token=" + facebookToken + "&fields=id,email")
	if fbResp.StatusCode != iris.StatusOK {
		return
	}
	var fbRespBody []byte
	defer fbResp.Body.Close()
	fbRespBody, err = ioutil.ReadAll(fbResp.Body)
	fbRespResult := facebookLoginBody{}
	err = json.Unmarshal(fbRespBody, &fbRespResult)
	facebookID = fbRespResult.ID
	email = fbRespResult.Email

	var result user.User
	err = col.Find(bson.M{"$or": []bson.M{bson.M{"facebookId": facebookID}, bson.M{"email": email}}}).One(&result)
	if err != nil {
		// Insert new user if not exists
		result = user.New()
		result.FacebookID = facebookID
		result.FacebookToken = facebookToken
		result.Email = email
		col.Insert(&result)
	} else {
		if len(result.FacebookID) == 0 {
			ResponseErrorMessage(ctx, "Facebook ID was not bound to game ID")
			return
		}
		// Login success, so update login token
		updateQueryData := user.NewEmpty()
		updateQueryData.FacebookToken = facebookToken
		updateQueryData.UpdateLogin()
		updateQuery := bson.M{"$set": updateQueryData}
		err = col.UpdateId(result.ID, updateQuery)
		if ResponseSystemError(ctx, err) {
			return
		}
	}
	// Set response body
	response.ID = result.ID
	response.LoginToken = result.LoginToken

	// Responding
	ctx.JSON(iris.StatusOK, &response)
}
