# WorkspaceGitDiff

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | **bool** | False when the diff cannot be read; reason explains why and diff is empty. | 
**Reason** | Pointer to **string** |  | [optional] 
**Path** | **string** |  | 
**Diff** | **string** | Unified diff of staged and unstaged changes. Empty when the file does not differ. | 
**Binary** | **bool** | True for a file Git does not show line by line. | 
**Truncated** | **bool** | True when the diff is longer than what is returned. | 

## Methods

### NewWorkspaceGitDiff

`func NewWorkspaceGitDiff(available bool, path string, diff string, binary bool, truncated bool, ) *WorkspaceGitDiff`

NewWorkspaceGitDiff instantiates a new WorkspaceGitDiff object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceGitDiffWithDefaults

`func NewWorkspaceGitDiffWithDefaults() *WorkspaceGitDiff`

NewWorkspaceGitDiffWithDefaults instantiates a new WorkspaceGitDiff object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *WorkspaceGitDiff) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *WorkspaceGitDiff) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *WorkspaceGitDiff) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetReason

`func (o *WorkspaceGitDiff) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *WorkspaceGitDiff) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *WorkspaceGitDiff) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *WorkspaceGitDiff) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetPath

`func (o *WorkspaceGitDiff) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *WorkspaceGitDiff) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *WorkspaceGitDiff) SetPath(v string)`

SetPath sets Path field to given value.


### GetDiff

`func (o *WorkspaceGitDiff) GetDiff() string`

GetDiff returns the Diff field if non-nil, zero value otherwise.

### GetDiffOk

`func (o *WorkspaceGitDiff) GetDiffOk() (*string, bool)`

GetDiffOk returns a tuple with the Diff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiff

`func (o *WorkspaceGitDiff) SetDiff(v string)`

SetDiff sets Diff field to given value.


### GetBinary

`func (o *WorkspaceGitDiff) GetBinary() bool`

GetBinary returns the Binary field if non-nil, zero value otherwise.

### GetBinaryOk

`func (o *WorkspaceGitDiff) GetBinaryOk() (*bool, bool)`

GetBinaryOk returns a tuple with the Binary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinary

`func (o *WorkspaceGitDiff) SetBinary(v bool)`

SetBinary sets Binary field to given value.


### GetTruncated

`func (o *WorkspaceGitDiff) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *WorkspaceGitDiff) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *WorkspaceGitDiff) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


