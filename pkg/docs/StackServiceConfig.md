# StackServiceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**StackServiceId** | **int32** |  | 
**Name** | **string** |  | 
**Config** | **string** |  | 
**Disabled** | **bool** |  | 

## Methods

### NewStackServiceConfig

`func NewStackServiceConfig(id int32, stackServiceId int32, name string, config string, disabled bool, ) *StackServiceConfig`

NewStackServiceConfig instantiates a new StackServiceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackServiceConfigWithDefaults

`func NewStackServiceConfigWithDefaults() *StackServiceConfig`

NewStackServiceConfigWithDefaults instantiates a new StackServiceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StackServiceConfig) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StackServiceConfig) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StackServiceConfig) SetId(v int32)`

SetId sets Id field to given value.


### GetStackServiceId

`func (o *StackServiceConfig) GetStackServiceId() int32`

GetStackServiceId returns the StackServiceId field if non-nil, zero value otherwise.

### GetStackServiceIdOk

`func (o *StackServiceConfig) GetStackServiceIdOk() (*int32, bool)`

GetStackServiceIdOk returns a tuple with the StackServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackServiceId

`func (o *StackServiceConfig) SetStackServiceId(v int32)`

SetStackServiceId sets StackServiceId field to given value.


### GetName

`func (o *StackServiceConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StackServiceConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StackServiceConfig) SetName(v string)`

SetName sets Name field to given value.


### GetConfig

`func (o *StackServiceConfig) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *StackServiceConfig) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *StackServiceConfig) SetConfig(v string)`

SetConfig sets Config field to given value.


### GetDisabled

`func (o *StackServiceConfig) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *StackServiceConfig) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *StackServiceConfig) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


