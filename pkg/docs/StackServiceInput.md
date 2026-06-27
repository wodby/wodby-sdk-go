# StackServiceInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Replicas** | Pointer to **NullableInt32** |  | [optional] 
**Required** | Pointer to **NullableBool** |  | [optional] 
**Disabled** | Pointer to **NullableBool** |  | [optional] 
**Main** | Pointer to **NullableBool** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**BuildSource** | Pointer to [**BuildSourceInput**](BuildSourceInput.md) |  | [optional] 

## Methods

### NewStackServiceInput

`func NewStackServiceInput() *StackServiceInput`

NewStackServiceInput instantiates a new StackServiceInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackServiceInputWithDefaults

`func NewStackServiceInputWithDefaults() *StackServiceInput`

NewStackServiceInputWithDefaults instantiates a new StackServiceInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReplicas

`func (o *StackServiceInput) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *StackServiceInput) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *StackServiceInput) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *StackServiceInput) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### SetReplicasNil

`func (o *StackServiceInput) SetReplicasNil(b bool)`

 SetReplicasNil sets the value for Replicas to be an explicit nil

### UnsetReplicas
`func (o *StackServiceInput) UnsetReplicas()`

UnsetReplicas ensures that no value is present for Replicas, not even an explicit nil
### GetRequired

`func (o *StackServiceInput) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *StackServiceInput) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *StackServiceInput) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *StackServiceInput) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### SetRequiredNil

`func (o *StackServiceInput) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *StackServiceInput) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil
### GetDisabled

`func (o *StackServiceInput) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *StackServiceInput) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *StackServiceInput) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *StackServiceInput) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### SetDisabledNil

`func (o *StackServiceInput) SetDisabledNil(b bool)`

 SetDisabledNil sets the value for Disabled to be an explicit nil

### UnsetDisabled
`func (o *StackServiceInput) UnsetDisabled()`

UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
### GetMain

`func (o *StackServiceInput) GetMain() bool`

GetMain returns the Main field if non-nil, zero value otherwise.

### GetMainOk

`func (o *StackServiceInput) GetMainOk() (*bool, bool)`

GetMainOk returns a tuple with the Main field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMain

`func (o *StackServiceInput) SetMain(v bool)`

SetMain sets Main field to given value.

### HasMain

`func (o *StackServiceInput) HasMain() bool`

HasMain returns a boolean if a field has been set.

### SetMainNil

`func (o *StackServiceInput) SetMainNil(b bool)`

 SetMainNil sets the value for Main to be an explicit nil

### UnsetMain
`func (o *StackServiceInput) UnsetMain()`

UnsetMain ensures that no value is present for Main, not even an explicit nil
### GetTitle

`func (o *StackServiceInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StackServiceInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StackServiceInput) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *StackServiceInput) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *StackServiceInput) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *StackServiceInput) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetBuildSource

`func (o *StackServiceInput) GetBuildSource() BuildSourceInput`

GetBuildSource returns the BuildSource field if non-nil, zero value otherwise.

### GetBuildSourceOk

`func (o *StackServiceInput) GetBuildSourceOk() (*BuildSourceInput, bool)`

GetBuildSourceOk returns a tuple with the BuildSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildSource

`func (o *StackServiceInput) SetBuildSource(v BuildSourceInput)`

SetBuildSource sets BuildSource field to given value.

### HasBuildSource

`func (o *StackServiceInput) HasBuildSource() bool`

HasBuildSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


