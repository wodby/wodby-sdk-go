# TeamMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**OrgMembershipId** | **int32** |  | 
**Role** | **string** |  | 

## Methods

### NewTeamMember

`func NewTeamMember(id int32, orgMembershipId int32, role string, ) *TeamMember`

NewTeamMember instantiates a new TeamMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamMemberWithDefaults

`func NewTeamMemberWithDefaults() *TeamMember`

NewTeamMemberWithDefaults instantiates a new TeamMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TeamMember) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamMember) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamMember) SetId(v int32)`

SetId sets Id field to given value.


### GetOrgMembershipId

`func (o *TeamMember) GetOrgMembershipId() int32`

GetOrgMembershipId returns the OrgMembershipId field if non-nil, zero value otherwise.

### GetOrgMembershipIdOk

`func (o *TeamMember) GetOrgMembershipIdOk() (*int32, bool)`

GetOrgMembershipIdOk returns a tuple with the OrgMembershipId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgMembershipId

`func (o *TeamMember) SetOrgMembershipId(v int32)`

SetOrgMembershipId sets OrgMembershipId field to given value.


### GetRole

`func (o *TeamMember) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *TeamMember) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *TeamMember) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


