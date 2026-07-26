# AppServiceVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**AppServiceId** | **int32** |  | 
**Name** | **string** |  | 
**Path** | **string** |  | 
**Shared** | **bool** |  | 
**ReadOnly** | **bool** |  | 
**Size** | **int32** |  | 
**ConfiguredStorageClassName** | Pointer to **NullableString** |  | [optional] 
**EffectiveStorageClassNames** | **[]string** |  | 
**StorageClassStatus** | **string** |  | 
**StorageClassSelectable** | **bool** |  | 
**FromVolumeId** | Pointer to **NullableInt32** |  | [optional] 
**StorageAppServiceId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewAppServiceVolume

`func NewAppServiceVolume(id int32, appServiceId int32, name string, path string, shared bool, readOnly bool, size int32, effectiveStorageClassNames []string, storageClassStatus string, storageClassSelectable bool, ) *AppServiceVolume`

NewAppServiceVolume instantiates a new AppServiceVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceVolumeWithDefaults

`func NewAppServiceVolumeWithDefaults() *AppServiceVolume`

NewAppServiceVolumeWithDefaults instantiates a new AppServiceVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppServiceVolume) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppServiceVolume) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppServiceVolume) SetId(v int32)`

SetId sets Id field to given value.


### GetAppServiceId

`func (o *AppServiceVolume) GetAppServiceId() int32`

GetAppServiceId returns the AppServiceId field if non-nil, zero value otherwise.

### GetAppServiceIdOk

`func (o *AppServiceVolume) GetAppServiceIdOk() (*int32, bool)`

GetAppServiceIdOk returns a tuple with the AppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceId

`func (o *AppServiceVolume) SetAppServiceId(v int32)`

SetAppServiceId sets AppServiceId field to given value.


### GetName

`func (o *AppServiceVolume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppServiceVolume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppServiceVolume) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *AppServiceVolume) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AppServiceVolume) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AppServiceVolume) SetPath(v string)`

SetPath sets Path field to given value.


### GetShared

`func (o *AppServiceVolume) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *AppServiceVolume) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *AppServiceVolume) SetShared(v bool)`

SetShared sets Shared field to given value.


### GetReadOnly

`func (o *AppServiceVolume) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *AppServiceVolume) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *AppServiceVolume) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.


### GetSize

`func (o *AppServiceVolume) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AppServiceVolume) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AppServiceVolume) SetSize(v int32)`

SetSize sets Size field to given value.


### GetConfiguredStorageClassName

`func (o *AppServiceVolume) GetConfiguredStorageClassName() string`

GetConfiguredStorageClassName returns the ConfiguredStorageClassName field if non-nil, zero value otherwise.

### GetConfiguredStorageClassNameOk

`func (o *AppServiceVolume) GetConfiguredStorageClassNameOk() (*string, bool)`

GetConfiguredStorageClassNameOk returns a tuple with the ConfiguredStorageClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredStorageClassName

`func (o *AppServiceVolume) SetConfiguredStorageClassName(v string)`

SetConfiguredStorageClassName sets ConfiguredStorageClassName field to given value.

### HasConfiguredStorageClassName

`func (o *AppServiceVolume) HasConfiguredStorageClassName() bool`

HasConfiguredStorageClassName returns a boolean if a field has been set.

### SetConfiguredStorageClassNameNil

`func (o *AppServiceVolume) SetConfiguredStorageClassNameNil(b bool)`

 SetConfiguredStorageClassNameNil sets the value for ConfiguredStorageClassName to be an explicit nil

### UnsetConfiguredStorageClassName
`func (o *AppServiceVolume) UnsetConfiguredStorageClassName()`

UnsetConfiguredStorageClassName ensures that no value is present for ConfiguredStorageClassName, not even an explicit nil
### GetEffectiveStorageClassNames

`func (o *AppServiceVolume) GetEffectiveStorageClassNames() []string`

GetEffectiveStorageClassNames returns the EffectiveStorageClassNames field if non-nil, zero value otherwise.

### GetEffectiveStorageClassNamesOk

`func (o *AppServiceVolume) GetEffectiveStorageClassNamesOk() (*[]string, bool)`

GetEffectiveStorageClassNamesOk returns a tuple with the EffectiveStorageClassNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveStorageClassNames

`func (o *AppServiceVolume) SetEffectiveStorageClassNames(v []string)`

SetEffectiveStorageClassNames sets EffectiveStorageClassNames field to given value.


### GetStorageClassStatus

`func (o *AppServiceVolume) GetStorageClassStatus() string`

GetStorageClassStatus returns the StorageClassStatus field if non-nil, zero value otherwise.

### GetStorageClassStatusOk

`func (o *AppServiceVolume) GetStorageClassStatusOk() (*string, bool)`

GetStorageClassStatusOk returns a tuple with the StorageClassStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClassStatus

`func (o *AppServiceVolume) SetStorageClassStatus(v string)`

SetStorageClassStatus sets StorageClassStatus field to given value.


### GetStorageClassSelectable

`func (o *AppServiceVolume) GetStorageClassSelectable() bool`

GetStorageClassSelectable returns the StorageClassSelectable field if non-nil, zero value otherwise.

### GetStorageClassSelectableOk

`func (o *AppServiceVolume) GetStorageClassSelectableOk() (*bool, bool)`

GetStorageClassSelectableOk returns a tuple with the StorageClassSelectable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClassSelectable

`func (o *AppServiceVolume) SetStorageClassSelectable(v bool)`

SetStorageClassSelectable sets StorageClassSelectable field to given value.


### GetFromVolumeId

`func (o *AppServiceVolume) GetFromVolumeId() int32`

GetFromVolumeId returns the FromVolumeId field if non-nil, zero value otherwise.

### GetFromVolumeIdOk

`func (o *AppServiceVolume) GetFromVolumeIdOk() (*int32, bool)`

GetFromVolumeIdOk returns a tuple with the FromVolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromVolumeId

`func (o *AppServiceVolume) SetFromVolumeId(v int32)`

SetFromVolumeId sets FromVolumeId field to given value.

### HasFromVolumeId

`func (o *AppServiceVolume) HasFromVolumeId() bool`

HasFromVolumeId returns a boolean if a field has been set.

### SetFromVolumeIdNil

`func (o *AppServiceVolume) SetFromVolumeIdNil(b bool)`

 SetFromVolumeIdNil sets the value for FromVolumeId to be an explicit nil

### UnsetFromVolumeId
`func (o *AppServiceVolume) UnsetFromVolumeId()`

UnsetFromVolumeId ensures that no value is present for FromVolumeId, not even an explicit nil
### GetStorageAppServiceId

`func (o *AppServiceVolume) GetStorageAppServiceId() int32`

GetStorageAppServiceId returns the StorageAppServiceId field if non-nil, zero value otherwise.

### GetStorageAppServiceIdOk

`func (o *AppServiceVolume) GetStorageAppServiceIdOk() (*int32, bool)`

GetStorageAppServiceIdOk returns a tuple with the StorageAppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAppServiceId

`func (o *AppServiceVolume) SetStorageAppServiceId(v int32)`

SetStorageAppServiceId sets StorageAppServiceId field to given value.

### HasStorageAppServiceId

`func (o *AppServiceVolume) HasStorageAppServiceId() bool`

HasStorageAppServiceId returns a boolean if a field has been set.

### SetStorageAppServiceIdNil

`func (o *AppServiceVolume) SetStorageAppServiceIdNil(b bool)`

 SetStorageAppServiceIdNil sets the value for StorageAppServiceId to be an explicit nil

### UnsetStorageAppServiceId
`func (o *AppServiceVolume) UnsetStorageAppServiceId()`

UnsetStorageAppServiceId ensures that no value is present for StorageAppServiceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


