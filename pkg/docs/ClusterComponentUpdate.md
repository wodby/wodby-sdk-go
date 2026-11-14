# ClusterComponentUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | cilium upgrades the Cilium release on a K3s cluster; provider_addon asks the provider to update an add-on. | 
**Name** | **string** | Identifies the update within its cluster. | 
**CurrentVersion** | Pointer to **string** | Absent when the update replaces a copy the provider does not manage yet. | [optional] 
**TargetVersion** | **string** |  | 
**Components** | **[]string** | The names of the cluster components this update changes. | 
**Note** | Pointer to **string** | What to expect while the update runs. | [optional] 

## Methods

### NewClusterComponentUpdate

`func NewClusterComponentUpdate(kind string, name string, targetVersion string, components []string, ) *ClusterComponentUpdate`

NewClusterComponentUpdate instantiates a new ClusterComponentUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterComponentUpdateWithDefaults

`func NewClusterComponentUpdateWithDefaults() *ClusterComponentUpdate`

NewClusterComponentUpdateWithDefaults instantiates a new ClusterComponentUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ClusterComponentUpdate) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ClusterComponentUpdate) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ClusterComponentUpdate) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *ClusterComponentUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterComponentUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterComponentUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetCurrentVersion

`func (o *ClusterComponentUpdate) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *ClusterComponentUpdate) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *ClusterComponentUpdate) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *ClusterComponentUpdate) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetTargetVersion

`func (o *ClusterComponentUpdate) GetTargetVersion() string`

GetTargetVersion returns the TargetVersion field if non-nil, zero value otherwise.

### GetTargetVersionOk

`func (o *ClusterComponentUpdate) GetTargetVersionOk() (*string, bool)`

GetTargetVersionOk returns a tuple with the TargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersion

`func (o *ClusterComponentUpdate) SetTargetVersion(v string)`

SetTargetVersion sets TargetVersion field to given value.


### GetComponents

`func (o *ClusterComponentUpdate) GetComponents() []string`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ClusterComponentUpdate) GetComponentsOk() (*[]string, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ClusterComponentUpdate) SetComponents(v []string)`

SetComponents sets Components field to given value.


### GetNote

`func (o *ClusterComponentUpdate) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ClusterComponentUpdate) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ClusterComponentUpdate) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *ClusterComponentUpdate) HasNote() bool`

HasNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


