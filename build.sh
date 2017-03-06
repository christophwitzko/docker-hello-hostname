#!/bin/bash

set -e

docker run --rm -v "$PWD":/go/src/hello-hostname -w /go/src/hello-hostname -e CGO_ENABLED=0 golang:1-alpine go build -a --installsuffix cgo -ldflags="-s -w"
docker build -t christophwitzko/hello-hostname .
