# ProjectMembershipRoleInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | **string** |  | 
**CustomRoleId** | Pointer to **int32** | Required when role is custom; a custom role of the project&#39;s organization. | [optional] 

## Methods

### NewProjectMembershipRoleInput

`func NewProjectMembershipRoleInput(role string, ) *ProjectMembershipRoleInput`

NewProjectMembershipRoleInput instantiates a new ProjectMembershipRoleInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectMembershipRoleInputWithDefaults

`func NewProjectMembershipRoleInputWithDefaults() *ProjectMembershipRoleInput`

NewProjectMembershipRoleInputWithDefaults instantiates a new ProjectMembershipRoleInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *ProjectMembershipRoleInput) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ProjectMembershipRoleInput) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ProjectMembershipRoleInput) SetRole(v string)`

SetRole sets Role field to given value.


### GetCustomRoleId

`func (o *ProjectMembershipRoleInput) GetCustomRoleId() int32`

GetCustomRoleId returns the CustomRoleId field if non-nil, zero value otherwise.

### GetCustomRoleIdOk

`func (o *ProjectMembershipRoleInput) GetCustomRoleIdOk() (*int32, bool)`

GetCustomRoleIdOk returns a tuple with the CustomRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRoleId

`func (o *ProjectMembershipRoleInput) SetCustomRoleId(v int32)`

SetCustomRoleId sets CustomRoleId field to given value.

### HasCustomRoleId

`func (o *ProjectMembershipRoleInput) HasCustomRoleId() bool`

HasCustomRoleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


