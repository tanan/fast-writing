#!/bin/sh

app="fast-writing-db"
tag=${1:-"latest"}
image_name="gcr.io/anan-project/fast-writing/${app}:${tag}"

[[ `docker ps -a -f name=${app} -q | wc -l` -eq 1 ]] && docker rm -f ${app}
docker run -d -p 3306:3306 \
    -e MYSQL_ROOT_PASSWORD=contoro \
    --name ${app} \
    ${image_name}
