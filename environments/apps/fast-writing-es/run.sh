#!/bin/sh

app="fast-writing-es"
tag=${1:-"latest"}
image_name="asia-northeast1-docker.pkg.dev/anan-project/fast-writing/${app}:${tag}"

[[ `docker ps -a -f name=${app} -q | wc -l` -eq 1 ]] && docker rm -f ${app}
docker run -d -p 9200:9200 \
    -p 9300:9300 \
    -e "discovery.type=single-node" \
    -e "xpack.security.enabled=false" \
    --name ${app} \
    ${image_name}