#!/bin/bash

app="fast-writing-es"
tag=${1:-"latest"}
image_name="asia-northeast1-docker.pkg.dev/anan-project/fast-writing/${app}:${tag}"

build () {
    docker build -t "${image_name}" .
}

push () {
    docker push $image_name
}

# build && push
build