## Requirement
* Golang 1.8 or above
* MongoDB
* Godep (https://github.com/tools/godep) -> restore or following packages
* Mgo (http://labix.org/mgo)
* Iris (https://github.com/kataras/iris)
* UUID (https://github.com/satori/go.uuid)

## Dependency Management
* Install godep: go get github.com/tools/godep
* Acessing to `./src/suriyun.com/suriyun/whitedoor`
* To restore dependencies: godep restore
* To save dependencies: godep save ./... 

## Project setup
You can create symlink from `./src/suriyun.com` to your $GOPATH
Then build from `./` with `go build` and run `{build-file-name}`

## How to test
After launch and build you can try to:
* Register new user with path http://localhost:6201/register with POST method and post data: username(string), password(string), email(string)
* Login with path http://localhost:6201/login with POST method and post data: username(string), password(string)
* login with facebook with path http://localhost:6201/loginWithFacebook with POST method and post data: facebookToken(string)
* bind Facebook account with path http://localhost:6201/bindFacebookAccount with POST method and post data: id(string), loginToken(string), facebookToken(string)
* unbind Facebook account with path http://localhost:6201/unbindFacebookAccount with POST method and post data: id(string), loginToken(string)
