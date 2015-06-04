#!/usr/bin/env bash

mongod &
$PWD/oauth "$@"