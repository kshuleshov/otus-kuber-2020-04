#!/bin/bash

ip=$(./get-ingress-ip.sh)

curl -v \
  -H 'kbn-xsrf: true' \
  -H 'Content-Type: application/json' \
  -d '{"type":["dashboard","index-pattern"],"excludeExportDetails": true,"includeReferencesDeep":true}' \
  http://kibana.$ip.xip.io/api/saved_objects/_export \
  > kibana.export.ndjson
