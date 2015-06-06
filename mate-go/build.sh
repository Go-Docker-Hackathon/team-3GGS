#!/bin/bash

<<<<<<< HEAD
ROOT_DIR=`pwd`
=======
ROOT_DIR="/gopath/app/team-3GGS/mate-go"
>>>>>>> ea134be5271c1f4f3e93a0a9c13f91ee2f4c6ccb
cd $ROOT_DIR
export GOPATH=$ROOT_DIR
export GOBIN=$GOPATH/bin
echo $ROOT_DIR
echo $GOPATH
echo $GOBIN

go get github.com/garyburd/redigo/redis

go install main

exec ./bin/main
