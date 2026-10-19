# TeamInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **int32** | Used when creating a team. Optional for API-key requests; defaults to the API key&#39;s organization. | [optional] 
**Name** | **string** |  | 
**Title** | **string** |  | 
**Members** | [**[]TeamMemberInput**](TeamMemberInput.md) | The complete member list. | 

## Methods

### NewTeamInput

`func NewTeamInput(name string, title string, members []TeamMemberInput, ) *TeamInput`

NewTeamInput instantiates a new TeamInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamInputWithDefaults

`func NewTeamInputWithDefaults() *TeamInput`

NewTeamInputWithDefaults instantiates a new TeamInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *TeamInput) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *TeamInput) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *TeamInput) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *TeamInput) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetName

`func (o *TeamInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamInput) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *TeamInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TeamInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TeamInput) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetMembers

`func (o *TeamInput) GetMembers() []TeamMemberInput`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *TeamInput) GetMembersOk() (*[]TeamMemberInput, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *TeamInput) SetMembers(v []TeamMemberInput)`

SetMembers sets Members field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


