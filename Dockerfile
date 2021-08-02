# syntax = docker/dockerfile:1-experimental
FROM golang:1.14.3-alpine AS build
WORKDIR /src/
ENV CGO_ENABLED=0
COPY go.* .
RUN go mod download
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build \
go build -o /bin/demo .
FROM scratch
COPY --from=build /bin/demo /bin/demo
ENTRYPOINT ["/bin/demo"]