#!/usr/bin/env bash
> assemble.hex
GOOS=linux GOARCH=amd64 go tool compile -S -N -L main.go > assemble.hex