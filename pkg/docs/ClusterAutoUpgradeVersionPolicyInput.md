# ClusterAutoUpgradeVersionPolicyInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowSameVersion** | Pointer to **NullableBool** | Allow newer infrastructure app stack revisions that keep the same stable semantic version. Ignored for cluster-level infrastructure versions. | [optional] 
**AllowPatch** | Pointer to **NullableBool** |  | [optional] 
**AllowMinor** | Pointer to **NullableBool** |  | [optional] 
**AllowMajor** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewClusterAutoUpgradeVersionPolicyInput

`func NewClusterAutoUpgradeVersionPolicyInput() *ClusterAutoUpgradeVersionPolicyInput`

NewClusterAutoUpgradeVersionPolicyInput instantiates a new ClusterAutoUpgradeVersionPolicyInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterAutoUpgradeVersionPolicyInputWithDefaults

`func NewClusterAutoUpgradeVersionPolicyInputWithDefaults() *ClusterAutoUpgradeVersionPolicyInput`

NewClusterAutoUpgradeVersionPolicyInputWithDefaults instantiates a new ClusterAutoUpgradeVersionPolicyInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowSameVersion

`func (o *ClusterAutoUpgradeVersionPolicyInput) GetAllowSameVersion() bool`

GetAllowSameVersion returns the AllowSameVersion field if non-nil, zero value otherwise.

### GetAllowSameVersionOk

`func (o *ClusterAutoUpgradeVersionPolicyInput) GetAllowSameVersionOk() (*bool, bool)`

GetAllowSameVersionOk returns a tuple with the AllowSameVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSameVersion

`func (o *ClusterAutoUpgradeVersionPolicyInput) SetAllowSameVersion(v bool)`

SetAllowSameVersion sets AllowSameVersion field to given value.

### HasAllowSameVersion

`func (o *ClusterAutoUpgradeVersionPolicyInput) HasAllowSameVersion() bool`

HasAllowSameVersion returns a boolean if a field has been set.

### SetAllowSameVersionNil

`func (o *ClusterAutoUpgradeVersionPolicyInput) SetAllowSameVersionNil(b bool)`

 SetAllowSameVersionNil sets the value for AllowSameVersion to be an explicit nil

### UnsetAllowSameVersion
`func (o *ClusterAutoUpgradeVersionPolicyInput) UnsetAllowSameVersion()`

UnsetAllowSameVersion ensures that no value is present for AllowSameVersion, not even an explicit nil
### GetAllowPatch

`func (o *ClusterAutoUpgradeVersionPolicyInput) GetAllowPatch() bool`

GetAllowPatch returns the AllowPatch field if non-nil, zero value otherwise.

### GetAllowPatchOk

`func (o *ClusterAutoUpgradeVersionPolicyInput) GetAllowPatchOk() (*bool, bool)`

GetAllowPatchOk returns a tuple with the AllowPatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPatch

`func (o *ClusterAutoUpgradeVersionPolicyInput) SetAllowPatch(v bool)`

SetAllowPatch sets AllowPatch field to given value.

### HasAllowPatch

`func (o *ClusterAutoUpgradeVersionPolicyInput) HasAllowPatch() bool`

HasAllowPatch returns a boolean if a field has been set.

### SetAllowPatchNil

`func (o *ClusterAutoUpgradeVersionPolicyInput) SetAllowPatchNil(b bool)`

 SetAllowPatchNil sets the value for AllowPatch to be an explicit nil

### UnsetAllowPatch
`func (o *ClusterAutoUpgradeVersionPolicyInput) UnsetAllowPatch()`

UnsetAllowPatch ensures that no value is present for AllowPatch, not even an explicit nil
### GetAllowMinor

`func (o *ClusterAutoUpgradeVersionPolicyInput) GetAllowMinor() bool`

GetAllowMinor returns the AllowMinor field if non-nil, zero value otherwise.

### GetAllowMinorOk

`func (o *ClusterAutoUpgradeVersionPolicyInput) GetAllowMinorOk() (*bool, bool)`

GetAllowMinorOk returns a tuple with the AllowMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMinor

`func (o *ClusterAutoUpgradeVersionPolicyInput) SetAllowMinor(v bool)`

SetAllowMinor sets AllowMinor field to given value.

### HasAllowMinor

`func (o *ClusterAutoUpgradeVersionPolicyInput) HasAllowMinor() bool`

HasAllowMinor returns a boolean if a field has been set.

### SetAllowMinorNil

`func (o *ClusterAutoUpgradeVersionPolicyInput) SetAllowMinorNil(b bool)`

 SetAllowMinorNil sets the value for AllowMinor to be an explicit nil

### UnsetAllowMinor
`func (o *ClusterAutoUpgradeVersionPolicyInput) UnsetAllowMinor()`

UnsetAllowMinor ensures that no value is present for AllowMinor, not even an explicit nil
### GetAllowMajor

`func (o *ClusterAutoUpgradeVersionPolicyInput) GetAllowMajor() bool`

GetAllowMajor returns the AllowMajor field if non-nil, zero value otherwise.

### GetAllowMajorOk

`func (o *ClusterAutoUpgradeVersionPolicyInput) GetAllowMajorOk() (*bool, bool)`

GetAllowMajorOk returns a tuple with the AllowMajor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMajor

`func (o *ClusterAutoUpgradeVersionPolicyInput) SetAllowMajor(v bool)`

SetAllowMajor sets AllowMajor field to given value.

### HasAllowMajor

`func (o *ClusterAutoUpgradeVersionPolicyInput) HasAllowMajor() bool`

HasAllowMajor returns a boolean if a field has been set.

### SetAllowMajorNil

`func (o *ClusterAutoUpgradeVersionPolicyInput) SetAllowMajorNil(b bool)`

 SetAllowMajorNil sets the value for AllowMajor to be an explicit nil

### UnsetAllowMajor
`func (o *ClusterAutoUpgradeVersionPolicyInput) UnsetAllowMajor()`

UnsetAllowMajor ensures that no value is present for AllowMajor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


