#!/bin/bash

ROOT_DIR=`pwd`
cd $ROOT_DIR
export GOPATH=$ROOT_DIR
export GOBIN=$GOPATH/bin
echo $ROOT_DIR
echo $GOPATH
echo $GOBIN

go get github.com/garyburd/redigo/redis
go install main

exec  ./bin/main
