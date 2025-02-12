# Deploying the Demo Application

This section defines kubernetes manifests to deploy:

- The services defined in `06-demo-application`
- Postgresql (via [a helm chart](https://github.com/bitnami/charts/tree/main/bitnami/postgresql)
- Traefik ingress controller (and config to route traffic to the application)

## Layout

The following shows the layout of this directory:

```
.
├── api-golang
│    ├── Deployment.yaml
│    ├── IngressRoute.yaml
│    ├── Secret.yml
│    └── Service.yaml
├── client
│    ├── ConfigMap.yaml
│    ├── Deployment.yaml
│    ├── IngressRoute.yaml
│    └── Service.yaml
├── common
│    ├── Middleware.yaml
│    ├── Namespace.yaml
│    └── Taskfile.yaml
├── create-cluster
│    ├── kind-bind-mount-1
│    │   └── hello-from-host
│    ├── kind-config.yaml.TEMPLATE
│    ├── README.md
│    └── Taskfile.yaml
├── load-generator
│    ├── ConfigMap.yaml
│    └── Deployment.yaml
├── Namespace.yaml
├── postgresql
│    ├── Job.db-migrator.yaml
│    ├── Secret.db-password.yaml
│    ├── Taskfile.yaml
│    └── values.yaml
├── README.md
└── Taskfile.yaml

```

## Breakdown

Helm is used to install:

1. `Postgresql`
2. `Traefik` (ingress controller)

Each service effectively gets broken down into:

1. `Deployment`: Contains the application
2. `Secret`: Contains database credentials
3. `Service`: Internal load balancer routing traffic to the deployment pods.
4. `IngressRoute`: Configures Ingress controller to route traffic to the proper service

The database migration is structured as a Kubernetes `Job`.

## Tasks

The top level `Taskfile` provides all necessary commands to deploy the application:

```bash
$ task --list-all
task: Available tasks for this project:
* prepare:                                         
* start:                                           
* stop:                                            
* api-golang:apply:                                Apply kubernetes resource manifests: api-golang
* client:apply:                                    Apply kubernetes resource manifests: client
* cluster:start:                                   Start a cluster with KinD
* cluster:stop:                                    Stop and clean up cluster
* common:apply-namespace:                          Apply Kubernetes Namespace
* common:apply-traefik-middleware:                 Deploy Traefik middleware
* common:deploy-traefik:                           Deploy Traefik using Helm
* load-generator:apply:                            Apply kubernetes resource manifests: load-generator
* load-generator:create-image-pull-secret:         Create image pull secret to pull from private registry
* postgresql:apply-initial-db-migration-job:       Run init.sql script against the DB
* postgresql:install-postgres:                     Deploy PostgreSQL using Helm


```

__commands__

- these commands will get you going.  For more, read the [Taskfile.yaml](Taskfile.yaml))
```bash
task prepare
task start
task stop
```

- now let's check out what happened

```bash
################################################
# make an alias to make things easier
# otherwise write out `kubectl` instead of `k` below
alias k=kubectl

################################################
k get service
NAME         TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)    AGE
api-golang   ClusterIP   10.96.23.203   <none>        8000/TCP   5m25s
client       ClusterIP   10.96.230.56   <none>        8080/TCP   5m24s

################################################
k get pods -n kube-system
NAME                                         READY   STATUS    RESTARTS   AGE
coredns-668d6bf9bc-j9gpc                     1/1     Running   0          15m
coredns-668d6bf9bc-zppdj                     1/1     Running   0          15m
etcd-kind-control-plane                      1/1     Running   0          16m
kindnet-fdtzq                                1/1     Running   0          15m
kindnet-rjjfm                                1/1     Running   0          15m
kindnet-w7j47                                1/1     Running   0          15m
kube-apiserver-kind-control-plane            1/1     Running   0          16m
kube-controller-manager-kind-control-plane   1/1     Running   0          16m
kube-proxy-7bqvw                             1/1     Running   0          15m
kube-proxy-95thj                             1/1     Running   0          15m
kube-proxy-m972h                             1/1     Running   0          15m
kube-scheduler-kind-control-plane            1/1     Running   0          16m

k get pods
NAME                             READY   STATUS      RESTARTS        AGE
api-golang-594bf5d656-7trjm      1/1     Running     3 (5m46s ago)   6m24s
client-686588487f-8s58p          1/1     Running     0               6m23s
db-migrator-585rw                0/1     Completed   3               6m27s
load-generator-986db98fb-4sch9   1/1     Running     0               99s

################################################
k get pods -o wide
NAME                             READY   STATUS      RESTARTS        AGE     IP           NODE           NOMINATED NODE   READINESS GATES
api-golang-594bf5d656-7trjm      1/1     Running     3 (9m51s ago)   10m     10.244.1.3   kind-worker2   <none>           <none>
client-686588487f-8s58p          1/1     Running     0               10m     10.244.1.4   kind-worker2   <none>           <none>
db-migrator-585rw                0/1     Completed   3               10m     10.244.2.2   kind-worker    <none>           <none>
load-generator-986db98fb-4sch9   1/1     Running     0               5m44s   10.244.2.4   kind-worker    <none>           <none>

```

