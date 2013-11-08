#!/bin/bash
go run pic.go | cut -d ":" -f2 | base64 --decode > test.png
