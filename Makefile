export GO111MODULE=on

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GODOTENV=godotenv -f test.env
VERSION ?= $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
PROJECT := $(shell basename "$(PWD)")

BINARY_NAME=${PROJECT}
BINARY_UNIX=$(BINARY_NAME)_unix

SSH_PRIVATE_KEY=`cat ~/.ssh/id_rsa`

all: test build

run:
	$(GODOTENV) $(GOCMD) run -ldflags "-X main.version=${VERSION}" pkg/cmd/main.go

launch:
	$(GODOTENV) $(GOCMD) run -ldflags "-X main.version=${VERSION}" cmd/${PROJECT}.go -launch

build:
		$(GOBUILD) -o build/${PROJECT} -ldflags "-X ${PROJECT}.version=${VERSION}" cmd/${PROJECT}.go
test:
		$(GODOTENV) $(GOTEST) -v ./...
clean:
		rm -rf ./build

deps:
	go mod download


dockerbuild:
	docker build --build-arg VERSION=${VERSION} --build-arg PROJECT=${PROJECT} --build-arg SSH_PRIVATE_KEY="${SSH_PRIVATE_KEY}" -t gcr.io/eve-vpn/${PROJECT}:latest -t gcr.io/eve-vpn/${PROJECT}:${VERSION} .

dockerpush:
	docker push gcr.io/eve-vpn/${PROJECT}:${VERSION}
	docker push gcr.io/eve-vpn/${PROJECT}:latest