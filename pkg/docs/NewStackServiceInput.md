# NewStackServiceInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StackId** | **int32** |  | 
**ServiceId** | **int32** |  | 
**Name** | **string** |  | 
**Title** | **string** |  | 
**Required** | **bool** |  | 
**Replicas** | **int32** |  | 

## Methods

### NewNewStackServiceInput

`func NewNewStackServiceInput(stackId int32, serviceId int32, name string, title string, required bool, replicas int32, ) *NewStackServiceInput`

NewNewStackServiceInput instantiates a new NewStackServiceInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewStackServiceInputWithDefaults

`func NewNewStackServiceInputWithDefaults() *NewStackServiceInput`

NewNewStackServiceInputWithDefaults instantiates a new NewStackServiceInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStackId

`func (o *NewStackServiceInput) GetStackId() int32`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *NewStackServiceInput) GetStackIdOk() (*int32, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *NewStackServiceInput) SetStackId(v int32)`

SetStackId sets StackId field to given value.


### GetServiceId

`func (o *NewStackServiceInput) GetServiceId() int32`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *NewStackServiceInput) GetServiceIdOk() (*int32, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *NewStackServiceInput) SetServiceId(v int32)`

SetServiceId sets ServiceId field to given value.


### GetName

`func (o *NewStackServiceInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewStackServiceInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewStackServiceInput) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *NewStackServiceInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewStackServiceInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NewStackServiceInput) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetRequired

`func (o *NewStackServiceInput) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *NewStackServiceInput) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *NewStackServiceInput) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetReplicas

`func (o *NewStackServiceInput) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *NewStackServiceInput) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *NewStackServiceInput) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


