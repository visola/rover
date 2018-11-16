#!/bin/bash

rm -Rf build
mkdir build

# Build all packages to run in this computer
go build -o build/uploader cmd/uploader/main.go

# Build all the packages that go in the RPi
export GOOS=linux
export GOARCH=arm
export GOARM=6

go build -o build/agent cmd/agent/main.go
