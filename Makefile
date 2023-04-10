.PHONY:

lint:
	golangci-lint run

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot