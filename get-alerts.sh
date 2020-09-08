#!/bin/sh

kubectl get prometheusrules -A -o go-template-file=get-alerts.gotmpl \
| sed -e 's/{{[^{}]\+}}/{}/g' -e 's/{{.\+}}/{}/g' \
| sed -e :a -e '$!N;s/\([^|]\)\n/\1<br>/;ta' -e 'P;D'
