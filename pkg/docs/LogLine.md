# LogLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SequenceId** | **int32** |  | 
**Level** | **string** |  | 
**Message** | **string** |  | 

## Methods

### NewLogLine

`func NewLogLine(sequenceId int32, level string, message string, ) *LogLine`

NewLogLine instantiates a new LogLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogLineWithDefaults

`func NewLogLineWithDefaults() *LogLine`

NewLogLineWithDefaults instantiates a new LogLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSequenceId

`func (o *LogLine) GetSequenceId() int32`

GetSequenceId returns the SequenceId field if non-nil, zero value otherwise.

### GetSequenceIdOk

`func (o *LogLine) GetSequenceIdOk() (*int32, bool)`

GetSequenceIdOk returns a tuple with the SequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceId

`func (o *LogLine) SetSequenceId(v int32)`

SetSequenceId sets SequenceId field to given value.


### GetLevel

`func (o *LogLine) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *LogLine) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *LogLine) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetMessage

`func (o *LogLine) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LogLine) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LogLine) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


