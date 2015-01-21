#!/bin/bash
#
# Deletes the compute instance for skia-snapshot-maker.
#
set -x

source vm_config.sh

gcloud compute instances delete \
  --project=$PROJECT_ID \
  --delete-disks "all" \
  --zone=$ZONE \
  $INSTANCE_NAME

