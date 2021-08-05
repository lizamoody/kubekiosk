# syntax = docker/dockerfile:1-experimental
FROM golang:1.16.3-alpine AS build
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.mod .
COPY go.sum .
COPY src/ .
RUN go mod download
RUN go build -ldflags="-w -s" -a -installsuffix cgo -o /example .
FROM scratch
COPY --from=build /src .
ENTRYPOINT [ "/src"]