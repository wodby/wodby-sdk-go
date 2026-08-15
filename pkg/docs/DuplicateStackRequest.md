# DuplicateStackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. | [optional] 
**ProjectId** | Pointer to **NullableInt32** |  | [optional] 
**SourceRevId** | Pointer to **NullableInt32** | Optional immutable source stack revision to duplicate. It must belong to the stack in the request path. | [optional] 
**Settings** | Pointer to [**CopyStackSettingsInput**](CopyStackSettingsInput.md) |  | [optional] 

## Methods

### NewDuplicateStackRequest

`func NewDuplicateStackRequest() *DuplicateStackRequest`

NewDuplicateStackRequest instantiates a new DuplicateStackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDuplicateStackRequestWithDefaults

`func NewDuplicateStackRequestWithDefaults() *DuplicateStackRequest`

NewDuplicateStackRequestWithDefaults instantiates a new DuplicateStackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *DuplicateStackRequest) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *DuplicateStackRequest) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *DuplicateStackRequest) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *DuplicateStackRequest) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProjectId

`func (o *DuplicateStackRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DuplicateStackRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DuplicateStackRequest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *DuplicateStackRequest) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *DuplicateStackRequest) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *DuplicateStackRequest) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetSourceRevId

`func (o *DuplicateStackRequest) GetSourceRevId() int32`

GetSourceRevId returns the SourceRevId field if non-nil, zero value otherwise.

### GetSourceRevIdOk

`func (o *DuplicateStackRequest) GetSourceRevIdOk() (*int32, bool)`

GetSourceRevIdOk returns a tuple with the SourceRevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRevId

`func (o *DuplicateStackRequest) SetSourceRevId(v int32)`

SetSourceRevId sets SourceRevId field to given value.

### HasSourceRevId

`func (o *DuplicateStackRequest) HasSourceRevId() bool`

HasSourceRevId returns a boolean if a field has been set.

### SetSourceRevIdNil

`func (o *DuplicateStackRequest) SetSourceRevIdNil(b bool)`

 SetSourceRevIdNil sets the value for SourceRevId to be an explicit nil

### UnsetSourceRevId
`func (o *DuplicateStackRequest) UnsetSourceRevId()`

UnsetSourceRevId ensures that no value is present for SourceRevId, not even an explicit nil
### GetSettings

`func (o *DuplicateStackRequest) GetSettings() CopyStackSettingsInput`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *DuplicateStackRequest) GetSettingsOk() (*CopyStackSettingsInput, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *DuplicateStackRequest) SetSettings(v CopyStackSettingsInput)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *DuplicateStackRequest) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


