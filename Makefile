.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -timeout 30s ./internal/app/apiserver


.DEFAULT_GOAL := build

