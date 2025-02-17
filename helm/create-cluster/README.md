# Create Cluster

This directory contains configurations and commands for setting up your cluster.

For development purposes, we use KinD - kubernetes clusters locally within Docker.
It is great for testing and development purposes (and doesn't cost any additional money). 

Nearly all examples in the course can be run within this cluster, except anything that requires a public DNS (you'll need to modify your /etc/hosts to fake it locally).


## Quick start

```bash
# start up your devbox session 
devbox shell

task --list-all
    task: Available tasks for this project:
    * kind:01-generate-config:               Generate kind config with local absolute paths for PV mounts
    * kind:02-create-cluster:                Create a Kubernetes cluster using kind
    * kind:03-run-cloud-provider-kind:       Run sigs.k8s.io/cloud-provider-kind@latest to enable load balancer services with KinD
    * kind:04-delete-cluster:                Delete and existing a kind Kubernetes cluster

```

__start__

```bash
task kind:01-generate-config
task kind:02-create-cluster
```

__emulate load balancer__
```bash
task kind:03-run-cloud-provider-kind
```

__clean up__

```bash
task kind:04-delete-cluster
```
