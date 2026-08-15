# UpdateAppAccessInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | **string** |  | 
**Settings** | Pointer to [**[]AppAccessSettingInput**](AppAccessSettingInput.md) |  | [optional] 
**Endpoints** | [**[]AppAccessEndpointInput**](AppAccessEndpointInput.md) |  | 

## Methods

### NewUpdateAppAccessInput

`func NewUpdateAppAccessInput(scope string, endpoints []AppAccessEndpointInput, ) *UpdateAppAccessInput`

NewUpdateAppAccessInput instantiates a new UpdateAppAccessInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAppAccessInputWithDefaults

`func NewUpdateAppAccessInputWithDefaults() *UpdateAppAccessInput`

NewUpdateAppAccessInputWithDefaults instantiates a new UpdateAppAccessInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *UpdateAppAccessInput) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *UpdateAppAccessInput) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *UpdateAppAccessInput) SetScope(v string)`

SetScope sets Scope field to given value.


### GetSettings

`func (o *UpdateAppAccessInput) GetSettings() []AppAccessSettingInput`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *UpdateAppAccessInput) GetSettingsOk() (*[]AppAccessSettingInput, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *UpdateAppAccessInput) SetSettings(v []AppAccessSettingInput)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *UpdateAppAccessInput) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetEndpoints

`func (o *UpdateAppAccessInput) GetEndpoints() []AppAccessEndpointInput`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *UpdateAppAccessInput) GetEndpointsOk() (*[]AppAccessEndpointInput, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *UpdateAppAccessInput) SetEndpoints(v []AppAccessEndpointInput)`

SetEndpoints sets Endpoints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


