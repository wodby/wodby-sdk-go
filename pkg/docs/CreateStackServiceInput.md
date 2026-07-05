# CreateStackServiceInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StackId** | **int32** |  | 
**ServiceId** | **int32** |  | 
**Name** | **string** |  | 
**Title** | **string** |  | 
**Required** | **bool** |  | 
**Replicas** | **int32** |  | 
**ServiceRevPinned** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewCreateStackServiceInput

`func NewCreateStackServiceInput(stackId int32, serviceId int32, name string, title string, required bool, replicas int32, ) *CreateStackServiceInput`

NewCreateStackServiceInput instantiates a new CreateStackServiceInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStackServiceInputWithDefaults

`func NewCreateStackServiceInputWithDefaults() *CreateStackServiceInput`

NewCreateStackServiceInputWithDefaults instantiates a new CreateStackServiceInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStackId

`func (o *CreateStackServiceInput) GetStackId() int32`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *CreateStackServiceInput) GetStackIdOk() (*int32, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *CreateStackServiceInput) SetStackId(v int32)`

SetStackId sets StackId field to given value.


### GetServiceId

`func (o *CreateStackServiceInput) GetServiceId() int32`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CreateStackServiceInput) GetServiceIdOk() (*int32, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CreateStackServiceInput) SetServiceId(v int32)`

SetServiceId sets ServiceId field to given value.


### GetName

`func (o *CreateStackServiceInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateStackServiceInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateStackServiceInput) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *CreateStackServiceInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateStackServiceInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateStackServiceInput) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetRequired

`func (o *CreateStackServiceInput) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *CreateStackServiceInput) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *CreateStackServiceInput) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetReplicas

`func (o *CreateStackServiceInput) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *CreateStackServiceInput) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *CreateStackServiceInput) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.


### GetServiceRevPinned

`func (o *CreateStackServiceInput) GetServiceRevPinned() bool`

GetServiceRevPinned returns the ServiceRevPinned field if non-nil, zero value otherwise.

### GetServiceRevPinnedOk

`func (o *CreateStackServiceInput) GetServiceRevPinnedOk() (*bool, bool)`

GetServiceRevPinnedOk returns a tuple with the ServiceRevPinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRevPinned

`func (o *CreateStackServiceInput) SetServiceRevPinned(v bool)`

SetServiceRevPinned sets ServiceRevPinned field to given value.

### HasServiceRevPinned

`func (o *CreateStackServiceInput) HasServiceRevPinned() bool`

HasServiceRevPinned returns a boolean if a field has been set.

### SetServiceRevPinnedNil

`func (o *CreateStackServiceInput) SetServiceRevPinnedNil(b bool)`

 SetServiceRevPinnedNil sets the value for ServiceRevPinned to be an explicit nil

### UnsetServiceRevPinned
`func (o *CreateStackServiceInput) UnsetServiceRevPinned()`

UnsetServiceRevPinned ensures that no value is present for ServiceRevPinned, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


