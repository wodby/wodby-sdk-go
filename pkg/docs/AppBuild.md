# AppBuild

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Number** | **int32** | App environment build-launch number. Build records created by the same Wodby CI launch share this number. | 
**Status** | **string** |  | 
**AppInstanceId** | **int32** |  | 
**AppServiceId** | **int32** |  | 
**TaskId** | Pointer to **NullableInt32** |  | [optional] 
**Task** | Pointer to [**NullableTask**](Task.md) |  | [optional] 
**AppServiceBuilds** | [**[]AppServiceBuild**](AppServiceBuild.md) |  | 
**GitRefType** | **string** |  | 
**GitRef** | **string** |  | 
**CommitHash** | **string** |  | 
**CommitMessage** | **string** |  | 
**AuthorName** | Pointer to **NullableString** |  | [optional] 
**AuthorEmail** | Pointer to **NullableString** |  | [optional] 
**Workflow** | **string** | CI workflow that produced an externally built image. Empty for Wodby CI builds. | 
**StackRevId** | Pointer to **NullableInt32** | Stack revision the build was made for. | [optional] 
**GitRepoId** | Pointer to **NullableInt32** |  | [optional] 
**CiIntegrationId** | Pointer to **NullableInt32** |  | [optional] 
**RegistryIntegrationId** | Pointer to **NullableInt32** |  | [optional] 
**ImageVoidStatus** | **string** | Whether the build&#39;s images were voided and can no longer be deployed. | 
**CanVoidImages** | **bool** | True when the build&#39;s images can be voided now. | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**StartedAt** | Pointer to **NullableTime** |  | [optional] 
**EndedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewAppBuild

`func NewAppBuild(id int32, number int32, status string, appInstanceId int32, appServiceId int32, appServiceBuilds []AppServiceBuild, gitRefType string, gitRef string, commitHash string, commitMessage string, workflow string, imageVoidStatus string, canVoidImages bool, createdAt time.Time, updatedAt time.Time, ) *AppBuild`

NewAppBuild instantiates a new AppBuild object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppBuildWithDefaults

`func NewAppBuildWithDefaults() *AppBuild`

NewAppBuildWithDefaults instantiates a new AppBuild object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppBuild) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppBuild) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppBuild) SetId(v int32)`

SetId sets Id field to given value.


### GetNumber

`func (o *AppBuild) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *AppBuild) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *AppBuild) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetStatus

`func (o *AppBuild) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppBuild) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppBuild) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAppInstanceId

`func (o *AppBuild) GetAppInstanceId() int32`

GetAppInstanceId returns the AppInstanceId field if non-nil, zero value otherwise.

### GetAppInstanceIdOk

`func (o *AppBuild) GetAppInstanceIdOk() (*int32, bool)`

GetAppInstanceIdOk returns a tuple with the AppInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceId

`func (o *AppBuild) SetAppInstanceId(v int32)`

SetAppInstanceId sets AppInstanceId field to given value.


### GetAppServiceId

`func (o *AppBuild) GetAppServiceId() int32`

GetAppServiceId returns the AppServiceId field if non-nil, zero value otherwise.

### GetAppServiceIdOk

`func (o *AppBuild) GetAppServiceIdOk() (*int32, bool)`

GetAppServiceIdOk returns a tuple with the AppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceId

`func (o *AppBuild) SetAppServiceId(v int32)`

SetAppServiceId sets AppServiceId field to given value.


### GetTaskId

`func (o *AppBuild) GetTaskId() int32`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *AppBuild) GetTaskIdOk() (*int32, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *AppBuild) SetTaskId(v int32)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *AppBuild) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskIdNil

`func (o *AppBuild) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *AppBuild) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
### GetTask

`func (o *AppBuild) GetTask() Task`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *AppBuild) GetTaskOk() (*Task, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *AppBuild) SetTask(v Task)`

SetTask sets Task field to given value.

### HasTask

`func (o *AppBuild) HasTask() bool`

HasTask returns a boolean if a field has been set.

### SetTaskNil

`func (o *AppBuild) SetTaskNil(b bool)`

 SetTaskNil sets the value for Task to be an explicit nil

### UnsetTask
`func (o *AppBuild) UnsetTask()`

UnsetTask ensures that no value is present for Task, not even an explicit nil
### GetAppServiceBuilds

`func (o *AppBuild) GetAppServiceBuilds() []AppServiceBuild`

GetAppServiceBuilds returns the AppServiceBuilds field if non-nil, zero value otherwise.

### GetAppServiceBuildsOk

`func (o *AppBuild) GetAppServiceBuildsOk() (*[]AppServiceBuild, bool)`

GetAppServiceBuildsOk returns a tuple with the AppServiceBuilds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceBuilds

`func (o *AppBuild) SetAppServiceBuilds(v []AppServiceBuild)`

SetAppServiceBuilds sets AppServiceBuilds field to given value.


### GetGitRefType

`func (o *AppBuild) GetGitRefType() string`

GetGitRefType returns the GitRefType field if non-nil, zero value otherwise.

### GetGitRefTypeOk

`func (o *AppBuild) GetGitRefTypeOk() (*string, bool)`

GetGitRefTypeOk returns a tuple with the GitRefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRefType

`func (o *AppBuild) SetGitRefType(v string)`

SetGitRefType sets GitRefType field to given value.


### GetGitRef

`func (o *AppBuild) GetGitRef() string`

GetGitRef returns the GitRef field if non-nil, zero value otherwise.

### GetGitRefOk

`func (o *AppBuild) GetGitRefOk() (*string, bool)`

GetGitRefOk returns a tuple with the GitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRef

`func (o *AppBuild) SetGitRef(v string)`

SetGitRef sets GitRef field to given value.


### GetCommitHash

`func (o *AppBuild) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *AppBuild) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *AppBuild) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.


### GetCommitMessage

`func (o *AppBuild) GetCommitMessage() string`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *AppBuild) GetCommitMessageOk() (*string, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *AppBuild) SetCommitMessage(v string)`

SetCommitMessage sets CommitMessage field to given value.


### GetAuthorName

`func (o *AppBuild) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *AppBuild) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *AppBuild) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.

### HasAuthorName

`func (o *AppBuild) HasAuthorName() bool`

HasAuthorName returns a boolean if a field has been set.

### SetAuthorNameNil

`func (o *AppBuild) SetAuthorNameNil(b bool)`

 SetAuthorNameNil sets the value for AuthorName to be an explicit nil

### UnsetAuthorName
`func (o *AppBuild) UnsetAuthorName()`

UnsetAuthorName ensures that no value is present for AuthorName, not even an explicit nil
### GetAuthorEmail

`func (o *AppBuild) GetAuthorEmail() string`

GetAuthorEmail returns the AuthorEmail field if non-nil, zero value otherwise.

### GetAuthorEmailOk

`func (o *AppBuild) GetAuthorEmailOk() (*string, bool)`

GetAuthorEmailOk returns a tuple with the AuthorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorEmail

`func (o *AppBuild) SetAuthorEmail(v string)`

SetAuthorEmail sets AuthorEmail field to given value.

### HasAuthorEmail

`func (o *AppBuild) HasAuthorEmail() bool`

HasAuthorEmail returns a boolean if a field has been set.

### SetAuthorEmailNil

`func (o *AppBuild) SetAuthorEmailNil(b bool)`

 SetAuthorEmailNil sets the value for AuthorEmail to be an explicit nil

### UnsetAuthorEmail
`func (o *AppBuild) UnsetAuthorEmail()`

UnsetAuthorEmail ensures that no value is present for AuthorEmail, not even an explicit nil
### GetWorkflow

`func (o *AppBuild) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *AppBuild) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *AppBuild) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.


### GetStackRevId

`func (o *AppBuild) GetStackRevId() int32`

GetStackRevId returns the StackRevId field if non-nil, zero value otherwise.

### GetStackRevIdOk

`func (o *AppBuild) GetStackRevIdOk() (*int32, bool)`

GetStackRevIdOk returns a tuple with the StackRevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackRevId

`func (o *AppBuild) SetStackRevId(v int32)`

SetStackRevId sets StackRevId field to given value.

### HasStackRevId

`func (o *AppBuild) HasStackRevId() bool`

HasStackRevId returns a boolean if a field has been set.

### SetStackRevIdNil

`func (o *AppBuild) SetStackRevIdNil(b bool)`

 SetStackRevIdNil sets the value for StackRevId to be an explicit nil

### UnsetStackRevId
`func (o *AppBuild) UnsetStackRevId()`

UnsetStackRevId ensures that no value is present for StackRevId, not even an explicit nil
### GetGitRepoId

`func (o *AppBuild) GetGitRepoId() int32`

GetGitRepoId returns the GitRepoId field if non-nil, zero value otherwise.

### GetGitRepoIdOk

`func (o *AppBuild) GetGitRepoIdOk() (*int32, bool)`

GetGitRepoIdOk returns a tuple with the GitRepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepoId

`func (o *AppBuild) SetGitRepoId(v int32)`

SetGitRepoId sets GitRepoId field to given value.

### HasGitRepoId

`func (o *AppBuild) HasGitRepoId() bool`

HasGitRepoId returns a boolean if a field has been set.

### SetGitRepoIdNil

`func (o *AppBuild) SetGitRepoIdNil(b bool)`

 SetGitRepoIdNil sets the value for GitRepoId to be an explicit nil

### UnsetGitRepoId
`func (o *AppBuild) UnsetGitRepoId()`

UnsetGitRepoId ensures that no value is present for GitRepoId, not even an explicit nil
### GetCiIntegrationId

`func (o *AppBuild) GetCiIntegrationId() int32`

GetCiIntegrationId returns the CiIntegrationId field if non-nil, zero value otherwise.

### GetCiIntegrationIdOk

`func (o *AppBuild) GetCiIntegrationIdOk() (*int32, bool)`

GetCiIntegrationIdOk returns a tuple with the CiIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiIntegrationId

`func (o *AppBuild) SetCiIntegrationId(v int32)`

SetCiIntegrationId sets CiIntegrationId field to given value.

### HasCiIntegrationId

`func (o *AppBuild) HasCiIntegrationId() bool`

HasCiIntegrationId returns a boolean if a field has been set.

### SetCiIntegrationIdNil

`func (o *AppBuild) SetCiIntegrationIdNil(b bool)`

 SetCiIntegrationIdNil sets the value for CiIntegrationId to be an explicit nil

### UnsetCiIntegrationId
`func (o *AppBuild) UnsetCiIntegrationId()`

UnsetCiIntegrationId ensures that no value is present for CiIntegrationId, not even an explicit nil
### GetRegistryIntegrationId

`func (o *AppBuild) GetRegistryIntegrationId() int32`

GetRegistryIntegrationId returns the RegistryIntegrationId field if non-nil, zero value otherwise.

### GetRegistryIntegrationIdOk

`func (o *AppBuild) GetRegistryIntegrationIdOk() (*int32, bool)`

GetRegistryIntegrationIdOk returns a tuple with the RegistryIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryIntegrationId

`func (o *AppBuild) SetRegistryIntegrationId(v int32)`

SetRegistryIntegrationId sets RegistryIntegrationId field to given value.

### HasRegistryIntegrationId

`func (o *AppBuild) HasRegistryIntegrationId() bool`

HasRegistryIntegrationId returns a boolean if a field has been set.

### SetRegistryIntegrationIdNil

`func (o *AppBuild) SetRegistryIntegrationIdNil(b bool)`

 SetRegistryIntegrationIdNil sets the value for RegistryIntegrationId to be an explicit nil

### UnsetRegistryIntegrationId
`func (o *AppBuild) UnsetRegistryIntegrationId()`

UnsetRegistryIntegrationId ensures that no value is present for RegistryIntegrationId, not even an explicit nil
### GetImageVoidStatus

`func (o *AppBuild) GetImageVoidStatus() string`

GetImageVoidStatus returns the ImageVoidStatus field if non-nil, zero value otherwise.

### GetImageVoidStatusOk

`func (o *AppBuild) GetImageVoidStatusOk() (*string, bool)`

GetImageVoidStatusOk returns a tuple with the ImageVoidStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageVoidStatus

`func (o *AppBuild) SetImageVoidStatus(v string)`

SetImageVoidStatus sets ImageVoidStatus field to given value.


### GetCanVoidImages

`func (o *AppBuild) GetCanVoidImages() bool`

GetCanVoidImages returns the CanVoidImages field if non-nil, zero value otherwise.

### GetCanVoidImagesOk

`func (o *AppBuild) GetCanVoidImagesOk() (*bool, bool)`

GetCanVoidImagesOk returns a tuple with the CanVoidImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanVoidImages

`func (o *AppBuild) SetCanVoidImages(v bool)`

SetCanVoidImages sets CanVoidImages field to given value.


### GetCreatedAt

`func (o *AppBuild) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppBuild) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppBuild) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AppBuild) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AppBuild) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AppBuild) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetStartedAt

`func (o *AppBuild) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *AppBuild) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *AppBuild) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *AppBuild) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *AppBuild) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *AppBuild) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetEndedAt

`func (o *AppBuild) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *AppBuild) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *AppBuild) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *AppBuild) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### SetEndedAtNil

`func (o *AppBuild) SetEndedAtNil(b bool)`

 SetEndedAtNil sets the value for EndedAt to be an explicit nil

### UnsetEndedAt
`func (o *AppBuild) UnsetEndedAt()`

UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


