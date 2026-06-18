# NewProjectInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **int32** |  | 
**Name** | **string** |  | 
**Title** | **string** |  | 
**TeamIds** | Pointer to **[]int32** |  | [optional] 
**OrgMembershipIds** | Pointer to **[]int32** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 

## Methods

### NewNewProjectInput

`func NewNewProjectInput(orgId int32, name string, title string, ) *NewProjectInput`

NewNewProjectInput instantiates a new NewProjectInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewProjectInputWithDefaults

`func NewNewProjectInputWithDefaults() *NewProjectInput`

NewNewProjectInputWithDefaults instantiates a new NewProjectInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *NewProjectInput) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *NewProjectInput) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *NewProjectInput) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.


### GetName

`func (o *NewProjectInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewProjectInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewProjectInput) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *NewProjectInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewProjectInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NewProjectInput) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetTeamIds

`func (o *NewProjectInput) GetTeamIds() []int32`

GetTeamIds returns the TeamIds field if non-nil, zero value otherwise.

### GetTeamIdsOk

`func (o *NewProjectInput) GetTeamIdsOk() (*[]int32, bool)`

GetTeamIdsOk returns a tuple with the TeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIds

`func (o *NewProjectInput) SetTeamIds(v []int32)`

SetTeamIds sets TeamIds field to given value.

### HasTeamIds

`func (o *NewProjectInput) HasTeamIds() bool`

HasTeamIds returns a boolean if a field has been set.

### GetOrgMembershipIds

`func (o *NewProjectInput) GetOrgMembershipIds() []int32`

GetOrgMembershipIds returns the OrgMembershipIds field if non-nil, zero value otherwise.

### GetOrgMembershipIdsOk

`func (o *NewProjectInput) GetOrgMembershipIdsOk() (*[]int32, bool)`

GetOrgMembershipIdsOk returns a tuple with the OrgMembershipIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgMembershipIds

`func (o *NewProjectInput) SetOrgMembershipIds(v []int32)`

SetOrgMembershipIds sets OrgMembershipIds field to given value.

### HasOrgMembershipIds

`func (o *NewProjectInput) HasOrgMembershipIds() bool`

HasOrgMembershipIds returns a boolean if a field has been set.

### GetRole

`func (o *NewProjectInput) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *NewProjectInput) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *NewProjectInput) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *NewProjectInput) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


