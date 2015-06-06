#!/bin/bash

ROOT_DIR=`pwd`
export GOPATH=$ROOT_DIR
export GOBIN=$GOPATH/bin
echo $GOPATH
echo $GOBIN

go get github.com/garyburd/redigo/redis
go install main

exec "$@"
