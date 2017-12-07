#!/bin/bash

# kubectl create -f php.yaml
kubectl create -f nginx.yaml
kubectl create -f fpm.yaml
kubectl create -f mysql.yaml
kubectl create -f redis.yaml

kubectl get po --namespace=lnmp-base
