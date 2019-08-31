#!/usr/bin/env bash

set -eu
set -o pipefail

cd $SERVER_SRC_DIR

go vet ./...

diff <(golint ./...) <(printf "")

go test ./...

cd $TRAVIS_BUILD_DIR
