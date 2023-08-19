#!/bin/sh

#go env -w GOPROXY=https://goproxy.cn,direct \
#    && go mod download \
#    && go build main.go \
#    && ls -al

export GOPROXY=https://goproxy.cn

go env -w GOFLAGS="-buildvcs=false"

go install github.com/cosmtrek/air@latest

go mod download

# ls -al && go run main.go

ls -al && air -c .air.toml
