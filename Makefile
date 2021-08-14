ifeq ($(OS), Windows_NT)
build:
	go build -o my-resource.dll -buildmode=c-shared && move .\my-resource.dll .\resources\go
build-module:
	cd altv-go-module && build.bat && move .\bin\Release\go-module.dll ..\modules
else
build:
	export CGO_LDFLAGS="-g -02 -ldl"
	go build -o my-resource.so -buildmode=c-shared
	rm my-resource.h
endif