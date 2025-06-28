
BIN_DIR = $(shell pwd)/bin

.PHONY: all
all: lint test

.PHONY: clean
clean: tools-clean
	rm -rf dependencies

dependencies: go.mod go.sum
	go mod download
	touch dependencies

.PHONY: lint
lint: bin/golangci-lint
	bin/golangci-lint config verify
	bin/golangci-lint run

.PHONY: test
test: dependencies
	go test -timeout 5m ./...

.PHONY: tools
tools: bin/buf bin/golangci-lint bin/protoc-gen-buf-breaking bin/protoc-gen-buf-lint bin/protoc-gen-go

.PHONY: tools-clean
tools-clean:
	rm -rf bin
	rm tools/dependencies

tools/dependencies: tools/go.mod tools/go.sum tools/tools.go
	cd tools && go mod download
	touch tools/dependencies

bin/buf: tools/dependencies
	cd tools && go build -o ../bin/buf github.com/bufbuild/buf/cmd/buf

bin/golangci-lint: tools/dependencies
	cd tools && go build -o ../bin/golangci-lint github.com/golangci/golangci-lint/v2/cmd/golangci-lint

bin/protoc-gen-buf-breaking: tools/dependencies
	cd tools && go build -o ../bin/protoc-gen-buf-breaking github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking

bin/protoc-gen-buf-lint: tools/dependencies
	cd tools && go build -o ../bin/protoc-gen-buf-lint github.com/bufbuild/buf/cmd/protoc-gen-buf-lint

bin/protoc-gen-go: tools/dependencies
	cd tools && go build -o ../bin/protoc-gen-go google.golang.org/protobuf/cmd/protoc-gen-go

.PHONY: proto
proto: bin/buf bin/protoc-gen-buf-breaking bin/protoc-gen-buf-lint bin/protoc-gen-go
	PATH="$(BIN_DIR):$$PATH" buf lint
	rm -rf internal/testtypes
	PATH="$(BIN_DIR):$$PATH" buf generate
