ifeq ($(OS), Windows_NT)
build:
	go build -o my-resource.dll -buildmode=c-shared
else
build:
	export CGO_LDFLAGS="-g -02 -ldl"
	go build -o my-resource.so -buildmode=c-shared
endif