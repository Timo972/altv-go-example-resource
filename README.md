# alt:V Go Resource Example
Go module: https://github.com/shockdev04/altv-go-module
Go package: https://pkg.go.dev/github.com/shockdev04/altv-go-pkg/alt

## Get started
- Clone this repo 
- Add resouce to server.cfg
- Build the resource with the following command:  
if you are on windows:  
``go build -o my-resource.dll -buildmode=c-shared``  
if you are on linux:  
``export CGO_LDFLAGS="-g -02 -ldl"``  
``go build -o my-resource.so -buildmode=c-shared``  
- copy resource.dll and resource.cfg into the resource folder

Things should work as in default go.