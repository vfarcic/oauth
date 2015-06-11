#!/usr/bin/env bash

mongod &
sed -i -e "s/localhost:8080/$DOMAIN:$PORT/g" /app/components/oauth/who-am-i.html
$PWD/oauth "$@"