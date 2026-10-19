# WorkspaceGitStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | **bool** | False when the state cannot be read; reason explains why and the other fields are omitted. | 
**Reason** | Pointer to **string** |  | [optional] 
**Branch** | Pointer to **string** | Current branch; omitted for a detached HEAD. | [optional] 
**Upstream** | Pointer to **string** | Upstream branch that the current branch tracks. | [optional] 
**Ahead** | Pointer to **int32** | Commits on the current branch that its upstream does not have. | [optional] 
**Behind** | Pointer to **int32** | Commits on the upstream that the current branch does not have. | [optional] 
**Commit** | Pointer to **string** | HEAD commit; omitted before the first commit. | [optional] 
**CommitSubject** | Pointer to **string** |  | [optional] 
**CommittedAt** | Pointer to **time.Time** |  | [optional] 
**Changes** | [**[]WorkspaceGitChange**](WorkspaceGitChange.md) | Uncommitted changes, including untracked files. | 
**Truncated** | **bool** | True when more changes exist than listed. | 
**CheckedAt** | **time.Time** |  | 

## Methods

### NewWorkspaceGitStatus

`func NewWorkspaceGitStatus(available bool, changes []WorkspaceGitChange, truncated bool, checkedAt time.Time, ) *WorkspaceGitStatus`

NewWorkspaceGitStatus instantiates a new WorkspaceGitStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceGitStatusWithDefaults

`func NewWorkspaceGitStatusWithDefaults() *WorkspaceGitStatus`

NewWorkspaceGitStatusWithDefaults instantiates a new WorkspaceGitStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *WorkspaceGitStatus) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *WorkspaceGitStatus) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *WorkspaceGitStatus) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetReason

`func (o *WorkspaceGitStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *WorkspaceGitStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *WorkspaceGitStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *WorkspaceGitStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetBranch

`func (o *WorkspaceGitStatus) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *WorkspaceGitStatus) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *WorkspaceGitStatus) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *WorkspaceGitStatus) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetUpstream

`func (o *WorkspaceGitStatus) GetUpstream() string`

GetUpstream returns the Upstream field if non-nil, zero value otherwise.

### GetUpstreamOk

`func (o *WorkspaceGitStatus) GetUpstreamOk() (*string, bool)`

GetUpstreamOk returns a tuple with the Upstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstream

`func (o *WorkspaceGitStatus) SetUpstream(v string)`

SetUpstream sets Upstream field to given value.

### HasUpstream

`func (o *WorkspaceGitStatus) HasUpstream() bool`

HasUpstream returns a boolean if a field has been set.

### GetAhead

`func (o *WorkspaceGitStatus) GetAhead() int32`

GetAhead returns the Ahead field if non-nil, zero value otherwise.

### GetAheadOk

`func (o *WorkspaceGitStatus) GetAheadOk() (*int32, bool)`

GetAheadOk returns a tuple with the Ahead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAhead

`func (o *WorkspaceGitStatus) SetAhead(v int32)`

SetAhead sets Ahead field to given value.

### HasAhead

`func (o *WorkspaceGitStatus) HasAhead() bool`

HasAhead returns a boolean if a field has been set.

### GetBehind

`func (o *WorkspaceGitStatus) GetBehind() int32`

GetBehind returns the Behind field if non-nil, zero value otherwise.

### GetBehindOk

`func (o *WorkspaceGitStatus) GetBehindOk() (*int32, bool)`

GetBehindOk returns a tuple with the Behind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehind

`func (o *WorkspaceGitStatus) SetBehind(v int32)`

SetBehind sets Behind field to given value.

### HasBehind

`func (o *WorkspaceGitStatus) HasBehind() bool`

HasBehind returns a boolean if a field has been set.

### GetCommit

`func (o *WorkspaceGitStatus) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *WorkspaceGitStatus) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *WorkspaceGitStatus) SetCommit(v string)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *WorkspaceGitStatus) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetCommitSubject

`func (o *WorkspaceGitStatus) GetCommitSubject() string`

GetCommitSubject returns the CommitSubject field if non-nil, zero value otherwise.

### GetCommitSubjectOk

`func (o *WorkspaceGitStatus) GetCommitSubjectOk() (*string, bool)`

GetCommitSubjectOk returns a tuple with the CommitSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSubject

`func (o *WorkspaceGitStatus) SetCommitSubject(v string)`

SetCommitSubject sets CommitSubject field to given value.

### HasCommitSubject

`func (o *WorkspaceGitStatus) HasCommitSubject() bool`

HasCommitSubject returns a boolean if a field has been set.

### GetCommittedAt

`func (o *WorkspaceGitStatus) GetCommittedAt() time.Time`

GetCommittedAt returns the CommittedAt field if non-nil, zero value otherwise.

### GetCommittedAtOk

`func (o *WorkspaceGitStatus) GetCommittedAtOk() (*time.Time, bool)`

GetCommittedAtOk returns a tuple with the CommittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommittedAt

`func (o *WorkspaceGitStatus) SetCommittedAt(v time.Time)`

SetCommittedAt sets CommittedAt field to given value.

### HasCommittedAt

`func (o *WorkspaceGitStatus) HasCommittedAt() bool`

HasCommittedAt returns a boolean if a field has been set.

### GetChanges

`func (o *WorkspaceGitStatus) GetChanges() []WorkspaceGitChange`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *WorkspaceGitStatus) GetChangesOk() (*[]WorkspaceGitChange, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *WorkspaceGitStatus) SetChanges(v []WorkspaceGitChange)`

SetChanges sets Changes field to given value.


### GetTruncated

`func (o *WorkspaceGitStatus) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *WorkspaceGitStatus) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *WorkspaceGitStatus) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.


### GetCheckedAt

`func (o *WorkspaceGitStatus) GetCheckedAt() time.Time`

GetCheckedAt returns the CheckedAt field if non-nil, zero value otherwise.

### GetCheckedAtOk

`func (o *WorkspaceGitStatus) GetCheckedAtOk() (*time.Time, bool)`

GetCheckedAtOk returns a tuple with the CheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckedAt

`func (o *WorkspaceGitStatus) SetCheckedAt(v time.Time)`

SetCheckedAt sets CheckedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


