#!/bin/sh

#go env -w GOPROXY=https://goproxy.cn,direct \
#    && go mod download \
#    && go build main.go \
#    && ls -al

export GOPROXY=https://goproxy.cn

go mod vendor

go ls -al && go run main.go
