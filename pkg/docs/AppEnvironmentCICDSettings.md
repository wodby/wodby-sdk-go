# AppEnvironmentCICDSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppEnvironmentId** | **int32** |  | 
**CiIntegrationId** | **int32** | Effective CI integration ID. Zero selects the built-in Wodby CI service. | 
**RegistryIntegrationId** | **int32** | Effective registry integration ID. Zero selects the built-in Wodby registry service. | 
**RegistryRepository** | **string** |  | 

## Methods

### NewAppEnvironmentCICDSettings

`func NewAppEnvironmentCICDSettings(appEnvironmentId int32, ciIntegrationId int32, registryIntegrationId int32, registryRepository string, ) *AppEnvironmentCICDSettings`

NewAppEnvironmentCICDSettings instantiates a new AppEnvironmentCICDSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppEnvironmentCICDSettingsWithDefaults

`func NewAppEnvironmentCICDSettingsWithDefaults() *AppEnvironmentCICDSettings`

NewAppEnvironmentCICDSettingsWithDefaults instantiates a new AppEnvironmentCICDSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppEnvironmentId

`func (o *AppEnvironmentCICDSettings) GetAppEnvironmentId() int32`

GetAppEnvironmentId returns the AppEnvironmentId field if non-nil, zero value otherwise.

### GetAppEnvironmentIdOk

`func (o *AppEnvironmentCICDSettings) GetAppEnvironmentIdOk() (*int32, bool)`

GetAppEnvironmentIdOk returns a tuple with the AppEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEnvironmentId

`func (o *AppEnvironmentCICDSettings) SetAppEnvironmentId(v int32)`

SetAppEnvironmentId sets AppEnvironmentId field to given value.


### GetCiIntegrationId

`func (o *AppEnvironmentCICDSettings) GetCiIntegrationId() int32`

GetCiIntegrationId returns the CiIntegrationId field if non-nil, zero value otherwise.

### GetCiIntegrationIdOk

`func (o *AppEnvironmentCICDSettings) GetCiIntegrationIdOk() (*int32, bool)`

GetCiIntegrationIdOk returns a tuple with the CiIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiIntegrationId

`func (o *AppEnvironmentCICDSettings) SetCiIntegrationId(v int32)`

SetCiIntegrationId sets CiIntegrationId field to given value.


### GetRegistryIntegrationId

`func (o *AppEnvironmentCICDSettings) GetRegistryIntegrationId() int32`

GetRegistryIntegrationId returns the RegistryIntegrationId field if non-nil, zero value otherwise.

### GetRegistryIntegrationIdOk

`func (o *AppEnvironmentCICDSettings) GetRegistryIntegrationIdOk() (*int32, bool)`

GetRegistryIntegrationIdOk returns a tuple with the RegistryIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryIntegrationId

`func (o *AppEnvironmentCICDSettings) SetRegistryIntegrationId(v int32)`

SetRegistryIntegrationId sets RegistryIntegrationId field to given value.


### GetRegistryRepository

`func (o *AppEnvironmentCICDSettings) GetRegistryRepository() string`

GetRegistryRepository returns the RegistryRepository field if non-nil, zero value otherwise.

### GetRegistryRepositoryOk

`func (o *AppEnvironmentCICDSettings) GetRegistryRepositoryOk() (*string, bool)`

GetRegistryRepositoryOk returns a tuple with the RegistryRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryRepository

`func (o *AppEnvironmentCICDSettings) SetRegistryRepository(v string)`

SetRegistryRepository sets RegistryRepository field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


