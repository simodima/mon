#!/bin/bash

ABSOLUTE_PROJECT_PATH=$(git rev-parse --show-toplevel)
export USER_ID=$(id -u -r)

docker-compose rm -sfv 
docker-compose up -d 