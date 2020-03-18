 # Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOFMT=$(GOCMD) fmt 
GOGET=$(GOCMD) get
BINARY_NAME=goborg
BINARY_UNIX=$(BINARY_NAME)_unix

all: fmt test build

fmt:
		$(GOFMT) ./...

test: deps
		$(GOTEST) -v ./...

build: deps
		$(GOBUILD) -o $(BINARY_NAME) -v ./cmd

clean: 
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
		rm -f $(BINARY_UNIX)

run: build
		./$(BINARY_NAME)

deps:
		$(GOGET) github.com/stretchr/testify/assert
		$(GOGET) github.com/docopt/docopt.go
		$(GOGET) github.com/davecgh/go-spew/spew
		$(GOGET) go.uber.org/zap
		$(GOGET) github.com/google/uuid
		$(GOGET) github.com/gozilla/mux

# Cross compilation
build-linux:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v ./cmd

docker-build:
		docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v

