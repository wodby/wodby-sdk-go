# TaskStepLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**StreamId** | Pointer to **NullableInt32** |  | [optional] 
**Url** | Pointer to **NullableString** | Temporary URL for persisted logs when URL delivery is selected or auto-selected. | [optional] 
**Lines** | [**[]LogLine**](LogLine.md) |  | 

## Methods

### NewTaskStepLogs

`func NewTaskStepLogs(status string, lines []LogLine, ) *TaskStepLogs`

NewTaskStepLogs instantiates a new TaskStepLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskStepLogsWithDefaults

`func NewTaskStepLogsWithDefaults() *TaskStepLogs`

NewTaskStepLogsWithDefaults instantiates a new TaskStepLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *TaskStepLogs) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskStepLogs) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskStepLogs) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStreamId

`func (o *TaskStepLogs) GetStreamId() int32`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *TaskStepLogs) GetStreamIdOk() (*int32, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *TaskStepLogs) SetStreamId(v int32)`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *TaskStepLogs) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.

### SetStreamIdNil

`func (o *TaskStepLogs) SetStreamIdNil(b bool)`

 SetStreamIdNil sets the value for StreamId to be an explicit nil

### UnsetStreamId
`func (o *TaskStepLogs) UnsetStreamId()`

UnsetStreamId ensures that no value is present for StreamId, not even an explicit nil
### GetUrl

`func (o *TaskStepLogs) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TaskStepLogs) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TaskStepLogs) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TaskStepLogs) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TaskStepLogs) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TaskStepLogs) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetLines

`func (o *TaskStepLogs) GetLines() []LogLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *TaskStepLogs) GetLinesOk() (*[]LogLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *TaskStepLogs) SetLines(v []LogLine)`

SetLines sets Lines field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


