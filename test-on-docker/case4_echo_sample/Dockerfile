FROM golang:1.12-alpine as build

ENV GOPATH /go
ENV GO111MODULE on
ENV WORKDIR $GOPATH/src/github.com/nakamurakyo/go-practice/test-on-docker/case4_echo_sample

WORKDIR $WORKDIR
ADD . $WORKDIR
RUN apk add --no-cache git \
  && go mod vendor \
  && go build -o app cmd/app/main.go \
  && go get github.com/oxequa/realize


FROM alpine

WORKDIR /app

COPY --from=build /go/app/app .

RUN addgroup go \
  && adduser -D -G go go \
  && chown -R go:go /app/app

CMD ["./app"]