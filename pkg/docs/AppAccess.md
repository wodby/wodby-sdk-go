# AppAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Mode** | **string** |  | 
**Scope** | **string** |  | 
**Status** | **string** |  | 
**Settings** | [**[]AppAccessSetting**](AppAccessSetting.md) |  | 
**PublicRoutesSuppressed** | **bool** |  | 
**EffectiveUrl** | Pointer to **NullableString** |  | [optional] 
**LastError** | Pointer to **NullableString** |  | [optional] 
**IntegrationId** | **int32** |  | 
**Endpoints** | [**[]AppAccessEndpoint**](AppAccessEndpoint.md) |  | 
**Resources** | [**[]AppAccessResource**](AppAccessResource.md) |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewAppAccess

`func NewAppAccess(id int32, mode string, scope string, status string, settings []AppAccessSetting, publicRoutesSuppressed bool, integrationId int32, endpoints []AppAccessEndpoint, resources []AppAccessResource, createdAt time.Time, updatedAt time.Time, ) *AppAccess`

NewAppAccess instantiates a new AppAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppAccessWithDefaults

`func NewAppAccessWithDefaults() *AppAccess`

NewAppAccessWithDefaults instantiates a new AppAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppAccess) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppAccess) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppAccess) SetId(v int32)`

SetId sets Id field to given value.


### GetMode

`func (o *AppAccess) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *AppAccess) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *AppAccess) SetMode(v string)`

SetMode sets Mode field to given value.


### GetScope

`func (o *AppAccess) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AppAccess) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AppAccess) SetScope(v string)`

SetScope sets Scope field to given value.


### GetStatus

`func (o *AppAccess) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppAccess) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppAccess) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSettings

`func (o *AppAccess) GetSettings() []AppAccessSetting`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *AppAccess) GetSettingsOk() (*[]AppAccessSetting, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *AppAccess) SetSettings(v []AppAccessSetting)`

SetSettings sets Settings field to given value.


### GetPublicRoutesSuppressed

`func (o *AppAccess) GetPublicRoutesSuppressed() bool`

GetPublicRoutesSuppressed returns the PublicRoutesSuppressed field if non-nil, zero value otherwise.

### GetPublicRoutesSuppressedOk

`func (o *AppAccess) GetPublicRoutesSuppressedOk() (*bool, bool)`

GetPublicRoutesSuppressedOk returns a tuple with the PublicRoutesSuppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicRoutesSuppressed

`func (o *AppAccess) SetPublicRoutesSuppressed(v bool)`

SetPublicRoutesSuppressed sets PublicRoutesSuppressed field to given value.


### GetEffectiveUrl

`func (o *AppAccess) GetEffectiveUrl() string`

GetEffectiveUrl returns the EffectiveUrl field if non-nil, zero value otherwise.

### GetEffectiveUrlOk

`func (o *AppAccess) GetEffectiveUrlOk() (*string, bool)`

GetEffectiveUrlOk returns a tuple with the EffectiveUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveUrl

`func (o *AppAccess) SetEffectiveUrl(v string)`

SetEffectiveUrl sets EffectiveUrl field to given value.

### HasEffectiveUrl

`func (o *AppAccess) HasEffectiveUrl() bool`

HasEffectiveUrl returns a boolean if a field has been set.

### SetEffectiveUrlNil

`func (o *AppAccess) SetEffectiveUrlNil(b bool)`

 SetEffectiveUrlNil sets the value for EffectiveUrl to be an explicit nil

### UnsetEffectiveUrl
`func (o *AppAccess) UnsetEffectiveUrl()`

UnsetEffectiveUrl ensures that no value is present for EffectiveUrl, not even an explicit nil
### GetLastError

`func (o *AppAccess) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *AppAccess) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *AppAccess) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *AppAccess) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *AppAccess) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *AppAccess) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetIntegrationId

`func (o *AppAccess) GetIntegrationId() int32`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *AppAccess) GetIntegrationIdOk() (*int32, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *AppAccess) SetIntegrationId(v int32)`

SetIntegrationId sets IntegrationId field to given value.


### GetEndpoints

`func (o *AppAccess) GetEndpoints() []AppAccessEndpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *AppAccess) GetEndpointsOk() (*[]AppAccessEndpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *AppAccess) SetEndpoints(v []AppAccessEndpoint)`

SetEndpoints sets Endpoints field to given value.


### GetResources

`func (o *AppAccess) GetResources() []AppAccessResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AppAccess) GetResourcesOk() (*[]AppAccessResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AppAccess) SetResources(v []AppAccessResource)`

SetResources sets Resources field to given value.


### GetCreatedAt

`func (o *AppAccess) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppAccess) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppAccess) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AppAccess) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AppAccess) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AppAccess) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


