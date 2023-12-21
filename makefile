# Go parameters
GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOINSTALL=$(GOCMD) install
BINARY_NAME=bin/wc-go

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	$(GOTEST) -v ./... --cover

install:
	$(GOINSTALL)

run:
	$(GORUN) main.go -c sample.txt

all: test run