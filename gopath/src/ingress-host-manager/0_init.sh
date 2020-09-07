#!/bin/sh

. ./.gopath

operator-sdk create api --group extensions --version v1beta1 --kind Ingress --resource=false --controller=true --make=false
