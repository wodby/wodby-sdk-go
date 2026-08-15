# NewAppInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. | [optional] 
**Name** | **string** |  | 
**Title** | Pointer to **string** | Defaults to name when omitted. | [optional] 
**InstanceName** | **string** |  | 
**InstanceTitle** | Pointer to **string** | Defaults to instanceName when omitted. | [optional] 
**Domain** | Pointer to **string** | Defaults to instanceName.name.orgDomain when omitted. | [optional] 
**ProjectId** | Pointer to **NullableInt32** |  | [optional] 
**StackRevId** | **int32** |  | 
**Services** | Pointer to [**[]CreateAppServiceInput**](CreateAppServiceInput.md) | Defaults to the stack revision&#39;s service defaults when omitted. | [optional] 
**ClusterId** | Pointer to **NullableInt32** |  | [optional] 
**EnvId** | **int32** |  | 
**CiIntegrationId** | Pointer to **NullableInt32** | Omit or use null to inherit the organization default, use 0 for the built-in CI service, or use an accessible CI integration ID. A project-owned integration must be shared with the app&#39;s project. | [optional] 
**RegistryIntegrationId** | Pointer to **NullableInt32** | Omit or use null to inherit the organization default, use 0 for the built-in registry, or use an accessible registry integration ID. A project-owned integration must be shared with the app&#39;s project. | [optional] 
**DeferInitialDeployment** | Pointer to **bool** | Defers the automatic initial build and deployment while preserving app instance initialization. Intended for automation that configures the instance before explicitly starting its first build. | [optional] [default to false]
**Settings** | Pointer to [**AppInstanceSettingsInput**](AppInstanceSettingsInput.md) |  | [optional] 
**Access** | Pointer to [**NewAppInstanceAccessInput**](NewAppInstanceAccessInput.md) |  | [optional] 

## Methods

### NewNewAppInput

`func NewNewAppInput(name string, instanceName string, stackRevId int32, envId int32, ) *NewAppInput`

NewNewAppInput instantiates a new NewAppInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewAppInputWithDefaults

`func NewNewAppInputWithDefaults() *NewAppInput`

NewNewAppInputWithDefaults instantiates a new NewAppInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *NewAppInput) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *NewAppInput) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *NewAppInput) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *NewAppInput) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

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

### HasTitle

`func (o *NewAppInput) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

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

### HasInstanceTitle

`func (o *NewAppInput) HasInstanceTitle() bool`

HasInstanceTitle returns a boolean if a field has been set.

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

### HasDomain

`func (o *NewAppInput) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetProjectId

`func (o *NewAppInput) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *NewAppInput) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *NewAppInput) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *NewAppInput) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *NewAppInput) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *NewAppInput) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetStackRevId

`func (o *NewAppInput) GetStackRevId() int32`

GetStackRevId returns the StackRevId field if non-nil, zero value otherwise.

### GetStackRevIdOk

`func (o *NewAppInput) GetStackRevIdOk() (*int32, bool)`

GetStackRevIdOk returns a tuple with the StackRevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackRevId

`func (o *NewAppInput) SetStackRevId(v int32)`

SetStackRevId sets StackRevId field to given value.


### GetServices

`func (o *NewAppInput) GetServices() []CreateAppServiceInput`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *NewAppInput) GetServicesOk() (*[]CreateAppServiceInput, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *NewAppInput) SetServices(v []CreateAppServiceInput)`

SetServices sets Services field to given value.

### HasServices

`func (o *NewAppInput) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetClusterId

`func (o *NewAppInput) GetClusterId() int32`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *NewAppInput) GetClusterIdOk() (*int32, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *NewAppInput) SetClusterId(v int32)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *NewAppInput) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *NewAppInput) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *NewAppInput) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetEnvId

`func (o *NewAppInput) GetEnvId() int32`

GetEnvId returns the EnvId field if non-nil, zero value otherwise.

### GetEnvIdOk

`func (o *NewAppInput) GetEnvIdOk() (*int32, bool)`

GetEnvIdOk returns a tuple with the EnvId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvId

`func (o *NewAppInput) SetEnvId(v int32)`

SetEnvId sets EnvId field to given value.


### GetCiIntegrationId

`func (o *NewAppInput) GetCiIntegrationId() int32`

GetCiIntegrationId returns the CiIntegrationId field if non-nil, zero value otherwise.

### GetCiIntegrationIdOk

`func (o *NewAppInput) GetCiIntegrationIdOk() (*int32, bool)`

GetCiIntegrationIdOk returns a tuple with the CiIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiIntegrationId

`func (o *NewAppInput) SetCiIntegrationId(v int32)`

SetCiIntegrationId sets CiIntegrationId field to given value.

### HasCiIntegrationId

`func (o *NewAppInput) HasCiIntegrationId() bool`

HasCiIntegrationId returns a boolean if a field has been set.

### SetCiIntegrationIdNil

`func (o *NewAppInput) SetCiIntegrationIdNil(b bool)`

 SetCiIntegrationIdNil sets the value for CiIntegrationId to be an explicit nil

### UnsetCiIntegrationId
`func (o *NewAppInput) UnsetCiIntegrationId()`

UnsetCiIntegrationId ensures that no value is present for CiIntegrationId, not even an explicit nil
### GetRegistryIntegrationId

`func (o *NewAppInput) GetRegistryIntegrationId() int32`

GetRegistryIntegrationId returns the RegistryIntegrationId field if non-nil, zero value otherwise.

### GetRegistryIntegrationIdOk

`func (o *NewAppInput) GetRegistryIntegrationIdOk() (*int32, bool)`

GetRegistryIntegrationIdOk returns a tuple with the RegistryIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryIntegrationId

`func (o *NewAppInput) SetRegistryIntegrationId(v int32)`

SetRegistryIntegrationId sets RegistryIntegrationId field to given value.

### HasRegistryIntegrationId

`func (o *NewAppInput) HasRegistryIntegrationId() bool`

HasRegistryIntegrationId returns a boolean if a field has been set.

### SetRegistryIntegrationIdNil

`func (o *NewAppInput) SetRegistryIntegrationIdNil(b bool)`

 SetRegistryIntegrationIdNil sets the value for RegistryIntegrationId to be an explicit nil

### UnsetRegistryIntegrationId
`func (o *NewAppInput) UnsetRegistryIntegrationId()`

UnsetRegistryIntegrationId ensures that no value is present for RegistryIntegrationId, not even an explicit nil
### GetDeferInitialDeployment

`func (o *NewAppInput) GetDeferInitialDeployment() bool`

GetDeferInitialDeployment returns the DeferInitialDeployment field if non-nil, zero value otherwise.

### GetDeferInitialDeploymentOk

`func (o *NewAppInput) GetDeferInitialDeploymentOk() (*bool, bool)`

GetDeferInitialDeploymentOk returns a tuple with the DeferInitialDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferInitialDeployment

`func (o *NewAppInput) SetDeferInitialDeployment(v bool)`

SetDeferInitialDeployment sets DeferInitialDeployment field to given value.

### HasDeferInitialDeployment

`func (o *NewAppInput) HasDeferInitialDeployment() bool`

HasDeferInitialDeployment returns a boolean if a field has been set.

### GetSettings

`func (o *NewAppInput) GetSettings() AppInstanceSettingsInput`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *NewAppInput) GetSettingsOk() (*AppInstanceSettingsInput, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *NewAppInput) SetSettings(v AppInstanceSettingsInput)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *NewAppInput) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetAccess

`func (o *NewAppInput) GetAccess() NewAppInstanceAccessInput`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *NewAppInput) GetAccessOk() (*NewAppInstanceAccessInput, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *NewAppInput) SetAccess(v NewAppInstanceAccessInput)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *NewAppInput) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


