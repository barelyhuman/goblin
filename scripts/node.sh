#!/bin/bash
set -e

wget -q https://nodejs.org/dist/latest-v14.x/node-v14.21.1-linux-x64.tar.xz
tar xf node-v14.21.1-linux-x64.tar.xz
rm -f *.tar *.xz

export PATH=$PATH:"$(pwd)/node-v14.21.1-linux-x64/bin"

cd www 

npm install -g pnpm
pnpm install
pnpm build

cd ..

cp -r ./www/build ./static

go build -o ./server cmd/goblin-api/main.go

rm -rf node-v12.22.12-linux-x64

