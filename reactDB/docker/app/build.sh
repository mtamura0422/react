#!/bin/sh

cd src && npm run build
cd ..

IS_PUSH=${1:-0}
REGISTRY=gcr.io/app-labo-02/exsrc
VERSION=latest

docker build -t ${REGISTRY}:${VERSION} .

if [ ${IS_PUSH} -eq 1 ]; then
  gcloud docker -- push ${REGISTRY}:${VERSION}
fi

