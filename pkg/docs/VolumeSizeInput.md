# VolumeSizeInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Size** | **int32** |  | 
**StorageClassName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewVolumeSizeInput

`func NewVolumeSizeInput(name string, size int32, ) *VolumeSizeInput`

NewVolumeSizeInput instantiates a new VolumeSizeInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeSizeInputWithDefaults

`func NewVolumeSizeInputWithDefaults() *VolumeSizeInput`

NewVolumeSizeInputWithDefaults instantiates a new VolumeSizeInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VolumeSizeInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumeSizeInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumeSizeInput) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *VolumeSizeInput) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VolumeSizeInput) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VolumeSizeInput) SetSize(v int32)`

SetSize sets Size field to given value.


### GetStorageClassName

`func (o *VolumeSizeInput) GetStorageClassName() string`

GetStorageClassName returns the StorageClassName field if non-nil, zero value otherwise.

### GetStorageClassNameOk

`func (o *VolumeSizeInput) GetStorageClassNameOk() (*string, bool)`

GetStorageClassNameOk returns a tuple with the StorageClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClassName

`func (o *VolumeSizeInput) SetStorageClassName(v string)`

SetStorageClassName sets StorageClassName field to given value.

### HasStorageClassName

`func (o *VolumeSizeInput) HasStorageClassName() bool`

HasStorageClassName returns a boolean if a field has been set.

### SetStorageClassNameNil

`func (o *VolumeSizeInput) SetStorageClassNameNil(b bool)`

 SetStorageClassNameNil sets the value for StorageClassName to be an explicit nil

### UnsetStorageClassName
`func (o *VolumeSizeInput) UnsetStorageClassName()`

UnsetStorageClassName ensures that no value is present for StorageClassName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


