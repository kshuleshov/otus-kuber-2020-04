#!/bin/bash

docker run --rm -ti \
 -v $(pwd):$(pwd) \
 --workdir $(pwd) \
 kshuleshov/otus-kuber-2020-04_gcloud-helmfile ./delete-cluster.docker.sh
