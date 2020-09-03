#!/bin/bash

ip=$(./get-ingress-ip.sh)

curl -X POST -v \
  -H "kbn-xsrf: true" \
  --form file=@kibana.export.ndjson \
  http://kibana.$ip.xip.io/api/saved_objects/_import?overwrite=true
