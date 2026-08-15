# AppAccessEndpointInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppPortId** | **int32** |  | 
**Host** | Pointer to **NullableString** |  | [optional] 
**Primary** | **bool** |  | 

## Methods

### NewAppAccessEndpointInput

`func NewAppAccessEndpointInput(appPortId int32, primary bool, ) *AppAccessEndpointInput`

NewAppAccessEndpointInput instantiates a new AppAccessEndpointInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppAccessEndpointInputWithDefaults

`func NewAppAccessEndpointInputWithDefaults() *AppAccessEndpointInput`

NewAppAccessEndpointInputWithDefaults instantiates a new AppAccessEndpointInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppPortId

`func (o *AppAccessEndpointInput) GetAppPortId() int32`

GetAppPortId returns the AppPortId field if non-nil, zero value otherwise.

### GetAppPortIdOk

`func (o *AppAccessEndpointInput) GetAppPortIdOk() (*int32, bool)`

GetAppPortIdOk returns a tuple with the AppPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPortId

`func (o *AppAccessEndpointInput) SetAppPortId(v int32)`

SetAppPortId sets AppPortId field to given value.


### GetHost

`func (o *AppAccessEndpointInput) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *AppAccessEndpointInput) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *AppAccessEndpointInput) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *AppAccessEndpointInput) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *AppAccessEndpointInput) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *AppAccessEndpointInput) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPrimary

`func (o *AppAccessEndpointInput) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *AppAccessEndpointInput) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *AppAccessEndpointInput) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


