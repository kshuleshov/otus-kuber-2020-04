#!/bin/bash

get_ingress_ip="kubectl get svc ingress-nginx-controller -n ingress-nginx -o jsonpath={.status.loadBalancer.ingress[0].ip}"

ip=$($get_ingress_ip)
[ -z $ip ] && sleep 10 && ip=$($get_ingress_ip)
[ -z $ip ] && sleep 30 && ip=$($get_ingress_ip)

echo -n $ip
