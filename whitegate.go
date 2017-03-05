package whitegate

import (
    "strconv"

    "gopkg.in/kataras/iris.v6"
    "gopkg.in/kataras/iris.v6/adaptors/httprouter"
    "gopkg.in/mgo.v2"

    "fmt"

    "github.com/insthync/whitegate/apis"
    "github.com/insthync/whitegate/configurations"
    "github.com/insthync/whitegate/databases/mongodb"
)

var appConfig *configurations.AppConfiguration
var dbSession *mgo.Session

// Start ... Start the server api
func Start() {
    defer clean()
    var err error
    appConfig, err = configurations.ReadConfig()

    if err != nil {
        panic(err)
    }

    dbSession, err = mongodb.Init(appConfig)
    if err != nil {
        panic(err)
    }

    app := iris.New()
    app.Adapt(iris.DevLogger()) // adapt a logger which prints all errors to the os.Stdout
    app.Adapt(httprouter.New()) // adapt the adaptors/httprouter or adaptors/gorillamux

    app.Post("/login", apis.Login)
    app.Post("/loginWithFacebook", apis.LoginWithFacebook)
    app.Post("/register", apis.Register)
    app.Post("/bindFacebookAccount", apis.BindFacebookAccount)
    app.Post("/unbindFacebookAccount", apis.UnbindFacebookAccount)
    app.Post("/unbindFacebookAccount", apis.UnbindFacebookAccount)

    // Start the server
    app.Listen(":" + strconv.Itoa(appConfig.Port))
}

func clean() {
    fmt.Println("Cleaning...")
    if dbSession != nil {
        dbSession.Close()
    }
}
