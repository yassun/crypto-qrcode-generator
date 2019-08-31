#!/usr/bin/env bash

set -eu
set -o pipefail

cd $SERVER_SRC_DIR

go get -u golang.org/x/lint/golint

go mod download

cd $TRAVIS_BUILD_DIR
