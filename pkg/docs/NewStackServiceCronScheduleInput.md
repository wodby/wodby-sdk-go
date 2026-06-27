# NewStackServiceCronScheduleInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Title** | **string** |  | 
**Crontab** | **string** |  | 
**Command** | **string** |  | 
**Workload** | Pointer to **NullableString** |  | [optional] 
**Disabled** | Pointer to **NullableBool** |  | [optional] 
**EnvType** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNewStackServiceCronScheduleInput

`func NewNewStackServiceCronScheduleInput(name string, title string, crontab string, command string, ) *NewStackServiceCronScheduleInput`

NewNewStackServiceCronScheduleInput instantiates a new NewStackServiceCronScheduleInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewStackServiceCronScheduleInputWithDefaults

`func NewNewStackServiceCronScheduleInputWithDefaults() *NewStackServiceCronScheduleInput`

NewNewStackServiceCronScheduleInputWithDefaults instantiates a new NewStackServiceCronScheduleInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NewStackServiceCronScheduleInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewStackServiceCronScheduleInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewStackServiceCronScheduleInput) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *NewStackServiceCronScheduleInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewStackServiceCronScheduleInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NewStackServiceCronScheduleInput) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetCrontab

`func (o *NewStackServiceCronScheduleInput) GetCrontab() string`

GetCrontab returns the Crontab field if non-nil, zero value otherwise.

### GetCrontabOk

`func (o *NewStackServiceCronScheduleInput) GetCrontabOk() (*string, bool)`

GetCrontabOk returns a tuple with the Crontab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrontab

`func (o *NewStackServiceCronScheduleInput) SetCrontab(v string)`

SetCrontab sets Crontab field to given value.


### GetCommand

`func (o *NewStackServiceCronScheduleInput) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *NewStackServiceCronScheduleInput) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *NewStackServiceCronScheduleInput) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetWorkload

`func (o *NewStackServiceCronScheduleInput) GetWorkload() string`

GetWorkload returns the Workload field if non-nil, zero value otherwise.

### GetWorkloadOk

`func (o *NewStackServiceCronScheduleInput) GetWorkloadOk() (*string, bool)`

GetWorkloadOk returns a tuple with the Workload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkload

`func (o *NewStackServiceCronScheduleInput) SetWorkload(v string)`

SetWorkload sets Workload field to given value.

### HasWorkload

`func (o *NewStackServiceCronScheduleInput) HasWorkload() bool`

HasWorkload returns a boolean if a field has been set.

### SetWorkloadNil

`func (o *NewStackServiceCronScheduleInput) SetWorkloadNil(b bool)`

 SetWorkloadNil sets the value for Workload to be an explicit nil

### UnsetWorkload
`func (o *NewStackServiceCronScheduleInput) UnsetWorkload()`

UnsetWorkload ensures that no value is present for Workload, not even an explicit nil
### GetDisabled

`func (o *NewStackServiceCronScheduleInput) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *NewStackServiceCronScheduleInput) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *NewStackServiceCronScheduleInput) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *NewStackServiceCronScheduleInput) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### SetDisabledNil

`func (o *NewStackServiceCronScheduleInput) SetDisabledNil(b bool)`

 SetDisabledNil sets the value for Disabled to be an explicit nil

### UnsetDisabled
`func (o *NewStackServiceCronScheduleInput) UnsetDisabled()`

UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
### GetEnvType

`func (o *NewStackServiceCronScheduleInput) GetEnvType() string`

GetEnvType returns the EnvType field if non-nil, zero value otherwise.

### GetEnvTypeOk

`func (o *NewStackServiceCronScheduleInput) GetEnvTypeOk() (*string, bool)`

GetEnvTypeOk returns a tuple with the EnvType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvType

`func (o *NewStackServiceCronScheduleInput) SetEnvType(v string)`

SetEnvType sets EnvType field to given value.

### HasEnvType

`func (o *NewStackServiceCronScheduleInput) HasEnvType() bool`

HasEnvType returns a boolean if a field has been set.

### SetEnvTypeNil

`func (o *NewStackServiceCronScheduleInput) SetEnvTypeNil(b bool)`

 SetEnvTypeNil sets the value for EnvType to be an explicit nil

### UnsetEnvType
`func (o *NewStackServiceCronScheduleInput) UnsetEnvType()`

UnsetEnvType ensures that no value is present for EnvType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


