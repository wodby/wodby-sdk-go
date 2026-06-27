# NewAppInstanceInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | **int32** |  | 
**InstanceName** | **string** |  | 
**InstanceTitle** | **string** |  | 
**Domain** | **string** |  | 
**StackRevId** | **int32** |  | 
**Services** | [**[]CreateAppServiceInput**](CreateAppServiceInput.md) |  | 
**ClusterId** | Pointer to **NullableInt32** |  | [optional] 
**EnvId** | **int32** |  | 
**CiIntegrationId** | Pointer to **NullableInt32** |  | [optional] 
**RegistryIntegrationId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewNewAppInstanceInput

`func NewNewAppInstanceInput(appId int32, instanceName string, instanceTitle string, domain string, stackRevId int32, services []CreateAppServiceInput, envId int32, ) *NewAppInstanceInput`

NewNewAppInstanceInput instantiates a new NewAppInstanceInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewAppInstanceInputWithDefaults

`func NewNewAppInstanceInputWithDefaults() *NewAppInstanceInput`

NewNewAppInstanceInputWithDefaults instantiates a new NewAppInstanceInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *NewAppInstanceInput) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *NewAppInstanceInput) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *NewAppInstanceInput) SetAppId(v int32)`

SetAppId sets AppId field to given value.


### GetInstanceName

`func (o *NewAppInstanceInput) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *NewAppInstanceInput) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *NewAppInstanceInput) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.


### GetInstanceTitle

`func (o *NewAppInstanceInput) GetInstanceTitle() string`

GetInstanceTitle returns the InstanceTitle field if non-nil, zero value otherwise.

### GetInstanceTitleOk

`func (o *NewAppInstanceInput) GetInstanceTitleOk() (*string, bool)`

GetInstanceTitleOk returns a tuple with the InstanceTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTitle

`func (o *NewAppInstanceInput) SetInstanceTitle(v string)`

SetInstanceTitle sets InstanceTitle field to given value.


### GetDomain

`func (o *NewAppInstanceInput) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *NewAppInstanceInput) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *NewAppInstanceInput) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetStackRevId

`func (o *NewAppInstanceInput) GetStackRevId() int32`

GetStackRevId returns the StackRevId field if non-nil, zero value otherwise.

### GetStackRevIdOk

`func (o *NewAppInstanceInput) GetStackRevIdOk() (*int32, bool)`

GetStackRevIdOk returns a tuple with the StackRevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackRevId

`func (o *NewAppInstanceInput) SetStackRevId(v int32)`

SetStackRevId sets StackRevId field to given value.


### GetServices

`func (o *NewAppInstanceInput) GetServices() []CreateAppServiceInput`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *NewAppInstanceInput) GetServicesOk() (*[]CreateAppServiceInput, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *NewAppInstanceInput) SetServices(v []CreateAppServiceInput)`

SetServices sets Services field to given value.


### GetClusterId

`func (o *NewAppInstanceInput) GetClusterId() int32`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *NewAppInstanceInput) GetClusterIdOk() (*int32, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *NewAppInstanceInput) SetClusterId(v int32)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *NewAppInstanceInput) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *NewAppInstanceInput) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *NewAppInstanceInput) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetEnvId

`func (o *NewAppInstanceInput) GetEnvId() int32`

GetEnvId returns the EnvId field if non-nil, zero value otherwise.

### GetEnvIdOk

`func (o *NewAppInstanceInput) GetEnvIdOk() (*int32, bool)`

GetEnvIdOk returns a tuple with the EnvId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvId

`func (o *NewAppInstanceInput) SetEnvId(v int32)`

SetEnvId sets EnvId field to given value.


### GetCiIntegrationId

`func (o *NewAppInstanceInput) GetCiIntegrationId() int32`

GetCiIntegrationId returns the CiIntegrationId field if non-nil, zero value otherwise.

### GetCiIntegrationIdOk

`func (o *NewAppInstanceInput) GetCiIntegrationIdOk() (*int32, bool)`

GetCiIntegrationIdOk returns a tuple with the CiIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiIntegrationId

`func (o *NewAppInstanceInput) SetCiIntegrationId(v int32)`

SetCiIntegrationId sets CiIntegrationId field to given value.

### HasCiIntegrationId

`func (o *NewAppInstanceInput) HasCiIntegrationId() bool`

HasCiIntegrationId returns a boolean if a field has been set.

### SetCiIntegrationIdNil

`func (o *NewAppInstanceInput) SetCiIntegrationIdNil(b bool)`

 SetCiIntegrationIdNil sets the value for CiIntegrationId to be an explicit nil

### UnsetCiIntegrationId
`func (o *NewAppInstanceInput) UnsetCiIntegrationId()`

UnsetCiIntegrationId ensures that no value is present for CiIntegrationId, not even an explicit nil
### GetRegistryIntegrationId

`func (o *NewAppInstanceInput) GetRegistryIntegrationId() int32`

GetRegistryIntegrationId returns the RegistryIntegrationId field if non-nil, zero value otherwise.

### GetRegistryIntegrationIdOk

`func (o *NewAppInstanceInput) GetRegistryIntegrationIdOk() (*int32, bool)`

GetRegistryIntegrationIdOk returns a tuple with the RegistryIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryIntegrationId

`func (o *NewAppInstanceInput) SetRegistryIntegrationId(v int32)`

SetRegistryIntegrationId sets RegistryIntegrationId field to given value.

### HasRegistryIntegrationId

`func (o *NewAppInstanceInput) HasRegistryIntegrationId() bool`

HasRegistryIntegrationId returns a boolean if a field has been set.

### SetRegistryIntegrationIdNil

`func (o *NewAppInstanceInput) SetRegistryIntegrationIdNil(b bool)`

 SetRegistryIntegrationIdNil sets the value for RegistryIntegrationId to be an explicit nil

### UnsetRegistryIntegrationId
`func (o *NewAppInstanceInput) UnsetRegistryIntegrationId()`

UnsetRegistryIntegrationId ensures that no value is present for RegistryIntegrationId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


