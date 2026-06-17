# NewAppInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgID** | **int32** |  | 
**Name** | **string** |  | 
**Title** | **string** |  | 
**InstanceName** | **string** |  | 
**InstanceTitle** | **string** |  | 
**Domain** | **string** |  | 
**ProjectID** | Pointer to **NullableInt32** |  | [optional] 
**StackRevID** | **int32** |  | 
**Services** | [**[]NewAppServiceInput**](NewAppServiceInput.md) |  | 
**ClusterID** | Pointer to **NullableInt32** |  | [optional] 
**NewCluster** | Pointer to [**NewManagedClusterInput**](NewManagedClusterInput.md) |  | [optional] 
**EnvID** | **int32** |  | 
**CiIntegrationID** | Pointer to **NullableInt32** |  | [optional] 
**RegistryIntegrationID** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewNewAppInput

`func NewNewAppInput(orgID int32, name string, title string, instanceName string, instanceTitle string, domain string, stackRevID int32, services []NewAppServiceInput, envID int32, ) *NewAppInput`

NewNewAppInput instantiates a new NewAppInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewAppInputWithDefaults

`func NewNewAppInputWithDefaults() *NewAppInput`

NewNewAppInputWithDefaults instantiates a new NewAppInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgID

`func (o *NewAppInput) GetOrgID() int32`

GetOrgID returns the OrgID field if non-nil, zero value otherwise.

### GetOrgIDOk

`func (o *NewAppInput) GetOrgIDOk() (*int32, bool)`

GetOrgIDOk returns a tuple with the OrgID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgID

`func (o *NewAppInput) SetOrgID(v int32)`

SetOrgID sets OrgID field to given value.


### GetName

`func (o *NewAppInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewAppInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewAppInput) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *NewAppInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewAppInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NewAppInput) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetInstanceName

`func (o *NewAppInput) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *NewAppInput) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *NewAppInput) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.


### GetInstanceTitle

`func (o *NewAppInput) GetInstanceTitle() string`

GetInstanceTitle returns the InstanceTitle field if non-nil, zero value otherwise.

### GetInstanceTitleOk

`func (o *NewAppInput) GetInstanceTitleOk() (*string, bool)`

GetInstanceTitleOk returns a tuple with the InstanceTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTitle

`func (o *NewAppInput) SetInstanceTitle(v string)`

SetInstanceTitle sets InstanceTitle field to given value.


### GetDomain

`func (o *NewAppInput) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *NewAppInput) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *NewAppInput) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetProjectID

`func (o *NewAppInput) GetProjectID() int32`

GetProjectID returns the ProjectID field if non-nil, zero value otherwise.

### GetProjectIDOk

`func (o *NewAppInput) GetProjectIDOk() (*int32, bool)`

GetProjectIDOk returns a tuple with the ProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectID

`func (o *NewAppInput) SetProjectID(v int32)`

SetProjectID sets ProjectID field to given value.

### HasProjectID

`func (o *NewAppInput) HasProjectID() bool`

HasProjectID returns a boolean if a field has been set.

### SetProjectIDNil

`func (o *NewAppInput) SetProjectIDNil(b bool)`

 SetProjectIDNil sets the value for ProjectID to be an explicit nil

### UnsetProjectID
`func (o *NewAppInput) UnsetProjectID()`

UnsetProjectID ensures that no value is present for ProjectID, not even an explicit nil
### GetStackRevID

`func (o *NewAppInput) GetStackRevID() int32`

GetStackRevID returns the StackRevID field if non-nil, zero value otherwise.

### GetStackRevIDOk

`func (o *NewAppInput) GetStackRevIDOk() (*int32, bool)`

GetStackRevIDOk returns a tuple with the StackRevID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackRevID

`func (o *NewAppInput) SetStackRevID(v int32)`

SetStackRevID sets StackRevID field to given value.


### GetServices

`func (o *NewAppInput) GetServices() []NewAppServiceInput`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *NewAppInput) GetServicesOk() (*[]NewAppServiceInput, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *NewAppInput) SetServices(v []NewAppServiceInput)`

SetServices sets Services field to given value.


### GetClusterID

`func (o *NewAppInput) GetClusterID() int32`

GetClusterID returns the ClusterID field if non-nil, zero value otherwise.

### GetClusterIDOk

`func (o *NewAppInput) GetClusterIDOk() (*int32, bool)`

GetClusterIDOk returns a tuple with the ClusterID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterID

`func (o *NewAppInput) SetClusterID(v int32)`

SetClusterID sets ClusterID field to given value.

### HasClusterID

`func (o *NewAppInput) HasClusterID() bool`

HasClusterID returns a boolean if a field has been set.

### SetClusterIDNil

`func (o *NewAppInput) SetClusterIDNil(b bool)`

 SetClusterIDNil sets the value for ClusterID to be an explicit nil

### UnsetClusterID
`func (o *NewAppInput) UnsetClusterID()`

UnsetClusterID ensures that no value is present for ClusterID, not even an explicit nil
### GetNewCluster

`func (o *NewAppInput) GetNewCluster() NewManagedClusterInput`

GetNewCluster returns the NewCluster field if non-nil, zero value otherwise.

### GetNewClusterOk

`func (o *NewAppInput) GetNewClusterOk() (*NewManagedClusterInput, bool)`

GetNewClusterOk returns a tuple with the NewCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewCluster

`func (o *NewAppInput) SetNewCluster(v NewManagedClusterInput)`

SetNewCluster sets NewCluster field to given value.

### HasNewCluster

`func (o *NewAppInput) HasNewCluster() bool`

HasNewCluster returns a boolean if a field has been set.

### GetEnvID

`func (o *NewAppInput) GetEnvID() int32`

GetEnvID returns the EnvID field if non-nil, zero value otherwise.

### GetEnvIDOk

`func (o *NewAppInput) GetEnvIDOk() (*int32, bool)`

GetEnvIDOk returns a tuple with the EnvID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvID

`func (o *NewAppInput) SetEnvID(v int32)`

SetEnvID sets EnvID field to given value.


### GetCiIntegrationID

`func (o *NewAppInput) GetCiIntegrationID() int32`

GetCiIntegrationID returns the CiIntegrationID field if non-nil, zero value otherwise.

### GetCiIntegrationIDOk

`func (o *NewAppInput) GetCiIntegrationIDOk() (*int32, bool)`

GetCiIntegrationIDOk returns a tuple with the CiIntegrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiIntegrationID

`func (o *NewAppInput) SetCiIntegrationID(v int32)`

SetCiIntegrationID sets CiIntegrationID field to given value.

### HasCiIntegrationID

`func (o *NewAppInput) HasCiIntegrationID() bool`

HasCiIntegrationID returns a boolean if a field has been set.

### SetCiIntegrationIDNil

`func (o *NewAppInput) SetCiIntegrationIDNil(b bool)`

 SetCiIntegrationIDNil sets the value for CiIntegrationID to be an explicit nil

### UnsetCiIntegrationID
`func (o *NewAppInput) UnsetCiIntegrationID()`

UnsetCiIntegrationID ensures that no value is present for CiIntegrationID, not even an explicit nil
### GetRegistryIntegrationID

`func (o *NewAppInput) GetRegistryIntegrationID() int32`

GetRegistryIntegrationID returns the RegistryIntegrationID field if non-nil, zero value otherwise.

### GetRegistryIntegrationIDOk

`func (o *NewAppInput) GetRegistryIntegrationIDOk() (*int32, bool)`

GetRegistryIntegrationIDOk returns a tuple with the RegistryIntegrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryIntegrationID

`func (o *NewAppInput) SetRegistryIntegrationID(v int32)`

SetRegistryIntegrationID sets RegistryIntegrationID field to given value.

### HasRegistryIntegrationID

`func (o *NewAppInput) HasRegistryIntegrationID() bool`

HasRegistryIntegrationID returns a boolean if a field has been set.

### SetRegistryIntegrationIDNil

`func (o *NewAppInput) SetRegistryIntegrationIDNil(b bool)`

 SetRegistryIntegrationIDNil sets the value for RegistryIntegrationID to be an explicit nil

### UnsetRegistryIntegrationID
`func (o *NewAppInput) UnsetRegistryIntegrationID()`

UnsetRegistryIntegrationID ensures that no value is present for RegistryIntegrationID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


