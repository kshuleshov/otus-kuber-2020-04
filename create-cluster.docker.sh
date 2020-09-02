#!/bin/bash -ex

export GCP_PROJECT=${GCP_PROJECT:=stoked-folder-278518}
GCP_SERVICE_KEY_FILE=${GCP_SERVICE_KEY_FILE:=stoked-folder-278518-7a22e15d932d.json}
export GKE_CLUSTER=${GKE_CLUSTER:=cluster-1}
export GKE_ZONE=${GKE_ZONE:=europe-north1-a}

gcloud version

# Google Cloud service accounts
if [ -n "$GCP_SERVICE_KEY" ]
then
  GCP_SERVICE_KEY_FILE=gcloud-service-key.json
  echo $GCP_SERVICE_KEY > ${GCP_SERVICE_KEY_FILE} 
fi
gcloud auth activate-service-account --key-file ${GCP_SERVICE_KEY_FILE}

gcloud beta container --project "$GCP_PROJECT" clusters list

if (gcloud beta container --project "$GCP_PROJECT" clusters list --zone "$GKE_ZONE" | grep -q "$GKE_CLUSTER")
then
  gcloud beta container --project "$GCP_PROJECT" clusters get-credentials "$GKE_CLUSTER" --zone "$GKE_ZONE"
else
  ./1-create-cluster.sh
fi

./2-install-infra.sh
