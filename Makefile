VERSION         = latest
BIN             = ./bin/app
BUILD_TIME      = $(shell date -u '+%Y%m%d%H%M%S')
GIT_COMMIT      = $(shell git rev-parse --short HEAD)
IMAGE_NAME      = shaowenchen/go-repo-template:${VERSION}

define LDFLAGS
"-X 'github.com/shaowenchen/go-repo-template/main.GitCommit=${GIT_COMMIT}' \
-X 'github.com/shaowenchen/go-repo-template/main.BuildTime=${BUILD_TIME}'"
endef

format:
	go fmt $(go list ./... | grep -v /vendor/)
	go mod tidy
	go mod vendor

run:
	go run main.go -c default.toml

binary:
	go build -ldflags ${LDFLAGS} -o $(BIN) ./main.go

image:
	docker build -t ${IMAGE_NAME} -f ./Dockerfile .

container:
	docker run -it --rm -p 80:80 ${IMAGE_NAME}

push:
	docker push ${IMAGE_NAME}

clear:
	rm -rf ./bin/*
