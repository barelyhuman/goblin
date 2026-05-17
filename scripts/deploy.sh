#!/usr/bin/env bash

docker build . --platform=linux/amd64 -t goblin:latest
docker save goblin:latest | gzip > goblin-latest.tar.gz

PREPARE_COMMANDS="""
set -euxo pipefail
mkdir -p ~/goblin
"""

ssh root@goblin.run "/bin/bash -c '$PREPARE_COMMANDS'"

rsync --progress goblin-latest.tar.gz root@goblin.run:~/goblin/

COMMANDS="""
set -euxo pipefail
cd ~/goblin
docker image load < goblin-latest.tar.gz
docker stop \$(docker container ls --all --filter=ancestor="goblin:latest" --format "{{.ID}}")
docker run -d -e 'ORIGIN_URL=http://goblin.run'  -p '3000:3000' -v='./:/usr/bin/app' goblin:latest
"""

ssh root@goblin.run "/bin/bash -c '$COMMANDS'"
