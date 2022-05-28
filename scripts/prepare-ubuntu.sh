#!/usr/bin/env bash

set -euxo pipefail

# install caddy 
sudo apt install -y debian-keyring debian-archive-keyring apt-transport-https
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/gpg.key' | sudo gpg --dearmor -o /usr/share/keyrings/caddy-stable-archive-keyring.gpg
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/debian.deb.txt' | sudo tee /etc/apt/sources.list.d/caddy-stable.list
sudo apt update -y 
sudo apt install caddy -y 


# setup node and yarn 
wget -qO- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh | bash
. ~/.nvm/nvm.sh
nvm install
nvm use
npm i -g yarn 
npm i -g pm2

# install golang and snap
sudo apt install snapd
sudo snap install go --channel=1.18/stable --classic


