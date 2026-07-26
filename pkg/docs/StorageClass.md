# StorageClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Provisioner** | **string** |  | 
**ReclaimPolicy** | Pointer to **NullableString** |  | [optional] 
**AllowVolumeExpansion** | **bool** |  | 
**MountOptions** | **[]string** |  | 
**VolumeBindingMode** | Pointer to **NullableString** |  | [optional] 
**IsDefault** | **bool** |  | 
**Selectable** | **bool** |  | 

## Methods

### NewStorageClass

`func NewStorageClass(name string, provisioner string, allowVolumeExpansion bool, mountOptions []string, isDefault bool, selectable bool, ) *StorageClass`

NewStorageClass instantiates a new StorageClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageClassWithDefaults

`func NewStorageClassWithDefaults() *StorageClass`

NewStorageClassWithDefaults instantiates a new StorageClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StorageClass) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageClass) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageClass) SetName(v string)`

SetName sets Name field to given value.


### GetProvisioner

`func (o *StorageClass) GetProvisioner() string`

GetProvisioner returns the Provisioner field if non-nil, zero value otherwise.

### GetProvisionerOk

`func (o *StorageClass) GetProvisionerOk() (*string, bool)`

GetProvisionerOk returns a tuple with the Provisioner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioner

`func (o *StorageClass) SetProvisioner(v string)`

SetProvisioner sets Provisioner field to given value.


### GetReclaimPolicy

`func (o *StorageClass) GetReclaimPolicy() string`

GetReclaimPolicy returns the ReclaimPolicy field if non-nil, zero value otherwise.

### GetReclaimPolicyOk

`func (o *StorageClass) GetReclaimPolicyOk() (*string, bool)`

GetReclaimPolicyOk returns a tuple with the ReclaimPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReclaimPolicy

`func (o *StorageClass) SetReclaimPolicy(v string)`

SetReclaimPolicy sets ReclaimPolicy field to given value.

### HasReclaimPolicy

`func (o *StorageClass) HasReclaimPolicy() bool`

HasReclaimPolicy returns a boolean if a field has been set.

### SetReclaimPolicyNil

`func (o *StorageClass) SetReclaimPolicyNil(b bool)`

 SetReclaimPolicyNil sets the value for ReclaimPolicy to be an explicit nil

### UnsetReclaimPolicy
`func (o *StorageClass) UnsetReclaimPolicy()`

UnsetReclaimPolicy ensures that no value is present for ReclaimPolicy, not even an explicit nil
### GetAllowVolumeExpansion

`func (o *StorageClass) GetAllowVolumeExpansion() bool`

GetAllowVolumeExpansion returns the AllowVolumeExpansion field if non-nil, zero value otherwise.

### GetAllowVolumeExpansionOk

`func (o *StorageClass) GetAllowVolumeExpansionOk() (*bool, bool)`

GetAllowVolumeExpansionOk returns a tuple with the AllowVolumeExpansion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowVolumeExpansion

`func (o *StorageClass) SetAllowVolumeExpansion(v bool)`

SetAllowVolumeExpansion sets AllowVolumeExpansion field to given value.


### GetMountOptions

`func (o *StorageClass) GetMountOptions() []string`

GetMountOptions returns the MountOptions field if non-nil, zero value otherwise.

### GetMountOptionsOk

`func (o *StorageClass) GetMountOptionsOk() (*[]string, bool)`

GetMountOptionsOk returns a tuple with the MountOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountOptions

`func (o *StorageClass) SetMountOptions(v []string)`

SetMountOptions sets MountOptions field to given value.


### GetVolumeBindingMode

`func (o *StorageClass) GetVolumeBindingMode() string`

GetVolumeBindingMode returns the VolumeBindingMode field if non-nil, zero value otherwise.

### GetVolumeBindingModeOk

`func (o *StorageClass) GetVolumeBindingModeOk() (*string, bool)`

GetVolumeBindingModeOk returns a tuple with the VolumeBindingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeBindingMode

`func (o *StorageClass) SetVolumeBindingMode(v string)`

SetVolumeBindingMode sets VolumeBindingMode field to given value.

### HasVolumeBindingMode

`func (o *StorageClass) HasVolumeBindingMode() bool`

HasVolumeBindingMode returns a boolean if a field has been set.

### SetVolumeBindingModeNil

`func (o *StorageClass) SetVolumeBindingModeNil(b bool)`

 SetVolumeBindingModeNil sets the value for VolumeBindingMode to be an explicit nil

### UnsetVolumeBindingMode
`func (o *StorageClass) UnsetVolumeBindingMode()`

UnsetVolumeBindingMode ensures that no value is present for VolumeBindingMode, not even an explicit nil
### GetIsDefault

`func (o *StorageClass) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *StorageClass) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *StorageClass) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetSelectable

`func (o *StorageClass) GetSelectable() bool`

GetSelectable returns the Selectable field if non-nil, zero value otherwise.

### GetSelectableOk

`func (o *StorageClass) GetSelectableOk() (*bool, bool)`

GetSelectableOk returns a tuple with the Selectable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectable

`func (o *StorageClass) SetSelectable(v bool)`

SetSelectable sets Selectable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


