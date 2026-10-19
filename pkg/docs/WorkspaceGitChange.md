# WorkspaceGitChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** |  | 
**OriginalPath** | Pointer to **string** | Previous path of a renamed or copied file. | [optional] 
**Kind** | **string** |  | 
**Staged** | **bool** | The index differs from HEAD for this path. | 
**Unstaged** | **bool** | The working tree differs from the index for this path. | 

## Methods

### NewWorkspaceGitChange

`func NewWorkspaceGitChange(path string, kind string, staged bool, unstaged bool, ) *WorkspaceGitChange`

NewWorkspaceGitChange instantiates a new WorkspaceGitChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceGitChangeWithDefaults

`func NewWorkspaceGitChangeWithDefaults() *WorkspaceGitChange`

NewWorkspaceGitChangeWithDefaults instantiates a new WorkspaceGitChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *WorkspaceGitChange) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *WorkspaceGitChange) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *WorkspaceGitChange) SetPath(v string)`

SetPath sets Path field to given value.


### GetOriginalPath

`func (o *WorkspaceGitChange) GetOriginalPath() string`

GetOriginalPath returns the OriginalPath field if non-nil, zero value otherwise.

### GetOriginalPathOk

`func (o *WorkspaceGitChange) GetOriginalPathOk() (*string, bool)`

GetOriginalPathOk returns a tuple with the OriginalPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPath

`func (o *WorkspaceGitChange) SetOriginalPath(v string)`

SetOriginalPath sets OriginalPath field to given value.

### HasOriginalPath

`func (o *WorkspaceGitChange) HasOriginalPath() bool`

HasOriginalPath returns a boolean if a field has been set.

### GetKind

`func (o *WorkspaceGitChange) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *WorkspaceGitChange) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *WorkspaceGitChange) SetKind(v string)`

SetKind sets Kind field to given value.


### GetStaged

`func (o *WorkspaceGitChange) GetStaged() bool`

GetStaged returns the Staged field if non-nil, zero value otherwise.

### GetStagedOk

`func (o *WorkspaceGitChange) GetStagedOk() (*bool, bool)`

GetStagedOk returns a tuple with the Staged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaged

`func (o *WorkspaceGitChange) SetStaged(v bool)`

SetStaged sets Staged field to given value.


### GetUnstaged

`func (o *WorkspaceGitChange) GetUnstaged() bool`

GetUnstaged returns the Unstaged field if non-nil, zero value otherwise.

### GetUnstagedOk

`func (o *WorkspaceGitChange) GetUnstagedOk() (*bool, bool)`

GetUnstagedOk returns a tuple with the Unstaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnstaged

`func (o *WorkspaceGitChange) SetUnstaged(v bool)`

SetUnstaged sets Unstaged field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


