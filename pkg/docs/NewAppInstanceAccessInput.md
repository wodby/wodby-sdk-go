# NewAppInstanceAccessInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationId** | **int32** |  | 
**Mode** | **string** |  | 
**Scope** | **string** |  | 
**Settings** | Pointer to [**[]AppAccessSettingInput**](AppAccessSettingInput.md) |  | [optional] 
**Host** | Pointer to **NullableString** | Required only when the selected provider uses a customer-assigned hostname. | [optional] 
**Endpoints** | Pointer to [**[]NewAppInstanceAccessEndpointInput**](NewAppInstanceAccessEndpointInput.md) | HTTP endpoints selected during creation. Used only with SELECTED_ENDPOINTS scope; older clients may omit it to select the main endpoint. | [optional] 

## Methods

### NewNewAppInstanceAccessInput

`func NewNewAppInstanceAccessInput(integrationId int32, mode string, scope string, ) *NewAppInstanceAccessInput`

NewNewAppInstanceAccessInput instantiates a new NewAppInstanceAccessInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewAppInstanceAccessInputWithDefaults

`func NewNewAppInstanceAccessInputWithDefaults() *NewAppInstanceAccessInput`

NewNewAppInstanceAccessInputWithDefaults instantiates a new NewAppInstanceAccessInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrationId

`func (o *NewAppInstanceAccessInput) GetIntegrationId() int32`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *NewAppInstanceAccessInput) GetIntegrationIdOk() (*int32, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *NewAppInstanceAccessInput) SetIntegrationId(v int32)`

SetIntegrationId sets IntegrationId field to given value.


### GetMode

`func (o *NewAppInstanceAccessInput) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *NewAppInstanceAccessInput) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *NewAppInstanceAccessInput) SetMode(v string)`

SetMode sets Mode field to given value.


### GetScope

`func (o *NewAppInstanceAccessInput) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *NewAppInstanceAccessInput) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *NewAppInstanceAccessInput) SetScope(v string)`

SetScope sets Scope field to given value.


### GetSettings

`func (o *NewAppInstanceAccessInput) GetSettings() []AppAccessSettingInput`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *NewAppInstanceAccessInput) GetSettingsOk() (*[]AppAccessSettingInput, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *NewAppInstanceAccessInput) SetSettings(v []AppAccessSettingInput)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *NewAppInstanceAccessInput) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetHost

`func (o *NewAppInstanceAccessInput) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *NewAppInstanceAccessInput) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *NewAppInstanceAccessInput) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *NewAppInstanceAccessInput) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *NewAppInstanceAccessInput) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *NewAppInstanceAccessInput) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetEndpoints

`func (o *NewAppInstanceAccessInput) GetEndpoints() []NewAppInstanceAccessEndpointInput`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *NewAppInstanceAccessInput) GetEndpointsOk() (*[]NewAppInstanceAccessEndpointInput, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *NewAppInstanceAccessInput) SetEndpoints(v []NewAppInstanceAccessEndpointInput)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *NewAppInstanceAccessInput) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


