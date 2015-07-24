#!/bin/bash

BASE="$(dirname $(cd "$(dirname "$BASH_SOURCE[0]}")" && pwd))"

cd $BASE/src
godep go build -o $BASE/bin/serve
