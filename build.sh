#!/bin/zsh
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./built/muddler_linux .
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./built/muddler.exe .
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./built/muddler_darwin .