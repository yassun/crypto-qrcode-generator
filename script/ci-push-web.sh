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

gcloud docker -- push $GCP_REGION/$GCP_PROJECT_ID/crypto-qrcode-generator-web:latest
