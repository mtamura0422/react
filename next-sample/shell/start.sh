#!/bin/sh

cd /usr/src/appllication
yarn install
npm run dev

while true; do sleep 1; done
