#!/bin/zsh
set -e
mkdir -p target
cd excel_extractor
GO111MODULE=on go build -o ../target/excel_data_extractor main.go
cd ..
