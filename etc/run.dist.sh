#!/usr/bin/env bash

# resolve real path to script including symlinks or other hijinks
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do # resolve $SOURCE until the file is no longer a symlink
  TARGET="$(readlink "$SOURCE")"
  if [[ ${TARGET} == /* ]]; then
    echo "SOURCE '$SOURCE' is an absolute symlink to '$TARGET'"
    SOURCE="$TARGET"
  else
    BIN="$( dirname "$SOURCE" )"
    echo "SOURCE '$SOURCE' is a relative symlink to '$TARGET' (relative to '$BIN')"
    SOURCE="$BIN/$TARGET" # if $SOURCE was a relative symlink, we need to resolve it relative to the path where the symlink file was located
  fi
done
USER="$( whoami )"
GROUP="$( users )"
RBIN="$( dirname "$SOURCE" )"
BIN="$( cd -P "$( dirname "$SOURCE" )" && pwd )"
DIR="$( cd -P "$BIN/../" && pwd )"
ETC="$( cd -P "$DIR/etc" && pwd )"
PKG="$( cd -P "$DIR/pkg" && pwd )"
SRC="$( cd -P "$DIR/src" && pwd )"
WEB="$( cd -P "$DIR/web" && pwd )"
JS_SRC="$( cd -P "$WEB/todo" && pwd )"
NODEJS_VERSION="v14.0.0"
GO_VERSION="1.14.2"
export NVM_DIR="$HOME/.nvm"

#build the application
GOPATH=$DIR GOBIN=$BIN $BIN/build.sh

#run the binary
BIN=${BIN} DIR=${DIR} ETC=${ETC} PKG=${PKG} SRC=${SRC} WEB=${WEB} \
ENV=${ENV} DEBUG=${DEBUG} \
DB_NAME=__DB_NAME__ \
DB_USER=__DB_USER__ \
DB_PASSWORD=__DB_PASSWORD__ \
DB_HOST=__DB_HOST__ \
DB_PORT=__DB_PORT__ \
PORT=__PORT__ \
$BIN/todo