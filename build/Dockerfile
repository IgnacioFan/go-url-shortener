FROM golang:1.19.0 as builder
WORKDIR /app
COPY . .
RUN go build -o app .
