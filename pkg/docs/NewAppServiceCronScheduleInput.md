# NewAppServiceCronScheduleInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Stable cron schedule identity. When omitted or blank, the server generates a unique name. | [optional] 
**Title** | **string** |  | 
**Crontab** | **string** |  | 
**Command** | **string** |  | 
**Workload** | Pointer to **NullableString** |  | [optional] 
**Disabled** | Pointer to **NullableBool** | Creates the schedule disabled. Disabled schedules do not require cron feature access. | [optional] 

## Methods

### NewNewAppServiceCronScheduleInput

`func NewNewAppServiceCronScheduleInput(title string, crontab string, command string, ) *NewAppServiceCronScheduleInput`

NewNewAppServiceCronScheduleInput instantiates a new NewAppServiceCronScheduleInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewAppServiceCronScheduleInputWithDefaults

`func NewNewAppServiceCronScheduleInputWithDefaults() *NewAppServiceCronScheduleInput`

NewNewAppServiceCronScheduleInputWithDefaults instantiates a new NewAppServiceCronScheduleInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NewAppServiceCronScheduleInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewAppServiceCronScheduleInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewAppServiceCronScheduleInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NewAppServiceCronScheduleInput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *NewAppServiceCronScheduleInput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *NewAppServiceCronScheduleInput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTitle

`func (o *NewAppServiceCronScheduleInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewAppServiceCronScheduleInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NewAppServiceCronScheduleInput) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetCrontab

`func (o *NewAppServiceCronScheduleInput) GetCrontab() string`

GetCrontab returns the Crontab field if non-nil, zero value otherwise.

### GetCrontabOk

`func (o *NewAppServiceCronScheduleInput) GetCrontabOk() (*string, bool)`

GetCrontabOk returns a tuple with the Crontab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrontab

`func (o *NewAppServiceCronScheduleInput) SetCrontab(v string)`

SetCrontab sets Crontab field to given value.


### GetCommand

`func (o *NewAppServiceCronScheduleInput) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *NewAppServiceCronScheduleInput) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *NewAppServiceCronScheduleInput) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetWorkload

`func (o *NewAppServiceCronScheduleInput) GetWorkload() string`

GetWorkload returns the Workload field if non-nil, zero value otherwise.

### GetWorkloadOk

`func (o *NewAppServiceCronScheduleInput) GetWorkloadOk() (*string, bool)`

GetWorkloadOk returns a tuple with the Workload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkload

`func (o *NewAppServiceCronScheduleInput) SetWorkload(v string)`

SetWorkload sets Workload field to given value.

### HasWorkload

`func (o *NewAppServiceCronScheduleInput) HasWorkload() bool`

HasWorkload returns a boolean if a field has been set.

### SetWorkloadNil

`func (o *NewAppServiceCronScheduleInput) SetWorkloadNil(b bool)`

 SetWorkloadNil sets the value for Workload to be an explicit nil

### UnsetWorkload
`func (o *NewAppServiceCronScheduleInput) UnsetWorkload()`

UnsetWorkload ensures that no value is present for Workload, not even an explicit nil
### GetDisabled

`func (o *NewAppServiceCronScheduleInput) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *NewAppServiceCronScheduleInput) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *NewAppServiceCronScheduleInput) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *NewAppServiceCronScheduleInput) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### SetDisabledNil

`func (o *NewAppServiceCronScheduleInput) SetDisabledNil(b bool)`

 SetDisabledNil sets the value for Disabled to be an explicit nil

### UnsetDisabled
`func (o *NewAppServiceCronScheduleInput) UnsetDisabled()`

UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


