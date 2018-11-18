#!/usr/bin/env bash

kubectl apply -f setup/rbac.yaml
helm init --service-account tiller
