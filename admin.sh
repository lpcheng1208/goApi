#!/bin/bash

SERVER="goApiServer"
BASE_DIR=$PWD
INTERVAL=2

RUN_ENV='prod'
case "$2" in
  'test')
  RUN_ENV='test'
  ;;
  'devl')
  RUN_ENV='devl'
  ;;
  'prod')
  RUN_ENV='prod'
  ;;

esac


function start()
{
#    if [ "`pgrep $SERVER -U $UID`" != "" ];then
#        echo "$SERVER already running"
#        exit 1
#    fi
    nohup $BASE_DIR/$SERVER --env=$RUN_ENV >/dev/null 2>&1 &
    echo "sleeping..." &&  sleep $INTERVAL

    # check status
    if [ "`pgrep $SERVER -U $UID`" == "" ];then
        echo "$SERVER start failed"
        exit 1
    else
      echo "start success"
    fi
}

function status()
{
    if [ "`pgrep $SERVER -U $UID`" != "" ];then
        echo $SERVER is running
    else
        echo $SERVER is not running
    fi
}

function stop()
{
    if [ "`pgrep $SERVER -U $UID`" != "" ];then
        kill `pgrep $SERVER -U $UID`
    fi

#    echo "sleeping..." &&  sleep $INTERVAL
#
#    if [ "`pgrep $SERVER -U $UID`" != "" ];then
#        echo "$SERVER stop failed"
#        exit 1
#    else
#        echo "stop success"
#    fi
}

function version()
{
  echo "$SERVER v1.0.1"
}


# 根据本地环境编译
function build()
{
    echo "$SERVER building..."
    buildResult=`go build -o "$SERVER" 2>&1`
    if [ -z "$buildResult" ]; then
        echo "$SERVER build done..."
    else
      echo "stop build"
      exit 1
    fi

}


# 根据Linux环境编译
function buildLinux()
{
    echo "$SERVER building..."
    buildResult=`GOOS=linux GOARCH=amd64 go build -o "$SERVER" 2>&1`
    if [ -z "$buildResult" ]; then
        echo "$SERVER build done..."
    else
      echo "stop build"
      exit 1
    fi

}


case "$1" in
    'start')
    start
    ;;
    'stop')
    stop
    ;;
    'status')
    status
    ;;
    'version')
    version
    ;;
    'build')
    build
    ;;
    'buildLinux')
    buildLinux
    ;;
    'restart')
    stop && build && start
    ;;
    'restartLinux')
    stop && buildLinux && start
    ;;
    *)
    echo "usage: $0 {start prod|stop|restart prod|restartLinux prod|status|version|build|buildLinux}"
    exit 1
    ;;
esac