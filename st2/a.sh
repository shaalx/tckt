#!/bin/sh

rm -rf target.md
go run main.go
md -r -f target.md
