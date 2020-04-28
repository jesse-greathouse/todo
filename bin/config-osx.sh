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
RUN_SCRIPT="${BIN}/run.sh"
SERVICE_RUN_SCRIPT="${BIN}/run-ubuntu-service.sh"
RUN_SCRIPT_TPL="${ETC}/run.dist.sh"

printf "\n"
printf "\n"
printf "=================================================================\n"
printf "Hello, "${USER}".  This will create your site's run script\n"
printf "=================================================================\n"
printf "\n"
printf "Enter your website port [80]: "
read PORT
if  [ "${PORT}" == "" ]; then
    PORT="80"
fi
printf "Enter your database host [127.0.0.1]: "
read DB_HOST
if  [ "${DB_HOST}" == "" ]; then
    DB_HOST="127.0.0.1"
fi
printf "Enter your database name [todo]: "
read DB_NAME
if  [ "${DB_NAME}" == "" ]; then
    DB_NAME="todo"
fi
printf "Enter your database user [todo]: "
read DB_USER
if  [ "${DB_USER}" == "" ]; then
    DB_USER="todo"
fi
printf "Enter your database password [todo]: "
read DB_PASSWORD
if  [ "${DB_PASSWORD}" == "" ]; then
    DB_PASSWORD="todo"
fi
printf "Enter your database port [3306]: "
read DB_PORT
if  [ "${DB_PORT}" == "" ]; then
    DB_PORT="3306"
fi

printf "\n"
printf "You have entered the following configuration: \n"
printf "\n"
printf "Web Port: ${PORT} \n"
printf "Database Host: ${DB_HOST} \n"
printf "Database Name: ${DB_NAME} \n"
printf "Database User: ${DB_USER} \n"
printf "Database Password: ${DB_PASSWORD} \n"
printf "Database Port: ${DB_PORT} \n"
printf "\n"
printf "Is this correct (y or n): "
read -n 1 CORRECT
printf "\n"

if  [ "${CORRECT}" == "y" ]; then

    if [ -f ${RUN_SCRIPT} ]; then
       rm ${RUN_SCRIPT}
    fi
    cp ${RUN_SCRIPT_TPL} ${RUN_SCRIPT}

    sed -i '' '1 a\
            PATH=/usr/local/bin:$PATH' ${RUN_SCRIPT}
    sed -i '' s/__PORT__/"${PORT}"/g ${RUN_SCRIPT}
    sed -i '' s/__DB_HOST__/"${DB_HOST}"/g ${RUN_SCRIPT}
    sed -i '' s/__DB_NAME__/"${DB_NAME}"/g ${RUN_SCRIPT}
    sed -i '' s/__DB_USER__/"${DB_USER}"/g ${RUN_SCRIPT}
    sed -i '' s/__DB_PASSWORD__/"${DB_PASSWORD}"/g ${RUN_SCRIPT}
    sed -i '' s/__DB_PORT__/"${DB_PORT}"/g ${RUN_SCRIPT}
    chmod +x ${RUN_SCRIPT}

printf "\n"
printf "\n"
printf "\n"
printf "================================================================\n"

    printf "Your run script has been created at: \n"
    printf "${RUN_SCRIPT}\n"
    printf "\n"
else
    printf "Please run this script again to enter the correct configuration. \n"
    printf "\n"
    printf "================================================================\n"
    exit 1
fi