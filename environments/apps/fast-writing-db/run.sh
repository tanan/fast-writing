#!/bin/sh

app="contoro-db"
tag=${1:-"latest"}
docker_network=${2:-"contoro"}
image_name="gcr.io/anan-project/contoro/${app}:${tag}"

[[ `docker ps -a -f name=${app} -q | wc -l` -eq 1 ]] && docker rm -f ${app}
docker run -d -p 3306:3306 \
    -e MYSQL_ROOT_PASSWORD=contoro \
    --network=${docker_network} \
    --name ${app} \
    ${image_name}
