#!/bin/sh
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app
docker build -t smileleo/lease-example:latest .
docker push smileleo/lease-example:latest