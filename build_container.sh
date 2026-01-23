#! /bin/bash

PUSH=false
REGION='us-west1'
TAG_BASE=''
VERSION=$(git rev-parse --short HEAD)"

while getopts ":pr:t:v:" flag; do
    case "${flag}" in
        p) PUSH=true ;;
        r) REGION=${OPTARG} ;;
        t) TAG_BASE=${OPTARG}
        v) VERSION=${OPTARG}
    esac
done

TAG="${REGION}-docker.pkg.dev/${TAG_BASE}:${VERSION}
docker buildx build -t $TAG --platform linux/amd64 .
echo "Built image with tag: $TAG"

if [ "$PUSH" = true ]; then
    docker push $TAG
else
    echo "Run 'docker push $TAG' to push the image to the container registry."
fi
