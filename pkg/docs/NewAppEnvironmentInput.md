# NewAppEnvironmentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | **int32** |  | 
**EnvironmentName** | **string** |  | 
**EnvironmentTitle** | Pointer to **string** | Defaults to environmentName when omitted. | [optional] 
**EnvironmentType** | **string** |  | 
**Domain** | Pointer to **string** | Defaults to environmentName.appName.orgDomain when omitted. | [optional] 
**StackRevId** | **int32** |  | 
**Services** | Pointer to [**[]CreateAppServiceInput**](CreateAppServiceInput.md) | Defaults to the stack revision&#39;s service defaults when omitted. | [optional] 
**ClusterId** | Pointer to **NullableInt32** |  | [optional] 
**CiIntegrationId** | Pointer to **NullableInt32** |  | [optional] 
**RegistryIntegrationId** | Pointer to **NullableInt32** |  | [optional] 
**DeferInitialDeployment** | Pointer to **bool** |  | [optional] [default to false]
**Settings** | Pointer to [**AppEnvironmentSettingsInput**](AppEnvironmentSettingsInput.md) |  | [optional] 
**Access** | Pointer to [**NewAppEnvironmentAccessInput**](NewAppEnvironmentAccessInput.md) |  | [optional] 

## Methods

### NewNewAppEnvironmentInput

`func NewNewAppEnvironmentInput(appId int32, environmentName string, environmentType string, stackRevId int32, ) *NewAppEnvironmentInput`

NewNewAppEnvironmentInput instantiates a new NewAppEnvironmentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewAppEnvironmentInputWithDefaults

`func NewNewAppEnvironmentInputWithDefaults() *NewAppEnvironmentInput`

NewNewAppEnvironmentInputWithDefaults instantiates a new NewAppEnvironmentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *NewAppEnvironmentInput) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *NewAppEnvironmentInput) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *NewAppEnvironmentInput) SetAppId(v int32)`

SetAppId sets AppId field to given value.


### GetEnvironmentName

`func (o *NewAppEnvironmentInput) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *NewAppEnvironmentInput) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *NewAppEnvironmentInput) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.


### GetEnvironmentTitle

`func (o *NewAppEnvironmentInput) GetEnvironmentTitle() string`

GetEnvironmentTitle returns the EnvironmentTitle field if non-nil, zero value otherwise.

### GetEnvironmentTitleOk

`func (o *NewAppEnvironmentInput) GetEnvironmentTitleOk() (*string, bool)`

GetEnvironmentTitleOk returns a tuple with the EnvironmentTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentTitle

`func (o *NewAppEnvironmentInput) SetEnvironmentTitle(v string)`

SetEnvironmentTitle sets EnvironmentTitle field to given value.

### HasEnvironmentTitle

`func (o *NewAppEnvironmentInput) HasEnvironmentTitle() bool`

HasEnvironmentTitle returns a boolean if a field has been set.

### GetEnvironmentType

`func (o *NewAppEnvironmentInput) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *NewAppEnvironmentInput) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *NewAppEnvironmentInput) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.


### GetDomain

`func (o *NewAppEnvironmentInput) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *NewAppEnvironmentInput) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *NewAppEnvironmentInput) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *NewAppEnvironmentInput) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetStackRevId

`func (o *NewAppEnvironmentInput) GetStackRevId() int32`

GetStackRevId returns the StackRevId field if non-nil, zero value otherwise.

### GetStackRevIdOk

`func (o *NewAppEnvironmentInput) GetStackRevIdOk() (*int32, bool)`

GetStackRevIdOk returns a tuple with the StackRevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackRevId

`func (o *NewAppEnvironmentInput) SetStackRevId(v int32)`

SetStackRevId sets StackRevId field to given value.


### GetServices

`func (o *NewAppEnvironmentInput) GetServices() []CreateAppServiceInput`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *NewAppEnvironmentInput) GetServicesOk() (*[]CreateAppServiceInput, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *NewAppEnvironmentInput) SetServices(v []CreateAppServiceInput)`

SetServices sets Services field to given value.

### HasServices

`func (o *NewAppEnvironmentInput) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetClusterId

`func (o *NewAppEnvironmentInput) GetClusterId() int32`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *NewAppEnvironmentInput) GetClusterIdOk() (*int32, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *NewAppEnvironmentInput) SetClusterId(v int32)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *NewAppEnvironmentInput) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *NewAppEnvironmentInput) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *NewAppEnvironmentInput) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetCiIntegrationId

`func (o *NewAppEnvironmentInput) GetCiIntegrationId() int32`

GetCiIntegrationId returns the CiIntegrationId field if non-nil, zero value otherwise.

### GetCiIntegrationIdOk

`func (o *NewAppEnvironmentInput) GetCiIntegrationIdOk() (*int32, bool)`

GetCiIntegrationIdOk returns a tuple with the CiIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiIntegrationId

`func (o *NewAppEnvironmentInput) SetCiIntegrationId(v int32)`

SetCiIntegrationId sets CiIntegrationId field to given value.

### HasCiIntegrationId

`func (o *NewAppEnvironmentInput) HasCiIntegrationId() bool`

HasCiIntegrationId returns a boolean if a field has been set.

### SetCiIntegrationIdNil

`func (o *NewAppEnvironmentInput) SetCiIntegrationIdNil(b bool)`

 SetCiIntegrationIdNil sets the value for CiIntegrationId to be an explicit nil

### UnsetCiIntegrationId
`func (o *NewAppEnvironmentInput) UnsetCiIntegrationId()`

UnsetCiIntegrationId ensures that no value is present for CiIntegrationId, not even an explicit nil
### GetRegistryIntegrationId

`func (o *NewAppEnvironmentInput) GetRegistryIntegrationId() int32`

GetRegistryIntegrationId returns the RegistryIntegrationId field if non-nil, zero value otherwise.

### GetRegistryIntegrationIdOk

`func (o *NewAppEnvironmentInput) GetRegistryIntegrationIdOk() (*int32, bool)`

GetRegistryIntegrationIdOk returns a tuple with the RegistryIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryIntegrationId

`func (o *NewAppEnvironmentInput) SetRegistryIntegrationId(v int32)`

SetRegistryIntegrationId sets RegistryIntegrationId field to given value.

### HasRegistryIntegrationId

`func (o *NewAppEnvironmentInput) HasRegistryIntegrationId() bool`

HasRegistryIntegrationId returns a boolean if a field has been set.

### SetRegistryIntegrationIdNil

`func (o *NewAppEnvironmentInput) SetRegistryIntegrationIdNil(b bool)`

 SetRegistryIntegrationIdNil sets the value for RegistryIntegrationId to be an explicit nil

### UnsetRegistryIntegrationId
`func (o *NewAppEnvironmentInput) UnsetRegistryIntegrationId()`

UnsetRegistryIntegrationId ensures that no value is present for RegistryIntegrationId, not even an explicit nil
### GetDeferInitialDeployment

`func (o *NewAppEnvironmentInput) GetDeferInitialDeployment() bool`

GetDeferInitialDeployment returns the DeferInitialDeployment field if non-nil, zero value otherwise.

### GetDeferInitialDeploymentOk

`func (o *NewAppEnvironmentInput) GetDeferInitialDeploymentOk() (*bool, bool)`

GetDeferInitialDeploymentOk returns a tuple with the DeferInitialDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferInitialDeployment

`func (o *NewAppEnvironmentInput) SetDeferInitialDeployment(v bool)`

SetDeferInitialDeployment sets DeferInitialDeployment field to given value.

### HasDeferInitialDeployment

`func (o *NewAppEnvironmentInput) HasDeferInitialDeployment() bool`

HasDeferInitialDeployment returns a boolean if a field has been set.

### GetSettings

`func (o *NewAppEnvironmentInput) GetSettings() AppEnvironmentSettingsInput`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *NewAppEnvironmentInput) GetSettingsOk() (*AppEnvironmentSettingsInput, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *NewAppEnvironmentInput) SetSettings(v AppEnvironmentSettingsInput)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *NewAppEnvironmentInput) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetAccess

`func (o *NewAppEnvironmentInput) GetAccess() NewAppEnvironmentAccessInput`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *NewAppEnvironmentInput) GetAccessOk() (*NewAppEnvironmentAccessInput, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *NewAppEnvironmentInput) SetAccess(v NewAppEnvironmentAccessInput)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *NewAppEnvironmentInput) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


