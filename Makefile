GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=lucid
EXAMPLE_DIR=example

all: deps build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
clean:
	$(GOCLEAN)
	rm -rf BINARY_NAME
deps:
	$(GOGET) -u gopkg.in/yaml.v2
	$(GOGET) -u github.com/spf13/cobra/cobra
	$(GOGET) -u nanomsg.org/go-mangos
	$(GOGET) -u github.com/golang/protobuf/protoc-gen-go
	$(GOGET) -u github.com/golang/protobuf/proto


