.PHONY: all test build lint

ldflags="-X github.com/gugahoi/dogwatch/cmd.version=snapshot"

all: test

build:
	go build -ldflags ${ldflags} -o build/dogwatch-macos

test:
	go test -v -race ./...

lint:
	gometalinter.v1 --deadline 120s --enable-all --vendor ./...
