#!/bin/zsh
set -e
mkdir -p target
cd api
GO111MODULE=on go build -o ../target/nrc_api main.go
cd ..
