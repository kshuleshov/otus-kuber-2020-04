#!/bin/bash -ex

export GCP_PROJECT=stoked-folder-278518
GCP_SERVICE_KEY_FILE=stoked-folder-278518-7a22e15d932d.json
export GKE_CLUSTER=cluster-1
export GKE_ZONE=europe-north1-a

gcloud beta container --project "$GCP_PROJECT" clusters get-credentials "${GKE_CLUSTER}" --zone "${GKE_ZONE}"
