FROM golang:1.12-alpine as build

ENV GOPATH /go

ENV WORKDIR ${GOPATH}/src/github.com/aaa/bbb

RUN mkdir -p /go/src/github.com/aaa/bbb

ADD . ${WORKDIR}
WORKDIR ${WORKDIR}

RUN apk add --no-cache git \
  && go build ${WORKDIR}/cmd/app \
  && go get github.com/oxequa/realize

FROM alpine

WORKDIR /app

COPY --from=build /go/app/app .

RUN addgroup go \
  && adduser -D -G go go \
  && chown -R go:go /app/app

CMD ["./app"]