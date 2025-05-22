# Makefile for Go project

APP_NAME := blog
GIT_COMMIT := $(shell git rev-parse --short HEAD)

.PHONY: all build clean test run

all: clean build test

build:
	go build -o ./bin/$(APP_NAME) cmd/main.go

run: build
	./bin/$(APP_NAME)

test:
	go test -v -coverprofile coverage.out ./...

clean:
	-@rm -rf ./bin
	go clean
	go mod tidy

docker-up:
	cd ./ops && docker-compose up -d

docker-down:
	cd ./ops && docker-compose down

docker:
	docker build -t $(APP_NAME):latest .
	docker tag $(APP_NAME):latest $(APP_NAME):$(GIT_COMMIT)
	docker tag $(APP_NAME):latest $(APP_NAME):latest
	docker push $(APP_NAME):$(GIT_COMMIT)
	docker push $(APP_NAME):latest