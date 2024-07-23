# Development environment using kind
Using `kind` we can start a lightweight environment for **assisted-service** development.  
The **assisted-service** does not require all the components of Openshift Container Platform, so we can use this environment to save time and resources.  

## Requirements
We usually use the **podman** provider with `kind`, so it must be installed in our Operating System.  
Nowadays, **podman** is available in the packaging system of all major GNU/Linux distributions.  
Ir order to manage the kind cluster, we will also need to have [kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl9) or [oc](https://docs.openshift.com/container-platform/4.8/cli_reference/openshift_cli/getting-started-cli.html) installed.

## Build the environment
This repository provides some scripts to install `kind`, **assisted-service** and their dependencies.  
To create our development environment:
```shell
<repository-dir>/hack/kind/dev-assisted-service.sh
```
When the script finishes, we will have an environment ready to work.  
The credentials for the new cluster will saved in the file `${HOME}/.kube/kind-assisted-service`.  
We can validate if it is working by checking the running pods:
```shell
$ export KUBECONFIG="${HOME}/.kube/kind-assisted-service"

$ kubectl get nodes
NAME                                    STATUS   ROLES           AGE     VERSION
assisted-installer-kind-control-plane   Ready    control-plane   6m23s   v1.30.0

$ kubectl get pods -A
NAMESPACE                   NAME                                                            READY   STATUS    RESTARTS   AGE
assisted-installer          agentinstalladmission-67c949b84-7l2hr                           1/1     Running   0          6m48s
assisted-installer          agentinstalladmission-67c949b84-jchzp                           1/1     Running   0          6m48s
assisted-installer          assisted-image-service-0                                        1/1     Running   0          6m48s
assisted-installer          assisted-service-5b68cbdfc6-rg4nf                               2/2     Running   0          6m48s
baremetal-operator-system   baremetal-operator-controller-manager-5546cbc489-m2vvq          2/2     Running   0          7m10s
baremetal-operator-system   ironic-659b44f9c8-mghwb                                         3/3     Running   0          10m
cert-manager                cert-manager-cainjector-9d956987c-tzds4                         1/1     Running   0          11m
cert-manager                cert-manager-fdd97855b-2v8nh                                    1/1     Running   0          11m
cert-manager                cert-manager-webhook-9f799c7d7-t7cvr                            1/1     Running   0          11m
kube-system                 coredns-7db6d8ff4d-jmw84                                        1/1     Running   0          11m
kube-system                 coredns-7db6d8ff4d-rgr8r                                        1/1     Running   0          11m
kube-system                 etcd-assisted-installer-kind-control-plane                      1/1     Running   0          11m
kube-system                 kindnet-mnhts                                                   1/1     Running   0          11m
kube-system                 kube-apiserver-assisted-installer-kind-control-plane            1/1     Running   0          11m
kube-system                 kube-controller-manager-assisted-installer-kind-control-plane   1/1     Running   0          11m
kube-system                 kube-proxy-cgv7b                                                1/1     Running   0          11m
kube-system                 kube-scheduler-assisted-installer-kind-control-plane            1/1     Running   0          11m
local-path-storage          local-path-provisioner-988d74bc-2ww6k                           1/1     Running   0          11m
```

## assisted-service configuration
In this environment, the configuration of the assisted-service is managed by a `ConfigMap`.  
It is located in the `assisted-installer` namespace and is called `assisted-service`.  
We can see or manage the options by editing it:
```shell
$ kubectl edit cm -n assisted-installer assisted-service
```
Remember to restart the **assisted-service** pods after the changes:
```shell
$ kubectl rollout restart deploy -n assisted-installer assisted-service
```

## Disable the REST API authentication
If we want to disable the REST API authentication of the **assisted-service** for our tests we can change the option `AUTH_TYPE` in the configuration and restart the deployment to apply the changes:
```shell
$ kubectl patch cm/assisted-service -n assisted-installer -p '{"data":{"AUTH_TYPE":"none"}}'
$ kubectl rollout restart deploy -n assisted-installer assisted-service
```
After this change we can send REST requests as explained [here](../../user-guide/rest-api-getting-started.md) without using the authentication bearer.  
Our development environment exposes the REST API in the port **8090**.

## Replace assisted-service image
During our development perhaps we would like to replace the current `assisted-service` image with our own custom image.  
To replace the **assisted-service** image that we are using we can do it by changing the `image` in the `assisted-service` deployment:
```shell
kubectl patch deploy/assisted-service \
  -n assisted-installer \
  --type merge \
  -p '{"spec":{"template":{"spec":{"containers":[{"name":"assisted-service","image":"<YOUR IMAGE>:<TAG>"}]}}}}'
```

## Delete the environment
To delete the environment simply execute the `kind.sh` script with the `delete` parameter:
```shell
<repository-dir>/hack/kind/kind.sh delete
```

## Additional working environments
This document covers the setup of a **assisted-service** development scenario using previous work provided by the [sylva-poc project](https://github.com/jianzzha/sylva-poc), but it's not the only option.

It is also possible to configure an **assisted-service** development environment using the infrastructure operator. This setup is described in the document [operator-on-kind](../operator-on-kind.md).

And, if you only need to start up the assisted-service application, for small tests without integration with other operators, you can also do it using podman following the instructions written in [deploy/podman](../../../deploy/podman/README.md).