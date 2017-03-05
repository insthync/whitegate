package apis

import (
	"gopkg.in/kataras/iris.v6"
	mgo "gopkg.in/mgo.v2"
)

// ResponseError ... struct for error responses
type ResponseError struct {
	ErrorMessage string `json:"error"`
}

// ResponseSystemError ... Function to response error to clients, return true if error occurs
func ResponseSystemError(ctx *iris.Context, err error) bool {
	if err != nil {
		ctx.JSON(iris.StatusInternalServerError, ResponseError{ErrorMessage: err.Error()})
		return true
	}
	return false
}

// ResponseLoginTokenError ... Function to response error to clients, return true if error occurs
func ResponseLoginTokenError(ctx *iris.Context, err error) bool {
	if err != nil && err.Error() == mgo.ErrNotFound.Error() {
		ctx.JSON(iris.StatusUnauthorized, ResponseError{ErrorMessage: "Invalid Login session."})
		return true
	}
	return ResponseSystemError(ctx, err)
}

// ResponseErrorMessage ... Function to response error to clients
func ResponseErrorMessage(ctx *iris.Context, message string) {
	ctx.JSON(iris.StatusInternalServerError, ResponseError{ErrorMessage: message})
}
