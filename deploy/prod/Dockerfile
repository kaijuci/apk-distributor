FROM golang:1.22-alpine AS build
LABEL org.opencontainers.image.authors="sumeet@kaiju.ci"

ARG SHA="unknown"
ARG RELEASE_VERSION="0.0"

ENV BUILD_SHA=$SHA
ENV BUILD_VERSION=$RELEASE_VERSION

WORKDIR /app
ADD . .

# build
RUN CGO_CFLAGS_ALLOW='-Xpreprocessor' \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=on \
    go build \
    -ldflags "-X github.com/kaijuci/apk-publisher/vars.SHA=$BUILD_SHA" \
    -ldflags "-X github.com/kaijuci/apk-publisher/vars.Version=$BUILD_VERSION" \
    -o apkpublish \
    cmd/apk-distributor/main.go

FROM scratch

WORKDIR /app
COPY --from=build /app/apkpublish .

ENTRYPOINT ["/app/apkpublish"]