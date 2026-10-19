# IgnoreStackServiceRevisionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StackServiceId** | **int32** | Root stack service whose update to skip. Its derivatives skip it too. | 
**ServiceRevId** | **int32** | Service revision to stop offering. | 

## Methods

### NewIgnoreStackServiceRevisionInput

`func NewIgnoreStackServiceRevisionInput(stackServiceId int32, serviceRevId int32, ) *IgnoreStackServiceRevisionInput`

NewIgnoreStackServiceRevisionInput instantiates a new IgnoreStackServiceRevisionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIgnoreStackServiceRevisionInputWithDefaults

`func NewIgnoreStackServiceRevisionInputWithDefaults() *IgnoreStackServiceRevisionInput`

NewIgnoreStackServiceRevisionInputWithDefaults instantiates a new IgnoreStackServiceRevisionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStackServiceId

`func (o *IgnoreStackServiceRevisionInput) GetStackServiceId() int32`

GetStackServiceId returns the StackServiceId field if non-nil, zero value otherwise.

### GetStackServiceIdOk

`func (o *IgnoreStackServiceRevisionInput) GetStackServiceIdOk() (*int32, bool)`

GetStackServiceIdOk returns a tuple with the StackServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackServiceId

`func (o *IgnoreStackServiceRevisionInput) SetStackServiceId(v int32)`

SetStackServiceId sets StackServiceId field to given value.


### GetServiceRevId

`func (o *IgnoreStackServiceRevisionInput) GetServiceRevId() int32`

GetServiceRevId returns the ServiceRevId field if non-nil, zero value otherwise.

### GetServiceRevIdOk

`func (o *IgnoreStackServiceRevisionInput) GetServiceRevIdOk() (*int32, bool)`

GetServiceRevIdOk returns a tuple with the ServiceRevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRevId

`func (o *IgnoreStackServiceRevisionInput) SetServiceRevId(v int32)`

SetServiceRevId sets ServiceRevId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


