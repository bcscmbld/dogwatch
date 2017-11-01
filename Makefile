all: test

build:
	CGO_ENABLED=0 go build -o dogwatch

test:
	go test -cover -race ./...

lint:
	gometalinter.v1 --deadline 120s --enable-all ./...
