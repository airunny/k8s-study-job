#!/bin/sh
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app
docker build -t smileleo/mutating-webhook:latest .
docker push smileleo/mutating-webhook:latest