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

## Example request
```json
{
  "kind": "AdmissionReview",
  "apiVersion": "admission.k8s.io/v1beta1",
  "request": {
    "uid": "1635536b-922f-4ee1-804f-b8bcaa0c7630",
    "kind": {
      "group": "",
      "version": "v1",
      "kind": "Pod"
    },
    "resource": {
      "group": "",
      "version": "v1",
      "resource": "pods"
    },
    "requestKind": {
      "group": "",
      "version": "v1",
      "kind": "Pod"
    },
    "requestResource": {
      "group": "",
      "version": "v1",
      "resource": "pods"
    },
    "name": "nginx",
    "namespace": "default",
    "operation": "CREATE",
    "userInfo": {
      "username": "kubernetes-admin",
      "groups": [
        "system:masters",
        "system:authenticated"
      ]
    },
    "object": {
      "kind": "Pod",
      "apiVersion": "v1",
      "metadata": {
        "name": "nginx",
        "namespace": "default",
        "uid": "b8648ba0-f503-4e0f-912b-4c10bebdae55",
        "creationTimestamp": "2022-10-09T20:21:41Z",
        "labels": {
          "run": "nginx"
        },
        "annotations": {
            "example-mutating-admission-webhook": "foo"
        }
      },
      "spec": {
        "containers": [
          {
            "name": "nginx",
            "image": "nginx",
            "resources": {}
          }
        ]
      }

    },
    "options": {
      "kind": "CreateOptions",
      "apiVersion": "meta.k8s.io/v1",
      "fieldManager": "kubectl-run"
    }
  }
}
```