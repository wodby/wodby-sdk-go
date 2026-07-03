# ClusterAutoUpgradeVersionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowPatch** | **bool** |  | 
**AllowMinor** | **bool** |  | 
**AllowMajor** | **bool** |  | 

## Methods

### NewClusterAutoUpgradeVersionPolicy

`func NewClusterAutoUpgradeVersionPolicy(allowPatch bool, allowMinor bool, allowMajor bool, ) *ClusterAutoUpgradeVersionPolicy`

NewClusterAutoUpgradeVersionPolicy instantiates a new ClusterAutoUpgradeVersionPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterAutoUpgradeVersionPolicyWithDefaults

`func NewClusterAutoUpgradeVersionPolicyWithDefaults() *ClusterAutoUpgradeVersionPolicy`

NewClusterAutoUpgradeVersionPolicyWithDefaults instantiates a new ClusterAutoUpgradeVersionPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


