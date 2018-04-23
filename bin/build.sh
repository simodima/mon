#!/bin/bash

docker-compose exec -T app env GOARCH=amd64 GOOS=linux go build -o /dist/mon main.go
