FROM golang:1.12-alpine3.9 AS build

WORKDIR /go/src/go-api-project-template

RUN GO111MODULE=on go mod vendor \
    && go build -o build/api-server ./cmd/api-server

COPY . .

FROM alpine:3.8 AS app

COPY --from=build /go/src/go-api-project-template/build/api-server /usr/local/bin/api-server

EXPOSE 8080

CMD ["api-server"]
