#!/bin/sh

#export DOCKER_BUILDKIT=1
IMAGE=kshuleshov/otus-kuber-2020-04_ingress-host-manager:latest

docker build $@ \
  --progress plain \
  --tag $IMAGE \
  . #&& docker push $IMAGE
