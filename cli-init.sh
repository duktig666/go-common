#!/bin/bash

echo "exec 'go mod tidy -compat=1.18'"
go mod tidy -compat=1.18

echo "exec 'go build -o couples-cli'"
go build -o couples-cli
