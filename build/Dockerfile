# Build stage
FROM golang:1.19.0-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app

# Release stage
FROM alpine:latest AS release

RUN apk --no-cache add ca-certificates bash

WORKDIR /app

COPY --from=build /app/app .
COPY /deployment/config/ /app/deployment/config/
COPY /deployment/nginx /app/deployment/nginx/

EXPOSE 3000
