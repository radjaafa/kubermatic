package common

import (
	"fmt"
	"strings"

	"github.com/docker/distribution/reference"
	"go.uber.org/zap"

	operatorv1alpha1 "github.com/kubermatic/kubermatic/api/pkg/crd/operator/v1alpha1"
	"github.com/kubermatic/kubermatic/api/pkg/resources"

	certmanagerv1alpha2 "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1alpha2"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/utils/pointer"
)

const (
	defaultPProfEndpoint  = ":6600"
	defaultNodePortRange  = "30000-32767"
	defaultEtcdVolumeSize = "5Gi"
	defaultAuthClientID   = "kubermatic"
)

var (
	kubernetesDefaultAddons = []string{
		"canal",
		"csi",
		"dns",
		"kube-proxy",
		"openvpn",
		"rbac",
		"kubelet-configmap",
		"default-storage-class",
		"nodelocal-dns-cache",
		"pod-security-policy",
		"logrotate",
	}

	defaultAccessibleAddons = []string{
		"node-exporter",
	}

	defaultUIResources = corev1.ResourceRequirements{
		Requests: corev1.ResourceList{
			corev1.ResourceCPU:    resource.MustParse("100m"),
			corev1.ResourceMemory: resource.MustParse("64Mi"),
		},
		Limits: corev1.ResourceList{
			corev1.ResourceCPU:    resource.MustParse("250m"),
			corev1.ResourceMemory: resource.MustParse("128Mi"),
		},
	}

	defaultAPIResources = corev1.ResourceRequirements{
		Requests: corev1.ResourceList{
			corev1.ResourceCPU:    resource.MustParse("100m"),
			corev1.ResourceMemory: resource.MustParse("512Mi"),
		},
		Limits: corev1.ResourceList{
			corev1.ResourceCPU:    resource.MustParse("250m"),
			corev1.ResourceMemory: resource.MustParse("1Gi"),
		},
	}

	defaultMasterControllerMgrResources = corev1.ResourceRequirements{
		Requests: corev1.ResourceList{
			corev1.ResourceCPU:    resource.MustParse("50m"),
			corev1.ResourceMemory: resource.MustParse("128Mi"),
		},
		Limits: corev1.ResourceList{
			corev1.ResourceCPU:    resource.MustParse("100m"),
			corev1.ResourceMemory: resource.MustParse("256Mi"),
		},
	}

	defaultSeedControllerMgrResources = corev1.ResourceRequirements{
		Requests: corev1.ResourceList{
			corev1.ResourceCPU:    resource.MustParse("200m"),
			corev1.ResourceMemory: resource.MustParse("512Mi"),
		},
		Limits: corev1.ResourceList{
			corev1.ResourceCPU:    resource.MustParse("500m"),
			corev1.ResourceMemory: resource.MustParse("1Gi"),
		},
	}
)

func DefaultConfiguration(config *operatorv1alpha1.KubermaticConfiguration, logger *zap.SugaredLogger) (*operatorv1alpha1.KubermaticConfiguration, error) {
	logger.Debug("Applying defaults to Kubermatic configuration")

	copy := config.DeepCopy()

	if copy.Spec.ExposeStrategy == "" {
		copy.Spec.ExposeStrategy = operatorv1alpha1.NodePortStrategy
		logger.Debugw("Defaulting field", "field", "exposeStrategy", "value", copy.Spec.ExposeStrategy)
	}

	if copy.Spec.SeedController.BackupStoreContainer == "" {
		copy.Spec.SeedController.BackupStoreContainer = strings.TrimSpace(defaultBackupStoreContainer)
		logger.Debugw("Defaulting field", "field", "seedController.backupStoreContainer")
	}

	if copy.Spec.SeedController.BackupCleanupContainer == "" {
		copy.Spec.SeedController.BackupCleanupContainer = strings.TrimSpace(defaultBackupCleanupContainer)
		logger.Debugw("Defaulting field", "field", "seedController.backupCleanupContainer")
	}

	if copy.Spec.API.PProfEndpoint == nil {
		copy.Spec.API.PProfEndpoint = pointer.StringPtr(defaultPProfEndpoint)
		logger.Debugw("Defaulting field", "field", "api.pprofEndpoint", "value", *copy.Spec.API.PProfEndpoint)
	}

	if copy.Spec.SeedController.PProfEndpoint == nil {
		copy.Spec.SeedController.PProfEndpoint = pointer.StringPtr(defaultPProfEndpoint)
		logger.Debugw("Defaulting field", "field", "seedController.pprofEndpoint", "value", *copy.Spec.SeedController.PProfEndpoint)
	}

	if copy.Spec.MasterController.PProfEndpoint == nil {
		copy.Spec.MasterController.PProfEndpoint = pointer.StringPtr(defaultPProfEndpoint)
		logger.Debugw("Defaulting field", "field", "masterController.pprofEndpoint", "value", *copy.Spec.MasterController.PProfEndpoint)
	}

	if len(copy.Spec.UserCluster.Addons.Kubernetes.Default) == 0 && copy.Spec.UserCluster.Addons.Kubernetes.DefaultManifests == "" {
		copy.Spec.UserCluster.Addons.Kubernetes.Default = kubernetesDefaultAddons
		logger.Debugw("Defaulting field", "field", "userCluster.addons.kubernetes.default", "value", copy.Spec.UserCluster.Addons.Kubernetes.Default)
	}

	if len(copy.Spec.UserCluster.Addons.Openshift.Default) == 0 && copy.Spec.UserCluster.Addons.Openshift.DefaultManifests == "" {
		copy.Spec.UserCluster.Addons.Openshift.DefaultManifests = strings.TrimSpace(defaultOpenshiftAddons)
		logger.Debugw("Defaulting field", "field", "userCluster.Addons.Openshift.DefaultManifests")
	}

	if len(copy.Spec.API.AccessibleAddons) == 0 {
		copy.Spec.API.AccessibleAddons = defaultAccessibleAddons
		logger.Debugw("Defaulting field", "field", "copy.Spec.API.AccessibleAddons", "value", copy.Spec.API.AccessibleAddons)
	}

	if copy.Spec.UserCluster.NodePortRange == "" {
		copy.Spec.UserCluster.NodePortRange = defaultNodePortRange
		logger.Debugw("Defaulting field", "field", "userCluster.nodePortRange", "value", copy.Spec.UserCluster.NodePortRange)
	}

	if copy.Spec.UserCluster.EtcdVolumeSize == "" {
		copy.Spec.UserCluster.EtcdVolumeSize = defaultEtcdVolumeSize
		logger.Debugw("Defaulting field", "field", "userCluster.etcdVolumeSize", "value", copy.Spec.UserCluster.EtcdVolumeSize)
	}

	// cert-manager's default is Issuer, but since we do not create an Issuer,
	// it does not make sense to force to change the configuration for the
	// default case
	if copy.Spec.CertificateIssuer.Kind == "" {
		copy.Spec.CertificateIssuer.Kind = certmanagerv1alpha2.ClusterIssuerKind
		logger.Debugw("Defaulting field", "field", "certificateIssuer.kind", "value", copy.Spec.CertificateIssuer.Kind)
	}

	if copy.Spec.UI.Config == "" {
		copy.Spec.UI.Config = strings.TrimSpace(defaultUIConfig)
		logger.Debugw("Defaulting field", "field", "ui.config", "value", copy.Spec.UI.Config)
	}

	if copy.Spec.MasterFiles == nil {
		copy.Spec.MasterFiles = map[string]string{}
	}

	if copy.Spec.MasterFiles["versions.yaml"] == "" {
		copy.Spec.MasterFiles["versions.yaml"] = strings.TrimSpace(defaultVersionsYAML)
		logger.Debugw("Defaulting field", "field", "masterFiles.'versions.yaml'")
	}

	if copy.Spec.MasterFiles["updates.yaml"] == "" {
		copy.Spec.MasterFiles["updates.yaml"] = strings.TrimSpace(defaultUpdatesYAML)
		logger.Debugw("Defaulting field", "field", "masterFiles.'updates.yaml'")
	}

	auth := copy.Spec.Auth

	if auth.ClientID == "" {
		auth.ClientID = defaultAuthClientID
		logger.Debugw("Defaulting field", "field", "auth.clientID", "value", auth.ClientID)
	}

	if auth.IssuerClientID == "" {
		auth.IssuerClientID = fmt.Sprintf("%sIssuer", auth.ClientID)
		logger.Debugw("Defaulting field", "field", "auth.issuerClientID", "value", auth.IssuerClientID)
	}

	if auth.TokenIssuer == "" && copy.Spec.Domain != "" {
		auth.TokenIssuer = fmt.Sprintf("https://%s/dex", copy.Spec.Domain)
		logger.Debugw("Defaulting field", "field", "auth.tokenIssuer", "value", auth.TokenIssuer)
	}

	if auth.IssuerRedirectURL == "" && copy.Spec.Domain != "" {
		auth.IssuerRedirectURL = fmt.Sprintf("https://%s/api/v1/kubeconfig", copy.Spec.Domain)
		logger.Debugw("Defaulting field", "field", "auth.issuerRedirectURL", "value", auth.IssuerRedirectURL)
	}

	copy.Spec.Auth = auth

	if err := defaultDockerRepo(&copy.Spec.API.DockerRepository, resources.DefaultKubermaticImage, "api.dockerRepository", logger); err != nil {
		return copy, err
	}

	if err := defaultDockerRepo(&copy.Spec.UI.DockerRepository, resources.DefaultDashboardImage, "ui.dockerRepository", logger); err != nil {
		return copy, err
	}

	if err := defaultDockerRepo(&copy.Spec.MasterController.DockerRepository, resources.DefaultKubermaticImage, "masterController.dockerRepository", logger); err != nil {
		return copy, err
	}

	if err := defaultDockerRepo(&copy.Spec.SeedController.DockerRepository, resources.DefaultKubermaticImage, "seedController.dockerRepository", logger); err != nil {
		return copy, err
	}

	if err := defaultDockerRepo(&copy.Spec.UserCluster.KubermaticDockerRepository, resources.DefaultKubermaticImage, "userCluster.addons.kubermaticDockerRepository", logger); err != nil {
		return copy, err
	}

	if err := defaultDockerRepo(&copy.Spec.UserCluster.DNATControllerDockerRepository, resources.DefaultDNATControllerImage, "userCluster.addons.dnatControllerDockerRepository", logger); err != nil {
		return copy, err
	}

	if err := defaultDockerRepo(&copy.Spec.UserCluster.Addons.Kubernetes.DockerRepository, resources.DefaultKubernetesAddonImage, "userCluster.addons.kubernetes.dockerRepository", logger); err != nil {
		return copy, err
	}

	if err := defaultDockerRepo(&copy.Spec.UserCluster.Addons.Openshift.DockerRepository, resources.DefaultOpenshiftAddonImage, "userCluster.addons.openshift.dockerRepository", logger); err != nil {
		return copy, err
	}

	if err := defaultResources(&copy.Spec.UI.Resources, defaultUIResources, "ui.resources", logger); err != nil {
		return copy, err
	}

	if err := defaultResources(&copy.Spec.API.Resources, defaultAPIResources, "api.resources", logger); err != nil {
		return copy, err
	}

	if err := defaultResources(&copy.Spec.SeedController.Resources, defaultSeedControllerMgrResources, "seedController.resources", logger); err != nil {
		return copy, err
	}

	if err := defaultResources(&copy.Spec.MasterController.Resources, defaultMasterControllerMgrResources, "masterController.resources", logger); err != nil {
		return copy, err
	}

	return copy, nil
}

func defaultDockerRepo(repo *string, defaultRepo string, key string, logger *zap.SugaredLogger) error {
	if *repo == "" {
		*repo = defaultRepo
		logger.Debugw("Defaulting Docker repository", "field", key, "value", defaultRepo)

		return nil
	}

	ref, err := reference.Parse(*repo)
	if err != nil {
		return fmt.Errorf("invalid docker repository '%s' configured for %s: %v", *repo, key, err)
	}

	if _, ok := ref.(reference.Tagged); ok {
		return fmt.Errorf("it is not allowed to specify an image tag for the %s repository", key)
	}

	return nil
}

func defaultResources(settings *corev1.ResourceRequirements, defaults corev1.ResourceRequirements, key string, logger *zap.SugaredLogger) error {
	// this should never happen as the resources are not pointers in a KubermaticConfiguration
	if settings == nil {
		return nil
	}

	if err := defaultResourceList(&settings.Requests, defaults.Requests, key+".requests", logger); err != nil {
		return fmt.Errorf("failed to default requests: %v", err)
	}

	if err := defaultResourceList(&settings.Limits, defaults.Limits, key+".limits", logger); err != nil {
		return fmt.Errorf("failed to default limits: %v", err)
	}

	return nil
}

func defaultResourceList(list *corev1.ResourceList, defaults corev1.ResourceList, key string, logger *zap.SugaredLogger) error {
	if list == nil || *list == nil {
		*list = defaults
		logger.Debugw("Defaulting resource constraints", "field", key, "memory", defaults.Memory(), "cpu", defaults.Cpu())
		return nil
	}

	for _, name := range []corev1.ResourceName{corev1.ResourceMemory, corev1.ResourceCPU} {
		(*list)[name] = defaults[name]
		logger.Debugw("Defaulting resource constraint", "field", key+"."+name.String(), "value", (*list)[name])
	}

	return nil
}

const defaultBackupStoreContainer = `
name: store-container
image: quay.io/kubermatic/s3-storer:v0.1.4
command:
- /bin/sh
- -c
- |
  set -euo pipefail
  s3-storeuploader store --endpoint minio.minio.svc.cluster.local:9000 --bucket kubermatic-etcd-backups --create-bucket --prefix $CLUSTER --file /backup/snapshot.db
  s3-storeuploader delete-old-revisions --endpoint minio.minio.svc.cluster.local:9000 --bucket kubermatic-etcd-backups --prefix $CLUSTER --file /backup/snapshot.db --max-revisions 20
env:
- name: ACCESS_KEY_ID
  valueFrom:
    secretKeyRef:
      name: s3-credentials
      key: ACCESS_KEY_ID
- name: SECRET_ACCESS_KEY
  valueFrom:
    secretKeyRef:
      name: s3-credentials
      key: SECRET_ACCESS_KEY
volumeMounts:
- name: etcd-backup
  mountPath: /backup
`

const defaultBackupCleanupContainer = `
name: cleanup-container
image: quay.io/kubermatic/s3-storer:v0.1.4
command:
- /bin/sh
- -c
- |
  set -euo pipefail
  s3-storeuploader delete-all --endpoint minio.minio.svc.cluster.local:9000 --bucket kubermatic-etcd-backups --prefix $CLUSTER
env:
- name: ACCESS_KEY_ID
  valueFrom:
    secretKeyRef:
      name: s3-credentials
      key: ACCESS_KEY_ID
- name: SECRET_ACCESS_KEY
  valueFrom:
    secretKeyRef:
      name: s3-credentials
      key: SECRET_ACCESS_KEY
`

const defaultUIConfig = `
{
  "share_kubeconfig": false
}`

const defaultVersionsYAML = `
versions:
# Kubernetes 1.14
- version: "v1.14.8"
  default: false
- version: "v1.14.9"
  default: false
# Kubernetes 1.15
- version: "v1.15.5"
  default: false
- version: "v1.15.6"
  default: true
# Kubernetes 1.16
- version: "v1.16.2"
  default: false
- version: "v1.16.3"
  default: false
# Kubernetes 1.17
- version: "v1.17.0"
  default: false
# OpenShift 4.1.9
- version: "v4.1.9"
  default: false
  type: "openshift"
# OpenShift 4.1.18
- version: "v4.1.18"
  default: true
  type: "openshift"
`

const defaultUpdatesYAML = `
### Updates file
#
# Contains a list of allowed updated
#
# Each update may optionally contain 'automatic: true' in which case the
# controlplane of all clusters whose version matches the 'from' directive
# will get updated to the 'to' version. If 'automatic: true' is set, the
# 'to' version must be a version and not a version range.
#
# All 'to' versions must be configured in the 'versions.yaml'.
#
# Also, updates may contan 'automaticNodeUpdate: true', in which case
# Nodes will get updates as well. 'automaticNodeUpdate: true' sets
# 'automatic: true' as well if not yet the case, because Nodes may not have
# a newer version than the controlplane.
#
####
updates:
# ======= 1.12 =======
# Allow to next minor release
- from: 1.12.*
  to: 1.13.*
  automatic: false

# ======= 1.13 =======
# CVE-2019-11247, CVE-2019-11249, CVE-2019-9512, CVE-2019-9514
- from: <= 1.13.9, >= 1.13.0
  to: 1.13.10
  automatic: true
# Allow to next minor release
- from: 1.13.*
  to: 1.14.*
  automatic: false

# ======= 1.14 =======
# Allow to change to any patch version
- from: 1.14.*
  to: 1.14.*
  automatic: false
# CVE-2019-11247, CVE-2019-11249, CVE-2019-9512, CVE-2019-9514, CVE-2019-11253
- from: <= 1.14.7, >= 1.14.0
  to: 1.14.8
  automatic: true
# Allow to next minor release
- from: 1.14.*
  to: 1.15.*
  automatic: false

# ======= 1.15 =======
# Allow to change to any patch version
- from: 1.15.*
  to: 1.15.*
  automatic: false
# CVE-2019-11247, CVE-2019-11249, CVE-2019-9512, CVE-2019-9514, CVE-2019-11253
- from: <= 1.15.4, >= 1.15.0
  to: 1.15.5
  automatic: true
# Allow to next minor release
- from: 1.15.*
  to: 1.16.*
  automatic: false

# ======= 1.16 =======
# Allow to change to any patch version
- from: 1.16.*
  to: 1.16.*
  automatic: false
# CVE-2019-11253
- from: <= 1.16.1, >= 1.16.0
  to: 1.16.2
  automatic: true
# Allow to next minor release
- from: 1.16.*
  to: 1.17.*
  automatic: false

# ======= 1.17 =======
# Allow to change to any patch version
- from: 1.17.*
  to: 1.17.*
  automatic: false
# Allow to next minor release
- from: 1.16.*
  to: 1.18.*
  automatic: false

# ======= Openshift 4.1 =======
# Allow to change to any patch version
- from: 4.1.*
  to: 4.1.*
  automatic: false
  type: openshift
# Allow to next minor release
- from: 4.1.*
  to: 2.2.*
  automatic: false
  type: openshift
`

const defaultOpenshiftAddons = `
apiVersion: v1
kind: List
items:
- apiVersion: kubermatic.k8s.io/v1
  kind: Addon
  metadata:
    name: crd
- apiVersion: kubermatic.k8s.io/v1
  kind: Addon
  metadata:
    name: default-storage-class
- apiVersion: kubermatic.k8s.io/v1
  kind: Addon
  metadata:
    name: logrotate
- apiVersion: kubermatic.k8s.io/v1
  kind: Addon
  metadata:
    name: network
  spec:
    requiredResourceTypes:
    - Group: config.openshift.io
      Kind: Network
      Version: v1
- apiVersion: kubermatic.k8s.io/v1
  kind: Addon
  metadata:
    name: openvpn
- apiVersion: kubermatic.k8s.io/v1
  kind: Addon
  metadata:
    name: rbac
- apiVersion: kubermatic.k8s.io/v1
  kind: Addon
  metadata:
    name: registry
  spec:
    requiredResourceTypes:
    - Group: cloudcredential.openshift.io
      Kind: CredentialsRequest
      Version: v1
`