# ImportStacksFromGitInput

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
**AvailableToAllProjects** | Pointer to **NullableBool** | Lets every project in the organization use the imported stacks without a project share. Only organization-owned stacks, imported without projectId, can be available to all projects. | [optional] 

## Methods

### NewImportStacksFromGitInput

`func NewImportStacksFromGitInput(integrationId int32, remoteGitRepoId string, gitRef string, gitRefType string, ) *ImportStacksFromGitInput`

NewImportStacksFromGitInput instantiates a new ImportStacksFromGitInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportStacksFromGitInputWithDefaults

`func NewImportStacksFromGitInputWithDefaults() *ImportStacksFromGitInput`

NewImportStacksFromGitInputWithDefaults instantiates a new ImportStacksFromGitInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *ImportStacksFromGitInput) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ImportStacksFromGitInput) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ImportStacksFromGitInput) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ImportStacksFromGitInput) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProjectId

`func (o *ImportStacksFromGitInput) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ImportStacksFromGitInput) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ImportStacksFromGitInput) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ImportStacksFromGitInput) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *ImportStacksFromGitInput) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *ImportStacksFromGitInput) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetIntegrationId

`func (o *ImportStacksFromGitInput) GetIntegrationId() int32`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *ImportStacksFromGitInput) GetIntegrationIdOk() (*int32, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *ImportStacksFromGitInput) SetIntegrationId(v int32)`

SetIntegrationId sets IntegrationId field to given value.


### GetRemoteGitRepoId

`func (o *ImportStacksFromGitInput) GetRemoteGitRepoId() string`

GetRemoteGitRepoId returns the RemoteGitRepoId field if non-nil, zero value otherwise.

### GetRemoteGitRepoIdOk

`func (o *ImportStacksFromGitInput) GetRemoteGitRepoIdOk() (*string, bool)`

GetRemoteGitRepoIdOk returns a tuple with the RemoteGitRepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteGitRepoId

`func (o *ImportStacksFromGitInput) SetRemoteGitRepoId(v string)`

SetRemoteGitRepoId sets RemoteGitRepoId field to given value.


### GetGitRef

`func (o *ImportStacksFromGitInput) GetGitRef() string`

GetGitRef returns the GitRef field if non-nil, zero value otherwise.

### GetGitRefOk

`func (o *ImportStacksFromGitInput) GetGitRefOk() (*string, bool)`

GetGitRefOk returns a tuple with the GitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRef

`func (o *ImportStacksFromGitInput) SetGitRef(v string)`

SetGitRef sets GitRef field to given value.


### GetGitRefType

`func (o *ImportStacksFromGitInput) GetGitRefType() string`

GetGitRefType returns the GitRefType field if non-nil, zero value otherwise.

### GetGitRefTypeOk

`func (o *ImportStacksFromGitInput) GetGitRefTypeOk() (*string, bool)`

GetGitRefTypeOk returns a tuple with the GitRefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRefType

`func (o *ImportStacksFromGitInput) SetGitRefType(v string)`

SetGitRefType sets GitRefType field to given value.


### GetAutoUpdate

`func (o *ImportStacksFromGitInput) GetAutoUpdate() GitAutoUpdateSettingsInput`

GetAutoUpdate returns the AutoUpdate field if non-nil, zero value otherwise.

### GetAutoUpdateOk

`func (o *ImportStacksFromGitInput) GetAutoUpdateOk() (*GitAutoUpdateSettingsInput, bool)`

GetAutoUpdateOk returns a tuple with the AutoUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpdate

`func (o *ImportStacksFromGitInput) SetAutoUpdate(v GitAutoUpdateSettingsInput)`

SetAutoUpdate sets AutoUpdate field to given value.

### HasAutoUpdate

`func (o *ImportStacksFromGitInput) HasAutoUpdate() bool`

HasAutoUpdate returns a boolean if a field has been set.

### GetAvailableToAllProjects

`func (o *ImportStacksFromGitInput) GetAvailableToAllProjects() bool`

GetAvailableToAllProjects returns the AvailableToAllProjects field if non-nil, zero value otherwise.

### GetAvailableToAllProjectsOk

`func (o *ImportStacksFromGitInput) GetAvailableToAllProjectsOk() (*bool, bool)`

GetAvailableToAllProjectsOk returns a tuple with the AvailableToAllProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableToAllProjects

`func (o *ImportStacksFromGitInput) SetAvailableToAllProjects(v bool)`

SetAvailableToAllProjects sets AvailableToAllProjects field to given value.

### HasAvailableToAllProjects

`func (o *ImportStacksFromGitInput) HasAvailableToAllProjects() bool`

HasAvailableToAllProjects returns a boolean if a field has been set.

### SetAvailableToAllProjectsNil

`func (o *ImportStacksFromGitInput) SetAvailableToAllProjectsNil(b bool)`

 SetAvailableToAllProjectsNil sets the value for AvailableToAllProjects to be an explicit nil

### UnsetAvailableToAllProjects
`func (o *ImportStacksFromGitInput) UnsetAvailableToAllProjects()`

UnsetAvailableToAllProjects ensures that no value is present for AvailableToAllProjects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


