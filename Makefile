GO ?= go
COVERAGE_MIN ?= 80
COVERAGE_FILE ?= coverage.out

.PHONY: build test test-coverage lint install clean

build:
	$(GO) build -o bin/gendiff ./cmd/gendiff

test:
	$(GO) test ./...

test-coverage:
	$(GO) test -coverprofile=$(COVERAGE_FILE) $$(go list ./... | grep -v '/cmd/')
	@coverage=$$($(GO) tool cover -func=$(COVERAGE_FILE) | awk '/^total:/ { gsub("%", "", $$3); print $$3 }'); \
	awk -v coverage="$$coverage" -v min="$(COVERAGE_MIN)" 'BEGIN { \
		if (coverage < min) { \
			printf "Coverage %.1f%% is below minimum %.1f%%\n", coverage, min; \
			exit 1; \
		} \
		printf "Coverage %.1f%%\n", coverage; \
	}'

lint:
	golangci-lint run

install:
	$(GO) install ./cmd/gendiff

clean:
	rm -rf bin coverage.out
