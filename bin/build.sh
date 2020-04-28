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
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"

cd $BIN
GOPATH=$DIR GOBIN=$BIN go build $BIN/todo.go
cd $DIR

cd $JS_SRC
nvm use $NODEJS_VERSION
ng build
cd $DIR