# AppEnvironmentHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cron** | [**AppEnvironmentCronHealth**](AppEnvironmentCronHealth.md) |  | 
**Backups** | [**AppEnvironmentBackupHealth**](AppEnvironmentBackupHealth.md) |  | 

## Methods

### NewAppEnvironmentHealth

`func NewAppEnvironmentHealth(cron AppEnvironmentCronHealth, backups AppEnvironmentBackupHealth, ) *AppEnvironmentHealth`

NewAppEnvironmentHealth instantiates a new AppEnvironmentHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppEnvironmentHealthWithDefaults

`func NewAppEnvironmentHealthWithDefaults() *AppEnvironmentHealth`

NewAppEnvironmentHealthWithDefaults instantiates a new AppEnvironmentHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCron

`func (o *AppEnvironmentHealth) GetCron() AppEnvironmentCronHealth`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *AppEnvironmentHealth) GetCronOk() (*AppEnvironmentCronHealth, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *AppEnvironmentHealth) SetCron(v AppEnvironmentCronHealth)`

SetCron sets Cron field to given value.


### GetBackups

`func (o *AppEnvironmentHealth) GetBackups() AppEnvironmentBackupHealth`

GetBackups returns the Backups field if non-nil, zero value otherwise.

### GetBackupsOk

`func (o *AppEnvironmentHealth) GetBackupsOk() (*AppEnvironmentBackupHealth, bool)`

GetBackupsOk returns a tuple with the Backups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackups

`func (o *AppEnvironmentHealth) SetBackups(v AppEnvironmentBackupHealth)`

SetBackups sets Backups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


