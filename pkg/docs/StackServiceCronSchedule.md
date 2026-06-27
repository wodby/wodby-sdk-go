# StackServiceCronSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**StackServiceId** | **int32** |  | 
**Name** | **string** |  | 
**Title** | **string** |  | 
**Crontab** | **string** |  | 
**Command** | **string** |  | 
**Workload** | Pointer to **NullableString** |  | [optional] 
**Disabled** | **bool** |  | 
**EnvType** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewStackServiceCronSchedule

`func NewStackServiceCronSchedule(id int32, stackServiceId int32, name string, title string, crontab string, command string, disabled bool, createdAt time.Time, updatedAt time.Time, ) *StackServiceCronSchedule`

NewStackServiceCronSchedule instantiates a new StackServiceCronSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackServiceCronScheduleWithDefaults

`func NewStackServiceCronScheduleWithDefaults() *StackServiceCronSchedule`

NewStackServiceCronScheduleWithDefaults instantiates a new StackServiceCronSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StackServiceCronSchedule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StackServiceCronSchedule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StackServiceCronSchedule) SetId(v int32)`

SetId sets Id field to given value.


### GetStackServiceId

`func (o *StackServiceCronSchedule) GetStackServiceId() int32`

GetStackServiceId returns the StackServiceId field if non-nil, zero value otherwise.

### GetStackServiceIdOk

`func (o *StackServiceCronSchedule) GetStackServiceIdOk() (*int32, bool)`

GetStackServiceIdOk returns a tuple with the StackServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackServiceId

`func (o *StackServiceCronSchedule) SetStackServiceId(v int32)`

SetStackServiceId sets StackServiceId field to given value.


### GetName

`func (o *StackServiceCronSchedule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StackServiceCronSchedule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StackServiceCronSchedule) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *StackServiceCronSchedule) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StackServiceCronSchedule) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StackServiceCronSchedule) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetCrontab

`func (o *StackServiceCronSchedule) GetCrontab() string`

GetCrontab returns the Crontab field if non-nil, zero value otherwise.

### GetCrontabOk

`func (o *StackServiceCronSchedule) GetCrontabOk() (*string, bool)`

GetCrontabOk returns a tuple with the Crontab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrontab

`func (o *StackServiceCronSchedule) SetCrontab(v string)`

SetCrontab sets Crontab field to given value.


### GetCommand

`func (o *StackServiceCronSchedule) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *StackServiceCronSchedule) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *StackServiceCronSchedule) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetWorkload

`func (o *StackServiceCronSchedule) GetWorkload() string`

GetWorkload returns the Workload field if non-nil, zero value otherwise.

### GetWorkloadOk

`func (o *StackServiceCronSchedule) GetWorkloadOk() (*string, bool)`

GetWorkloadOk returns a tuple with the Workload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkload

`func (o *StackServiceCronSchedule) SetWorkload(v string)`

SetWorkload sets Workload field to given value.

### HasWorkload

`func (o *StackServiceCronSchedule) HasWorkload() bool`

HasWorkload returns a boolean if a field has been set.

### SetWorkloadNil

`func (o *StackServiceCronSchedule) SetWorkloadNil(b bool)`

 SetWorkloadNil sets the value for Workload to be an explicit nil

### UnsetWorkload
`func (o *StackServiceCronSchedule) UnsetWorkload()`

UnsetWorkload ensures that no value is present for Workload, not even an explicit nil
### GetDisabled

`func (o *StackServiceCronSchedule) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *StackServiceCronSchedule) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *StackServiceCronSchedule) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetEnvType

`func (o *StackServiceCronSchedule) GetEnvType() string`

GetEnvType returns the EnvType field if non-nil, zero value otherwise.

### GetEnvTypeOk

`func (o *StackServiceCronSchedule) GetEnvTypeOk() (*string, bool)`

GetEnvTypeOk returns a tuple with the EnvType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvType

`func (o *StackServiceCronSchedule) SetEnvType(v string)`

SetEnvType sets EnvType field to given value.

### HasEnvType

`func (o *StackServiceCronSchedule) HasEnvType() bool`

HasEnvType returns a boolean if a field has been set.

### SetEnvTypeNil

`func (o *StackServiceCronSchedule) SetEnvTypeNil(b bool)`

 SetEnvTypeNil sets the value for EnvType to be an explicit nil

### UnsetEnvType
`func (o *StackServiceCronSchedule) UnsetEnvType()`

UnsetEnvType ensures that no value is present for EnvType, not even an explicit nil
### GetCreatedAt

`func (o *StackServiceCronSchedule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StackServiceCronSchedule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StackServiceCronSchedule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *StackServiceCronSchedule) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *StackServiceCronSchedule) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *StackServiceCronSchedule) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


