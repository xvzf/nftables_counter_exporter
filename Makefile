.PHONY: test

all: test build

test:
	go test -v --race ./...

build: build_amd64 build_arm build_arm64

build_amd64:
	GCO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o nft_exporter.amd64 ./cmd/nft_exporter.go

build_arm:
	GCO_ENABLED=0 GOOS=linux GOARCH=arm go build -v -o nft_exporter.arm ./cmd/nft_exporter.go

build_arm64:
	GCO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -v -o nft_exporter.arm64 ./cmd/nft_exporter.go