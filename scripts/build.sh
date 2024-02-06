#!/usr/bin/env bash

set -euxo pipefail


# if using darwin arm64, uncomment the next line
# make install
npm i
npm run build

# build the go server 
go build -o ./goblin-api ./cmd/goblin-api 
pm2 stop goblin-api
pm2 start goblin-api -- --env .env
