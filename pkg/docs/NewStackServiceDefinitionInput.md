# NewStackServiceDefinitionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **int32** |  | 
**Name** | **string** | Stack-local service name. | 
**Title** | **string** |  | 
**Required** | Pointer to **NullableBool** |  | [optional] 
**Replicas** | **int32** |  | 
**Main** | Pointer to **NullableBool** |  | [optional] 
**ServiceRevPinned** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewNewStackServiceDefinitionInput

`func NewNewStackServiceDefinitionInput(serviceId int32, name string, title string, replicas int32, ) *NewStackServiceDefinitionInput`

NewNewStackServiceDefinitionInput instantiates a new NewStackServiceDefinitionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewStackServiceDefinitionInputWithDefaults

`func NewNewStackServiceDefinitionInputWithDefaults() *NewStackServiceDefinitionInput`

NewNewStackServiceDefinitionInputWithDefaults instantiates a new NewStackServiceDefinitionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *NewStackServiceDefinitionInput) GetServiceId() int32`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *NewStackServiceDefinitionInput) GetServiceIdOk() (*int32, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *NewStackServiceDefinitionInput) SetServiceId(v int32)`

SetServiceId sets ServiceId field to given value.


### GetName

`func (o *NewStackServiceDefinitionInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewStackServiceDefinitionInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewStackServiceDefinitionInput) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *NewStackServiceDefinitionInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewStackServiceDefinitionInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NewStackServiceDefinitionInput) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetRequired

`func (o *NewStackServiceDefinitionInput) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *NewStackServiceDefinitionInput) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *NewStackServiceDefinitionInput) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *NewStackServiceDefinitionInput) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### SetRequiredNil

`func (o *NewStackServiceDefinitionInput) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *NewStackServiceDefinitionInput) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil
### GetReplicas

`func (o *NewStackServiceDefinitionInput) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *NewStackServiceDefinitionInput) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *NewStackServiceDefinitionInput) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.


### GetMain

`func (o *NewStackServiceDefinitionInput) GetMain() bool`

GetMain returns the Main field if non-nil, zero value otherwise.

### GetMainOk

`func (o *NewStackServiceDefinitionInput) GetMainOk() (*bool, bool)`

GetMainOk returns a tuple with the Main field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMain

`func (o *NewStackServiceDefinitionInput) SetMain(v bool)`

SetMain sets Main field to given value.

### HasMain

`func (o *NewStackServiceDefinitionInput) HasMain() bool`

HasMain returns a boolean if a field has been set.

### SetMainNil

`func (o *NewStackServiceDefinitionInput) SetMainNil(b bool)`

 SetMainNil sets the value for Main to be an explicit nil

### UnsetMain
`func (o *NewStackServiceDefinitionInput) UnsetMain()`

UnsetMain ensures that no value is present for Main, not even an explicit nil
### GetServiceRevPinned

`func (o *NewStackServiceDefinitionInput) GetServiceRevPinned() bool`

GetServiceRevPinned returns the ServiceRevPinned field if non-nil, zero value otherwise.

### GetServiceRevPinnedOk

`func (o *NewStackServiceDefinitionInput) GetServiceRevPinnedOk() (*bool, bool)`

GetServiceRevPinnedOk returns a tuple with the ServiceRevPinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRevPinned

`func (o *NewStackServiceDefinitionInput) SetServiceRevPinned(v bool)`

SetServiceRevPinned sets ServiceRevPinned field to given value.

### HasServiceRevPinned

`func (o *NewStackServiceDefinitionInput) HasServiceRevPinned() bool`

HasServiceRevPinned returns a boolean if a field has been set.

### SetServiceRevPinnedNil

`func (o *NewStackServiceDefinitionInput) SetServiceRevPinnedNil(b bool)`

 SetServiceRevPinnedNil sets the value for ServiceRevPinned to be an explicit nil

### UnsetServiceRevPinned
`func (o *NewStackServiceDefinitionInput) UnsetServiceRevPinned()`

UnsetServiceRevPinned ensures that no value is present for ServiceRevPinned, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


