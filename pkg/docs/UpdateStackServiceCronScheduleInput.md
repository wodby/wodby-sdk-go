# UpdateStackServiceCronScheduleInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **NullableBool** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Crontab** | Pointer to **NullableString** |  | [optional] 
**Command** | Pointer to **NullableString** |  | [optional] 
**Workload** | Pointer to **NullableString** |  | [optional] 
**EnvType** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateStackServiceCronScheduleInput

`func NewUpdateStackServiceCronScheduleInput() *UpdateStackServiceCronScheduleInput`

NewUpdateStackServiceCronScheduleInput instantiates a new UpdateStackServiceCronScheduleInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStackServiceCronScheduleInputWithDefaults

`func NewUpdateStackServiceCronScheduleInputWithDefaults() *UpdateStackServiceCronScheduleInput`

NewUpdateStackServiceCronScheduleInputWithDefaults instantiates a new UpdateStackServiceCronScheduleInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *UpdateStackServiceCronScheduleInput) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *UpdateStackServiceCronScheduleInput) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *UpdateStackServiceCronScheduleInput) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *UpdateStackServiceCronScheduleInput) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### SetDisabledNil

`func (o *UpdateStackServiceCronScheduleInput) SetDisabledNil(b bool)`

 SetDisabledNil sets the value for Disabled to be an explicit nil

### UnsetDisabled
`func (o *UpdateStackServiceCronScheduleInput) UnsetDisabled()`

UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
### GetTitle

`func (o *UpdateStackServiceCronScheduleInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateStackServiceCronScheduleInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateStackServiceCronScheduleInput) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateStackServiceCronScheduleInput) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *UpdateStackServiceCronScheduleInput) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *UpdateStackServiceCronScheduleInput) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetCrontab

`func (o *UpdateStackServiceCronScheduleInput) GetCrontab() string`

GetCrontab returns the Crontab field if non-nil, zero value otherwise.

### GetCrontabOk

`func (o *UpdateStackServiceCronScheduleInput) GetCrontabOk() (*string, bool)`

GetCrontabOk returns a tuple with the Crontab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrontab

`func (o *UpdateStackServiceCronScheduleInput) SetCrontab(v string)`

SetCrontab sets Crontab field to given value.

### HasCrontab

`func (o *UpdateStackServiceCronScheduleInput) HasCrontab() bool`

HasCrontab returns a boolean if a field has been set.

### SetCrontabNil

`func (o *UpdateStackServiceCronScheduleInput) SetCrontabNil(b bool)`

 SetCrontabNil sets the value for Crontab to be an explicit nil

### UnsetCrontab
`func (o *UpdateStackServiceCronScheduleInput) UnsetCrontab()`

UnsetCrontab ensures that no value is present for Crontab, not even an explicit nil
### GetCommand

`func (o *UpdateStackServiceCronScheduleInput) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *UpdateStackServiceCronScheduleInput) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *UpdateStackServiceCronScheduleInput) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *UpdateStackServiceCronScheduleInput) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### SetCommandNil

`func (o *UpdateStackServiceCronScheduleInput) SetCommandNil(b bool)`

 SetCommandNil sets the value for Command to be an explicit nil

### UnsetCommand
`func (o *UpdateStackServiceCronScheduleInput) UnsetCommand()`

UnsetCommand ensures that no value is present for Command, not even an explicit nil
### GetWorkload

`func (o *UpdateStackServiceCronScheduleInput) GetWorkload() string`

GetWorkload returns the Workload field if non-nil, zero value otherwise.

### GetWorkloadOk

`func (o *UpdateStackServiceCronScheduleInput) GetWorkloadOk() (*string, bool)`

GetWorkloadOk returns a tuple with the Workload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkload

`func (o *UpdateStackServiceCronScheduleInput) SetWorkload(v string)`

SetWorkload sets Workload field to given value.

### HasWorkload

`func (o *UpdateStackServiceCronScheduleInput) HasWorkload() bool`

HasWorkload returns a boolean if a field has been set.

### SetWorkloadNil

`func (o *UpdateStackServiceCronScheduleInput) SetWorkloadNil(b bool)`

 SetWorkloadNil sets the value for Workload to be an explicit nil

### UnsetWorkload
`func (o *UpdateStackServiceCronScheduleInput) UnsetWorkload()`

UnsetWorkload ensures that no value is present for Workload, not even an explicit nil
### GetEnvType

`func (o *UpdateStackServiceCronScheduleInput) GetEnvType() string`

GetEnvType returns the EnvType field if non-nil, zero value otherwise.

### GetEnvTypeOk

`func (o *UpdateStackServiceCronScheduleInput) GetEnvTypeOk() (*string, bool)`

GetEnvTypeOk returns a tuple with the EnvType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvType

`func (o *UpdateStackServiceCronScheduleInput) SetEnvType(v string)`

SetEnvType sets EnvType field to given value.

### HasEnvType

`func (o *UpdateStackServiceCronScheduleInput) HasEnvType() bool`

HasEnvType returns a boolean if a field has been set.

### SetEnvTypeNil

`func (o *UpdateStackServiceCronScheduleInput) SetEnvTypeNil(b bool)`

 SetEnvTypeNil sets the value for EnvType to be an explicit nil

### UnsetEnvType
`func (o *UpdateStackServiceCronScheduleInput) UnsetEnvType()`

UnsetEnvType ensures that no value is present for EnvType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


