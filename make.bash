#!/usr/bin/env bash

#go list ./... | grep -v vendor/ | xargs -L1 go generate
go list ./... | grep -v vendor/ | xargs -L1 go fmt
#go get -u ./...
go mod tidy
export GOBIN=$(pwd)/bin
go install -ldflags "-s -w" github.com/runningmaster/sc/cmd/sc
