# AutomationTimeWindowInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Set to false to remove the configured time window. | 
**Start** | Pointer to **NullableString** |  | [optional] 
**End** | Pointer to **NullableString** |  | [optional] 
**TimeZone** | Pointer to **NullableString** |  | [optional] [default to "UTC"]
**Days** | Pointer to **[]string** | Selected weekdays. Omit for every day. Overnight windows use the day on which the window starts. | [optional] 

## Methods

### NewAutomationTimeWindowInput

`func NewAutomationTimeWindowInput(enabled bool, ) *AutomationTimeWindowInput`

NewAutomationTimeWindowInput instantiates a new AutomationTimeWindowInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationTimeWindowInputWithDefaults

`func NewAutomationTimeWindowInputWithDefaults() *AutomationTimeWindowInput`

NewAutomationTimeWindowInputWithDefaults instantiates a new AutomationTimeWindowInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AutomationTimeWindowInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AutomationTimeWindowInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AutomationTimeWindowInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetStart

`func (o *AutomationTimeWindowInput) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *AutomationTimeWindowInput) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *AutomationTimeWindowInput) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *AutomationTimeWindowInput) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStartNil

`func (o *AutomationTimeWindowInput) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *AutomationTimeWindowInput) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil
### GetEnd

`func (o *AutomationTimeWindowInput) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *AutomationTimeWindowInput) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *AutomationTimeWindowInput) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *AutomationTimeWindowInput) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *AutomationTimeWindowInput) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *AutomationTimeWindowInput) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetTimeZone

`func (o *AutomationTimeWindowInput) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *AutomationTimeWindowInput) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *AutomationTimeWindowInput) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *AutomationTimeWindowInput) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### SetTimeZoneNil

`func (o *AutomationTimeWindowInput) SetTimeZoneNil(b bool)`

 SetTimeZoneNil sets the value for TimeZone to be an explicit nil

### UnsetTimeZone
`func (o *AutomationTimeWindowInput) UnsetTimeZone()`

UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil
### GetDays

`func (o *AutomationTimeWindowInput) GetDays() []string`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *AutomationTimeWindowInput) GetDaysOk() (*[]string, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *AutomationTimeWindowInput) SetDays(v []string)`

SetDays sets Days field to given value.

### HasDays

`func (o *AutomationTimeWindowInput) HasDays() bool`

HasDays returns a boolean if a field has been set.

### SetDaysNil

`func (o *AutomationTimeWindowInput) SetDaysNil(b bool)`

 SetDaysNil sets the value for Days to be an explicit nil

### UnsetDays
`func (o *AutomationTimeWindowInput) UnsetDays()`

UnsetDays ensures that no value is present for Days, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


