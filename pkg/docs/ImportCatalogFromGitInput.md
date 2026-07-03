# ImportCatalogFromGitInput

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

## Methods

### NewImportCatalogFromGitInput

`func NewImportCatalogFromGitInput(integrationId int32, remoteGitRepoId string, gitRef string, gitRefType string, ) *ImportCatalogFromGitInput`

NewImportCatalogFromGitInput instantiates a new ImportCatalogFromGitInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportCatalogFromGitInputWithDefaults

`func NewImportCatalogFromGitInputWithDefaults() *ImportCatalogFromGitInput`

NewImportCatalogFromGitInputWithDefaults instantiates a new ImportCatalogFromGitInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *ImportCatalogFromGitInput) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ImportCatalogFromGitInput) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ImportCatalogFromGitInput) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ImportCatalogFromGitInput) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProjectId

`func (o *ImportCatalogFromGitInput) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ImportCatalogFromGitInput) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ImportCatalogFromGitInput) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ImportCatalogFromGitInput) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *ImportCatalogFromGitInput) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *ImportCatalogFromGitInput) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetIntegrationId

`func (o *ImportCatalogFromGitInput) GetIntegrationId() int32`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *ImportCatalogFromGitInput) GetIntegrationIdOk() (*int32, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *ImportCatalogFromGitInput) SetIntegrationId(v int32)`

SetIntegrationId sets IntegrationId field to given value.


### GetRemoteGitRepoId

`func (o *ImportCatalogFromGitInput) GetRemoteGitRepoId() string`

GetRemoteGitRepoId returns the RemoteGitRepoId field if non-nil, zero value otherwise.

### GetRemoteGitRepoIdOk

`func (o *ImportCatalogFromGitInput) GetRemoteGitRepoIdOk() (*string, bool)`

GetRemoteGitRepoIdOk returns a tuple with the RemoteGitRepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteGitRepoId

`func (o *ImportCatalogFromGitInput) SetRemoteGitRepoId(v string)`

SetRemoteGitRepoId sets RemoteGitRepoId field to given value.


### GetGitRef

`func (o *ImportCatalogFromGitInput) GetGitRef() string`

GetGitRef returns the GitRef field if non-nil, zero value otherwise.

### GetGitRefOk

`func (o *ImportCatalogFromGitInput) GetGitRefOk() (*string, bool)`

GetGitRefOk returns a tuple with the GitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRef

`func (o *ImportCatalogFromGitInput) SetGitRef(v string)`

SetGitRef sets GitRef field to given value.


### GetGitRefType

`func (o *ImportCatalogFromGitInput) GetGitRefType() string`

GetGitRefType returns the GitRefType field if non-nil, zero value otherwise.

### GetGitRefTypeOk

`func (o *ImportCatalogFromGitInput) GetGitRefTypeOk() (*string, bool)`

GetGitRefTypeOk returns a tuple with the GitRefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRefType

`func (o *ImportCatalogFromGitInput) SetGitRefType(v string)`

SetGitRefType sets GitRefType field to given value.


### GetAutoUpdate

`func (o *ImportCatalogFromGitInput) GetAutoUpdate() GitAutoUpdateSettingsInput`

GetAutoUpdate returns the AutoUpdate field if non-nil, zero value otherwise.

### GetAutoUpdateOk

`func (o *ImportCatalogFromGitInput) GetAutoUpdateOk() (*GitAutoUpdateSettingsInput, bool)`

GetAutoUpdateOk returns a tuple with the AutoUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpdate

`func (o *ImportCatalogFromGitInput) SetAutoUpdate(v GitAutoUpdateSettingsInput)`

SetAutoUpdate sets AutoUpdate field to given value.

### HasAutoUpdate

`func (o *ImportCatalogFromGitInput) HasAutoUpdate() bool`

HasAutoUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


