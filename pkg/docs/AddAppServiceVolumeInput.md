# AddAppServiceVolumeInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Size** | Pointer to **NullableInt32** |  | [optional] 
**StorageClassName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAddAppServiceVolumeInput

`func NewAddAppServiceVolumeInput(name string, ) *AddAppServiceVolumeInput`

NewAddAppServiceVolumeInput instantiates a new AddAppServiceVolumeInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAppServiceVolumeInputWithDefaults

`func NewAddAppServiceVolumeInputWithDefaults() *AddAppServiceVolumeInput`

NewAddAppServiceVolumeInputWithDefaults instantiates a new AddAppServiceVolumeInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddAppServiceVolumeInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddAppServiceVolumeInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddAppServiceVolumeInput) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *AddAppServiceVolumeInput) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AddAppServiceVolumeInput) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AddAppServiceVolumeInput) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *AddAppServiceVolumeInput) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *AddAppServiceVolumeInput) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *AddAppServiceVolumeInput) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetStorageClassName

`func (o *AddAppServiceVolumeInput) GetStorageClassName() string`

GetStorageClassName returns the StorageClassName field if non-nil, zero value otherwise.

### GetStorageClassNameOk

`func (o *AddAppServiceVolumeInput) GetStorageClassNameOk() (*string, bool)`

GetStorageClassNameOk returns a tuple with the StorageClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClassName

`func (o *AddAppServiceVolumeInput) SetStorageClassName(v string)`

SetStorageClassName sets StorageClassName field to given value.

### HasStorageClassName

`func (o *AddAppServiceVolumeInput) HasStorageClassName() bool`

HasStorageClassName returns a boolean if a field has been set.

### SetStorageClassNameNil

`func (o *AddAppServiceVolumeInput) SetStorageClassNameNil(b bool)`

 SetStorageClassNameNil sets the value for StorageClassName to be an explicit nil

### UnsetStorageClassName
`func (o *AddAppServiceVolumeInput) UnsetStorageClassName()`

UnsetStorageClassName ensures that no value is present for StorageClassName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


