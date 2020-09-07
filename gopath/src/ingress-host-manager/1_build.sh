#!/bin/sh

. ./.gopath

make docker-build IMG=kshuleshov/otus-kuber-2020-04_ingress-host-manager:latest SHELL=/bin/bash
