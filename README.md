## Dependency Management
* Install godep: go get github.com/tools/godep
* Acessing to `./src/suriyun.com/suriyun/whitedoor`
* To restore dependencies: godep restore
* To save dependencies: godep save ./... 

## Project setup
You can create symlink from `./src/suriyun.com` to your $GOPATH
Then build from `./` with `go build` and run `{build-file-name}`
