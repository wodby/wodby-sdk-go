# ImportServicesFromGitInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. | [optional] 
**ProjectId** | Pointer to **NullableInt32** |  | [optional] 
**IntegrationId** | **int32** |  | 
**RemoteGitRepoId** | **string** |  | 
**GitRef** | **string** |  | 
**GitRefType** | **string** |  | 
**AutoUpdate** | Pointer to [**GitAutoUpdateSettingsInput**](GitAutoUpdateSettingsInput.md) |  | [optional] 
**RegistryIntegrationId** | Pointer to **NullableInt32** | Optional registry integration used when an imported service image requires authentication. | [optional] 

## Methods

### NewImportServicesFromGitInput

`func NewImportServicesFromGitInput(integrationId int32, remoteGitRepoId string, gitRef string, gitRefType string, ) *ImportServicesFromGitInput`

NewImportServicesFromGitInput instantiates a new ImportServicesFromGitInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportServicesFromGitInputWithDefaults

`func NewImportServicesFromGitInputWithDefaults() *ImportServicesFromGitInput`

NewImportServicesFromGitInputWithDefaults instantiates a new ImportServicesFromGitInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *ImportServicesFromGitInput) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ImportServicesFromGitInput) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ImportServicesFromGitInput) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ImportServicesFromGitInput) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProjectId

`func (o *ImportServicesFromGitInput) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ImportServicesFromGitInput) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ImportServicesFromGitInput) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ImportServicesFromGitInput) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *ImportServicesFromGitInput) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *ImportServicesFromGitInput) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetIntegrationId

`func (o *ImportServicesFromGitInput) GetIntegrationId() int32`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *ImportServicesFromGitInput) GetIntegrationIdOk() (*int32, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *ImportServicesFromGitInput) SetIntegrationId(v int32)`

SetIntegrationId sets IntegrationId field to given value.


### GetRemoteGitRepoId

`func (o *ImportServicesFromGitInput) GetRemoteGitRepoId() string`

GetRemoteGitRepoId returns the RemoteGitRepoId field if non-nil, zero value otherwise.

### GetRemoteGitRepoIdOk

`func (o *ImportServicesFromGitInput) GetRemoteGitRepoIdOk() (*string, bool)`

GetRemoteGitRepoIdOk returns a tuple with the RemoteGitRepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteGitRepoId

`func (o *ImportServicesFromGitInput) SetRemoteGitRepoId(v string)`

SetRemoteGitRepoId sets RemoteGitRepoId field to given value.


### GetGitRef

`func (o *ImportServicesFromGitInput) GetGitRef() string`

GetGitRef returns the GitRef field if non-nil, zero value otherwise.

### GetGitRefOk

`func (o *ImportServicesFromGitInput) GetGitRefOk() (*string, bool)`

GetGitRefOk returns a tuple with the GitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRef

`func (o *ImportServicesFromGitInput) SetGitRef(v string)`

SetGitRef sets GitRef field to given value.


### GetGitRefType

`func (o *ImportServicesFromGitInput) GetGitRefType() string`

GetGitRefType returns the GitRefType field if non-nil, zero value otherwise.

### GetGitRefTypeOk

`func (o *ImportServicesFromGitInput) GetGitRefTypeOk() (*string, bool)`

GetGitRefTypeOk returns a tuple with the GitRefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRefType

`func (o *ImportServicesFromGitInput) SetGitRefType(v string)`

SetGitRefType sets GitRefType field to given value.


### GetAutoUpdate

`func (o *ImportServicesFromGitInput) GetAutoUpdate() GitAutoUpdateSettingsInput`

GetAutoUpdate returns the AutoUpdate field if non-nil, zero value otherwise.

### GetAutoUpdateOk

`func (o *ImportServicesFromGitInput) GetAutoUpdateOk() (*GitAutoUpdateSettingsInput, bool)`

GetAutoUpdateOk returns a tuple with the AutoUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpdate

`func (o *ImportServicesFromGitInput) SetAutoUpdate(v GitAutoUpdateSettingsInput)`

SetAutoUpdate sets AutoUpdate field to given value.

### HasAutoUpdate

`func (o *ImportServicesFromGitInput) HasAutoUpdate() bool`

HasAutoUpdate returns a boolean if a field has been set.

### GetRegistryIntegrationId

`func (o *ImportServicesFromGitInput) GetRegistryIntegrationId() int32`

GetRegistryIntegrationId returns the RegistryIntegrationId field if non-nil, zero value otherwise.

### GetRegistryIntegrationIdOk

`func (o *ImportServicesFromGitInput) GetRegistryIntegrationIdOk() (*int32, bool)`

GetRegistryIntegrationIdOk returns a tuple with the RegistryIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryIntegrationId

`func (o *ImportServicesFromGitInput) SetRegistryIntegrationId(v int32)`

SetRegistryIntegrationId sets RegistryIntegrationId field to given value.

### HasRegistryIntegrationId

`func (o *ImportServicesFromGitInput) HasRegistryIntegrationId() bool`

HasRegistryIntegrationId returns a boolean if a field has been set.

### SetRegistryIntegrationIdNil

`func (o *ImportServicesFromGitInput) SetRegistryIntegrationIdNil(b bool)`

 SetRegistryIntegrationIdNil sets the value for RegistryIntegrationId to be an explicit nil

### UnsetRegistryIntegrationId
`func (o *ImportServicesFromGitInput) UnsetRegistryIntegrationId()`

UnsetRegistryIntegrationId ensures that no value is present for RegistryIntegrationId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


