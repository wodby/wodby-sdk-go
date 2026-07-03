# StackService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Title** | **string** |  | 
**Type** | **string** |  | 
**Main** | **bool** |  | 
**Disabled** | **bool** |  | 
**Required** | **bool** |  | 
**Replicas** | **int32** |  | 
**Outdated** | **bool** |  | 
**ServiceRevId** | **int32** |  | 
**ServiceRevName** | **string** |  | 
**ServiceRevTitle** | **string** |  | 
**ServiceRevVersion** | **string** |  | 
**BuildSourceIntegrationId** | Pointer to **NullableInt32** |  | [optional] 
**BuildSourceRemoteRepoId** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewStackService

`func NewStackService(id int32, name string, title string, type_ string, main bool, disabled bool, required bool, replicas int32, outdated bool, serviceRevId int32, serviceRevName string, serviceRevTitle string, serviceRevVersion string, createdAt time.Time, updatedAt time.Time, ) *StackService`

NewStackService instantiates a new StackService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackServiceWithDefaults

`func NewStackServiceWithDefaults() *StackService`

NewStackServiceWithDefaults instantiates a new StackService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StackService) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StackService) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StackService) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *StackService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StackService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StackService) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *StackService) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StackService) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StackService) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *StackService) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StackService) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StackService) SetType(v string)`

SetType sets Type field to given value.


### GetMain

`func (o *StackService) GetMain() bool`

GetMain returns the Main field if non-nil, zero value otherwise.

### GetMainOk

`func (o *StackService) GetMainOk() (*bool, bool)`

GetMainOk returns a tuple with the Main field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMain

`func (o *StackService) SetMain(v bool)`

SetMain sets Main field to given value.


### GetDisabled

`func (o *StackService) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *StackService) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *StackService) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetRequired

`func (o *StackService) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *StackService) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *StackService) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetReplicas

`func (o *StackService) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *StackService) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *StackService) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.


### GetOutdated

`func (o *StackService) GetOutdated() bool`

GetOutdated returns the Outdated field if non-nil, zero value otherwise.

### GetOutdatedOk

`func (o *StackService) GetOutdatedOk() (*bool, bool)`

GetOutdatedOk returns a tuple with the Outdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutdated

`func (o *StackService) SetOutdated(v bool)`

SetOutdated sets Outdated field to given value.


### GetServiceRevId

`func (o *StackService) GetServiceRevId() int32`

GetServiceRevId returns the ServiceRevId field if non-nil, zero value otherwise.

### GetServiceRevIdOk

`func (o *StackService) GetServiceRevIdOk() (*int32, bool)`

GetServiceRevIdOk returns a tuple with the ServiceRevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRevId

`func (o *StackService) SetServiceRevId(v int32)`

SetServiceRevId sets ServiceRevId field to given value.


### GetServiceRevName

`func (o *StackService) GetServiceRevName() string`

GetServiceRevName returns the ServiceRevName field if non-nil, zero value otherwise.

### GetServiceRevNameOk

`func (o *StackService) GetServiceRevNameOk() (*string, bool)`

GetServiceRevNameOk returns a tuple with the ServiceRevName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRevName

`func (o *StackService) SetServiceRevName(v string)`

SetServiceRevName sets ServiceRevName field to given value.


### GetServiceRevTitle

`func (o *StackService) GetServiceRevTitle() string`

GetServiceRevTitle returns the ServiceRevTitle field if non-nil, zero value otherwise.

### GetServiceRevTitleOk

`func (o *StackService) GetServiceRevTitleOk() (*string, bool)`

GetServiceRevTitleOk returns a tuple with the ServiceRevTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRevTitle

`func (o *StackService) SetServiceRevTitle(v string)`

SetServiceRevTitle sets ServiceRevTitle field to given value.


### GetServiceRevVersion

`func (o *StackService) GetServiceRevVersion() string`

GetServiceRevVersion returns the ServiceRevVersion field if non-nil, zero value otherwise.

### GetServiceRevVersionOk

`func (o *StackService) GetServiceRevVersionOk() (*string, bool)`

GetServiceRevVersionOk returns a tuple with the ServiceRevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRevVersion

`func (o *StackService) SetServiceRevVersion(v string)`

SetServiceRevVersion sets ServiceRevVersion field to given value.


### GetBuildSourceIntegrationId

`func (o *StackService) GetBuildSourceIntegrationId() int32`

GetBuildSourceIntegrationId returns the BuildSourceIntegrationId field if non-nil, zero value otherwise.

### GetBuildSourceIntegrationIdOk

`func (o *StackService) GetBuildSourceIntegrationIdOk() (*int32, bool)`

GetBuildSourceIntegrationIdOk returns a tuple with the BuildSourceIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildSourceIntegrationId

`func (o *StackService) SetBuildSourceIntegrationId(v int32)`

SetBuildSourceIntegrationId sets BuildSourceIntegrationId field to given value.

### HasBuildSourceIntegrationId

`func (o *StackService) HasBuildSourceIntegrationId() bool`

HasBuildSourceIntegrationId returns a boolean if a field has been set.

### SetBuildSourceIntegrationIdNil

`func (o *StackService) SetBuildSourceIntegrationIdNil(b bool)`

 SetBuildSourceIntegrationIdNil sets the value for BuildSourceIntegrationId to be an explicit nil

### UnsetBuildSourceIntegrationId
`func (o *StackService) UnsetBuildSourceIntegrationId()`

UnsetBuildSourceIntegrationId ensures that no value is present for BuildSourceIntegrationId, not even an explicit nil
### GetBuildSourceRemoteRepoId

`func (o *StackService) GetBuildSourceRemoteRepoId() string`

GetBuildSourceRemoteRepoId returns the BuildSourceRemoteRepoId field if non-nil, zero value otherwise.

### GetBuildSourceRemoteRepoIdOk

`func (o *StackService) GetBuildSourceRemoteRepoIdOk() (*string, bool)`

GetBuildSourceRemoteRepoIdOk returns a tuple with the BuildSourceRemoteRepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildSourceRemoteRepoId

`func (o *StackService) SetBuildSourceRemoteRepoId(v string)`

SetBuildSourceRemoteRepoId sets BuildSourceRemoteRepoId field to given value.

### HasBuildSourceRemoteRepoId

`func (o *StackService) HasBuildSourceRemoteRepoId() bool`

HasBuildSourceRemoteRepoId returns a boolean if a field has been set.

### SetBuildSourceRemoteRepoIdNil

`func (o *StackService) SetBuildSourceRemoteRepoIdNil(b bool)`

 SetBuildSourceRemoteRepoIdNil sets the value for BuildSourceRemoteRepoId to be an explicit nil

### UnsetBuildSourceRemoteRepoId
`func (o *StackService) UnsetBuildSourceRemoteRepoId()`

UnsetBuildSourceRemoteRepoId ensures that no value is present for BuildSourceRemoteRepoId, not even an explicit nil
### GetCreatedAt

`func (o *StackService) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StackService) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StackService) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *StackService) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *StackService) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *StackService) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


