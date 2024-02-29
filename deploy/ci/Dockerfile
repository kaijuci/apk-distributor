FROM golang:1.22 AS base
LABEL org.opencontainers.image.authors="sumeet@kaiju.ci"

# install deps
RUN apt-get update \
    && apt-get install -y \
    tree

WORKDIR /app
COPY go.mod go.sum ./

# fetch dependencies
RUN go mod download -x

FROM base AS go-build

ARG RELEASE_VERSION="0.0"

COPY . .

RUN tree .

# build
RUN CGO_CFLAGS_ALLOW='-Xpreprocessor' \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=on \
    go build \
    -ldflags "-X github.com/kaijuci/apk-publisher/vars.SHA=$(git rev-parse --short HEAD)" \
    -ldflags "-X github.com/kaijuci/apk-publisher/vars.Version=$RELEASE_VERSION" \
    -o apkpublish \
    cmd/apk-distributor/main.go

# test
RUN CGO_CFLAGS_ALLOW='-Xpreprocessor' \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=on \
    go test -v ./...

# echo version - mostly to satisfy kaniko
ENTRYPOINT ["/app/apkpublish"]