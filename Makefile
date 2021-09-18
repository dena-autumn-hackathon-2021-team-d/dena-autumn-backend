ENV_FILE := .env
ENV := $(shell cat $(ENV_FILE))

ENV_TEST_FILE := .env.test
ENV_TEST := $(shell cat $(ENV_TEST_FILE))

.PHONY:run
run:
	$(ENV) go run main.go

.PHONY:test
test:
	$(ENV_TEST) go test -count=1 ./...

.PHONY:test-with-coverage
test-with-coverage:
	$(ENV_TEST) go test -covermode=atomic -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o ./cover.html

.PHONY:gen
gen:
	go generate ./...

.PHONY:build
build:
	docker build . -t dena

.PHONY:docker-run
docker-run:
	docker run dena

.PHONY:lint
lint:
	golangci-lint run ./...
