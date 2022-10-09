# k8s-webhook-server

## Pre-requisites
- Docker
- Go
- K8s kind cluster

## Certificate generation
- Install [cert-manager](https://cert-manager.io/docs/installation/kubernetes/) in your cluster
```shell
make install-cert-manager
```
- Deploy the webhook server
```shell
make build
make deploy
```

## Running locally
```shell
make run-local
```