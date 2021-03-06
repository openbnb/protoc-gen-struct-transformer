VERSION=1.0.6-dev
BUILDTIME=$(shell date +"%Y-%m-%dT%T%z")
LDFLAGS= -ldflags '-X github.com/bold-commerce/protoc-gen-struct-transformer/generator.version=$(VERSION) -X github.com/bold-commerce/protoc-gen-struct-transformer/generator.buildTime=$(BUILDTIME)'

.PHONY: re-generate-example generate install build version setup

re-generate-example:
	protoc \
		--proto_path=$(GOPATH)/pkg/mod/github.com/gogo:. \
		--struct-transformer_out=package=transform,debug=false,helper-package=helpers,goimports=true:. \
		--gogofaster_out=Moptions/annotations.proto=github.com/bold-commerce/protoc-gen-struct-transformer/options:. \
		./example/message.proto

generate: version re-generate-example

install: setup
	go install $(LDFLAGS)

build: OUTPUT=.
build: setup
	go build $(LDFLAGS) -o $(OUTPUT)

version:
	protoc-gen-struct-transformer --version

setup:
	go mod download
	go mod verify
