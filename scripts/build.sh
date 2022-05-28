#!/usr/bin/env bash

set -euxo pipefail


# build the web project 
. ~/.nvm/nvm.sh
nvm use 
npm i -g yarn 
cd www 
yarn
yarn build 
cd  ..
ln -sf ./www/build ./static

# build the go server 
go build -o ./goblin-api ./cmd/goblin-api 
pm2 stop goblin-api
pm2 start goblin-api -- --env .env
