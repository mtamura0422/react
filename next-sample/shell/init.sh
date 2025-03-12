#!/bin/sh

cd /usr/src/
npm install -g typescript ts-node ts-node-dev
npx --yes create-next-app@latest --ts app
npm run dev

#while true; do sleep 1; done
