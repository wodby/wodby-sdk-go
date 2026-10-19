# TeamMemberInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgMembershipId** | **int32** |  | 
**Role** | **string** |  | 

## Methods

### NewTeamMemberInput

`func NewTeamMemberInput(orgMembershipId int32, role string, ) *TeamMemberInput`

NewTeamMemberInput instantiates a new TeamMemberInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamMemberInputWithDefaults

`func NewTeamMemberInputWithDefaults() *TeamMemberInput`

NewTeamMemberInputWithDefaults instantiates a new TeamMemberInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgMembershipId

`func (o *TeamMemberInput) GetOrgMembershipId() int32`

GetOrgMembershipId returns the OrgMembershipId field if non-nil, zero value otherwise.

### GetOrgMembershipIdOk

`func (o *TeamMemberInput) GetOrgMembershipIdOk() (*int32, bool)`

GetOrgMembershipIdOk returns a tuple with the OrgMembershipId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgMembershipId

`func (o *TeamMemberInput) SetOrgMembershipId(v int32)`

SetOrgMembershipId sets OrgMembershipId field to given value.


### GetRole

`func (o *TeamMemberInput) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *TeamMemberInput) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *TeamMemberInput) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


