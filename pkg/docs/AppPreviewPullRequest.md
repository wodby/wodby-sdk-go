# AppPreviewPullRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **int32** |  | 
**Title** | **string** |  | 
**Url** | **string** |  | 
**RemoteGitRepoId** | **string** | Remote repository of the pull request. Pass it when creating the preview environment. | 
**RepositoryUrl** | **string** |  | 
**HeadRef** | **string** |  | 
**BaseRef** | **string** |  | 
**Author** | **string** |  | 
**Draft** | **bool** |  | 
**FromFork** | **bool** | The head branch lives in another repository. Pull requests from forks never get preview environments. | 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**HasPreview** | **bool** | The pull request already has a preview environment, possibly one being deleted. | 
**PreviewAppEnvironmentId** | Pointer to **NullableInt32** |  | [optional] 
**SkipReason** | Pointer to **NullableString** | Why a preview environment cannot be created for this pull request now. Null when it can. | [optional] 

## Methods

### NewAppPreviewPullRequest

`func NewAppPreviewPullRequest(number int32, title string, url string, remoteGitRepoId string, repositoryUrl string, headRef string, baseRef string, author string, draft bool, fromFork bool, hasPreview bool, ) *AppPreviewPullRequest`

NewAppPreviewPullRequest instantiates a new AppPreviewPullRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPreviewPullRequestWithDefaults

`func NewAppPreviewPullRequestWithDefaults() *AppPreviewPullRequest`

NewAppPreviewPullRequestWithDefaults instantiates a new AppPreviewPullRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *AppPreviewPullRequest) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *AppPreviewPullRequest) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *AppPreviewPullRequest) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetTitle

`func (o *AppPreviewPullRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AppPreviewPullRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AppPreviewPullRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUrl

`func (o *AppPreviewPullRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AppPreviewPullRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AppPreviewPullRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetRemoteGitRepoId

`func (o *AppPreviewPullRequest) GetRemoteGitRepoId() string`

GetRemoteGitRepoId returns the RemoteGitRepoId field if non-nil, zero value otherwise.

### GetRemoteGitRepoIdOk

`func (o *AppPreviewPullRequest) GetRemoteGitRepoIdOk() (*string, bool)`

GetRemoteGitRepoIdOk returns a tuple with the RemoteGitRepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteGitRepoId

`func (o *AppPreviewPullRequest) SetRemoteGitRepoId(v string)`

SetRemoteGitRepoId sets RemoteGitRepoId field to given value.


### GetRepositoryUrl

`func (o *AppPreviewPullRequest) GetRepositoryUrl() string`

GetRepositoryUrl returns the RepositoryUrl field if non-nil, zero value otherwise.

### GetRepositoryUrlOk

`func (o *AppPreviewPullRequest) GetRepositoryUrlOk() (*string, bool)`

GetRepositoryUrlOk returns a tuple with the RepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryUrl

`func (o *AppPreviewPullRequest) SetRepositoryUrl(v string)`

SetRepositoryUrl sets RepositoryUrl field to given value.


### GetHeadRef

`func (o *AppPreviewPullRequest) GetHeadRef() string`

GetHeadRef returns the HeadRef field if non-nil, zero value otherwise.

### GetHeadRefOk

`func (o *AppPreviewPullRequest) GetHeadRefOk() (*string, bool)`

GetHeadRefOk returns a tuple with the HeadRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadRef

`func (o *AppPreviewPullRequest) SetHeadRef(v string)`

SetHeadRef sets HeadRef field to given value.


### GetBaseRef

`func (o *AppPreviewPullRequest) GetBaseRef() string`

GetBaseRef returns the BaseRef field if non-nil, zero value otherwise.

### GetBaseRefOk

`func (o *AppPreviewPullRequest) GetBaseRefOk() (*string, bool)`

GetBaseRefOk returns a tuple with the BaseRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseRef

`func (o *AppPreviewPullRequest) SetBaseRef(v string)`

SetBaseRef sets BaseRef field to given value.


### GetAuthor

`func (o *AppPreviewPullRequest) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *AppPreviewPullRequest) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *AppPreviewPullRequest) SetAuthor(v string)`

SetAuthor sets Author field to given value.


### GetDraft

`func (o *AppPreviewPullRequest) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *AppPreviewPullRequest) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *AppPreviewPullRequest) SetDraft(v bool)`

SetDraft sets Draft field to given value.


### GetFromFork

`func (o *AppPreviewPullRequest) GetFromFork() bool`

GetFromFork returns the FromFork field if non-nil, zero value otherwise.

### GetFromForkOk

`func (o *AppPreviewPullRequest) GetFromForkOk() (*bool, bool)`

GetFromForkOk returns a tuple with the FromFork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromFork

`func (o *AppPreviewPullRequest) SetFromFork(v bool)`

SetFromFork sets FromFork field to given value.


### GetUpdatedAt

`func (o *AppPreviewPullRequest) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AppPreviewPullRequest) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AppPreviewPullRequest) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AppPreviewPullRequest) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *AppPreviewPullRequest) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *AppPreviewPullRequest) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetHasPreview

`func (o *AppPreviewPullRequest) GetHasPreview() bool`

GetHasPreview returns the HasPreview field if non-nil, zero value otherwise.

### GetHasPreviewOk

`func (o *AppPreviewPullRequest) GetHasPreviewOk() (*bool, bool)`

GetHasPreviewOk returns a tuple with the HasPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPreview

`func (o *AppPreviewPullRequest) SetHasPreview(v bool)`

SetHasPreview sets HasPreview field to given value.


### GetPreviewAppEnvironmentId

`func (o *AppPreviewPullRequest) GetPreviewAppEnvironmentId() int32`

GetPreviewAppEnvironmentId returns the PreviewAppEnvironmentId field if non-nil, zero value otherwise.

### GetPreviewAppEnvironmentIdOk

`func (o *AppPreviewPullRequest) GetPreviewAppEnvironmentIdOk() (*int32, bool)`

GetPreviewAppEnvironmentIdOk returns a tuple with the PreviewAppEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewAppEnvironmentId

`func (o *AppPreviewPullRequest) SetPreviewAppEnvironmentId(v int32)`

SetPreviewAppEnvironmentId sets PreviewAppEnvironmentId field to given value.

### HasPreviewAppEnvironmentId

`func (o *AppPreviewPullRequest) HasPreviewAppEnvironmentId() bool`

HasPreviewAppEnvironmentId returns a boolean if a field has been set.

### SetPreviewAppEnvironmentIdNil

`func (o *AppPreviewPullRequest) SetPreviewAppEnvironmentIdNil(b bool)`

 SetPreviewAppEnvironmentIdNil sets the value for PreviewAppEnvironmentId to be an explicit nil

### UnsetPreviewAppEnvironmentId
`func (o *AppPreviewPullRequest) UnsetPreviewAppEnvironmentId()`

UnsetPreviewAppEnvironmentId ensures that no value is present for PreviewAppEnvironmentId, not even an explicit nil
### GetSkipReason

`func (o *AppPreviewPullRequest) GetSkipReason() string`

GetSkipReason returns the SkipReason field if non-nil, zero value otherwise.

### GetSkipReasonOk

`func (o *AppPreviewPullRequest) GetSkipReasonOk() (*string, bool)`

GetSkipReasonOk returns a tuple with the SkipReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipReason

`func (o *AppPreviewPullRequest) SetSkipReason(v string)`

SetSkipReason sets SkipReason field to given value.

### HasSkipReason

`func (o *AppPreviewPullRequest) HasSkipReason() bool`

HasSkipReason returns a boolean if a field has been set.

### SetSkipReasonNil

`func (o *AppPreviewPullRequest) SetSkipReasonNil(b bool)`

 SetSkipReasonNil sets the value for SkipReason to be an explicit nil

### UnsetSkipReason
`func (o *AppPreviewPullRequest) UnsetSkipReason()`

UnsetSkipReason ensures that no value is present for SkipReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


