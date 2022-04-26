# alt:V Go Resource Example
Go module: https://github.com/shockdev04/altv-go-module  
Go package: https://pkg.go.dev/github.com/shockdev04/altv-go-pkg/alt

## Get started
- Clone this repo 
- Build the resource with the following command:  
if you are on windows:  
``go build -o my-resource.dll -buildmode=c-shared``  
if you are on linux:  
``export CGO_LDFLAGS="-g -02 -ldl"``  
``go build -o my-resource.so -buildmode=c-shared``  
- copy resource.dll, client.js and resource.cfg into the resource folder
- Add resouce to server.cfg

## Notice 
To run this resource you need the official alt:V Chat resource

Things should work as in default go.

intern function not working
get config should be sent using protobuf dict & parsed by interface reflect value ptr
add missing vehicle api
add missing core api