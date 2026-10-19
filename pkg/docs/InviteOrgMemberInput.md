# InviteOrgMemberInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. | [optional] 
**Email** | **string** |  | 
**Role** | **string** |  | 
**TeamIds** | Pointer to **[]int32** |  | [optional] 
**ProjectIds** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewInviteOrgMemberInput

`func NewInviteOrgMemberInput(email string, role string, ) *InviteOrgMemberInput`

NewInviteOrgMemberInput instantiates a new InviteOrgMemberInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteOrgMemberInputWithDefaults

`func NewInviteOrgMemberInputWithDefaults() *InviteOrgMemberInput`

NewInviteOrgMemberInputWithDefaults instantiates a new InviteOrgMemberInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *InviteOrgMemberInput) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *InviteOrgMemberInput) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *InviteOrgMemberInput) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *InviteOrgMemberInput) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetEmail

`func (o *InviteOrgMemberInput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InviteOrgMemberInput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InviteOrgMemberInput) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRole

`func (o *InviteOrgMemberInput) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *InviteOrgMemberInput) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *InviteOrgMemberInput) SetRole(v string)`

SetRole sets Role field to given value.


### GetTeamIds

`func (o *InviteOrgMemberInput) GetTeamIds() []int32`

GetTeamIds returns the TeamIds field if non-nil, zero value otherwise.

### GetTeamIdsOk

`func (o *InviteOrgMemberInput) GetTeamIdsOk() (*[]int32, bool)`

GetTeamIdsOk returns a tuple with the TeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIds

`func (o *InviteOrgMemberInput) SetTeamIds(v []int32)`

SetTeamIds sets TeamIds field to given value.

### HasTeamIds

`func (o *InviteOrgMemberInput) HasTeamIds() bool`

HasTeamIds returns a boolean if a field has been set.

### GetProjectIds

`func (o *InviteOrgMemberInput) GetProjectIds() []int32`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *InviteOrgMemberInput) GetProjectIdsOk() (*[]int32, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *InviteOrgMemberInput) SetProjectIds(v []int32)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *InviteOrgMemberInput) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


