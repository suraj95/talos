#!/bin/bash

set -eou pipefail

export USER_DISKS_MOUNTS="/var/lib/extra,/var/lib/p1,/var/lib/p2"

source ./hack/test/e2e.sh

PROVISIONER=qemu
CLUSTER_NAME=e2e-${PROVISIONER}

case "${CI:-false}" in
  true)
    QEMU_FLAGS=""
    INSTALLER_TAG="${TAG}"
    ;;
  *)
    QEMU_FLAGS="--with-bootloader=false"
    INSTALLER_TAG="latest"
    ;;
esac

case "${CUSTOM_CNI_URL:-false}" in
  false)
    CUSTOM_CNI_FLAG=
    ;;
  *)
    CUSTOM_CNI_FLAG="--custom-cni-url=${CUSTOM_CNI_URL}"
    ;;
esac

case "${WITH_UEFI:-false}" in
  true)
    QEMU_FLAGS="${QEMU_FLAGS} --with-uefi"
    ;;
esac

function create_cluster {
  build_registry_mirrors

  "${TALOSCTL}" cluster create \
    --provisioner "${PROVISIONER}" \
    --name "${CLUSTER_NAME}" \
    --masters=3 \
    --mtu 1500 \
    --memory 2048 \
    --cpus 2.0 \
    --cidr 172.20.1.0/24 \
    --user-disk /var/lib/extra:100MB \
    --user-disk /var/lib/p1:100MB:/var/lib/p2:100MB \
    --install-image ${REGISTRY:-ghcr.io}/talos-systems/installer:${INSTALLER_TAG} \
    --with-init-node=false \
    --cni-bundle-url ${ARTIFACTS}/talosctl-cni-bundle-'${ARCH}'.tar.gz \
    --crashdump \
    ${REGISTRY_MIRROR_FLAGS} \
    ${QEMU_FLAGS} \
    ${CUSTOM_CNI_FLAG}

  "${TALOSCTL}" config node 172.20.1.2
}

function destroy_cluster() {
  "${TALOSCTL}" cluster destroy --name "${CLUSTER_NAME}" --provisioner "${PROVISIONER}"
}

create_cluster
get_kubeconfig
run_talos_integration_test
run_kubernetes_integration_test
destroy_cluster
