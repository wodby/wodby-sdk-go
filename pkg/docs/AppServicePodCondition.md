# AppServicePodCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The condition, such as PodScheduled, Initialized, ContainersReady or Ready. | 
**Status** | **string** | False or Unknown. | 
**Reason** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**LastTransitionTime** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewAppServicePodCondition

`func NewAppServicePodCondition(type_ string, status string, ) *AppServicePodCondition`

NewAppServicePodCondition instantiates a new AppServicePodCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServicePodConditionWithDefaults

`func NewAppServicePodConditionWithDefaults() *AppServicePodCondition`

NewAppServicePodConditionWithDefaults instantiates a new AppServicePodCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppServicePodCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppServicePodCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppServicePodCondition) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *AppServicePodCondition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppServicePodCondition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppServicePodCondition) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetReason

`func (o *AppServicePodCondition) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AppServicePodCondition) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AppServicePodCondition) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AppServicePodCondition) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *AppServicePodCondition) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *AppServicePodCondition) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetMessage

`func (o *AppServicePodCondition) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AppServicePodCondition) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AppServicePodCondition) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AppServicePodCondition) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *AppServicePodCondition) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *AppServicePodCondition) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetLastTransitionTime

`func (o *AppServicePodCondition) GetLastTransitionTime() time.Time`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *AppServicePodCondition) GetLastTransitionTimeOk() (*time.Time, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *AppServicePodCondition) SetLastTransitionTime(v time.Time)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *AppServicePodCondition) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### SetLastTransitionTimeNil

`func (o *AppServicePodCondition) SetLastTransitionTimeNil(b bool)`

 SetLastTransitionTimeNil sets the value for LastTransitionTime to be an explicit nil

### UnsetLastTransitionTime
`func (o *AppServicePodCondition) UnsetLastTransitionTime()`

UnsetLastTransitionTime ensures that no value is present for LastTransitionTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


