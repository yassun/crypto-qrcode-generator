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

# install SDK
export CLOUDSDK_CORE_DISABLE_PROMPTS=1
if [ ! -d "$HOME/google-cloud-sdk/bin" ]; then rm -rf "$HOME/google-cloud-sdk"; curl https://sdk.cloud.google.com | bash > /dev/null; fi
source $HOME/google-cloud-sdk/path.bash.inc

# set up env
gcloud config set project $GCP_PROJECT_ID
gcloud config set compute/zone $GCP_ZONE

# decode secret
openssl aes-256-cbc -K $encrypted_3c9334502a76_key -iv $encrypted_3c9334502a76_iv -in $GCP_SDK_SECRET.enc -out $GCP_SDK_SECRET -d

# auth GCP
gcloud auth activate-service-account --key-file $GCP_SDK_SECRET

# install kubectl
gcloud components install kubectl

# for Docker Registry
gcloud --quiet auth configure-docker

# for kube
gcloud container clusters get-credentials crypto-qrcode-generator
