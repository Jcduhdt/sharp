#!/bin/bash

set -ex

# 项目名称
MODULE_NAME=sharp
VERSION=1.16.5
WORKSPACE=$(cd "$(dirname $0)";pwd -P)
DEPLOY_DIR="output"
DATE=$(date "+%s")
GOPATH=/tmp/go-build${DATE}
GOPKG="${GOPATH}/pkg"

#env
export GOPATH
export GOROOT=/usr/local/go$VERSION
export PATH=${GOROOT}/bin:$GOPATH/bin:${PATH}:$GOBIN
export GOPROXY=https://goproxy.io,direct
export GOSUMDB=off

GOPKG="${GOPATH}/pkg"

if [ ! d $GOROOT ];then
  echo "ERROR!!! GO VERSION should more than 1.16.5"
  exit 1
fi

function build() {
    echo "building..." && make
    if [[ $? != 0 ]]; then
        echo -e "build failed"
        exit 1
    fi
    echo -e "build success"
}

build