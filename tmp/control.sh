#!/bin/bash

set -ex

# 以托管的方式启动服务
workspace=$(cd $(dirname $0) && pwd -p)
cd $workspace
module=sharp
app=$module

function set_env() {
    export LD_LIBRARY_PATH=$workspace/icu_lib
    cluster_name=$(cat .environment/service.env.txt)
    if [[ $cluster_name =~ ^.*sim.*$ ]]; then
        app=$module-cover
    fi
}

function run_command() {
    action=$1
    case $action in
      "start")
      set_env
    # 启动服务，以前台方式启动，否则无法托管
    exec &> >(while read line || [ -n "$line" ]; do echo "[$(date "+%Y-%m-%d %H:%M:%S")] $line";done) ./bin/$app
    ;;
    *)
    # 非法命令，以非0码退出
    echo "unknown command"
    exit 1
    ;;
  esac
}