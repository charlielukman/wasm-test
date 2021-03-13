BINARY='ascii_pic.wasm'
all: build

build:
	GOARCH=wasm GOOS=js go build -o ${BINARY} main.go
