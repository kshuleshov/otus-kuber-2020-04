# otus-kuber-2020-04

[![pipeline status](https://gitlab.com/kshuleshov/otus-kuber-2020-04/badges/master/pipeline.svg)](https://gitlab.com/kshuleshov/otus-kuber-2020-04/-/commits/master)
[![coverage report](https://gitlab.com/kshuleshov/otus-kuber-2020-04/badges/master/coverage.svg)](https://gitlab.com/kshuleshov/otus-kuber-2020-04/-/commits/master)

Project work for the course [Infrastructure Platform based on Kubernetes](https://otus.ru/learning/51674/).

Task: Minimum viable product of the infrastructure platform for the demo application.

| Directory | Description |
| --------- | ----------- |
| doc | Documentation |
| gcloud | GCP tool image sources |
| gopath/src/ingress-host-manager | Ingress host manager sources |
| helmfile.d | Resources for bootstrapping the infrastructure platform |
| microservices-demo/deploy | Demo application |
| public | Pages content |

## External resources

[GitLab CI/CD project](https://gitlab.com/kshuleshov/otus-kuber-2020-04)

[GitLab Pages with live environment](https://kshuleshov.gitlab.io/otus-kuber-2020-04/)

## Components
### Kubernetes

Google Cloud Platform/Kubernetes Engine

Helmfile

Ingress Host Manager

### Monitoring

Prometheus, Grafana, AlertManager

#### Alerts

[List of alerts](doc/alerts.md).

#### Dashboards

[List of available dashboards](doc/dashboards.md).

### Centralized logging

Elasticsearch, Fluent-bit, Kibana

### CI/CD Pipeline

[GitLab Pipelines](https://gitlab.com/kshuleshov/otus-kuber-2020-04/-/pipelines)

Flux, Helm Operator

### Demo application

[Sock Shop by Weaveworks](https://microservices-demo.github.io/)

Go to [GitLab](https://gitlab.com/kshuleshov/otus-kuber-2020-04) »
Operations »
Environments »
gcloud »
![Open live environment](doc/gitlab-external-link.png)

If there is no active environment, the cluster is currently stopped.

