# AppBuild

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Number** | **int32** | App-instance build-launch number. Build records created by the same Wodby CI launch share this number. | 
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
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**StartedAt** | Pointer to **NullableTime** |  | [optional] 
**EndedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewAppBuild

`func NewAppBuild(id int32, number int32, status string, appInstanceId int32, appServiceId int32, appServiceBuilds []AppServiceBuild, gitRefType string, gitRef string, commitHash string, commitMessage string, createdAt time.Time, updatedAt time.Time, ) *AppBuild`

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


