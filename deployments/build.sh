#!/usr/bin/bash

export PATH=$PATH:~/tmp/go/bin

go build -o singular.out ./cmd/singular/main.go
go build -o system.out ./cmd/system/main.go
go build -o diff_system.out ./cmd/diff_system/main.go
