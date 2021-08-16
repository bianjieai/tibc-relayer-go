#!/usr/bin/make -f

build:
ifeq ($(OS),Windows_NT)
	go build  -o build/relayer.exe .
else
	go build  -o build/relayer .
endif

build-linux: go.sum
	LEDGER_ENABLED=false GOOS=linux GOARCH=amd64 $(MAKE) build

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify

install:
	go build -o relayer && mv relayer $(GOPATH)/bin

format:
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs gofmt -w -s
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs misspell -w
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./lite/*/statik.go" -not -path "*.pb.go" | xargs goimports -w -local github.com/bianjieai/tibc-relayer-go


setup: build-linux
	@docker build -t relayer .
	@rm -rf ./build
