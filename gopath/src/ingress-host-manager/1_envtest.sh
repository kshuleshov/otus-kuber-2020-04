#!/bin/sh

docker run --rm -ti \
 -v $(pwd):$(pwd) \
 --workdir $(pwd) \
 golang:1.14 make envtest SHELL=/bin/bash
