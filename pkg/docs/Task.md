# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Title** | **string** |  | 
**ExecutionScope** | **string** |  | 
**Status** | **string** |  | 
**Progress** | **int32** |  | 
**Silent** | **bool** |  | 
**System** | **bool** |  | 
**UserId** | **int32** |  | 
**User** | Pointer to [**NullableUser**](User.md) |  | [optional] 
**OrgId** | Pointer to **NullableInt32** |  | [optional] 
**ProjectIds** | Pointer to **[]int32** |  | [optional] 
**AppId** | Pointer to **NullableInt32** |  | [optional] 
**AppInstanceId** | Pointer to **NullableInt32** |  | [optional] 
**ClusterId** | Pointer to **NullableInt32** |  | [optional] 
**IntegrationId** | Pointer to **NullableInt32** |  | [optional] 
**ServiceId** | Pointer to **NullableInt32** |  | [optional] 
**StackId** | Pointer to **NullableInt32** |  | [optional] 
**ProviderId** | Pointer to **NullableInt32** |  | [optional] 
**OriginTaskId** | Pointer to **NullableInt32** |  | [optional] 
**SpawnedTaskIds** | Pointer to **[]int32** |  | [optional] 
**RepeatedTaskId** | Pointer to **NullableInt32** |  | [optional] 
**Jobs** | [**[]TaskJob**](TaskJob.md) |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**StartedAt** | Pointer to **NullableTime** |  | [optional] 
**EndedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewTask

`func NewTask(id int32, name string, title string, executionScope string, status string, progress int32, silent bool, system bool, userId int32, jobs []TaskJob, createdAt time.Time, updatedAt time.Time, ) *Task`

NewTask instantiates a new Task object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWithDefaults

`func NewTaskWithDefaults() *Task`

NewTaskWithDefaults instantiates a new Task object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Task) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Task) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Task) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Task) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Task) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Task) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *Task) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Task) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Task) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetExecutionScope

`func (o *Task) GetExecutionScope() string`

GetExecutionScope returns the ExecutionScope field if non-nil, zero value otherwise.

### GetExecutionScopeOk

`func (o *Task) GetExecutionScopeOk() (*string, bool)`

GetExecutionScopeOk returns a tuple with the ExecutionScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionScope

`func (o *Task) SetExecutionScope(v string)`

SetExecutionScope sets ExecutionScope field to given value.


### GetStatus

`func (o *Task) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Task) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Task) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetProgress

`func (o *Task) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *Task) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *Task) SetProgress(v int32)`

SetProgress sets Progress field to given value.


### GetSilent

`func (o *Task) GetSilent() bool`

GetSilent returns the Silent field if non-nil, zero value otherwise.

### GetSilentOk

`func (o *Task) GetSilentOk() (*bool, bool)`

GetSilentOk returns a tuple with the Silent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSilent

`func (o *Task) SetSilent(v bool)`

SetSilent sets Silent field to given value.


### GetSystem

`func (o *Task) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *Task) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *Task) SetSystem(v bool)`

SetSystem sets System field to given value.


### GetUserId

`func (o *Task) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Task) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Task) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetUser

`func (o *Task) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Task) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Task) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *Task) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *Task) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *Task) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetOrgId

`func (o *Task) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Task) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Task) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Task) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *Task) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *Task) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
### GetProjectIds

`func (o *Task) GetProjectIds() []int32`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *Task) GetProjectIdsOk() (*[]int32, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *Task) SetProjectIds(v []int32)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *Task) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### GetAppId

`func (o *Task) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *Task) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *Task) SetAppId(v int32)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *Task) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### SetAppIdNil

`func (o *Task) SetAppIdNil(b bool)`

 SetAppIdNil sets the value for AppId to be an explicit nil

### UnsetAppId
`func (o *Task) UnsetAppId()`

UnsetAppId ensures that no value is present for AppId, not even an explicit nil
### GetAppInstanceId

`func (o *Task) GetAppInstanceId() int32`

GetAppInstanceId returns the AppInstanceId field if non-nil, zero value otherwise.

### GetAppInstanceIdOk

`func (o *Task) GetAppInstanceIdOk() (*int32, bool)`

GetAppInstanceIdOk returns a tuple with the AppInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceId

`func (o *Task) SetAppInstanceId(v int32)`

SetAppInstanceId sets AppInstanceId field to given value.

### HasAppInstanceId

`func (o *Task) HasAppInstanceId() bool`

HasAppInstanceId returns a boolean if a field has been set.

### SetAppInstanceIdNil

`func (o *Task) SetAppInstanceIdNil(b bool)`

 SetAppInstanceIdNil sets the value for AppInstanceId to be an explicit nil

### UnsetAppInstanceId
`func (o *Task) UnsetAppInstanceId()`

UnsetAppInstanceId ensures that no value is present for AppInstanceId, not even an explicit nil
### GetClusterId

`func (o *Task) GetClusterId() int32`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *Task) GetClusterIdOk() (*int32, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *Task) SetClusterId(v int32)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *Task) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *Task) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *Task) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetIntegrationId

`func (o *Task) GetIntegrationId() int32`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *Task) GetIntegrationIdOk() (*int32, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *Task) SetIntegrationId(v int32)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *Task) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### SetIntegrationIdNil

`func (o *Task) SetIntegrationIdNil(b bool)`

 SetIntegrationIdNil sets the value for IntegrationId to be an explicit nil

### UnsetIntegrationId
`func (o *Task) UnsetIntegrationId()`

UnsetIntegrationId ensures that no value is present for IntegrationId, not even an explicit nil
### GetServiceId

`func (o *Task) GetServiceId() int32`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *Task) GetServiceIdOk() (*int32, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *Task) SetServiceId(v int32)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *Task) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### SetServiceIdNil

`func (o *Task) SetServiceIdNil(b bool)`

 SetServiceIdNil sets the value for ServiceId to be an explicit nil

### UnsetServiceId
`func (o *Task) UnsetServiceId()`

UnsetServiceId ensures that no value is present for ServiceId, not even an explicit nil
### GetStackId

`func (o *Task) GetStackId() int32`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *Task) GetStackIdOk() (*int32, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *Task) SetStackId(v int32)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *Task) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### SetStackIdNil

`func (o *Task) SetStackIdNil(b bool)`

 SetStackIdNil sets the value for StackId to be an explicit nil

### UnsetStackId
`func (o *Task) UnsetStackId()`

UnsetStackId ensures that no value is present for StackId, not even an explicit nil
### GetProviderId

`func (o *Task) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *Task) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *Task) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *Task) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### SetProviderIdNil

`func (o *Task) SetProviderIdNil(b bool)`

 SetProviderIdNil sets the value for ProviderId to be an explicit nil

### UnsetProviderId
`func (o *Task) UnsetProviderId()`

UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
### GetOriginTaskId

`func (o *Task) GetOriginTaskId() int32`

GetOriginTaskId returns the OriginTaskId field if non-nil, zero value otherwise.

### GetOriginTaskIdOk

`func (o *Task) GetOriginTaskIdOk() (*int32, bool)`

GetOriginTaskIdOk returns a tuple with the OriginTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginTaskId

`func (o *Task) SetOriginTaskId(v int32)`

SetOriginTaskId sets OriginTaskId field to given value.

### HasOriginTaskId

`func (o *Task) HasOriginTaskId() bool`

HasOriginTaskId returns a boolean if a field has been set.

### SetOriginTaskIdNil

`func (o *Task) SetOriginTaskIdNil(b bool)`

 SetOriginTaskIdNil sets the value for OriginTaskId to be an explicit nil

### UnsetOriginTaskId
`func (o *Task) UnsetOriginTaskId()`

UnsetOriginTaskId ensures that no value is present for OriginTaskId, not even an explicit nil
### GetSpawnedTaskIds

`func (o *Task) GetSpawnedTaskIds() []int32`

GetSpawnedTaskIds returns the SpawnedTaskIds field if non-nil, zero value otherwise.

### GetSpawnedTaskIdsOk

`func (o *Task) GetSpawnedTaskIdsOk() (*[]int32, bool)`

GetSpawnedTaskIdsOk returns a tuple with the SpawnedTaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpawnedTaskIds

`func (o *Task) SetSpawnedTaskIds(v []int32)`

SetSpawnedTaskIds sets SpawnedTaskIds field to given value.

### HasSpawnedTaskIds

`func (o *Task) HasSpawnedTaskIds() bool`

HasSpawnedTaskIds returns a boolean if a field has been set.

### GetRepeatedTaskId

`func (o *Task) GetRepeatedTaskId() int32`

GetRepeatedTaskId returns the RepeatedTaskId field if non-nil, zero value otherwise.

### GetRepeatedTaskIdOk

`func (o *Task) GetRepeatedTaskIdOk() (*int32, bool)`

GetRepeatedTaskIdOk returns a tuple with the RepeatedTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatedTaskId

`func (o *Task) SetRepeatedTaskId(v int32)`

SetRepeatedTaskId sets RepeatedTaskId field to given value.

### HasRepeatedTaskId

`func (o *Task) HasRepeatedTaskId() bool`

HasRepeatedTaskId returns a boolean if a field has been set.

### SetRepeatedTaskIdNil

`func (o *Task) SetRepeatedTaskIdNil(b bool)`

 SetRepeatedTaskIdNil sets the value for RepeatedTaskId to be an explicit nil

### UnsetRepeatedTaskId
`func (o *Task) UnsetRepeatedTaskId()`

UnsetRepeatedTaskId ensures that no value is present for RepeatedTaskId, not even an explicit nil
### GetJobs

`func (o *Task) GetJobs() []TaskJob`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *Task) GetJobsOk() (*[]TaskJob, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *Task) SetJobs(v []TaskJob)`

SetJobs sets Jobs field to given value.


### GetCreatedAt

`func (o *Task) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Task) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Task) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Task) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Task) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Task) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetStartedAt

`func (o *Task) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Task) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Task) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *Task) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *Task) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *Task) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetEndedAt

`func (o *Task) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *Task) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *Task) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *Task) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### SetEndedAtNil

`func (o *Task) SetEndedAtNil(b bool)`

 SetEndedAtNil sets the value for EndedAt to be an explicit nil

### UnsetEndedAt
`func (o *Task) UnsetEndedAt()`

UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


