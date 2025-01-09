.PHONY: build api all

all: api

export CONFIG_DIR=./
export SKIP_LOG=8

build:
	go build -o ./build/arabian-snooker .

format:
	goimports -w .
	gofmt -s -w .

lint:
	golangci-lint run --timeout 10m

api:
	./build/arabian-snooker api