# NewProjectInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgID** | **int32** |  | 
**Name** | **string** |  | 
**Title** | **string** |  | 
**TeamIDs** | Pointer to **[]int32** |  | [optional] 
**OrgMembershipIDs** | Pointer to **[]int32** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 

## Methods

### NewNewProjectInput

`func NewNewProjectInput(orgID int32, name string, title string, ) *NewProjectInput`

NewNewProjectInput instantiates a new NewProjectInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewProjectInputWithDefaults

`func NewNewProjectInputWithDefaults() *NewProjectInput`

NewNewProjectInputWithDefaults instantiates a new NewProjectInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgID

`func (o *NewProjectInput) GetOrgID() int32`

GetOrgID returns the OrgID field if non-nil, zero value otherwise.

### GetOrgIDOk

`func (o *NewProjectInput) GetOrgIDOk() (*int32, bool)`

GetOrgIDOk returns a tuple with the OrgID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgID

`func (o *NewProjectInput) SetOrgID(v int32)`

SetOrgID sets OrgID field to given value.


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


### GetTeamIDs

`func (o *NewProjectInput) GetTeamIDs() []int32`

GetTeamIDs returns the TeamIDs field if non-nil, zero value otherwise.

### GetTeamIDsOk

`func (o *NewProjectInput) GetTeamIDsOk() (*[]int32, bool)`

GetTeamIDsOk returns a tuple with the TeamIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIDs

`func (o *NewProjectInput) SetTeamIDs(v []int32)`

SetTeamIDs sets TeamIDs field to given value.

### HasTeamIDs

`func (o *NewProjectInput) HasTeamIDs() bool`

HasTeamIDs returns a boolean if a field has been set.

### GetOrgMembershipIDs

`func (o *NewProjectInput) GetOrgMembershipIDs() []int32`

GetOrgMembershipIDs returns the OrgMembershipIDs field if non-nil, zero value otherwise.

### GetOrgMembershipIDsOk

`func (o *NewProjectInput) GetOrgMembershipIDsOk() (*[]int32, bool)`

GetOrgMembershipIDsOk returns a tuple with the OrgMembershipIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgMembershipIDs

`func (o *NewProjectInput) SetOrgMembershipIDs(v []int32)`

SetOrgMembershipIDs sets OrgMembershipIDs field to given value.

### HasOrgMembershipIDs

`func (o *NewProjectInput) HasOrgMembershipIDs() bool`

HasOrgMembershipIDs returns a boolean if a field has been set.

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


