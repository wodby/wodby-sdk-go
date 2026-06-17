# NewAppInstanceInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppID** | **int32** |  | 
**InstanceName** | **string** |  | 
**InstanceTitle** | **string** |  | 
**Domain** | **string** |  | 
**StackRevID** | **int32** |  | 
**Services** | [**[]NewAppServiceInput**](NewAppServiceInput.md) |  | 
**ClusterID** | Pointer to **NullableInt32** |  | [optional] 
**NewCluster** | Pointer to [**NewManagedClusterInput**](NewManagedClusterInput.md) |  | [optional] 
**EnvID** | **int32** |  | 
**CiIntegrationID** | Pointer to **NullableInt32** |  | [optional] 
**RegistryIntegrationID** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewNewAppInstanceInput

`func NewNewAppInstanceInput(appID int32, instanceName string, instanceTitle string, domain string, stackRevID int32, services []NewAppServiceInput, envID int32, ) *NewAppInstanceInput`

NewNewAppInstanceInput instantiates a new NewAppInstanceInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewAppInstanceInputWithDefaults

`func NewNewAppInstanceInputWithDefaults() *NewAppInstanceInput`

NewNewAppInstanceInputWithDefaults instantiates a new NewAppInstanceInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppID

`func (o *NewAppInstanceInput) GetAppID() int32`

GetAppID returns the AppID field if non-nil, zero value otherwise.

### GetAppIDOk

`func (o *NewAppInstanceInput) GetAppIDOk() (*int32, bool)`

GetAppIDOk returns a tuple with the AppID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppID

`func (o *NewAppInstanceInput) SetAppID(v int32)`

SetAppID sets AppID field to given value.


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


### GetStackRevID

`func (o *NewAppInstanceInput) GetStackRevID() int32`

GetStackRevID returns the StackRevID field if non-nil, zero value otherwise.

### GetStackRevIDOk

`func (o *NewAppInstanceInput) GetStackRevIDOk() (*int32, bool)`

GetStackRevIDOk returns a tuple with the StackRevID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackRevID

`func (o *NewAppInstanceInput) SetStackRevID(v int32)`

SetStackRevID sets StackRevID field to given value.


### GetServices

`func (o *NewAppInstanceInput) GetServices() []NewAppServiceInput`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *NewAppInstanceInput) GetServicesOk() (*[]NewAppServiceInput, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *NewAppInstanceInput) SetServices(v []NewAppServiceInput)`

SetServices sets Services field to given value.


### GetClusterID

`func (o *NewAppInstanceInput) GetClusterID() int32`

GetClusterID returns the ClusterID field if non-nil, zero value otherwise.

### GetClusterIDOk

`func (o *NewAppInstanceInput) GetClusterIDOk() (*int32, bool)`

GetClusterIDOk returns a tuple with the ClusterID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterID

`func (o *NewAppInstanceInput) SetClusterID(v int32)`

SetClusterID sets ClusterID field to given value.

### HasClusterID

`func (o *NewAppInstanceInput) HasClusterID() bool`

HasClusterID returns a boolean if a field has been set.

### SetClusterIDNil

`func (o *NewAppInstanceInput) SetClusterIDNil(b bool)`

 SetClusterIDNil sets the value for ClusterID to be an explicit nil

### UnsetClusterID
`func (o *NewAppInstanceInput) UnsetClusterID()`

UnsetClusterID ensures that no value is present for ClusterID, not even an explicit nil
### GetNewCluster

`func (o *NewAppInstanceInput) GetNewCluster() NewManagedClusterInput`

GetNewCluster returns the NewCluster field if non-nil, zero value otherwise.

### GetNewClusterOk

`func (o *NewAppInstanceInput) GetNewClusterOk() (*NewManagedClusterInput, bool)`

GetNewClusterOk returns a tuple with the NewCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewCluster

`func (o *NewAppInstanceInput) SetNewCluster(v NewManagedClusterInput)`

SetNewCluster sets NewCluster field to given value.

### HasNewCluster

`func (o *NewAppInstanceInput) HasNewCluster() bool`

HasNewCluster returns a boolean if a field has been set.

### GetEnvID

`func (o *NewAppInstanceInput) GetEnvID() int32`

GetEnvID returns the EnvID field if non-nil, zero value otherwise.

### GetEnvIDOk

`func (o *NewAppInstanceInput) GetEnvIDOk() (*int32, bool)`

GetEnvIDOk returns a tuple with the EnvID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvID

`func (o *NewAppInstanceInput) SetEnvID(v int32)`

SetEnvID sets EnvID field to given value.


### GetCiIntegrationID

`func (o *NewAppInstanceInput) GetCiIntegrationID() int32`

GetCiIntegrationID returns the CiIntegrationID field if non-nil, zero value otherwise.

### GetCiIntegrationIDOk

`func (o *NewAppInstanceInput) GetCiIntegrationIDOk() (*int32, bool)`

GetCiIntegrationIDOk returns a tuple with the CiIntegrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiIntegrationID

`func (o *NewAppInstanceInput) SetCiIntegrationID(v int32)`

SetCiIntegrationID sets CiIntegrationID field to given value.

### HasCiIntegrationID

`func (o *NewAppInstanceInput) HasCiIntegrationID() bool`

HasCiIntegrationID returns a boolean if a field has been set.

### SetCiIntegrationIDNil

`func (o *NewAppInstanceInput) SetCiIntegrationIDNil(b bool)`

 SetCiIntegrationIDNil sets the value for CiIntegrationID to be an explicit nil

### UnsetCiIntegrationID
`func (o *NewAppInstanceInput) UnsetCiIntegrationID()`

UnsetCiIntegrationID ensures that no value is present for CiIntegrationID, not even an explicit nil
### GetRegistryIntegrationID

`func (o *NewAppInstanceInput) GetRegistryIntegrationID() int32`

GetRegistryIntegrationID returns the RegistryIntegrationID field if non-nil, zero value otherwise.

### GetRegistryIntegrationIDOk

`func (o *NewAppInstanceInput) GetRegistryIntegrationIDOk() (*int32, bool)`

GetRegistryIntegrationIDOk returns a tuple with the RegistryIntegrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryIntegrationID

`func (o *NewAppInstanceInput) SetRegistryIntegrationID(v int32)`

SetRegistryIntegrationID sets RegistryIntegrationID field to given value.

### HasRegistryIntegrationID

`func (o *NewAppInstanceInput) HasRegistryIntegrationID() bool`

HasRegistryIntegrationID returns a boolean if a field has been set.

### SetRegistryIntegrationIDNil

`func (o *NewAppInstanceInput) SetRegistryIntegrationIDNil(b bool)`

 SetRegistryIntegrationIDNil sets the value for RegistryIntegrationID to be an explicit nil

### UnsetRegistryIntegrationID
`func (o *NewAppInstanceInput) UnsetRegistryIntegrationID()`

UnsetRegistryIntegrationID ensures that no value is present for RegistryIntegrationID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


