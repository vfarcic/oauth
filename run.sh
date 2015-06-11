#!/usr/bin/env bash

mongod &
sed -i -e "s/localhost:8080/$DOMAIN:$PORT/g" /app/components/oauth/oauth-who-am-i.html
sed -i -e "s/localhost:8080/$DOMAIN:$PORT/g" /app/components/oauth/oauth-providers.html
$PWD/oauth "$@"