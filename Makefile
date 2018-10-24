PATH := $(GOPATH)/bin:$(PATH)
GO111MODULE=off

all: build

deps:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure -vendor-only

build:
	go build .
