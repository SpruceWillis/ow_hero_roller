#! /bin/bash

docker build -t us-west1-docker.pkg.dev/ow-hero-roller-1741411768301/ow-hero-roller-docker/ow_hero_roller:$(git rev-parse --short HEAD) --platform linux/amd64 .