# NewStackInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. | [optional] 
**ProjectId** | Pointer to **NullableInt32** |  | [optional] 
**AvailableToAllProjects** | Pointer to **NullableBool** | Lets every project in the organization use the new resource without a project share. Only an organization-owned resource, created without projectId, can be available to all projects. | [optional] 
**Name** | **string** |  | 
**Title** | **string** |  | 
**Icon** | Pointer to **NullableString** |  | [optional] 
**Services** | [**[]NewStackServiceDefinitionInput**](NewStackServiceDefinitionInput.md) |  | 

## Methods

### NewNewStackInput

`func NewNewStackInput(name string, title string, services []NewStackServiceDefinitionInput, ) *NewStackInput`

NewNewStackInput instantiates a new NewStackInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewStackInputWithDefaults

`func NewNewStackInputWithDefaults() *NewStackInput`

NewNewStackInputWithDefaults instantiates a new NewStackInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *NewStackInput) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *NewStackInput) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *NewStackInput) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *NewStackInput) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProjectId

`func (o *NewStackInput) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *NewStackInput) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *NewStackInput) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *NewStackInput) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *NewStackInput) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *NewStackInput) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetAvailableToAllProjects

`func (o *NewStackInput) GetAvailableToAllProjects() bool`

GetAvailableToAllProjects returns the AvailableToAllProjects field if non-nil, zero value otherwise.

### GetAvailableToAllProjectsOk

`func (o *NewStackInput) GetAvailableToAllProjectsOk() (*bool, bool)`

GetAvailableToAllProjectsOk returns a tuple with the AvailableToAllProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableToAllProjects

`func (o *NewStackInput) SetAvailableToAllProjects(v bool)`

SetAvailableToAllProjects sets AvailableToAllProjects field to given value.

### HasAvailableToAllProjects

`func (o *NewStackInput) HasAvailableToAllProjects() bool`

HasAvailableToAllProjects returns a boolean if a field has been set.

### SetAvailableToAllProjectsNil

`func (o *NewStackInput) SetAvailableToAllProjectsNil(b bool)`

 SetAvailableToAllProjectsNil sets the value for AvailableToAllProjects to be an explicit nil

### UnsetAvailableToAllProjects
`func (o *NewStackInput) UnsetAvailableToAllProjects()`

UnsetAvailableToAllProjects ensures that no value is present for AvailableToAllProjects, not even an explicit nil
### GetName

`func (o *NewStackInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewStackInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewStackInput) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *NewStackInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewStackInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NewStackInput) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetIcon

`func (o *NewStackInput) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *NewStackInput) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *NewStackInput) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *NewStackInput) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *NewStackInput) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *NewStackInput) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetServices

`func (o *NewStackInput) GetServices() []NewStackServiceDefinitionInput`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *NewStackInput) GetServicesOk() (*[]NewStackServiceDefinitionInput, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *NewStackInput) SetServices(v []NewStackServiceDefinitionInput)`

SetServices sets Services field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


