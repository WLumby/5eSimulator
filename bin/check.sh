#!/usr/bin/env bash

set -e

cd ../
go vet
go test