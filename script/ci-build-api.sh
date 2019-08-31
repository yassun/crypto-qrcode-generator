#!/usr/bin/env bash

set -eu
set -o pipefail

if [ $TRAVIS_PULL_REQUEST != "false" ]; then
  echo "Skip build because this build is a pull request."
  exit 0
fi

if [ $TRAVIS_BRANCH != "master" ]; then
  echo "Skip build because this build is not master."
  exit 0
fi

cd $SERVER_SRC_DIR
docker build -t $GCP_REGION/$GCP_PROJECT_ID/crypto-qrcode-generator-server:latest -f Dockerfile.prod .
cd $TRAVIS_BUILD_DIR
