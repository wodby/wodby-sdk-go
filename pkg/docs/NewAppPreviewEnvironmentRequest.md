# NewAppPreviewEnvironmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **int32** | Pull request number. | 
**RemoteGitRepoId** | Pointer to **NullableString** | Remote repository of the pull request. Required when the base environment tracks several repositories. | [optional] 

## Methods

### NewNewAppPreviewEnvironmentRequest

`func NewNewAppPreviewEnvironmentRequest(number int32, ) *NewAppPreviewEnvironmentRequest`

NewNewAppPreviewEnvironmentRequest instantiates a new NewAppPreviewEnvironmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewAppPreviewEnvironmentRequestWithDefaults

`func NewNewAppPreviewEnvironmentRequestWithDefaults() *NewAppPreviewEnvironmentRequest`

NewNewAppPreviewEnvironmentRequestWithDefaults instantiates a new NewAppPreviewEnvironmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *NewAppPreviewEnvironmentRequest) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *NewAppPreviewEnvironmentRequest) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *NewAppPreviewEnvironmentRequest) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetRemoteGitRepoId

`func (o *NewAppPreviewEnvironmentRequest) GetRemoteGitRepoId() string`

GetRemoteGitRepoId returns the RemoteGitRepoId field if non-nil, zero value otherwise.

### GetRemoteGitRepoIdOk

`func (o *NewAppPreviewEnvironmentRequest) GetRemoteGitRepoIdOk() (*string, bool)`

GetRemoteGitRepoIdOk returns a tuple with the RemoteGitRepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteGitRepoId

`func (o *NewAppPreviewEnvironmentRequest) SetRemoteGitRepoId(v string)`

SetRemoteGitRepoId sets RemoteGitRepoId field to given value.

### HasRemoteGitRepoId

`func (o *NewAppPreviewEnvironmentRequest) HasRemoteGitRepoId() bool`

HasRemoteGitRepoId returns a boolean if a field has been set.

### SetRemoteGitRepoIdNil

`func (o *NewAppPreviewEnvironmentRequest) SetRemoteGitRepoIdNil(b bool)`

 SetRemoteGitRepoIdNil sets the value for RemoteGitRepoId to be an explicit nil

### UnsetRemoteGitRepoId
`func (o *NewAppPreviewEnvironmentRequest) UnsetRemoteGitRepoId()`

UnsetRemoteGitRepoId ensures that no value is present for RemoteGitRepoId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


