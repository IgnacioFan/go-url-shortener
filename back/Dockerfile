# Build stage
FROM golang:1.19.0-alpine AS build

WORKDIR /api

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o url_shortener

# Release stage
FROM alpine:latest AS release

RUN apk --no-cache add ca-certificates bash

WORKDIR /api

COPY --from=build /api/url_shortener .

EXPOSE 8000

CMD ["./url_shortener", "server", "-p", "8000"]
