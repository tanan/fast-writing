#!/bin/bash

# DOCKER_BUILDKIT=1
# export DOCKER_BUILDKIT

app="contoro-db"
tag=${1:-"latest"}
image_name="gcr.io/anan-project/contoro/${app}:${tag}"

# docker_dir=$(cd $(dirname $0); pwd)
# app_dir="$docker_dir/../../../apps/${app}"

build () {
    # cd $app_dir
    # docker build -f "${docker_dir}/Dockerfile" -t "${image_name}" .
    docker build -t "${image_name}" .
}

push () {
    docker push $image_name
}

#build && push
build
