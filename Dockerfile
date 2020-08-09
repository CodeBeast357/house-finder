#syntax=docker/dockerfile:experimental

FROM golang:alpine3.12 as builder

RUN mkdir /house-finder-src
WORKDIR /house-finder-src

COPY go.mod go.sum /house-finder-src/
RUN --mount=type=cache,target=/root/.cache/go-build,id=go-build-cache GOPROXY=https://proxy.golang.org go mod download

COPY src /house-finder-src/src/
RUN --mount=type=cache,target=/root/.cache/go-build,id=go-build-cache \
  GOPROXY=https://proxy.golang.org go build -o /main ./src/main.go

FROM alpine:3.12

COPY --from=builder /main /house-finder/
COPY src/house-finder/data /house-finder/data

ENTRYPOINT [ "/house-finder/main" ]
