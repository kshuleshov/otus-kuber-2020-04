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

| Group | Name | Message | Severity |
| ----- | ---- | ------- | -------- |
| alertmanager.rules | AlertmanagerConfigInconsistent | The configuration of the instances of the Alertmanager cluster `{}/{}` are out of sync.<br>{}<br>Configuration hash for pod {} is "{}"<br>{}<br> | critical |
| alertmanager.rules | AlertmanagerFailedReload | Reloading Alertmanager's configuration has failed for {}/{}. | warning |
| alertmanager.rules | AlertmanagerMembersInconsistent | Alertmanager has not found all other members of the cluster. | critical |
| etcd | etcdMembersDown | etcd cluster "{}": members are down ({}). | critical |
| etcd | etcdInsufficientMembers | etcd cluster "{}": insufficient members ({}). | critical |
| etcd | etcdNoLeader | etcd cluster "{}": member {} has no leader. | critical |
| etcd | etcdHighNumberOfLeaderChanges | etcd cluster "{}": {} leader changes within the last 15 minutes. Frequent elections may be a sign of insufficient resources, high network latency, or disruptions by other components and should be investigated. | warning |
| etcd | etcdHighNumberOfFailedGRPCRequests | etcd cluster "{}": {}% of requests for {} failed on etcd instance {}. | warning |
| etcd | etcdHighNumberOfFailedGRPCRequests | etcd cluster "{}": {}% of requests for {} failed on etcd instance {}. | critical |
| etcd | etcdGRPCRequestsSlow | etcd cluster "{}": gRPC requests to {} are taking {}s on etcd instance {}. | critical |
| etcd | etcdMemberCommunicationSlow | etcd cluster "{}": member communication with {} is taking {}s on etcd instance {}. | warning |
| etcd | etcdHighNumberOfFailedProposals | etcd cluster "{}": {} proposal failures within the last 30 minutes on etcd instance {}. | warning |
| etcd | etcdHighFsyncDurations | etcd cluster "{}": 99th percentile fync durations are {}s on etcd instance {}. | warning |
| etcd | etcdHighCommitDurations | etcd cluster "{}": 99th percentile commit durations {}s on etcd instance {}. | warning |
| etcd | etcdHighNumberOfFailedHTTPRequests | {}% of requests for {} failed on etcd instance {} | warning |
| etcd | etcdHighNumberOfFailedHTTPRequests | {}% of requests for {} failed on etcd instance {}. | critical |
| etcd | etcdHTTPRequestsSlow | etcd instance {} HTTP requests to {} are slow. | warning |
| general.rules | TargetDown | {}% of the {}/{} targets in {} namespace are down. | warning |
| general.rules | Watchdog | This is an alert meant to ensure that the entire alerting pipeline is functional.<br>This alert is always firing, therefore it should always be firing in Alertmanager<br>and always fire against a receiver. There are integrations with various notification<br>mechanisms that send a notification when this alert is not firing. For example the<br>"DeadMansSnitch" integration in PagerDuty.<br> | none |
| kube-apiserver-slos | KubeAPIErrorBudgetBurn | The API server is burning too much error budget | critical |
| kube-apiserver-slos | KubeAPIErrorBudgetBurn | The API server is burning too much error budget | critical |
| kube-apiserver-slos | KubeAPIErrorBudgetBurn | The API server is burning too much error budget | warning |
| kube-apiserver-slos | KubeAPIErrorBudgetBurn | The API server is burning too much error budget | warning |
| kube-state-metrics | KubeStateMetricsListErrors | kube-state-metrics is experiencing errors at an elevated rate in list operations. This is likely causing it to not be able to expose metrics about Kubernetes objects correctly or at all. | critical |
| kube-state-metrics | KubeStateMetricsWatchErrors | kube-state-metrics is experiencing errors at an elevated rate in watch operations. This is likely causing it to not be able to expose metrics about Kubernetes objects correctly or at all. | critical |
| kubernetes-apps | KubePodCrashLooping | Pod {}/{} ({}) is restarting {} times / 5 minutes. | warning |
| kubernetes-apps | KubePodNotReady | Pod {}/{} has been in a non-ready state for longer than 15 minutes. | warning |
| kubernetes-apps | KubeDeploymentGenerationMismatch | Deployment generation for {}/{} does not match, this indicates that the Deployment has failed but has not been rolled back. | warning |
| kubernetes-apps | KubeDeploymentReplicasMismatch | Deployment {}/{} has not matched the expected number of replicas for longer than 15 minutes. | warning |
| kubernetes-apps | KubeStatefulSetReplicasMismatch | StatefulSet {}/{} has not matched the expected number of replicas for longer than 15 minutes. | warning |
| kubernetes-apps | KubeStatefulSetGenerationMismatch | StatefulSet generation for {}/{} does not match, this indicates that the StatefulSet has failed but has not been rolled back. | warning |
| kubernetes-apps | KubeStatefulSetUpdateNotRolledOut | StatefulSet {}/{} update has not been rolled out. | warning |
| kubernetes-apps | KubeDaemonSetRolloutStuck | Only {} of the desired Pods of DaemonSet {}/{} are scheduled and ready. | warning |
| kubernetes-apps | KubeContainerWaiting | Pod {}/{} container {} has been in waiting state for longer than 1 hour. | warning |
| kubernetes-apps | KubeDaemonSetNotScheduled | {} Pods of DaemonSet {}/{} are not scheduled. | warning |
| kubernetes-apps | KubeDaemonSetMisScheduled | {} Pods of DaemonSet {}/{} are running where they are not supposed to run. | warning |
| kubernetes-apps | KubeJobCompletion | Job {}/{} is taking more than 12 hours to complete. | warning |
| kubernetes-apps | KubeJobFailed | Job {}/{} failed to complete. | warning |
| kubernetes-apps | KubeHpaReplicasMismatch | HPA {}/{} has not matched the desired number of replicas for longer than 15 minutes. | warning |
| kubernetes-apps | KubeHpaMaxedOut | HPA {}/{} has been running at max replicas for longer than 15 minutes. | warning |
| kubernetes-resources | KubeCPUOvercommit | Cluster has overcommitted CPU resource requests for Pods and cannot tolerate node failure. | warning |
| kubernetes-resources | KubeMemoryOvercommit | Cluster has overcommitted memory resource requests for Pods and cannot tolerate node failure. | warning |
| kubernetes-resources | KubeCPUQuotaOvercommit | Cluster has overcommitted CPU resource requests for Namespaces. | warning |
| kubernetes-resources | KubeMemoryQuotaOvercommit | Cluster has overcommitted memory resource requests for Namespaces. | warning |
| kubernetes-resources | KubeQuotaFullyUsed | Namespace {} is using {} of its {} quota. | info |
| kubernetes-resources | CPUThrottlingHigh | {} throttling of CPU in namespace {} for container {} in pod {}. | info |
| kubernetes-storage | KubePersistentVolumeFillingUp | The PersistentVolume claimed by {} in Namespace {} is only {} free. | critical |
| kubernetes-storage | KubePersistentVolumeFillingUp | Based on recent sampling, the PersistentVolume claimed by {} in Namespace {} is expected to fill up within four days. Currently {} is available. | warning |
| kubernetes-storage | KubePersistentVolumeErrors | The persistent volume {} has status {}. | critical |
| kubernetes-system | KubeVersionMismatch | There are {} different semantic versions of Kubernetes components running. | warning |
| kubernetes-system | KubeClientErrors | Kubernetes API server client '{}/{}' is experiencing {} errors.' | warning |
| kubernetes-system-apiserver | KubeClientCertificateExpiration | A client certificate used to authenticate to the apiserver is expiring in less than 7.0 days. | warning |
| kubernetes-system-apiserver | KubeClientCertificateExpiration | A client certificate used to authenticate to the apiserver is expiring in less than 24.0 hours. | critical |
| kubernetes-system-apiserver | AggregatedAPIErrors | An aggregated API {}/{} has reported errors. The number of errors have increased for it in the past five minutes. High values indicate that the availability of the service changes too often. | warning |
| kubernetes-system-apiserver | AggregatedAPIDown | An aggregated API {}/{} has been only {}% available over the last 5m. | warning |
| kubernetes-system-apiserver | KubeAPIDown | KubeAPI has disappeared from Prometheus target discovery. | critical |
| kubernetes-system-controller-manager | KubeControllerManagerDown | KubeControllerManager has disappeared from Prometheus target discovery. | critical |
| kubernetes-system-kubelet | KubeNodeNotReady | {} has been unready for more than 15 minutes. | warning |
| kubernetes-system-kubelet | KubeNodeUnreachable | {} is unreachable and some workloads may be rescheduled. | warning |
| kubernetes-system-kubelet | KubeletTooManyPods | Kubelet '{}' is running at {} of its Pod capacity. | warning |
| kubernetes-system-kubelet | KubeNodeReadinessFlapping | The readiness status of node {} has changed {} times in the last 15 minutes. | warning |
| kubernetes-system-kubelet | KubeletPlegDurationHigh | The Kubelet Pod Lifecycle Event Generator has a 99th percentile duration of {} seconds on node {}. | warning |
| kubernetes-system-kubelet | KubeletPodStartUpLatencyHigh | Kubelet Pod startup 99th percentile latency is {} seconds on node {}. | warning |
| kubernetes-system-kubelet | KubeletDown | Kubelet has disappeared from Prometheus target discovery. | critical |
| kubernetes-system-scheduler | KubeSchedulerDown | KubeScheduler has disappeared from Prometheus target discovery. | critical |
| node-exporter | NodeFilesystemSpaceFillingUp |  | warning |
| node-exporter | NodeFilesystemSpaceFillingUp |  | critical |
| node-exporter | NodeFilesystemAlmostOutOfSpace |  | warning |
| node-exporter | NodeFilesystemAlmostOutOfSpace |  | critical |
| node-exporter | NodeFilesystemFilesFillingUp |  | warning |
| node-exporter | NodeFilesystemFilesFillingUp |  | critical |
| node-exporter | NodeFilesystemAlmostOutOfFiles |  | warning |
| node-exporter | NodeFilesystemAlmostOutOfFiles |  | critical |
| node-exporter | NodeNetworkReceiveErrs |  | warning |
| node-exporter | NodeNetworkTransmitErrs |  | warning |
| node-exporter | NodeHighNumberConntrackEntriesUsed |  | warning |
| node-exporter | NodeTextFileCollectorScrapeError |  | warning |
| node-exporter | NodeClockSkewDetected | Clock on {} is out of sync by more than 300s. Ensure NTP is configured correctly on this host. | warning |
| node-exporter | NodeClockNotSynchronising | Clock on {} is not synchronising. Ensure NTP is configured on this host. | warning |
| node-network | NodeNetworkInterfaceFlapping | Network interface "{}" changing it's up status often on node-exporter {}/{}" | warning |
| prometheus | PrometheusBadConfig |  | critical |
| prometheus | PrometheusNotificationQueueRunningFull |  | warning |
| prometheus | PrometheusErrorSendingAlertsToSomeAlertmanagers |  | warning |
| prometheus | PrometheusErrorSendingAlertsToAnyAlertmanager |  | critical |
| prometheus | PrometheusNotConnectedToAlertmanagers |  | warning |
| prometheus | PrometheusTSDBReloadsFailing |  | warning |
| prometheus | PrometheusTSDBCompactionsFailing |  | warning |
| prometheus | PrometheusNotIngestingSamples |  | warning |
| prometheus | PrometheusDuplicateTimestamps |  | warning |
| prometheus | PrometheusOutOfOrderTimestamps |  | warning |
| prometheus | PrometheusRemoteStorageFailures |  | critical |
| prometheus | PrometheusRemoteWriteBehind |  | critical |
| prometheus | PrometheusRemoteWriteDesiredShards |  | warning |
| prometheus | PrometheusRuleFailures |  | critical |
| prometheus | PrometheusMissingRuleEvaluations |  | warning |
| prometheus-operator | PrometheusOperatorListErrors | Errors while performing List operations in controller {} in {} namespace. | warning |
| prometheus-operator | PrometheusOperatorWatchErrors | Errors while performing Watch operations in controller {} in {} namespace. | warning |
| prometheus-operator | PrometheusOperatorReconcileErrors | Errors while reconciling {} in {} Namespace. | warning |
| prometheus-operator | PrometheusOperatorNodeLookupErrors | Errors while reconciling Prometheus in {} Namespace. | warning |

#### Dashboards

| Dashboard | Description |
| --------- | --- |
| Elasticsearch | Elasticsearch detailed dashboard |
| etcd | |
| Flux Dashboard | FluxCD |
| Helm Operator Dashboard | FluxCD |
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
| Kubernetes / Networking / Pod | kubernetes-mixin |
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

### CI/CD Pipeline

[GitLab](https://gitlab.com/kshuleshov/otus-kuber-2020-04)

Flux, Helm Operator

### Demo application

[Sock Shop by Weaveworks](https://microservices-demo.github.io/)

Go to [GitLab](https://gitlab.com/kshuleshov/otus-kuber-2020-04) >
Operations >
Environments >
gcloud >
![Open live environment](gitlab-external-link-small.png)

If there is no active environment, the cluster is currently stopped.

