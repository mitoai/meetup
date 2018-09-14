#!/usr/bin/env bash

TAG="${VERSION:-latest}"

helm upgrade \
    -i \
    --wait \
    --set image.tag=${TAG} \
    weather \
    ./helm/weather
