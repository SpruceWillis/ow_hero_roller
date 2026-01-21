#! /bin/bash

PUSH=false
REGION='us-west1'
TAG_BASE=''

while getopts ":pr:t:" flag; do
    case "${flag}" in
        p) PUSH=true ;;
        r) REGION=${OPTARG} ;;
        t) TAG_BASE=${OPTARG}
    esac
done

TAG="${REGION}-docker.pkg.dev/${TAG_BASE}:$(git rev-parse --short HEAD)"
docker buildx build -t $TAG --platform linux/amd64 .
echo "Built image with tag: $TAG"

if [ "$PUSH" = true ]; then
    docker push $TAG
else
    echo "Run 'docker push $TAG' to push the image to the container registry."
fi
