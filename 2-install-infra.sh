#!/bin/bash -ex

HF_ENVIRONMENT=${HF_ENVIRONMENT:=default}

helmfile --environment "$HF_ENVIRONMENT" --log-level debug apply
