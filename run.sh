#!/usr/bin/env bash

mongod &
sed -i -e "s/localhost:8080/$DOMAIN:$PORT/g" /app/components/oauth/whoami.html
$PWD/oauth "$@"