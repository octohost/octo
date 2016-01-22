OCTO_VERSION="0.2-dev"
GIT_COMMIT=$(shell git rev-parse HEAD)
COMPILE_DATE=$(shell date -u +%Y%m%d.%H%M%S)
BUILD_FLAGS=-X main.CompileDate=$(COMPILE_DATE) -X main.GitCommit=$(GIT_COMMIT) -X main.Version=$(OCTO_VERSION)

all: build

deps:
	go get -u github.com/spf13/cobra
	go get -u github.com/progrium/basht
	go get -u github.com/CiscoCloud/consul-cli
	go get -u github.com/spf13/viper
	go get -u github.com/hashicorp/consul/api
	go get -u github.com/samalba/dockerclient

format:
	gofmt -w .

clean:
	rm -f bin/octo || true

build: clean
	go build -ldflags "$(BUILD_FLAGS)" -o bin/octo main.go

gziposx:
	gzip bin/octo
	mv bin/octo.gz bin/octo-$(OCTO_VERSION)-darwin.gz

linux: clean
	GOOS=linux GOARCH=amd64 go build -ldflags "$(BUILD_FLAGS)" -o bin/octo main.go

gziplinux:
	gzip bin/octo
	mv bin/octo.gz bin/octo-$(OCTO_VERSION)-linux-amd64.gz

release: clean build gziposx clean linux gziplinux clean

consul:
	consul agent -data-dir `mktemp -d` -bootstrap -server -bind=127.0.0.1 1>/dev/null &

consul_kill:
	pkill consul

unit:
	cd cmd && go test -v -cover

test: consul unit wercker

wercker: consul
	basht test/tests.bash
