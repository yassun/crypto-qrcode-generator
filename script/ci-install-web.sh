#!/usr/bin/env bash

set -eu
set -o pipefail

# Install Yarn (https://yarnpkg.com/en/docs/install-ci#travis-tab)
curl -o- -L https://yarnpkg.com/install.sh | bash
export PATH="$HOME/.yarn/bin:$PATH"

# Install package json
cd $WEB_SRC_DIR
yarn install
cd $TRAVIS_BUILD_DIR
