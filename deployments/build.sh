#!/usr/bin/bash

export PATH=$PATH:~/tmp/go/bin

go build -o simgular.out ./cmd/simgular/main.go
go build -o integrate.out ./cmd/integrate/main.go
