ARG TARGET_DIR=/app
ARG GOBIN=/.bin

FROM golang:1.20.1-alpine

ARG TARGET_DIR
ARG GOBIN

RUN apk add --update make git musl-dev gcc libc-dev binutils-gold

ENV GO111MODULE=on
ENV GOPRIVATE=gitlab.heronodes.io
ENV GOPATH=/go
WORKDIR /go/src/github.com/magmel48/social-network

COPY go.mod go.mod
COPY go.sum go.sum
COPY Makefile Makefile
ADD cmd cmd
ADD internal internal
RUN go mod download
RUN GOBIN=$GOBIN make install-tools
RUN TARGET_DIR=$TARGET_DIR make build

ENTRYPOINT make watch
