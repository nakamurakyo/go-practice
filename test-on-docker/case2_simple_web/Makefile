GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOFMT=$(GOCMD) fmt
GOGET=$(GOCMD) get
BINARY_NAME=mybinary

# phoneyが無いと同名のファイルがある時に実行しない
.PHONY: build run clean deps

build:
	$(GOFMT)
	$(GOBUILD) -o $(BINARY_NAME) -v

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

dep:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

##### dep init
# dep-init:
# 	dep init