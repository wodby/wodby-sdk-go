# AppServicePodContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Primary** | **bool** |  | 
**Ready** | **bool** |  | 
**Restarts** | **int32** |  | 
**State** | **string** |  | 
**Reason** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**ExitCode** | Pointer to **NullableInt32** |  | [optional] 
**StartedAt** | Pointer to **NullableTime** |  | [optional] 
**FinishedAt** | Pointer to **NullableTime** |  | [optional] 
**CpuUsage** | Pointer to **NullableFloat32** |  | [optional] 
**MemoryUsage** | Pointer to **NullableInt32** |  | [optional] 
**CpuRequest** | Pointer to **NullableFloat32** |  | [optional] 
**MemoryRequest** | Pointer to **NullableInt32** |  | [optional] 
**CpuLimit** | Pointer to **NullableFloat32** |  | [optional] 
**MemoryLimit** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewAppServicePodContainer

`func NewAppServicePodContainer(id string, name string, primary bool, ready bool, restarts int32, state string, ) *AppServicePodContainer`

NewAppServicePodContainer instantiates a new AppServicePodContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServicePodContainerWithDefaults

`func NewAppServicePodContainerWithDefaults() *AppServicePodContainer`

NewAppServicePodContainerWithDefaults instantiates a new AppServicePodContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppServicePodContainer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppServicePodContainer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppServicePodContainer) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AppServicePodContainer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppServicePodContainer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppServicePodContainer) SetName(v string)`

SetName sets Name field to given value.


### GetPrimary

`func (o *AppServicePodContainer) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *AppServicePodContainer) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *AppServicePodContainer) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.


### GetReady

`func (o *AppServicePodContainer) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *AppServicePodContainer) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *AppServicePodContainer) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetRestarts

`func (o *AppServicePodContainer) GetRestarts() int32`

GetRestarts returns the Restarts field if non-nil, zero value otherwise.

### GetRestartsOk

`func (o *AppServicePodContainer) GetRestartsOk() (*int32, bool)`

GetRestartsOk returns a tuple with the Restarts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestarts

`func (o *AppServicePodContainer) SetRestarts(v int32)`

SetRestarts sets Restarts field to given value.


### GetState

`func (o *AppServicePodContainer) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppServicePodContainer) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppServicePodContainer) SetState(v string)`

SetState sets State field to given value.


### GetReason

`func (o *AppServicePodContainer) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AppServicePodContainer) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AppServicePodContainer) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AppServicePodContainer) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *AppServicePodContainer) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *AppServicePodContainer) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetMessage

`func (o *AppServicePodContainer) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AppServicePodContainer) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AppServicePodContainer) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AppServicePodContainer) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *AppServicePodContainer) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *AppServicePodContainer) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetExitCode

`func (o *AppServicePodContainer) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *AppServicePodContainer) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *AppServicePodContainer) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *AppServicePodContainer) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### SetExitCodeNil

`func (o *AppServicePodContainer) SetExitCodeNil(b bool)`

 SetExitCodeNil sets the value for ExitCode to be an explicit nil

### UnsetExitCode
`func (o *AppServicePodContainer) UnsetExitCode()`

UnsetExitCode ensures that no value is present for ExitCode, not even an explicit nil
### GetStartedAt

`func (o *AppServicePodContainer) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *AppServicePodContainer) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *AppServicePodContainer) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *AppServicePodContainer) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *AppServicePodContainer) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *AppServicePodContainer) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetFinishedAt

`func (o *AppServicePodContainer) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *AppServicePodContainer) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *AppServicePodContainer) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *AppServicePodContainer) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### SetFinishedAtNil

`func (o *AppServicePodContainer) SetFinishedAtNil(b bool)`

 SetFinishedAtNil sets the value for FinishedAt to be an explicit nil

### UnsetFinishedAt
`func (o *AppServicePodContainer) UnsetFinishedAt()`

UnsetFinishedAt ensures that no value is present for FinishedAt, not even an explicit nil
### GetCpuUsage

`func (o *AppServicePodContainer) GetCpuUsage() float32`

GetCpuUsage returns the CpuUsage field if non-nil, zero value otherwise.

### GetCpuUsageOk

`func (o *AppServicePodContainer) GetCpuUsageOk() (*float32, bool)`

GetCpuUsageOk returns a tuple with the CpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsage

`func (o *AppServicePodContainer) SetCpuUsage(v float32)`

SetCpuUsage sets CpuUsage field to given value.

### HasCpuUsage

`func (o *AppServicePodContainer) HasCpuUsage() bool`

HasCpuUsage returns a boolean if a field has been set.

### SetCpuUsageNil

`func (o *AppServicePodContainer) SetCpuUsageNil(b bool)`

 SetCpuUsageNil sets the value for CpuUsage to be an explicit nil

### UnsetCpuUsage
`func (o *AppServicePodContainer) UnsetCpuUsage()`

UnsetCpuUsage ensures that no value is present for CpuUsage, not even an explicit nil
### GetMemoryUsage

`func (o *AppServicePodContainer) GetMemoryUsage() int32`

GetMemoryUsage returns the MemoryUsage field if non-nil, zero value otherwise.

### GetMemoryUsageOk

`func (o *AppServicePodContainer) GetMemoryUsageOk() (*int32, bool)`

GetMemoryUsageOk returns a tuple with the MemoryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsage

`func (o *AppServicePodContainer) SetMemoryUsage(v int32)`

SetMemoryUsage sets MemoryUsage field to given value.

### HasMemoryUsage

`func (o *AppServicePodContainer) HasMemoryUsage() bool`

HasMemoryUsage returns a boolean if a field has been set.

### SetMemoryUsageNil

`func (o *AppServicePodContainer) SetMemoryUsageNil(b bool)`

 SetMemoryUsageNil sets the value for MemoryUsage to be an explicit nil

### UnsetMemoryUsage
`func (o *AppServicePodContainer) UnsetMemoryUsage()`

UnsetMemoryUsage ensures that no value is present for MemoryUsage, not even an explicit nil
### GetCpuRequest

`func (o *AppServicePodContainer) GetCpuRequest() float32`

GetCpuRequest returns the CpuRequest field if non-nil, zero value otherwise.

### GetCpuRequestOk

`func (o *AppServicePodContainer) GetCpuRequestOk() (*float32, bool)`

GetCpuRequestOk returns a tuple with the CpuRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuRequest

`func (o *AppServicePodContainer) SetCpuRequest(v float32)`

SetCpuRequest sets CpuRequest field to given value.

### HasCpuRequest

`func (o *AppServicePodContainer) HasCpuRequest() bool`

HasCpuRequest returns a boolean if a field has been set.

### SetCpuRequestNil

`func (o *AppServicePodContainer) SetCpuRequestNil(b bool)`

 SetCpuRequestNil sets the value for CpuRequest to be an explicit nil

### UnsetCpuRequest
`func (o *AppServicePodContainer) UnsetCpuRequest()`

UnsetCpuRequest ensures that no value is present for CpuRequest, not even an explicit nil
### GetMemoryRequest

`func (o *AppServicePodContainer) GetMemoryRequest() int32`

GetMemoryRequest returns the MemoryRequest field if non-nil, zero value otherwise.

### GetMemoryRequestOk

`func (o *AppServicePodContainer) GetMemoryRequestOk() (*int32, bool)`

GetMemoryRequestOk returns a tuple with the MemoryRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryRequest

`func (o *AppServicePodContainer) SetMemoryRequest(v int32)`

SetMemoryRequest sets MemoryRequest field to given value.

### HasMemoryRequest

`func (o *AppServicePodContainer) HasMemoryRequest() bool`

HasMemoryRequest returns a boolean if a field has been set.

### SetMemoryRequestNil

`func (o *AppServicePodContainer) SetMemoryRequestNil(b bool)`

 SetMemoryRequestNil sets the value for MemoryRequest to be an explicit nil

### UnsetMemoryRequest
`func (o *AppServicePodContainer) UnsetMemoryRequest()`

UnsetMemoryRequest ensures that no value is present for MemoryRequest, not even an explicit nil
### GetCpuLimit

`func (o *AppServicePodContainer) GetCpuLimit() float32`

GetCpuLimit returns the CpuLimit field if non-nil, zero value otherwise.

### GetCpuLimitOk

`func (o *AppServicePodContainer) GetCpuLimitOk() (*float32, bool)`

GetCpuLimitOk returns a tuple with the CpuLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLimit

`func (o *AppServicePodContainer) SetCpuLimit(v float32)`

SetCpuLimit sets CpuLimit field to given value.

### HasCpuLimit

`func (o *AppServicePodContainer) HasCpuLimit() bool`

HasCpuLimit returns a boolean if a field has been set.

### SetCpuLimitNil

`func (o *AppServicePodContainer) SetCpuLimitNil(b bool)`

 SetCpuLimitNil sets the value for CpuLimit to be an explicit nil

### UnsetCpuLimit
`func (o *AppServicePodContainer) UnsetCpuLimit()`

UnsetCpuLimit ensures that no value is present for CpuLimit, not even an explicit nil
### GetMemoryLimit

`func (o *AppServicePodContainer) GetMemoryLimit() int32`

GetMemoryLimit returns the MemoryLimit field if non-nil, zero value otherwise.

### GetMemoryLimitOk

`func (o *AppServicePodContainer) GetMemoryLimitOk() (*int32, bool)`

GetMemoryLimitOk returns a tuple with the MemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimit

`func (o *AppServicePodContainer) SetMemoryLimit(v int32)`

SetMemoryLimit sets MemoryLimit field to given value.

### HasMemoryLimit

`func (o *AppServicePodContainer) HasMemoryLimit() bool`

HasMemoryLimit returns a boolean if a field has been set.

### SetMemoryLimitNil

`func (o *AppServicePodContainer) SetMemoryLimitNil(b bool)`

 SetMemoryLimitNil sets the value for MemoryLimit to be an explicit nil

### UnsetMemoryLimit
`func (o *AppServicePodContainer) UnsetMemoryLimit()`

UnsetMemoryLimit ensures that no value is present for MemoryLimit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


