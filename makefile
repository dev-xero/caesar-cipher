.DEFAULT_GOAL := run

fmt:
	go fmt ./...
.PHONY: fmt

vet: fmt
	go vet ./...
.PHONY: vet

build: vet
	go build -o ./bin/ ./...
.PHONY: build

run: build
	./bin/caesar-cipher
.PHONY: run
