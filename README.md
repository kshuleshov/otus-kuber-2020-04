# otus-kuber-2020-04

[![pipeline status](https://gitlab.com/kshuleshov/otus-kuber-2020-04/badges/master/pipeline.svg)](https://gitlab.com/kshuleshov/otus-kuber-2020-04/-/commits/master)
[![coverage report](https://gitlab.com/kshuleshov/otus-kuber-2020-04/badges/master/coverage.svg)](https://gitlab.com/kshuleshov/otus-kuber-2020-04/-/commits/master)

Project work for the course [Infrastructure Platform based on Kubernetes](https://otus.ru/learning/51674/).

Task: Minimum viable product of the infrastructure platform for the demo application.

| Directory | Description |
| --------- | ----------- |
| gcloud | GCP tool image sources |
| gopath/src/ingress-host-manager | Ingress host manager sources |
| helmfile.d | Resources for bootstrapping the infrastructure platform |
| microservices-demo/deploy | Demo application |

## External resources

[GitLab CI/CD project](https://gitlab.com/kshuleshov/otus-kuber-2020-04)

## Components
### Kubernetes

Google Cloud Platform/Kubernetes Engine

Helmfile

Ingress Host Manager

### Monitoring

Prometheus, Grafana, AlertManager

#### Alerts

https://alertmanager.0.0.0.0.xip.io/

#### Dashboards

https://grafana.0.0.0.0.xip.io/

| Dashboard | Description |
| --------- | --- |
| Elasticsearch | Elasticsearch detailed dashboard |
| etcd | |
| Flux Dashboard | |
| Helm Operator Dashboard | |
| Kubernetes / API server | kubernetes-mixin |
| Kubernetes / Compute Resources / Cluster | kubernetes-mixin |
| Kubernetes / Compute Resources / Namespace (Pods) | kubernetes-mixin |
| Kubernetes / Compute Resources / Namespace (Workloads) | kubernetes-mixin |
| Kubernetes / Compute Resources / Node (Pods) | kubernetes-mixin |
| Kubernetes / Compute Resources / Pod | kubernetes-mixin |
| Kubernetes / Compute Resources / Workload | kubernetes-mixin |
| Kubernetes / Controller Manager | kubernetes-mixin |
| Kubernetes / Kubelet | kubernetes-mixin |
| Kubernetes / Networking / Cluster | kubernetes-mixin |
| Kubernetes / Networking / Namespace (Pods) | kubernetes-mixin |
| Kubernetes / Networking / Namespace (Workload) | kubernetes-mixin |
| Kubernetes / Networking / | Pod kubernetes-mixin |
| Kubernetes / Networking / Workload | kubernetes-mixin |
| Kubernetes / Persistent Volumes | kubernetes-mixin |
| Kubernetes / Proxy | kubernetes-mixin |
| Kubernetes / Scheduler | kubernetes-mixin |
| Kubernetes / StatefulSets | kubernetes-mixin |
| Nodes | |
| Prometheus Overview | |
| USE Method / Cluster | |
| USE Method / Node | |

### Centralized logging

Elasticsearch, Fluent-bit, Kibana

https://kibana.0.0.0.0.xip.io/

### CI/CD Pipeline

[GitLab](https://gitlab.com/kshuleshov/otus-kuber-2020-04)

Flux, Helm Operator

### Demo application

[Sock Shop by Weaveworks](https://microservices-demo.github.io/)

https://sock-shop.0.0.0.0.xip.io/