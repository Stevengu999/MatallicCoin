#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
echo "metalicoin binary dir:" "$DIR"
pushd "$DIR" >/dev/null
go run cmd/metalicoin/metalicoin.go --gui-dir="${DIR}/src/gui/static/" $@
popd >/dev/null
