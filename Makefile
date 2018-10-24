PATH := $(GOPATH)/bin:$(PATH)
GO111MODULE=off

all: build

deps:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure -vendor-only

build:
	go build .

test:
	go test ./...

1.13:
ifndef GITHUB_TOKEN
	$(error GITHUB_TOKEN is not set)
endif
	@echo "[+] fetching the latest commit information..."
	$(eval LATEST:=$(shell curl -s https://api.github.com/repos/kubernetes/kubernetes/git/refs/heads/master | jq .object.sha))
	@echo "[+] generating notes"
	./release-notes \
		-start-sha="ad58349f7455a9ae8bb633e93ba0902a5cd6bc65" \
		-end-sha=${LATEST}
