.PHONY: test

all: test build

test:
	go test -v --race ./...

build:
	go build -v ./cmd/nft_exporter.go

build_linux:
	GCO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o nft_exporter.linux64 ./cmd/nft_exporter.go