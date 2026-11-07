# ProjectRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Custom role ID; absent for system roles. | [optional] 
**Role** | **string** | The system role, or custom. | 
**System** | **bool** |  | 
**Name** | **string** |  | 
**Title** | **string** |  | 
**Description** | **string** |  | 
**ProjectPermissions** | **[]string** |  | 
**EnvPermissions** | **map[string][]string** | Environment permissions by environment type (prod, staging, test, dev, feature). | 
**AssignableToTeams** | **bool** |  | 
**SupportsRuntimeSecretReveal** | **bool** | Whether memberships with this system role can turn on runtime secret reveal. | 
**MembershipCount** | **int32** | Project memberships holding a custom role. Visible to organization owners, admins and support. | 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewProjectRole

`func NewProjectRole(role string, system bool, name string, title string, description string, projectPermissions []string, envPermissions map[string][]string, assignableToTeams bool, supportsRuntimeSecretReveal bool, membershipCount int32, ) *ProjectRole`

NewProjectRole instantiates a new ProjectRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectRoleWithDefaults

`func NewProjectRoleWithDefaults() *ProjectRole`

NewProjectRoleWithDefaults instantiates a new ProjectRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectRole) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectRole) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectRole) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectRole) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProjectRole) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProjectRole) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetRole

`func (o *ProjectRole) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ProjectRole) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ProjectRole) SetRole(v string)`

SetRole sets Role field to given value.


### GetSystem

`func (o *ProjectRole) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *ProjectRole) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *ProjectRole) SetSystem(v bool)`

SetSystem sets System field to given value.


### GetName

`func (o *ProjectRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectRole) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *ProjectRole) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProjectRole) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProjectRole) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *ProjectRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectRole) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetProjectPermissions

`func (o *ProjectRole) GetProjectPermissions() []string`

GetProjectPermissions returns the ProjectPermissions field if non-nil, zero value otherwise.

### GetProjectPermissionsOk

`func (o *ProjectRole) GetProjectPermissionsOk() (*[]string, bool)`

GetProjectPermissionsOk returns a tuple with the ProjectPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectPermissions

`func (o *ProjectRole) SetProjectPermissions(v []string)`

SetProjectPermissions sets ProjectPermissions field to given value.


### GetEnvPermissions

`func (o *ProjectRole) GetEnvPermissions() map[string][]string`

GetEnvPermissions returns the EnvPermissions field if non-nil, zero value otherwise.

### GetEnvPermissionsOk

`func (o *ProjectRole) GetEnvPermissionsOk() (*map[string][]string, bool)`

GetEnvPermissionsOk returns a tuple with the EnvPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvPermissions

`func (o *ProjectRole) SetEnvPermissions(v map[string][]string)`

SetEnvPermissions sets EnvPermissions field to given value.


### GetAssignableToTeams

`func (o *ProjectRole) GetAssignableToTeams() bool`

GetAssignableToTeams returns the AssignableToTeams field if non-nil, zero value otherwise.

### GetAssignableToTeamsOk

`func (o *ProjectRole) GetAssignableToTeamsOk() (*bool, bool)`

GetAssignableToTeamsOk returns a tuple with the AssignableToTeams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignableToTeams

`func (o *ProjectRole) SetAssignableToTeams(v bool)`

SetAssignableToTeams sets AssignableToTeams field to given value.


### GetSupportsRuntimeSecretReveal

`func (o *ProjectRole) GetSupportsRuntimeSecretReveal() bool`

GetSupportsRuntimeSecretReveal returns the SupportsRuntimeSecretReveal field if non-nil, zero value otherwise.

### GetSupportsRuntimeSecretRevealOk

`func (o *ProjectRole) GetSupportsRuntimeSecretRevealOk() (*bool, bool)`

GetSupportsRuntimeSecretRevealOk returns a tuple with the SupportsRuntimeSecretReveal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsRuntimeSecretReveal

`func (o *ProjectRole) SetSupportsRuntimeSecretReveal(v bool)`

SetSupportsRuntimeSecretReveal sets SupportsRuntimeSecretReveal field to given value.


### GetMembershipCount

`func (o *ProjectRole) GetMembershipCount() int32`

GetMembershipCount returns the MembershipCount field if non-nil, zero value otherwise.

### GetMembershipCountOk

`func (o *ProjectRole) GetMembershipCountOk() (*int32, bool)`

GetMembershipCountOk returns a tuple with the MembershipCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipCount

`func (o *ProjectRole) SetMembershipCount(v int32)`

SetMembershipCount sets MembershipCount field to given value.


### GetCreatedAt

`func (o *ProjectRole) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectRole) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectRole) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProjectRole) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ProjectRole) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ProjectRole) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *ProjectRole) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProjectRole) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProjectRole) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProjectRole) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *ProjectRole) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ProjectRole) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


