# ResourceSharingInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnershipScope** | **string** | org to make the organization the owner, project to make one project the owner. | 
**OwnerProjectId** | Pointer to **NullableInt32** | Project that becomes the owner. Required when ownershipScope is project. | [optional] 
**SharedProjects** | [**[]ProjectShareInput**](ProjectShareInput.md) | The complete list of project shares. A project left out loses its share. The owning project always keeps write access. You need write access to every project whose share you add or remove. | 
**AvailableToAllProjects** | Pointer to **NullableBool** | Lets every project in the organization use the resource without a project share. Only organization-owned clusters, integrations, services and stacks can be available to all projects. Omit it to keep the current setting. | [optional] 

## Methods

### NewResourceSharingInput

`func NewResourceSharingInput(ownershipScope string, sharedProjects []ProjectShareInput, ) *ResourceSharingInput`

NewResourceSharingInput instantiates a new ResourceSharingInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSharingInputWithDefaults

`func NewResourceSharingInputWithDefaults() *ResourceSharingInput`

NewResourceSharingInputWithDefaults instantiates a new ResourceSharingInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnershipScope

`func (o *ResourceSharingInput) GetOwnershipScope() string`

GetOwnershipScope returns the OwnershipScope field if non-nil, zero value otherwise.

### GetOwnershipScopeOk

`func (o *ResourceSharingInput) GetOwnershipScopeOk() (*string, bool)`

GetOwnershipScopeOk returns a tuple with the OwnershipScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipScope

`func (o *ResourceSharingInput) SetOwnershipScope(v string)`

SetOwnershipScope sets OwnershipScope field to given value.


### GetOwnerProjectId

`func (o *ResourceSharingInput) GetOwnerProjectId() int32`

GetOwnerProjectId returns the OwnerProjectId field if non-nil, zero value otherwise.

### GetOwnerProjectIdOk

`func (o *ResourceSharingInput) GetOwnerProjectIdOk() (*int32, bool)`

GetOwnerProjectIdOk returns a tuple with the OwnerProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerProjectId

`func (o *ResourceSharingInput) SetOwnerProjectId(v int32)`

SetOwnerProjectId sets OwnerProjectId field to given value.

### HasOwnerProjectId

`func (o *ResourceSharingInput) HasOwnerProjectId() bool`

HasOwnerProjectId returns a boolean if a field has been set.

### SetOwnerProjectIdNil

`func (o *ResourceSharingInput) SetOwnerProjectIdNil(b bool)`

 SetOwnerProjectIdNil sets the value for OwnerProjectId to be an explicit nil

### UnsetOwnerProjectId
`func (o *ResourceSharingInput) UnsetOwnerProjectId()`

UnsetOwnerProjectId ensures that no value is present for OwnerProjectId, not even an explicit nil
### GetSharedProjects

`func (o *ResourceSharingInput) GetSharedProjects() []ProjectShareInput`

GetSharedProjects returns the SharedProjects field if non-nil, zero value otherwise.

### GetSharedProjectsOk

`func (o *ResourceSharingInput) GetSharedProjectsOk() (*[]ProjectShareInput, bool)`

GetSharedProjectsOk returns a tuple with the SharedProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedProjects

`func (o *ResourceSharingInput) SetSharedProjects(v []ProjectShareInput)`

SetSharedProjects sets SharedProjects field to given value.


### GetAvailableToAllProjects

`func (o *ResourceSharingInput) GetAvailableToAllProjects() bool`

GetAvailableToAllProjects returns the AvailableToAllProjects field if non-nil, zero value otherwise.

### GetAvailableToAllProjectsOk

`func (o *ResourceSharingInput) GetAvailableToAllProjectsOk() (*bool, bool)`

GetAvailableToAllProjectsOk returns a tuple with the AvailableToAllProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableToAllProjects

`func (o *ResourceSharingInput) SetAvailableToAllProjects(v bool)`

SetAvailableToAllProjects sets AvailableToAllProjects field to given value.

### HasAvailableToAllProjects

`func (o *ResourceSharingInput) HasAvailableToAllProjects() bool`

HasAvailableToAllProjects returns a boolean if a field has been set.

### SetAvailableToAllProjectsNil

`func (o *ResourceSharingInput) SetAvailableToAllProjectsNil(b bool)`

 SetAvailableToAllProjectsNil sets the value for AvailableToAllProjects to be an explicit nil

### UnsetAvailableToAllProjects
`func (o *ResourceSharingInput) UnsetAvailableToAllProjects()`

UnsetAvailableToAllProjects ensures that no value is present for AvailableToAllProjects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


