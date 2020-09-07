#!/bin/sh

. ./.gopath

make deploy.yaml IMG=kshuleshov/otus-kuber-2020-04_ingress-host-manager:latest
