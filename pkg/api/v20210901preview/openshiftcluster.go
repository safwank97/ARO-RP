package v20210901preview

import "time"

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

// OpenShiftClusterList represents a list of OpenShift clusters.
type OpenShiftClusterList struct {
	// The list of OpenShift clusters.
	OpenShiftClusters []*OpenShiftCluster `json:"value"`

	// The link used to get the next page of operations.
	NextLink string `json:"nextLink,omitempty"`
}

// OpenShiftCluster represents an Azure Red Hat OpenShift cluster.
type OpenShiftCluster struct {
	// The resource ID.
	ID string `json:"id,omitempty" mutable:"case"`

	// The resource name.
	Name string `json:"name,omitempty" mutable:"case"`

	// The resource type.
	Type string `json:"type,omitempty" mutable:"case"`

	// The resource location.
	Location string `json:"location,omitempty"`

	// SystemData - The system metadata relating to this resource
	SystemData *SystemData `json:"systemData,omitempty"`

	// The resource tags.
	Tags Tags `json:"tags,omitempty" mutable:"true"`

	// The cluster properties.
	Properties OpenShiftClusterProperties `json:"properties,omitempty"`
}

// Tags represents an OpenShift cluster's tags.
type Tags map[string]string

// OpenShiftClusterProperties represents an OpenShift cluster's properties.
type OpenShiftClusterProperties struct {
	// The cluster provisioning state.
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`

	// The cluster profile.
	ClusterProfile ClusterProfile `json:"clusterProfile,omitempty"`

	// The console profile.
	ConsoleProfile ConsoleProfile `json:"consoleProfile,omitempty"`

	// The cluster service principal profile.
	ServicePrincipalProfile *ServicePrincipalProfile `json:"servicePrincipalProfile,omitempty"`

	// The cluster network profile.
	NetworkProfile NetworkProfile `json:"networkProfile,omitempty"`

	// The cluster master profile.
	MasterProfile MasterProfile `json:"masterProfile,omitempty"`

	// The cluster worker profiles.
	WorkerProfiles []WorkerProfile `json:"workerProfiles,omitempty"`

	// The cluster API server profile.
	APIServerProfile APIServerProfile `json:"apiserverProfile,omitempty"`

	// The cluster ingress profiles.
	IngressProfiles []IngressProfile `json:"ingressProfiles,omitempty"`
}

// ProvisioningState represents a provisioning state.
type ProvisioningState string

// ProvisioningState constants.
const (
	ProvisioningStateCreating      ProvisioningState = "Creating"
	ProvisioningStateUpdating      ProvisioningState = "Updating"
	ProvisioningStateAdminUpdating ProvisioningState = "AdminUpdating"
	ProvisioningStateDeleting      ProvisioningState = "Deleting"
	ProvisioningStateSucceeded     ProvisioningState = "Succeeded"
	ProvisioningStateFailed        ProvisioningState = "Failed"
)

// ClusterProfile represents a cluster profile.
type ClusterProfile struct {
	// The pull secret for the cluster.
	PullSecret string `json:"pullSecret,omitempty"`

	// The domain for the cluster.
	Domain string `json:"domain,omitempty"`

	// The version of the cluster.
	Version string `json:"version,omitempty"`

	// The ID of the cluster resource group.
	ResourceGroupID string `json:"resourceGroupId,omitempty"`
}

// ConsoleProfile represents a console profile.
type ConsoleProfile struct {
	// The URL to access the cluster console.
	URL string `json:"url,omitempty"`
}

// ServicePrincipalProfile represents a service principal profile.
type ServicePrincipalProfile struct {
	// The client ID used for the cluster.
	ClientID string `json:"clientId,omitempty" mutable:"true"`

	// The client secret used for the cluster.
	ClientSecret string `json:"clientSecret,omitempty" mutable:"true"`
}

// SoftwareDefinedNetwork constants.
type SoftwareDefinedNetwork string

const (
	SoftwareDefinedNetworkOVNKubernetes SoftwareDefinedNetwork = "OVNKubernetes"
	SoftwareDefinedNetworkOpenShiftSDN  SoftwareDefinedNetwork = "OpenShiftSDN"
)

// NetworkProfile represents a network profile.
type NetworkProfile struct {
	// The software defined network (SDN) to use when installing the cluster.
	SoftwareDefinedNetwork SoftwareDefinedNetwork `json:"softwareDefinedNetwork,omitempty"`

	// The CIDR used for OpenShift/Kubernetes Pods.
	PodCIDR string `json:"podCidr,omitempty"`

	// The CIDR used for OpenShift/Kubernetes Services.
	ServiceCIDR string `json:"serviceCidr,omitempty"`
}

// EncryptionAtHost represents encryption at host state
type EncryptionAtHost string

// EncryptionAtHost constants
const (
	EncryptionAtHostEnabled  EncryptionAtHost = "Enabled"
	EncryptionAtHostDisabled EncryptionAtHost = "Disabled"
)

// MasterProfile represents a master profile.
type MasterProfile struct {
	// The size of the master VMs.
	VMSize VMSize `json:"vmSize,omitempty"`

	// The Azure resource ID of the master subnet.
	SubnetID string `json:"subnetId,omitempty"`

	// Whether master virtual machines are encrypted at host.
	EncryptionAtHost EncryptionAtHost `json:"encryptionAtHost,omitempty"`

	// The resource ID of an associated DiskEncryptionSet, if applicable.
	DiskEncryptionSetID string `json:"diskEncryptionSetId,omitempty"`
}

// VMSize represents a VM size.
type VMSize string

// VMSize constants
const (
	VMSizeStandardD2sV3 VMSize = "Standard_D2s_v3"

	VMSizeStandardD4asV4  VMSize = "Standard_D4as_v4"
	VMSizeStandardD8asV4  VMSize = "Standard_D8as_v4"
	VMSizeStandardD16asV4 VMSize = "Standard_D16as_v4"
	VMSizeStandardD32asV4 VMSize = "Standard_D32as_v4"

	VMSizeStandardD4sV3  VMSize = "Standard_D4s_v3"
	VMSizeStandardD8sV3  VMSize = "Standard_D8s_v3"
	VMSizeStandardD16sV3 VMSize = "Standard_D16s_v3"
	VMSizeStandardD32sV3 VMSize = "Standard_D32s_v3"

	VMSizeStandardE4sV3   VMSize = "Standard_E4s_v3"
	VMSizeStandardE8sV3   VMSize = "Standard_E8s_v3"
	VMSizeStandardE16sV3  VMSize = "Standard_E16s_v3"
	VMSizeStandardE32sV3  VMSize = "Standard_E32s_v3"
	VMSizeStandardE64isV3 VMSize = "Standard_E64is_v3"
	VMSizeStandardE64iV3  VMSize = "Standard_E64i_v3"

	VMSizeStandardF4sV2  VMSize = "Standard_F4s_v2"
	VMSizeStandardF8sV2  VMSize = "Standard_F8s_v2"
	VMSizeStandardF16sV2 VMSize = "Standard_F16s_v2"
	VMSizeStandardF32sV2 VMSize = "Standard_F32s_v2"
	VMSizeStandardF72sV2 VMSize = "Standard_F72s_v2"

	VMSizeStandardM128ms VMSize = "Standard_M128ms"
	VMSizeStandardG5     VMSize = "Standard_G5"
	VMSizeStandardGS5    VMSize = "Standard_GS5"
)

// WorkerProfile represents a worker profile.
type WorkerProfile struct {
	// The worker profile name.
	Name string `json:"name,omitempty"`

	// The size of the worker VMs.
	VMSize VMSize `json:"vmSize,omitempty"`

	// The disk size of the worker VMs.
	DiskSizeGB int `json:"diskSizeGB,omitempty"`

	// The Azure resource ID of the worker subnet.
	SubnetID string `json:"subnetId,omitempty"`

	// The number of worker VMs.
	Count int `json:"count,omitempty"`

	// Whether master virtual machines are encrypted at host.
	EncryptionAtHost EncryptionAtHost `json:"encryptionAtHost,omitempty"`

	// The resource ID of an associated DiskEncryptionSet, if applicable.
	DiskEncryptionSetID string `json:"diskEncryptionSetId,omitempty"`
}

// APIServerProfile represents an API server profile.
type APIServerProfile struct {
	// API server visibility.
	Visibility Visibility `json:"visibility,omitempty"`

	// The URL to access the cluster API server.
	URL string `json:"url,omitempty"`

	// The IP of the cluster API server.
	IP string `json:"ip,omitempty"`
}

// Visibility represents visibility.
type Visibility string

// Visibility constants
const (
	VisibilityPublic  Visibility = "Public"
	VisibilityPrivate Visibility = "Private"
)

// IngressProfile represents an ingress profile.
type IngressProfile struct {
	// The ingress profile name.
	Name string `json:"name,omitempty"`

	// Ingress visibility.
	Visibility Visibility `json:"visibility,omitempty"`

	// The IP of the ingress.
	IP string `json:"ip,omitempty"`
}

// CreatedByType by defines user type, which executed the request
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// SystemData metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The identity that created the resource.
	CreatedBy string `json:"createdBy,omitempty"`
	// The type of identity that created the resource. Possible values include: 'User', 'Application', 'ManagedIdentity', 'Key'
	CreatedByType CreatedByType `json:"createdByType,omitempty"`
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The identity that last modified the resource.
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`
	// The type of identity that last modified the resource. Possible values include: 'User', 'Application', 'ManagedIdentity', 'Key'
	LastModifiedByType CreatedByType `json:"lastModifiedByType,omitempty"`
	// The type of identity that last modified the resource.
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`
}
