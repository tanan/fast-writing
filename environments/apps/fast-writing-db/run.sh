#!/bin/sh

app="fast-writing-db"
tag=${1:-"latest"}
image_name="asia-northeast1-docker.pkg.dev/anan-project/fast-writing/${app}:${tag}"

[[ `docker ps -a -f name=${app} -q | wc -l` -eq 1 ]] && docker rm -f ${app}
docker run -d -p 3306:3306 \
    -e MYSQL_ROOT_PASSWORD=password \
    --name ${app} \
    ${image_name}
