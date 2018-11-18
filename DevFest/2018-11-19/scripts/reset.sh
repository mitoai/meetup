#!/usr/bin/env bash

helm delete --purge cert-manager weather-prod weather-beta

kubectl get customresourcedefinitions.apiextensions.k8s.io -o json | jq -r .items[].metadata.name | xargs kubectl delete customresourcedefinitions.apiextensions.k8s.io


