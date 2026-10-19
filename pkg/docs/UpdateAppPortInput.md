# UpdateAppPortInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicPort** | Pointer to **NullableInt32** | Public port on the cluster&#39;s address. | [optional] 

## Methods

### NewUpdateAppPortInput

`func NewUpdateAppPortInput() *UpdateAppPortInput`

NewUpdateAppPortInput instantiates a new UpdateAppPortInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAppPortInputWithDefaults

`func NewUpdateAppPortInputWithDefaults() *UpdateAppPortInput`

NewUpdateAppPortInputWithDefaults instantiates a new UpdateAppPortInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicPort

`func (o *UpdateAppPortInput) GetPublicPort() int32`

GetPublicPort returns the PublicPort field if non-nil, zero value otherwise.

### GetPublicPortOk

`func (o *UpdateAppPortInput) GetPublicPortOk() (*int32, bool)`

GetPublicPortOk returns a tuple with the PublicPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicPort

`func (o *UpdateAppPortInput) SetPublicPort(v int32)`

SetPublicPort sets PublicPort field to given value.

### HasPublicPort

`func (o *UpdateAppPortInput) HasPublicPort() bool`

HasPublicPort returns a boolean if a field has been set.

### SetPublicPortNil

`func (o *UpdateAppPortInput) SetPublicPortNil(b bool)`

 SetPublicPortNil sets the value for PublicPort to be an explicit nil

### UnsetPublicPort
`func (o *UpdateAppPortInput) UnsetPublicPort()`

UnsetPublicPort ensures that no value is present for PublicPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


