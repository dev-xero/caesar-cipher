.DEFAULT_GOAL := run

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o ./bin/ ./...

run: build
	./bin/caesar-cipher
