VERSION := $(shell git describe --tags --abbrev=0)
COMMIT := $(shell git show --format='%h' --quiet HEAD)
DATE := $(shell git show --format='%ct' --quiet HEAD)

# go tool dist list
GOOS ?= linux
GOARCH ?= amd64

.PHONY: help
help:
	@grep -h '##\ [[:alnum:]]' $(MAKEFILE_LIST) | sed -E 's/(.*):.*##(.*)/\1: \2/' | column -s: -t

.PHONY: build
build: ## build
	@mkdir -p dist
	@for goos in $(GOOS); do \
		for goarch in $(GOARCH); do \
			echo ">> building dist/go-cli_$${goos}_$${goarch}"; \
			GOOS=$$goos GOARCH=$$goarch go build -ldflags="-s -w -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(DATE)" -o "dist/go-cli_$${goos}_$${goarch}"; \
		done; \
	done

.PHONY: fmt
fmt: ## run gofmt
	gofmt -d .

.PHONY: lint
lint: ## run golangci-lint
	golangci-lint run ./...

.PHONY: test
test: ## run go test
	go test -v ./...

.PHONY: update
update: ## run go get -u
	go get -u

.PHONY: tidy
tidy: ## run go mod tidy
	go mod tidy
