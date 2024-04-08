VERSION         = latest
BIN             = ./bin/app
BUILD_TIME      = $(shell date -u '+%Y%m%d%H%M%S')
GIT_COMMIT      = $(shell git rev-parse --short HEAD)
IMAGE_NAME      = shaowenchen/go-repo-template:${VERSION}

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

define LDFLAGS
"-X 'github.com/shaowenchen/go-repo-template/main.GitCommit=${GIT_COMMIT}' \
-X 'github.com/shaowenchen/go-repo-template/main.BuildTime=${BUILD_TIME}'"
endef

doc:
	swag init -g ./cmd/server/main.go -o docs

format:
	go fmt $(shell go list ./... | grep -v /vendor/)
	go mod tidy
	go mod vendor

build-web:
	cd web && yarn && yarn build

gen:
	go run cmd/gen/main.go

table:
	go run cmd/table/main.go

run:
	go run cmd/server/main.go -c default.toml

build:
	go build -ldflags ${LDFLAGS} -o $(BIN) cmd/server/main.go

image:
	docker build -t ${IMAGE_NAME} -f ./Dockerfile .

container:
	docker run -it --rm -p 80:80 --env-file .env ${IMAGE_NAME}

push:
	docker push ${IMAGE_NAME}

clear:
	rm -rf ./bin/*
