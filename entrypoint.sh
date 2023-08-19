#!/bin/sh

#go env -w GOPROXY=https://goproxy.cn,direct \
#    && go mod download \
#    && go build main.go \
#    && ls -al

export GOPROXY=https://goproxy.cn

# go get github.com/pilu/fresh

go mod download

ls -al && go run main.go

# ls -al && fresh -c fresh.conf
