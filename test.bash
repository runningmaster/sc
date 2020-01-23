#!/usr/bin/env bash

#go list ./... | grep -v vendor/ | xargs -L1 go generate
go list ./... | grep -v vendor/ | xargs -L1 go test -race -bench . -benchmem -cover
