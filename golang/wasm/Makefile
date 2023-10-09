GOROOT = $(shell go env GOROOT)

build:
	GOOS=js GOARCH=wasm go build -o main.wasm

install:
	cp main.wasm ../../node/
