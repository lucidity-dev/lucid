GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=lucid
EXAMPLE_DIR=example

all: build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -rf BINARY_NAME
deps:
	$(GOGET) gopkg.in/yaml.v2
	$(GOGET) github.com/spf13/cobra/cobra


