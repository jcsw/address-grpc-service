# Go parameters
GO_CMD=go
DEP_CMD=dep
GO_GET=$(GO_CMD) get
GO_RUN=$(GO_CMD) run
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_TEST=$(GO_CMD) test
GO_VET=$(GO_CMD) vet
GO_COVER=$(GO_CMD) tool cover
GO_FMT=gofmt
GO_LINT=golint

BUILD_DIRECTORY=build
BINARY_NAME=address-grpc-service
BINARY_UNIX=$(BINARY_NAME)_unix
APP_INIT=./cmd/main.go

SRC=$(shell find . -type f -name '*.go' -not -path "./vendor/*")
PKG= $(shell go list ./... | grep -v /vendor/)

all: clean fmt vet lint test build

build:
	$(GO_BUILD) -o $(BINARY_DIRECTORY)/$(BINARY_NAME) -v $(APP_INIT)
test:
	$(GO_TEST) -v ./... -covermode=count -coverprofile=$(BUILD_DIRECTORY)/coverage.out
	$(GO_COVER) -html=$(BUILD_DIRECTORY)/coverage.out -o $(BUILD_DIRECTORY)/coverage.html
itest:
	$(GO_TEST) -v -tags=integration ./... -covermode=count -coverprofile=$(BUILD_DIRECTORY)/icoverage.out
	$(GO_COVER) -html=$(BUILD_DIRECTORY)/icoverage.out -o $(BUILD_DIRECTORY)/icoverage.html
vet:
	$(GO_VET) -v ./...
lint:
	$(GO_LINT) $(PKG)
fmt:
	$(GO_FMT) -w -d $(SRC)
clean:
	$(GO_CLEAN)
	rm -f $(BUILD_DIRECTORY)/*
run:
	$(GO_RUN) $(APP_INIT) -env=dev

run-deps:
	docker-compose --file docker/mongodb/docker-compose.yml up -d

stop-deps:
	docker-compose --file docker/mongodb/docker-compose.yml stop

deps:
	$(DEP_CMD) ensure

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO_BUILD) -o $(BUILD_DIRECTORY)/$(BINARY_UNIX) -v $(APP_INIT)