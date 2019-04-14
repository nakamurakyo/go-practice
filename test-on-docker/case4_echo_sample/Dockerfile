FROM golang:1.12-alpine3.9 AS build

WORKDIR /go/src/git.dmm.com/nakamurakyo/go-practice

RUN apk add --no-cache ca-certificates git make \
	&& go get -u github.com/oxequa/realize

COPY . .

RUN make mod \
	&& make build-default

CMD ["sh", "-c", "make mod && cd cmd/prctice-backend-server && realize start"]

FROM alpine:3.8 AS app

COPY --from=build /go/src/git.dmm.com/nakamurakyo/go-practice/build/practice-backend /usr/local/bin/practice-backend

RUN apk add --no-cache ca-certificates

EXPOSE 8080

CMD ["practice-backend"]