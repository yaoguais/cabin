#!/bin/bash

kubectl delete -f php.yaml
kubectl delete -f nginx.yaml
kubectl delete -f fpm.yaml
kubectl delete -f mysql.yaml
kubectl delete -f redis.yaml
