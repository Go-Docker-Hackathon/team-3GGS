#!/bin/bash

ROOT_DIR=$GOPATH
cd $ROOT_DIR
export GOPATH=$ROOT_DIR
export GOBIN=$GOPATH/bin
echo $ROOT_DIR
echo $GOPATH
echo $GOBIN

go get github.com/garyburd/redigo/redis

go install main

exec  /gopath/app/team-3GGs/mate-go/bin/main
