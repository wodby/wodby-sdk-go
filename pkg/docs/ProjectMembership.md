# ProjectMembership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**ProjectId** | **int32** |  | 
**OrgMembershipId** | Pointer to **NullableInt32** |  | [optional] 
**TeamId** | Pointer to **NullableInt32** |  | [optional] 
**Role** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewProjectMembership

`func NewProjectMembership(id int32, projectId int32, role string, createdAt time.Time, updatedAt time.Time, ) *ProjectMembership`

NewProjectMembership instantiates a new ProjectMembership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectMembershipWithDefaults

`func NewProjectMembershipWithDefaults() *ProjectMembership`

NewProjectMembershipWithDefaults instantiates a new ProjectMembership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectMembership) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectMembership) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectMembership) SetId(v int32)`

SetId sets Id field to given value.


### GetProjectId

`func (o *ProjectMembership) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ProjectMembership) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ProjectMembership) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetOrgMembershipId

`func (o *ProjectMembership) GetOrgMembershipId() int32`

GetOrgMembershipId returns the OrgMembershipId field if non-nil, zero value otherwise.

### GetOrgMembershipIdOk

`func (o *ProjectMembership) GetOrgMembershipIdOk() (*int32, bool)`

GetOrgMembershipIdOk returns a tuple with the OrgMembershipId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgMembershipId

`func (o *ProjectMembership) SetOrgMembershipId(v int32)`

SetOrgMembershipId sets OrgMembershipId field to given value.

### HasOrgMembershipId

`func (o *ProjectMembership) HasOrgMembershipId() bool`

HasOrgMembershipId returns a boolean if a field has been set.

### SetOrgMembershipIdNil

`func (o *ProjectMembership) SetOrgMembershipIdNil(b bool)`

 SetOrgMembershipIdNil sets the value for OrgMembershipId to be an explicit nil

### UnsetOrgMembershipId
`func (o *ProjectMembership) UnsetOrgMembershipId()`

UnsetOrgMembershipId ensures that no value is present for OrgMembershipId, not even an explicit nil
### GetTeamId

`func (o *ProjectMembership) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *ProjectMembership) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *ProjectMembership) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *ProjectMembership) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### SetTeamIdNil

`func (o *ProjectMembership) SetTeamIdNil(b bool)`

 SetTeamIdNil sets the value for TeamId to be an explicit nil

### UnsetTeamId
`func (o *ProjectMembership) UnsetTeamId()`

UnsetTeamId ensures that no value is present for TeamId, not even an explicit nil
### GetRole

`func (o *ProjectMembership) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ProjectMembership) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ProjectMembership) SetRole(v string)`

SetRole sets Role field to given value.


### GetCreatedAt

`func (o *ProjectMembership) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectMembership) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectMembership) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ProjectMembership) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProjectMembership) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProjectMembership) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


