# ProjectRoleInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **int32** | Used when creating a role. Optional for API-key requests; defaults to the API key&#39;s organization. | [optional] 
**Name** | **string** | Lowercase letters, digits, dashes or underscores; unique in the organization. | 
**Title** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**ProjectPermissions** | **[]string** |  | 
**EnvPermissions** | **map[string][]string** | Environment permissions by environment type (prod, staging, test, dev, feature). | 

## Methods

### NewProjectRoleInput

`func NewProjectRoleInput(name string, title string, projectPermissions []string, envPermissions map[string][]string, ) *ProjectRoleInput`

NewProjectRoleInput instantiates a new ProjectRoleInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectRoleInputWithDefaults

`func NewProjectRoleInputWithDefaults() *ProjectRoleInput`

NewProjectRoleInputWithDefaults instantiates a new ProjectRoleInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *ProjectRoleInput) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ProjectRoleInput) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ProjectRoleInput) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ProjectRoleInput) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetName

`func (o *ProjectRoleInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectRoleInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectRoleInput) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *ProjectRoleInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProjectRoleInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProjectRoleInput) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *ProjectRoleInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectRoleInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectRoleInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectRoleInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProjectPermissions

`func (o *ProjectRoleInput) GetProjectPermissions() []string`

GetProjectPermissions returns the ProjectPermissions field if non-nil, zero value otherwise.

### GetProjectPermissionsOk

`func (o *ProjectRoleInput) GetProjectPermissionsOk() (*[]string, bool)`

GetProjectPermissionsOk returns a tuple with the ProjectPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectPermissions

`func (o *ProjectRoleInput) SetProjectPermissions(v []string)`

SetProjectPermissions sets ProjectPermissions field to given value.


### GetEnvPermissions

`func (o *ProjectRoleInput) GetEnvPermissions() map[string][]string`

GetEnvPermissions returns the EnvPermissions field if non-nil, zero value otherwise.

### GetEnvPermissionsOk

`func (o *ProjectRoleInput) GetEnvPermissionsOk() (*map[string][]string, bool)`

GetEnvPermissionsOk returns a tuple with the EnvPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvPermissions

`func (o *ProjectRoleInput) SetEnvPermissions(v map[string][]string)`

SetEnvPermissions sets EnvPermissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


