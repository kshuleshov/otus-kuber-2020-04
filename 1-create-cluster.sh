#!/bin/bash -ex

GCP_PROJECT=${GCP_PROJECT:=stoked-folder-278518}
GKE_CLUSTER=${GKE_CLUSTER:=cluster-1}
GKE_ZONE=${GKE_ZONE:=europe-north1-a}

gcloud beta container --project "$GCP_PROJECT" \
  clusters create "$GKE_CLUSTER" \
  --addons HorizontalPodAutoscaling,HttpLoadBalancing \
  --cluster-version "1.15.12-gke.20" \
  --default-max-pods-per-node "110" \
  --disk-size "100" \
  --disk-type "pd-standard" \
  --enable-autorepair \
  --enable-autoupgrade \
  --enable-ip-alias \
  --image-type "COS" \
  --machine-type "e2-standard-2" \
  --max-surge-upgrade 1 \
  --max-unavailable-upgrade 0 \
  --metadata disable-legacy-endpoints=true \
  --network "projects/$GCP_PROJECT/global/networks/default" \
  --no-enable-basic-auth \
  --no-enable-master-authorized-networks \
  --no-enable-stackdriver-kubernetes \
  --num-nodes "1" \
  --scopes \
"https://www.googleapis.com/auth/devstorage.read_only",\
"https://www.googleapis.com/auth/logging.write",\
"https://www.googleapis.com/auth/monitoring",\
"https://www.googleapis.com/auth/servicecontrol",\
"https://www.googleapis.com/auth/service.management.readonly",\
"https://www.googleapis.com/auth/trace.append" \
  --subnetwork "projects/$GCP_PROJECT/regions/europe-north1/subnetworks/default" \
  --zone "$GKE_ZONE"

gcloud beta container --project "$GCP_PROJECT" \
  node-pools create "infra-pool" \
  --cluster "$GKE_CLUSTER" \
  --disk-size "100" \
  --disk-type "pd-standard" \
  --enable-autorepair \
  --enable-autoupgrade \
  --image-type "COS" \
  --machine-type "e2-standard-2" \
  --max-surge-upgrade 1 \
  --max-unavailable-upgrade 0 \
  --metadata disable-legacy-endpoints=true \
  --node-taints node-role=infra:NoSchedule \
  --node-version "1.15.12-gke.2" \
  --num-nodes "3" \
  --scopes \
"https://www.googleapis.com/auth/devstorage.read_only",\
"https://www.googleapis.com/auth/logging.write",\
"https://www.googleapis.com/auth/monitoring",\
"https://www.googleapis.com/auth/servicecontrol",\
"https://www.googleapis.com/auth/service.management.readonly",\
"https://www.googleapis.com/auth/trace.append" \
  --zone "$GKE_ZONE"
