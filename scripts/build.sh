#!/usr/bin/env bash

set -euxo pipefail

# build the web project 
cd www 
# if using darwin arm64, uncomment the next line
# make install
make installLinux
make build
cd ..

ln -sf ./www/dist ./static

# build the go server 
go build -o ./goblin-api ./cmd/goblin-api 
pm2 stop goblin-api
pm2 start goblin-api -- --env .env
