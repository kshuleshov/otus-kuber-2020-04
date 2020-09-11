#!/bin/bash -ex

export GCP_PROJECT=${GCP_PROJECT:=stoked-folder-278518}
GCP_SERVICE_KEY_FILE=${GCP_SERVICE_KEY_FILE:=stoked-folder-278518-7a22e15d932d.json}
export GKE_CLUSTER=${GKE_CLUSTER:=cluster-1}
export GKE_ZONE=${GKE_ZONE:=europe-north1-a}

gcloud version

# Google Cloud service accounts
gcloud auth activate-service-account --key-file ${GCP_SERVICE_KEY_FILE}

gcloud beta container --project "$GCP_PROJECT" clusters list
gcloud beta container --project "$GCP_PROJECT" clusters delete "$GKE_CLUSTER" --zone "$GKE_ZONE" --quiet

gcloud compute disks list --filter="name~^gke-$GKE_CLUSTER- zone~/$GKE_ZONE$ -users:*"
DISKS_NOT_IN_USE=`gcloud compute disks list --filter="name~^gke-$GKE_CLUSTER- zone~/$GKE_ZONE$ -users:*" --format='value(name)'`
gcloud compute disks delete $DISKS_NOT_IN_USE --quiet
