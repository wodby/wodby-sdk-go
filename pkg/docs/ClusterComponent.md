# ClusterComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Version** | **string** | The image tag or reported version as published; empty for an image referenced only by digest. | 
**Source** | **string** |  | 
**Manager** | **string** | Who can update the component. wodby is installed by Wodby outside an app, such as Cilium on K3S; infra_app runs in an infrastructure app; k3s comes with the K3S release; provider is run by the managed Kubernetes provider; host is the operating system of a server you run; and other was installed outside Wodby or is not recognized. | 
**Namespace** | **string** | Empty for node components, as are workloadKind, workload, container, image and imageDigest. | 
**WorkloadKind** | **string** |  | 
**Workload** | **string** |  | 
**Container** | **string** | Container names, separated by commas when a workload runs several from one image. | 
**Image** | **string** |  | 
**ImageDigest** | **string** |  | 
**Count** | **int32** | Running pods, or nodes for a node component. | 
**ObservedAt** | **time.Time** |  | 
**Status** | **string** | Whether the component needs anything. update_available means Wodby can apply a newer version, and latestVersion and updatePath say which version and how. bundled follows the cluster&#39;s Kubernetes version, infra_app is reported by its infrastructure app, and unknown means the check could not run or Wodby has nothing to offer. | 
**LatestVersion** | Pointer to **string** | The version an available update would apply. | [optional] 
**UpdatePath** | Pointer to **string** | The operation that updates the component. | [optional] 

## Methods

### NewClusterComponent

`func NewClusterComponent(name string, version string, source string, manager string, namespace string, workloadKind string, workload string, container string, image string, imageDigest string, count int32, observedAt time.Time, status string, ) *ClusterComponent`

NewClusterComponent instantiates a new ClusterComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterComponentWithDefaults

`func NewClusterComponentWithDefaults() *ClusterComponent`

NewClusterComponentWithDefaults instantiates a new ClusterComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClusterComponent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterComponent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterComponent) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *ClusterComponent) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ClusterComponent) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ClusterComponent) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetSource

`func (o *ClusterComponent) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ClusterComponent) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ClusterComponent) SetSource(v string)`

SetSource sets Source field to given value.


### GetManager

`func (o *ClusterComponent) GetManager() string`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *ClusterComponent) GetManagerOk() (*string, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *ClusterComponent) SetManager(v string)`

SetManager sets Manager field to given value.


### GetNamespace

`func (o *ClusterComponent) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ClusterComponent) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ClusterComponent) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetWorkloadKind

`func (o *ClusterComponent) GetWorkloadKind() string`

GetWorkloadKind returns the WorkloadKind field if non-nil, zero value otherwise.

### GetWorkloadKindOk

`func (o *ClusterComponent) GetWorkloadKindOk() (*string, bool)`

GetWorkloadKindOk returns a tuple with the WorkloadKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadKind

`func (o *ClusterComponent) SetWorkloadKind(v string)`

SetWorkloadKind sets WorkloadKind field to given value.


### GetWorkload

`func (o *ClusterComponent) GetWorkload() string`

GetWorkload returns the Workload field if non-nil, zero value otherwise.

### GetWorkloadOk

`func (o *ClusterComponent) GetWorkloadOk() (*string, bool)`

GetWorkloadOk returns a tuple with the Workload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkload

`func (o *ClusterComponent) SetWorkload(v string)`

SetWorkload sets Workload field to given value.


### GetContainer

`func (o *ClusterComponent) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *ClusterComponent) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *ClusterComponent) SetContainer(v string)`

SetContainer sets Container field to given value.


### GetImage

`func (o *ClusterComponent) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ClusterComponent) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ClusterComponent) SetImage(v string)`

SetImage sets Image field to given value.


### GetImageDigest

`func (o *ClusterComponent) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *ClusterComponent) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *ClusterComponent) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.


### GetCount

`func (o *ClusterComponent) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ClusterComponent) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ClusterComponent) SetCount(v int32)`

SetCount sets Count field to given value.


### GetObservedAt

`func (o *ClusterComponent) GetObservedAt() time.Time`

GetObservedAt returns the ObservedAt field if non-nil, zero value otherwise.

### GetObservedAtOk

`func (o *ClusterComponent) GetObservedAtOk() (*time.Time, bool)`

GetObservedAtOk returns a tuple with the ObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAt

`func (o *ClusterComponent) SetObservedAt(v time.Time)`

SetObservedAt sets ObservedAt field to given value.


### GetStatus

`func (o *ClusterComponent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterComponent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterComponent) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLatestVersion

`func (o *ClusterComponent) GetLatestVersion() string`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *ClusterComponent) GetLatestVersionOk() (*string, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *ClusterComponent) SetLatestVersion(v string)`

SetLatestVersion sets LatestVersion field to given value.

### HasLatestVersion

`func (o *ClusterComponent) HasLatestVersion() bool`

HasLatestVersion returns a boolean if a field has been set.

### GetUpdatePath

`func (o *ClusterComponent) GetUpdatePath() string`

GetUpdatePath returns the UpdatePath field if non-nil, zero value otherwise.

### GetUpdatePathOk

`func (o *ClusterComponent) GetUpdatePathOk() (*string, bool)`

GetUpdatePathOk returns a tuple with the UpdatePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatePath

`func (o *ClusterComponent) SetUpdatePath(v string)`

SetUpdatePath sets UpdatePath field to given value.

### HasUpdatePath

`func (o *ClusterComponent) HasUpdatePath() bool`

HasUpdatePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


