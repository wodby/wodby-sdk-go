# AppEnvironmentStackUpgradeInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deployment** | Pointer to **bool** | Build affected services when required and deploy the upgraded stack configuration. | [optional] [default to true]
**Versions** | **bool** |  | 
**Replicas** | **bool** |  | 
**Resources** | **bool** |  | 
**Integrations** | **bool** |  | 
**Services** | **bool** |  | 
**Settings** | **bool** |  | 
**Links** | **bool** |  | 
**Tokens** | **bool** |  | 
**Configs** | **bool** |  | 
**Cron** | **bool** |  | 
**Volumes** | **bool** |  | 
**Main** | **bool** |  | 

## Methods

### NewAppEnvironmentStackUpgradeInput

`func NewAppEnvironmentStackUpgradeInput(versions bool, replicas bool, resources bool, integrations bool, services bool, settings bool, links bool, tokens bool, configs bool, cron bool, volumes bool, main bool, ) *AppEnvironmentStackUpgradeInput`

NewAppEnvironmentStackUpgradeInput instantiates a new AppEnvironmentStackUpgradeInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppEnvironmentStackUpgradeInputWithDefaults

`func NewAppEnvironmentStackUpgradeInputWithDefaults() *AppEnvironmentStackUpgradeInput`

NewAppEnvironmentStackUpgradeInputWithDefaults instantiates a new AppEnvironmentStackUpgradeInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeployment

`func (o *AppEnvironmentStackUpgradeInput) GetDeployment() bool`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *AppEnvironmentStackUpgradeInput) GetDeploymentOk() (*bool, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *AppEnvironmentStackUpgradeInput) SetDeployment(v bool)`

SetDeployment sets Deployment field to given value.

### HasDeployment

`func (o *AppEnvironmentStackUpgradeInput) HasDeployment() bool`

HasDeployment returns a boolean if a field has been set.

### GetVersions

`func (o *AppEnvironmentStackUpgradeInput) GetVersions() bool`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *AppEnvironmentStackUpgradeInput) GetVersionsOk() (*bool, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *AppEnvironmentStackUpgradeInput) SetVersions(v bool)`

SetVersions sets Versions field to given value.


### GetReplicas

`func (o *AppEnvironmentStackUpgradeInput) GetReplicas() bool`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *AppEnvironmentStackUpgradeInput) GetReplicasOk() (*bool, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *AppEnvironmentStackUpgradeInput) SetReplicas(v bool)`

SetReplicas sets Replicas field to given value.


### GetResources

`func (o *AppEnvironmentStackUpgradeInput) GetResources() bool`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AppEnvironmentStackUpgradeInput) GetResourcesOk() (*bool, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AppEnvironmentStackUpgradeInput) SetResources(v bool)`

SetResources sets Resources field to given value.


### GetIntegrations

`func (o *AppEnvironmentStackUpgradeInput) GetIntegrations() bool`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *AppEnvironmentStackUpgradeInput) GetIntegrationsOk() (*bool, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *AppEnvironmentStackUpgradeInput) SetIntegrations(v bool)`

SetIntegrations sets Integrations field to given value.


### GetServices

`func (o *AppEnvironmentStackUpgradeInput) GetServices() bool`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *AppEnvironmentStackUpgradeInput) GetServicesOk() (*bool, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *AppEnvironmentStackUpgradeInput) SetServices(v bool)`

SetServices sets Services field to given value.


### GetSettings

`func (o *AppEnvironmentStackUpgradeInput) GetSettings() bool`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *AppEnvironmentStackUpgradeInput) GetSettingsOk() (*bool, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *AppEnvironmentStackUpgradeInput) SetSettings(v bool)`

SetSettings sets Settings field to given value.


### GetLinks

`func (o *AppEnvironmentStackUpgradeInput) GetLinks() bool`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppEnvironmentStackUpgradeInput) GetLinksOk() (*bool, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppEnvironmentStackUpgradeInput) SetLinks(v bool)`

SetLinks sets Links field to given value.


### GetTokens

`func (o *AppEnvironmentStackUpgradeInput) GetTokens() bool`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *AppEnvironmentStackUpgradeInput) GetTokensOk() (*bool, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *AppEnvironmentStackUpgradeInput) SetTokens(v bool)`

SetTokens sets Tokens field to given value.


### GetConfigs

`func (o *AppEnvironmentStackUpgradeInput) GetConfigs() bool`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *AppEnvironmentStackUpgradeInput) GetConfigsOk() (*bool, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *AppEnvironmentStackUpgradeInput) SetConfigs(v bool)`

SetConfigs sets Configs field to given value.


### GetCron

`func (o *AppEnvironmentStackUpgradeInput) GetCron() bool`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *AppEnvironmentStackUpgradeInput) GetCronOk() (*bool, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *AppEnvironmentStackUpgradeInput) SetCron(v bool)`

SetCron sets Cron field to given value.


### GetVolumes

`func (o *AppEnvironmentStackUpgradeInput) GetVolumes() bool`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *AppEnvironmentStackUpgradeInput) GetVolumesOk() (*bool, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *AppEnvironmentStackUpgradeInput) SetVolumes(v bool)`

SetVolumes sets Volumes field to given value.


### GetMain

`func (o *AppEnvironmentStackUpgradeInput) GetMain() bool`

GetMain returns the Main field if non-nil, zero value otherwise.

### GetMainOk

`func (o *AppEnvironmentStackUpgradeInput) GetMainOk() (*bool, bool)`

GetMainOk returns a tuple with the Main field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMain

`func (o *AppEnvironmentStackUpgradeInput) SetMain(v bool)`

SetMain sets Main field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


