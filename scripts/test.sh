#!/usr/bin/env sh

go generate ./...
go test -tags assert -v -timeout 30s ./...
