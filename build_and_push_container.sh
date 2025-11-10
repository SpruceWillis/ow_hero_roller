#! /bin/bash

TAG="us-west1-docker.pkg.dev/ow-hero-roller-1741411768301/ow-hero-roller-docker/ow_hero_roller:$(git rev-parse --short HEAD)"
docker build -t $TAG --platform linux/amd64 .
docker push $TAG
