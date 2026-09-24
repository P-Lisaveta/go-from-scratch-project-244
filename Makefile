GO ?= go

.PHONY: build test lint install clean

build:
	$(GO) build -o bin/gendiff ./cmd/gendiff

test:
	$(GO) test ./...

lint:
	golangci-lint run

install:
	$(GO) install ./cmd/gendiff

clean:
	rm -rf bin
