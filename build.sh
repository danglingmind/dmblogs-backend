#!/bin/bash

GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -a -installsuffix cgo -o app .

docker-compose up -d

