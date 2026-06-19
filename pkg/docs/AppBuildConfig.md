# AppBuildConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegistryHost** | **string** |  | 
**RegistryRepository** | **string** |  | 
**Services** | [**[]AppServiceBuildConfig**](AppServiceBuildConfig.md) |  | 

## Methods

### NewAppBuildConfig

`func NewAppBuildConfig(registryHost string, registryRepository string, services []AppServiceBuildConfig, ) *AppBuildConfig`

NewAppBuildConfig instantiates a new AppBuildConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppBuildConfigWithDefaults

`func NewAppBuildConfigWithDefaults() *AppBuildConfig`

NewAppBuildConfigWithDefaults instantiates a new AppBuildConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistryHost

`func (o *AppBuildConfig) GetRegistryHost() string`

GetRegistryHost returns the RegistryHost field if non-nil, zero value otherwise.

### GetRegistryHostOk

`func (o *AppBuildConfig) GetRegistryHostOk() (*string, bool)`

GetRegistryHostOk returns a tuple with the RegistryHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryHost

`func (o *AppBuildConfig) SetRegistryHost(v string)`

SetRegistryHost sets RegistryHost field to given value.


### GetRegistryRepository

`func (o *AppBuildConfig) GetRegistryRepository() string`

GetRegistryRepository returns the RegistryRepository field if non-nil, zero value otherwise.

### GetRegistryRepositoryOk

`func (o *AppBuildConfig) GetRegistryRepositoryOk() (*string, bool)`

GetRegistryRepositoryOk returns a tuple with the RegistryRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryRepository

`func (o *AppBuildConfig) SetRegistryRepository(v string)`

SetRegistryRepository sets RegistryRepository field to given value.


### GetServices

`func (o *AppBuildConfig) GetServices() []AppServiceBuildConfig`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *AppBuildConfig) GetServicesOk() (*[]AppServiceBuildConfig, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *AppBuildConfig) SetServices(v []AppServiceBuildConfig)`

SetServices sets Services field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


