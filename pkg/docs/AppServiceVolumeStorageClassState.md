# AppServiceVolumeStorageClassState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VolumeId** | **int32** |  | 
**ConfiguredStorageClassName** | Pointer to **NullableString** |  | [optional] 
**EffectiveStorageClassNames** | **[]string** |  | 
**Status** | **string** |  | 

## Methods

### NewAppServiceVolumeStorageClassState

`func NewAppServiceVolumeStorageClassState(volumeId int32, effectiveStorageClassNames []string, status string, ) *AppServiceVolumeStorageClassState`

NewAppServiceVolumeStorageClassState instantiates a new AppServiceVolumeStorageClassState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceVolumeStorageClassStateWithDefaults

`func NewAppServiceVolumeStorageClassStateWithDefaults() *AppServiceVolumeStorageClassState`

NewAppServiceVolumeStorageClassStateWithDefaults instantiates a new AppServiceVolumeStorageClassState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumeId

`func (o *AppServiceVolumeStorageClassState) GetVolumeId() int32`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *AppServiceVolumeStorageClassState) GetVolumeIdOk() (*int32, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *AppServiceVolumeStorageClassState) SetVolumeId(v int32)`

SetVolumeId sets VolumeId field to given value.


### GetConfiguredStorageClassName

`func (o *AppServiceVolumeStorageClassState) GetConfiguredStorageClassName() string`

GetConfiguredStorageClassName returns the ConfiguredStorageClassName field if non-nil, zero value otherwise.

### GetConfiguredStorageClassNameOk

`func (o *AppServiceVolumeStorageClassState) GetConfiguredStorageClassNameOk() (*string, bool)`

GetConfiguredStorageClassNameOk returns a tuple with the ConfiguredStorageClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredStorageClassName

`func (o *AppServiceVolumeStorageClassState) SetConfiguredStorageClassName(v string)`

SetConfiguredStorageClassName sets ConfiguredStorageClassName field to given value.

### HasConfiguredStorageClassName

`func (o *AppServiceVolumeStorageClassState) HasConfiguredStorageClassName() bool`

HasConfiguredStorageClassName returns a boolean if a field has been set.

### SetConfiguredStorageClassNameNil

`func (o *AppServiceVolumeStorageClassState) SetConfiguredStorageClassNameNil(b bool)`

 SetConfiguredStorageClassNameNil sets the value for ConfiguredStorageClassName to be an explicit nil

### UnsetConfiguredStorageClassName
`func (o *AppServiceVolumeStorageClassState) UnsetConfiguredStorageClassName()`

UnsetConfiguredStorageClassName ensures that no value is present for ConfiguredStorageClassName, not even an explicit nil
### GetEffectiveStorageClassNames

`func (o *AppServiceVolumeStorageClassState) GetEffectiveStorageClassNames() []string`

GetEffectiveStorageClassNames returns the EffectiveStorageClassNames field if non-nil, zero value otherwise.

### GetEffectiveStorageClassNamesOk

`func (o *AppServiceVolumeStorageClassState) GetEffectiveStorageClassNamesOk() (*[]string, bool)`

GetEffectiveStorageClassNamesOk returns a tuple with the EffectiveStorageClassNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveStorageClassNames

`func (o *AppServiceVolumeStorageClassState) SetEffectiveStorageClassNames(v []string)`

SetEffectiveStorageClassNames sets EffectiveStorageClassNames field to given value.


### GetStatus

`func (o *AppServiceVolumeStorageClassState) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppServiceVolumeStorageClassState) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppServiceVolumeStorageClassState) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


