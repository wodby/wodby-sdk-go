# ClusterAutoUpgradeVersionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowSameVersion** | **bool** | Allow newer infrastructure app stack revisions that keep the same stable semantic version. Ignored for cluster-level infrastructure versions. | 
**AllowPatch** | **bool** |  | 
**AllowMinor** | **bool** |  | 
**AllowMajor** | **bool** |  | 

## Methods

### NewClusterAutoUpgradeVersionPolicy

`func NewClusterAutoUpgradeVersionPolicy(allowSameVersion bool, allowPatch bool, allowMinor bool, allowMajor bool, ) *ClusterAutoUpgradeVersionPolicy`

NewClusterAutoUpgradeVersionPolicy instantiates a new ClusterAutoUpgradeVersionPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterAutoUpgradeVersionPolicyWithDefaults

`func NewClusterAutoUpgradeVersionPolicyWithDefaults() *ClusterAutoUpgradeVersionPolicy`

NewClusterAutoUpgradeVersionPolicyWithDefaults instantiates a new ClusterAutoUpgradeVersionPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowSameVersion

`func (o *ClusterAutoUpgradeVersionPolicy) GetAllowSameVersion() bool`

GetAllowSameVersion returns the AllowSameVersion field if non-nil, zero value otherwise.

### GetAllowSameVersionOk

`func (o *ClusterAutoUpgradeVersionPolicy) GetAllowSameVersionOk() (*bool, bool)`

GetAllowSameVersionOk returns a tuple with the AllowSameVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSameVersion

`func (o *ClusterAutoUpgradeVersionPolicy) SetAllowSameVersion(v bool)`

SetAllowSameVersion sets AllowSameVersion field to given value.


### GetAllowPatch

`func (o *ClusterAutoUpgradeVersionPolicy) GetAllowPatch() bool`

GetAllowPatch returns the AllowPatch field if non-nil, zero value otherwise.

### GetAllowPatchOk

`func (o *ClusterAutoUpgradeVersionPolicy) GetAllowPatchOk() (*bool, bool)`

GetAllowPatchOk returns a tuple with the AllowPatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPatch

`func (o *ClusterAutoUpgradeVersionPolicy) SetAllowPatch(v bool)`

SetAllowPatch sets AllowPatch field to given value.


### GetAllowMinor

`func (o *ClusterAutoUpgradeVersionPolicy) GetAllowMinor() bool`

GetAllowMinor returns the AllowMinor field if non-nil, zero value otherwise.

### GetAllowMinorOk

`func (o *ClusterAutoUpgradeVersionPolicy) GetAllowMinorOk() (*bool, bool)`

GetAllowMinorOk returns a tuple with the AllowMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMinor

`func (o *ClusterAutoUpgradeVersionPolicy) SetAllowMinor(v bool)`

SetAllowMinor sets AllowMinor field to given value.


### GetAllowMajor

`func (o *ClusterAutoUpgradeVersionPolicy) GetAllowMajor() bool`

GetAllowMajor returns the AllowMajor field if non-nil, zero value otherwise.

### GetAllowMajorOk

`func (o *ClusterAutoUpgradeVersionPolicy) GetAllowMajorOk() (*bool, bool)`

GetAllowMajorOk returns a tuple with the AllowMajor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMajor

`func (o *ClusterAutoUpgradeVersionPolicy) SetAllowMajor(v bool)`

SetAllowMajor sets AllowMajor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


