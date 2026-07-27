# AppInstanceHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cron** | [**AppInstanceCronHealth**](AppInstanceCronHealth.md) |  | 
**Backups** | [**AppInstanceBackupHealth**](AppInstanceBackupHealth.md) |  | 

## Methods

### NewAppInstanceHealth

`func NewAppInstanceHealth(cron AppInstanceCronHealth, backups AppInstanceBackupHealth, ) *AppInstanceHealth`

NewAppInstanceHealth instantiates a new AppInstanceHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppInstanceHealthWithDefaults

`func NewAppInstanceHealthWithDefaults() *AppInstanceHealth`

NewAppInstanceHealthWithDefaults instantiates a new AppInstanceHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCron

`func (o *AppInstanceHealth) GetCron() AppInstanceCronHealth`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *AppInstanceHealth) GetCronOk() (*AppInstanceCronHealth, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *AppInstanceHealth) SetCron(v AppInstanceCronHealth)`

SetCron sets Cron field to given value.


### GetBackups

`func (o *AppInstanceHealth) GetBackups() AppInstanceBackupHealth`

GetBackups returns the Backups field if non-nil, zero value otherwise.

### GetBackupsOk

`func (o *AppInstanceHealth) GetBackupsOk() (*AppInstanceBackupHealth, bool)`

GetBackupsOk returns a tuple with the Backups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackups

`func (o *AppInstanceHealth) SetBackups(v AppInstanceBackupHealth)`

SetBackups sets Backups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


