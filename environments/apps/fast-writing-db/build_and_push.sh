#!/bin/bash

app="fast-writing-db"
tag=${1:-"latest"}
image_name="gcr.io/anan-project/fast-writing/${app}:${tag}"

build () {
    docker build -t "${image_name}" .
}

push () {
    docker push $image_name
}

build && push
build
