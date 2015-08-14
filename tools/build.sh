#!/bin/bash

BASE="$(dirname $(cd "$(dirname "$BASH_SOURCE[0]}")" && pwd))"

cd $BASE/src
go build -o $BASE/bin/serve
