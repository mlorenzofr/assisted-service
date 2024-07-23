#!/usr/bin/env bash

CERTMANAGER_VERSION="1.15.1"
KIND_HELPER="$(dirname "${0}")/kind.sh"
export KUBECONFIG="${HOME}/.kube/kind-assisted-service"

# Check if the kind-helper script is present
if ! [ -f "${KIND_HELPER}" ]; then
  echo "Script ${KIND_HELPER} not found"
  exit 1
fi

# Install kind if it is not present
"${KIND_HELPER}" check || "${KIND_HELPER}" install

# Create kind cluster
"${KIND_HELPER}" create || exit 1

# cert-manager installation
kubectl apply -f "https://github.com/cert-manager/cert-manager/releases/download/v${CERTMANAGER_VERSION}/cert-manager.yaml"
echo "Waiting for cert-manager ..."
for app in webhook cainjector cert-manager; do
  while ! kubectl wait -n cert-manager -l app="${app}" --for=condition=Ready pods --timeout=300s > /dev/null 2>&1; do
    sleep 10
  done
done

# metal3 baremetal-operator ironic deployment
kubectl apply -k https://github.com/jianzzha/sylva-poc//manifests/ironic/
echo "Waiting for ironic ..."
while ! kubectl wait -n baremetal-operator-system -l name=ironic --for=condition=Ready pods --timeout=300s > /dev/null 2>&1; do
  sleep 10
done

# metal3 baremetal-operator
kubectl apply -k https://github.com/jianzzha/sylva-poc//manifests/bmo/
echo "Waiting for baremetal-operator ..."
while ! kubectl wait -n baremetal-operator-system -l webhook=metal3-io-v1alpha1-baremetalhost --for=condition=Ready pods --timeout=300s > /dev/null 2>&1; do
  sleep 10
done

# assisted-service installation
kubectl apply -k https://github.com/jianzzha/sylva-poc//manifests/

# expose assisted-service to kind
cat <<EOF | kubectl create -f -
---
apiVersion: v1
kind: Service
metadata:
  name: assisted-service-nodeport
  namespace: assisted-installer
spec:
  type: NodePort
  ports:
  - name: http
    nodePort: 30000
    port: 8090
  selector:
    app: assisted-service
EOF
