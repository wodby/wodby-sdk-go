# AppServiceCronSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**AppServiceId** | **int32** |  | 
**Title** | **string** |  | 
**Crontab** | **string** |  | 
**Command** | **string** |  | 
**Workload** | Pointer to **NullableString** |  | [optional] 
**Disabled** | **bool** |  | 
**EnvType** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewAppServiceCronSchedule

`func NewAppServiceCronSchedule(id int32, appServiceId int32, title string, crontab string, command string, disabled bool, createdAt time.Time, updatedAt time.Time, ) *AppServiceCronSchedule`

NewAppServiceCronSchedule instantiates a new AppServiceCronSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceCronScheduleWithDefaults

`func NewAppServiceCronScheduleWithDefaults() *AppServiceCronSchedule`

NewAppServiceCronScheduleWithDefaults instantiates a new AppServiceCronSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppServiceCronSchedule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppServiceCronSchedule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppServiceCronSchedule) SetId(v int32)`

SetId sets Id field to given value.


### GetAppServiceId

`func (o *AppServiceCronSchedule) GetAppServiceId() int32`

GetAppServiceId returns the AppServiceId field if non-nil, zero value otherwise.

### GetAppServiceIdOk

`func (o *AppServiceCronSchedule) GetAppServiceIdOk() (*int32, bool)`

GetAppServiceIdOk returns a tuple with the AppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceId

`func (o *AppServiceCronSchedule) SetAppServiceId(v int32)`

SetAppServiceId sets AppServiceId field to given value.


### GetTitle

`func (o *AppServiceCronSchedule) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AppServiceCronSchedule) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AppServiceCronSchedule) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetCrontab

`func (o *AppServiceCronSchedule) GetCrontab() string`

GetCrontab returns the Crontab field if non-nil, zero value otherwise.

### GetCrontabOk

`func (o *AppServiceCronSchedule) GetCrontabOk() (*string, bool)`

GetCrontabOk returns a tuple with the Crontab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrontab

`func (o *AppServiceCronSchedule) SetCrontab(v string)`

SetCrontab sets Crontab field to given value.


### GetCommand

`func (o *AppServiceCronSchedule) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *AppServiceCronSchedule) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *AppServiceCronSchedule) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetWorkload

`func (o *AppServiceCronSchedule) GetWorkload() string`

GetWorkload returns the Workload field if non-nil, zero value otherwise.

### GetWorkloadOk

`func (o *AppServiceCronSchedule) GetWorkloadOk() (*string, bool)`

GetWorkloadOk returns a tuple with the Workload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkload

`func (o *AppServiceCronSchedule) SetWorkload(v string)`

SetWorkload sets Workload field to given value.

### HasWorkload

`func (o *AppServiceCronSchedule) HasWorkload() bool`

HasWorkload returns a boolean if a field has been set.

### SetWorkloadNil

`func (o *AppServiceCronSchedule) SetWorkloadNil(b bool)`

 SetWorkloadNil sets the value for Workload to be an explicit nil

### UnsetWorkload
`func (o *AppServiceCronSchedule) UnsetWorkload()`

UnsetWorkload ensures that no value is present for Workload, not even an explicit nil
### GetDisabled

`func (o *AppServiceCronSchedule) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *AppServiceCronSchedule) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *AppServiceCronSchedule) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetEnvType

`func (o *AppServiceCronSchedule) GetEnvType() string`

GetEnvType returns the EnvType field if non-nil, zero value otherwise.

### GetEnvTypeOk

`func (o *AppServiceCronSchedule) GetEnvTypeOk() (*string, bool)`

GetEnvTypeOk returns a tuple with the EnvType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvType

`func (o *AppServiceCronSchedule) SetEnvType(v string)`

SetEnvType sets EnvType field to given value.

### HasEnvType

`func (o *AppServiceCronSchedule) HasEnvType() bool`

HasEnvType returns a boolean if a field has been set.

### SetEnvTypeNil

`func (o *AppServiceCronSchedule) SetEnvTypeNil(b bool)`

 SetEnvTypeNil sets the value for EnvType to be an explicit nil

### UnsetEnvType
`func (o *AppServiceCronSchedule) UnsetEnvType()`

UnsetEnvType ensures that no value is present for EnvType, not even an explicit nil
### GetCreatedAt

`func (o *AppServiceCronSchedule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppServiceCronSchedule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppServiceCronSchedule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AppServiceCronSchedule) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AppServiceCronSchedule) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AppServiceCronSchedule) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


