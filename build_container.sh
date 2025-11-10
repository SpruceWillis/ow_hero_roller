#! /bin/bash

TAG="us-west1-docker.pkg.dev/ow-hero-roller-1741411768301/ow-hero-roller-docker/ow_hero_roller:$(git rev-parse --short HEAD)"
docker buildx build -t $TAG --platform linux/amd64 .
echo "Built image with tag: $TAG"
echo "Run 'docker push $TAG' to push the image to the container registry."
