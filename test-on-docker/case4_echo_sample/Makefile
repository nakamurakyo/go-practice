BINARY = my-test
GOOS   = linux
GOARCH = amd64

# Goでビルドバージョン情報を参照できるように https://qiita.com/harukasan/items/37698ec799678c12e71d
LDFLAGS = -ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"

VERSION=0.0.1
COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
GIT_URL=local-git://
GIT_BRANCH=master

MYSQL_DATABASE=yatter
export MYSQL_DATABASE
MYSQL_USER=yatter
export MYSQL_USER
MYSQL_PASSWORD=yatter
export MYSQL_PASSWORD
MYSQL_HOST=localhost:3306
export MYSQL_HOST

all: clean fmt mod build-all

run:
	docker-compose up --build

build-default:
	go build ${LDFLAGS} -o build/${BINARY} ${CMD_DIR}

build-linux:
	GOOS=${GOOS} GOARCH  =${GOARCH} go build ${LDFLAGS} -o build/${BINARY}-linux-${GOARCH} ${CMD_DIR}

# コードの静的解析
vet:
	go vet ./...

# フォーマット
fmt:
	go fmt $$(go list ./... | grep -v /vendor/)

# vendorディレクトリに依存関係を保存 ※ onは常にmodule-aware modeで動作 
mod:
	 GO111MODULE=on go mod vendor

clean:
	rm -f ${TEST_REPORT}
	rm -f ${BINARY} ${BINARY}-*

docker-clean:
	docker-compose down -v

###=========================================================


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
