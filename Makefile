export GO111MODULE=on

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GODOTENV=godotenv -f dev.env
VERSION ?= $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
PROJECT := $(shell basename "$(PWD)")

BINARY_NAME=${PROJECT}
BINARY_UNIX=$(BINARY_NAME)_unix

SSH_PRIVATE_KEY=`cat ~/.ssh/id_rsa`

all: test build

run:
	godotenv -f .env $(GOCMD) run -ldflags "-X main.version=${VERSION}" cmd/client/main.go

dev:
	$(GODOTENV) $(GOCMD) run -ldflags "-X main.version=${VERSION}" cmd/client/main.go

build: mkdir-build cp
	$(GOBUILD) -o build/${PROJECT}.exe -ldflags "-X main.version=${VERSION}" cmd/client/main.go

test:
	$(GODOTENV) $(GOTEST) -v ./...
clean:
		rm -rf ./build

mkdir-build:
	mkdir build	
	
cp:
	cp audio.wav build/

deps:
	go mod download


dockerbuild:
	docker build --build-arg VERSION=${VERSION} --build-arg PROJECT=${PROJECT} --build-arg SSH_PRIVATE_KEY="${SSH_PRIVATE_KEY}" -t gcr.io/eve-vpn/${PROJECT}:latest -t gcr.io/eve-vpn/${PROJECT}:${VERSION} .

dockerpush:
	docker push gcr.io/eve-vpn/${PROJECT}:${VERSION}
	docker push gcr.io/eve-vpn/${PROJECT}:latest

.PHONY: build
