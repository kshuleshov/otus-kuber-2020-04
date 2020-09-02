#!/bin/bash -ex

#./1-create-cluster.sh
#./2-install-infra.sh

docker run --rm -ti \
 -v $(pwd):$(pwd) \
 --workdir $(pwd) \
 kshuleshov/otus-kuber-2020-04_gcloud-helmfile ./create-cluster.docker.sh #gcloud version
