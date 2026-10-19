# ResourceSharing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnershipScope** | **string** | org when the organization owns the resource, project when one project owns it. | 
**OwnerProjectId** | Pointer to **NullableInt32** | Project that owns the resource when ownershipScope is project. | [optional] 
**AvailableToAllProjects** | **bool** | True when every project in the organization can use the resource without a project share. Only organization-owned clusters, integrations, services and stacks can be available to all projects; it is always false for other resources. | 
**SharedProjects** | [**[]ResourceProjectShare**](ResourceProjectShare.md) | Projects the resource is shared with that you can see. The owning project has the owner role. | 

## Methods

### NewResourceSharing

`func NewResourceSharing(ownershipScope string, availableToAllProjects bool, sharedProjects []ResourceProjectShare, ) *ResourceSharing`

NewResourceSharing instantiates a new ResourceSharing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSharingWithDefaults

`func NewResourceSharingWithDefaults() *ResourceSharing`

NewResourceSharingWithDefaults instantiates a new ResourceSharing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnershipScope

`func (o *ResourceSharing) GetOwnershipScope() string`

GetOwnershipScope returns the OwnershipScope field if non-nil, zero value otherwise.

### GetOwnershipScopeOk

`func (o *ResourceSharing) GetOwnershipScopeOk() (*string, bool)`

GetOwnershipScopeOk returns a tuple with the OwnershipScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipScope

`func (o *ResourceSharing) SetOwnershipScope(v string)`

SetOwnershipScope sets OwnershipScope field to given value.


### GetOwnerProjectId

`func (o *ResourceSharing) GetOwnerProjectId() int32`

GetOwnerProjectId returns the OwnerProjectId field if non-nil, zero value otherwise.

### GetOwnerProjectIdOk

`func (o *ResourceSharing) GetOwnerProjectIdOk() (*int32, bool)`

GetOwnerProjectIdOk returns a tuple with the OwnerProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerProjectId

`func (o *ResourceSharing) SetOwnerProjectId(v int32)`

SetOwnerProjectId sets OwnerProjectId field to given value.

### HasOwnerProjectId

`func (o *ResourceSharing) HasOwnerProjectId() bool`

HasOwnerProjectId returns a boolean if a field has been set.

### SetOwnerProjectIdNil

`func (o *ResourceSharing) SetOwnerProjectIdNil(b bool)`

 SetOwnerProjectIdNil sets the value for OwnerProjectId to be an explicit nil

### UnsetOwnerProjectId
`func (o *ResourceSharing) UnsetOwnerProjectId()`

UnsetOwnerProjectId ensures that no value is present for OwnerProjectId, not even an explicit nil
### GetAvailableToAllProjects

`func (o *ResourceSharing) GetAvailableToAllProjects() bool`

GetAvailableToAllProjects returns the AvailableToAllProjects field if non-nil, zero value otherwise.

### GetAvailableToAllProjectsOk

`func (o *ResourceSharing) GetAvailableToAllProjectsOk() (*bool, bool)`

GetAvailableToAllProjectsOk returns a tuple with the AvailableToAllProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableToAllProjects

`func (o *ResourceSharing) SetAvailableToAllProjects(v bool)`

SetAvailableToAllProjects sets AvailableToAllProjects field to given value.


### GetSharedProjects

`func (o *ResourceSharing) GetSharedProjects() []ResourceProjectShare`

GetSharedProjects returns the SharedProjects field if non-nil, zero value otherwise.

### GetSharedProjectsOk

`func (o *ResourceSharing) GetSharedProjectsOk() (*[]ResourceProjectShare, bool)`

GetSharedProjectsOk returns a tuple with the SharedProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedProjects

`func (o *ResourceSharing) SetSharedProjects(v []ResourceProjectShare)`

SetSharedProjects sets SharedProjects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


