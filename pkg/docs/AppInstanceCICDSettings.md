# AppInstanceCICDSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppInstanceId** | **int32** |  | 
**CiIntegrationId** | **int32** | Effective CI integration ID. Zero selects the built-in Wodby CI service. | 
**RegistryIntegrationId** | **int32** | Effective registry integration ID. Zero selects the built-in Wodby registry service. | 
**RegistryRepository** | **string** |  | 

## Methods

### NewAppInstanceCICDSettings

`func NewAppInstanceCICDSettings(appInstanceId int32, ciIntegrationId int32, registryIntegrationId int32, registryRepository string, ) *AppInstanceCICDSettings`

NewAppInstanceCICDSettings instantiates a new AppInstanceCICDSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppInstanceCICDSettingsWithDefaults

`func NewAppInstanceCICDSettingsWithDefaults() *AppInstanceCICDSettings`

NewAppInstanceCICDSettingsWithDefaults instantiates a new AppInstanceCICDSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppInstanceId

`func (o *AppInstanceCICDSettings) GetAppInstanceId() int32`

GetAppInstanceId returns the AppInstanceId field if non-nil, zero value otherwise.

### GetAppInstanceIdOk

`func (o *AppInstanceCICDSettings) GetAppInstanceIdOk() (*int32, bool)`

GetAppInstanceIdOk returns a tuple with the AppInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceId

`func (o *AppInstanceCICDSettings) SetAppInstanceId(v int32)`

SetAppInstanceId sets AppInstanceId field to given value.


### GetCiIntegrationId

`func (o *AppInstanceCICDSettings) GetCiIntegrationId() int32`

GetCiIntegrationId returns the CiIntegrationId field if non-nil, zero value otherwise.

### GetCiIntegrationIdOk

`func (o *AppInstanceCICDSettings) GetCiIntegrationIdOk() (*int32, bool)`

GetCiIntegrationIdOk returns a tuple with the CiIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiIntegrationId

`func (o *AppInstanceCICDSettings) SetCiIntegrationId(v int32)`

SetCiIntegrationId sets CiIntegrationId field to given value.


### GetRegistryIntegrationId

`func (o *AppInstanceCICDSettings) GetRegistryIntegrationId() int32`

GetRegistryIntegrationId returns the RegistryIntegrationId field if non-nil, zero value otherwise.

### GetRegistryIntegrationIdOk

`func (o *AppInstanceCICDSettings) GetRegistryIntegrationIdOk() (*int32, bool)`

GetRegistryIntegrationIdOk returns a tuple with the RegistryIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryIntegrationId

`func (o *AppInstanceCICDSettings) SetRegistryIntegrationId(v int32)`

SetRegistryIntegrationId sets RegistryIntegrationId field to given value.


### GetRegistryRepository

`func (o *AppInstanceCICDSettings) GetRegistryRepository() string`

GetRegistryRepository returns the RegistryRepository field if non-nil, zero value otherwise.

### GetRegistryRepositoryOk

`func (o *AppInstanceCICDSettings) GetRegistryRepositoryOk() (*string, bool)`

GetRegistryRepositoryOk returns a tuple with the RegistryRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryRepository

`func (o *AppInstanceCICDSettings) SetRegistryRepository(v string)`

SetRegistryRepository sets RegistryRepository field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


