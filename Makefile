# Makefile configuration
.DEFAULT_GOAL := help
.PHONY: deps test vet build-only build travis help

deps: ## Download dependencies
	go get ./...
	go get github.com/stretchr/testify/assert
	go get github.com/fzipp/gocyclo
	go get golang.org/x/lint/golint

test: ## Run unit tests
	go test ./...

vet: ## Code check
	gofmt -s -w .
	go vet ./...
	gocyclo -over 20 .
	golint ./...

build-only: ## Compile binaries
	mkdir -p ./release/
	## Needed to build artifact for alpine linux architecture
	export GOOS=linux; export GOARCH=amd64; \
	go build -o ./release/charlie ./charlie/main.go; \
	go build -o ./release/runner ./runner/runner.go

build: deps vet test build-only ## Full build - download deps, check code, test and then compile

docker-only: ## Making docker image go-binary
	docker build . -t charlie-alpine:latest

docker: build docker-only ## Full build + dockerizing final go image

travis: deps vet test ## Runs all tasks for travis CI

help:
	@grep --extended-regexp '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
