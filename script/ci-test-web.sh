#!/usr/bin/env bash

set -eu
set -o pipefail

cd $WEB_SRC_DIR

yarn lint

cd $TRAVIS_BUILD_DIR
