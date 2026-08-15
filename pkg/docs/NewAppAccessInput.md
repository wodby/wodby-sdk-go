# NewAppAccessInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationId** | **int32** |  | 
**Mode** | **string** |  | 
**Scope** | **string** |  | 
**Settings** | Pointer to [**[]AppAccessSettingInput**](AppAccessSettingInput.md) |  | [optional] 
**Endpoints** | [**[]AppAccessEndpointInput**](AppAccessEndpointInput.md) |  | 

## Methods

### NewNewAppAccessInput

`func NewNewAppAccessInput(integrationId int32, mode string, scope string, endpoints []AppAccessEndpointInput, ) *NewAppAccessInput`

NewNewAppAccessInput instantiates a new NewAppAccessInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewAppAccessInputWithDefaults

`func NewNewAppAccessInputWithDefaults() *NewAppAccessInput`

NewNewAppAccessInputWithDefaults instantiates a new NewAppAccessInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrationId

`func (o *NewAppAccessInput) GetIntegrationId() int32`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *NewAppAccessInput) GetIntegrationIdOk() (*int32, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *NewAppAccessInput) SetIntegrationId(v int32)`

SetIntegrationId sets IntegrationId field to given value.


### GetMode

`func (o *NewAppAccessInput) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *NewAppAccessInput) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *NewAppAccessInput) SetMode(v string)`

SetMode sets Mode field to given value.


### GetScope

`func (o *NewAppAccessInput) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *NewAppAccessInput) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *NewAppAccessInput) SetScope(v string)`

SetScope sets Scope field to given value.


### GetSettings

`func (o *NewAppAccessInput) GetSettings() []AppAccessSettingInput`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *NewAppAccessInput) GetSettingsOk() (*[]AppAccessSettingInput, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *NewAppAccessInput) SetSettings(v []AppAccessSettingInput)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *NewAppAccessInput) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetEndpoints

`func (o *NewAppAccessInput) GetEndpoints() []AppAccessEndpointInput`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *NewAppAccessInput) GetEndpointsOk() (*[]AppAccessEndpointInput, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *NewAppAccessInput) SetEndpoints(v []AppAccessEndpointInput)`

SetEndpoints sets Endpoints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


