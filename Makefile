.PHONY: all test build lint

ldflags="-X github.com/gugahoi/dogwatch/cmd.version=1.0.0"

all: test

build:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -ldflags ${ldflags} -o build/dogwatch-macos
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags ${ldflags} -o build/dogwatch-windows
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags ${ldflags} -o build/dogwatch-linux

test:
	go test -v -race ./...

lint:
	gometalinter.v1 --deadline 120s --enable-all --vendor ./...
