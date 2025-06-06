#!/bin/zsh
set -e
mkdir -p target
cd database_populator
GO111MODULE=on go build -o ../target/database_populator main.go
cd ..
