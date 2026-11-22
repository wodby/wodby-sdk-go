# AppEnvironmentDeploymentSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunScriptsForOlderBuilds** | **bool** | Run repository pre- and post-deployment scripts when a deployment rolls an app service back to an older build. | 
**SkipRollback** | **bool** | Keep a failed rollout in place instead of restoring the previous release, for every deployment. | 

## Methods

### NewAppEnvironmentDeploymentSettings

`func NewAppEnvironmentDeploymentSettings(runScriptsForOlderBuilds bool, skipRollback bool, ) *AppEnvironmentDeploymentSettings`

NewAppEnvironmentDeploymentSettings instantiates a new AppEnvironmentDeploymentSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppEnvironmentDeploymentSettingsWithDefaults

`func NewAppEnvironmentDeploymentSettingsWithDefaults() *AppEnvironmentDeploymentSettings`

NewAppEnvironmentDeploymentSettingsWithDefaults instantiates a new AppEnvironmentDeploymentSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunScriptsForOlderBuilds

`func (o *AppEnvironmentDeploymentSettings) GetRunScriptsForOlderBuilds() bool`

GetRunScriptsForOlderBuilds returns the RunScriptsForOlderBuilds field if non-nil, zero value otherwise.

### GetRunScriptsForOlderBuildsOk

`func (o *AppEnvironmentDeploymentSettings) GetRunScriptsForOlderBuildsOk() (*bool, bool)`

GetRunScriptsForOlderBuildsOk returns a tuple with the RunScriptsForOlderBuilds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunScriptsForOlderBuilds

`func (o *AppEnvironmentDeploymentSettings) SetRunScriptsForOlderBuilds(v bool)`

SetRunScriptsForOlderBuilds sets RunScriptsForOlderBuilds field to given value.


### GetSkipRollback

`func (o *AppEnvironmentDeploymentSettings) GetSkipRollback() bool`

GetSkipRollback returns the SkipRollback field if non-nil, zero value otherwise.

### GetSkipRollbackOk

`func (o *AppEnvironmentDeploymentSettings) GetSkipRollbackOk() (*bool, bool)`

GetSkipRollbackOk returns a tuple with the SkipRollback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipRollback

`func (o *AppEnvironmentDeploymentSettings) SetSkipRollback(v bool)`

SetSkipRollback sets SkipRollback field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


