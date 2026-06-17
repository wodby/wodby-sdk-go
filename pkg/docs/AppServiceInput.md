# AppServiceInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Replicas** | Pointer to **NullableInt32** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**Disabled** | Pointer to **NullableBool** |  | [optional] 
**Main** | Pointer to **NullableBool** |  | [optional] 
**BuildSource** | Pointer to [**BuildSourceInput**](BuildSourceInput.md) |  | [optional] 

## Methods

### NewAppServiceInput

`func NewAppServiceInput() *AppServiceInput`

NewAppServiceInput instantiates a new AppServiceInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceInputWithDefaults

`func NewAppServiceInputWithDefaults() *AppServiceInput`

NewAppServiceInputWithDefaults instantiates a new AppServiceInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReplicas

`func (o *AppServiceInput) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *AppServiceInput) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *AppServiceInput) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *AppServiceInput) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### SetReplicasNil

`func (o *AppServiceInput) SetReplicasNil(b bool)`

 SetReplicasNil sets the value for Replicas to be an explicit nil

### UnsetReplicas
`func (o *AppServiceInput) UnsetReplicas()`

UnsetReplicas ensures that no value is present for Replicas, not even an explicit nil
### GetVersion

`func (o *AppServiceInput) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AppServiceInput) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AppServiceInput) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AppServiceInput) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *AppServiceInput) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *AppServiceInput) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetDisabled

`func (o *AppServiceInput) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *AppServiceInput) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *AppServiceInput) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *AppServiceInput) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### SetDisabledNil

`func (o *AppServiceInput) SetDisabledNil(b bool)`

 SetDisabledNil sets the value for Disabled to be an explicit nil

### UnsetDisabled
`func (o *AppServiceInput) UnsetDisabled()`

UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
### GetMain

`func (o *AppServiceInput) GetMain() bool`

GetMain returns the Main field if non-nil, zero value otherwise.

### GetMainOk

`func (o *AppServiceInput) GetMainOk() (*bool, bool)`

GetMainOk returns a tuple with the Main field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMain

`func (o *AppServiceInput) SetMain(v bool)`

SetMain sets Main field to given value.

### HasMain

`func (o *AppServiceInput) HasMain() bool`

HasMain returns a boolean if a field has been set.

### SetMainNil

`func (o *AppServiceInput) SetMainNil(b bool)`

 SetMainNil sets the value for Main to be an explicit nil

### UnsetMain
`func (o *AppServiceInput) UnsetMain()`

UnsetMain ensures that no value is present for Main, not even an explicit nil
### GetBuildSource

`func (o *AppServiceInput) GetBuildSource() BuildSourceInput`

GetBuildSource returns the BuildSource field if non-nil, zero value otherwise.

### GetBuildSourceOk

`func (o *AppServiceInput) GetBuildSourceOk() (*BuildSourceInput, bool)`

GetBuildSourceOk returns a tuple with the BuildSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildSource

`func (o *AppServiceInput) SetBuildSource(v BuildSourceInput)`

SetBuildSource sets BuildSource field to given value.

### HasBuildSource

`func (o *AppServiceInput) HasBuildSource() bool`

HasBuildSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


