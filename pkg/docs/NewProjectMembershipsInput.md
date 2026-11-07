# NewProjectMembershipsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int32** |  | 
**OrgMembershipIds** | Pointer to **[]int32** |  | [optional] 
**TeamIds** | Pointer to **[]int32** |  | [optional] 
**Role** | **string** |  | 
**CustomRoleId** | Pointer to **int32** | Required when role is custom; a custom role of the project&#39;s organization. | [optional] 

## Methods

### NewNewProjectMembershipsInput

`func NewNewProjectMembershipsInput(projectId int32, role string, ) *NewProjectMembershipsInput`

NewNewProjectMembershipsInput instantiates a new NewProjectMembershipsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewProjectMembershipsInputWithDefaults

`func NewNewProjectMembershipsInputWithDefaults() *NewProjectMembershipsInput`

NewNewProjectMembershipsInputWithDefaults instantiates a new NewProjectMembershipsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *NewProjectMembershipsInput) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *NewProjectMembershipsInput) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *NewProjectMembershipsInput) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetOrgMembershipIds

`func (o *NewProjectMembershipsInput) GetOrgMembershipIds() []int32`

GetOrgMembershipIds returns the OrgMembershipIds field if non-nil, zero value otherwise.

### GetOrgMembershipIdsOk

`func (o *NewProjectMembershipsInput) GetOrgMembershipIdsOk() (*[]int32, bool)`

GetOrgMembershipIdsOk returns a tuple with the OrgMembershipIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgMembershipIds

`func (o *NewProjectMembershipsInput) SetOrgMembershipIds(v []int32)`

SetOrgMembershipIds sets OrgMembershipIds field to given value.

### HasOrgMembershipIds

`func (o *NewProjectMembershipsInput) HasOrgMembershipIds() bool`

HasOrgMembershipIds returns a boolean if a field has been set.

### GetTeamIds

`func (o *NewProjectMembershipsInput) GetTeamIds() []int32`

GetTeamIds returns the TeamIds field if non-nil, zero value otherwise.

### GetTeamIdsOk

`func (o *NewProjectMembershipsInput) GetTeamIdsOk() (*[]int32, bool)`

GetTeamIdsOk returns a tuple with the TeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIds

`func (o *NewProjectMembershipsInput) SetTeamIds(v []int32)`

SetTeamIds sets TeamIds field to given value.

### HasTeamIds

`func (o *NewProjectMembershipsInput) HasTeamIds() bool`

HasTeamIds returns a boolean if a field has been set.

### GetRole

`func (o *NewProjectMembershipsInput) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *NewProjectMembershipsInput) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *NewProjectMembershipsInput) SetRole(v string)`

SetRole sets Role field to given value.


### GetCustomRoleId

`func (o *NewProjectMembershipsInput) GetCustomRoleId() int32`

GetCustomRoleId returns the CustomRoleId field if non-nil, zero value otherwise.

### GetCustomRoleIdOk

`func (o *NewProjectMembershipsInput) GetCustomRoleIdOk() (*int32, bool)`

GetCustomRoleIdOk returns a tuple with the CustomRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRoleId

`func (o *NewProjectMembershipsInput) SetCustomRoleId(v int32)`

SetCustomRoleId sets CustomRoleId field to given value.

### HasCustomRoleId

`func (o *NewProjectMembershipsInput) HasCustomRoleId() bool`

HasCustomRoleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


