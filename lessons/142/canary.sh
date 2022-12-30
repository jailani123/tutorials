#!/bin/bash

kubectl apply -f go-app/deploy/7-service-b-v1.yaml
kubectl apply -f go-app/deploy/8-service-b-v2.yaml
kubectl apply -f go-app/deploy/9-route.yaml
