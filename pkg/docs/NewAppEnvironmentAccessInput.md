# NewAppEnvironmentAccessInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationId** | **int32** |  | 
**Mode** | **string** |  | 
**Scope** | **string** |  | 
**Settings** | Pointer to [**[]AppAccessSettingInput**](AppAccessSettingInput.md) |  | [optional] 
**Host** | Pointer to **NullableString** | Required only when the selected provider uses a customer-assigned hostname. | [optional] 
**Endpoints** | Pointer to [**[]NewAppEnvironmentAccessEndpointInput**](NewAppEnvironmentAccessEndpointInput.md) | HTTP endpoints selected during creation. Used only with SELECTED_ENDPOINTS scope; older clients may omit it to select the main endpoint. | [optional] 

## Methods

### NewNewAppEnvironmentAccessInput

`func NewNewAppEnvironmentAccessInput(integrationId int32, mode string, scope string, ) *NewAppEnvironmentAccessInput`

NewNewAppEnvironmentAccessInput instantiates a new NewAppEnvironmentAccessInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewAppEnvironmentAccessInputWithDefaults

`func NewNewAppEnvironmentAccessInputWithDefaults() *NewAppEnvironmentAccessInput`

NewNewAppEnvironmentAccessInputWithDefaults instantiates a new NewAppEnvironmentAccessInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrationId

`func (o *NewAppEnvironmentAccessInput) GetIntegrationId() int32`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *NewAppEnvironmentAccessInput) GetIntegrationIdOk() (*int32, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *NewAppEnvironmentAccessInput) SetIntegrationId(v int32)`

SetIntegrationId sets IntegrationId field to given value.


### GetMode

`func (o *NewAppEnvironmentAccessInput) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *NewAppEnvironmentAccessInput) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *NewAppEnvironmentAccessInput) SetMode(v string)`

SetMode sets Mode field to given value.


### GetScope

`func (o *NewAppEnvironmentAccessInput) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *NewAppEnvironmentAccessInput) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *NewAppEnvironmentAccessInput) SetScope(v string)`

SetScope sets Scope field to given value.


### GetSettings

`func (o *NewAppEnvironmentAccessInput) GetSettings() []AppAccessSettingInput`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *NewAppEnvironmentAccessInput) GetSettingsOk() (*[]AppAccessSettingInput, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *NewAppEnvironmentAccessInput) SetSettings(v []AppAccessSettingInput)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *NewAppEnvironmentAccessInput) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetHost

`func (o *NewAppEnvironmentAccessInput) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *NewAppEnvironmentAccessInput) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *NewAppEnvironmentAccessInput) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *NewAppEnvironmentAccessInput) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *NewAppEnvironmentAccessInput) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *NewAppEnvironmentAccessInput) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetEndpoints

`func (o *NewAppEnvironmentAccessInput) GetEndpoints() []NewAppEnvironmentAccessEndpointInput`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *NewAppEnvironmentAccessInput) GetEndpointsOk() (*[]NewAppEnvironmentAccessEndpointInput, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *NewAppEnvironmentAccessInput) SetEndpoints(v []NewAppEnvironmentAccessEndpointInput)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *NewAppEnvironmentAccessInput) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


