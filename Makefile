.PHONY: test

test:
	@go test -v -short ./...

build:
	@go build ./...
