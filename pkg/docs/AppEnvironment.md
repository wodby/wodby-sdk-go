# AppEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Title** | **string** |  | 
**EnvironmentType** | **string** |  | 
**Status** | **string** |  | 
**Outdated** | **bool** |  | 
**MainDomain** | Pointer to **NullableString** |  | [optional] 
**MainRouteCert** | [**NullableAppEnvironmentMainRouteCert**](AppEnvironmentMainRouteCert.md) |  | 
**AppId** | **int32** |  | 
**ClusterId** | **int32** |  | 
**StackId** | **int32** |  | 
**StackRevId** | **int32** |  | 
**StackName** | **string** |  | 
**StackTitle** | **string** |  | 
**StackIcon** | **string** |  | 
**StackRevNumber** | **int32** |  | 
**StackVersion** | **string** |  | 
**Access** | Pointer to [**NullableAppAccess**](AppAccess.md) |  | [optional] 
**RoutingMode** | **string** |  | 
**RoutingPending** | **bool** |  | 
**MaintenanceMode** | **bool** |  | 
**MaintenanceModeActive** | **bool** |  | 
**ConfigurationReady** | **bool** |  | 
**ConfigurationIssues** | [**[]AppServiceConfigurationIssue**](AppServiceConfigurationIssue.md) |  | 
**Settings** | Pointer to [**AppEnvironmentSettings**](AppEnvironmentSettings.md) |  | [optional] 
**Health** | [**AppEnvironmentHealth**](AppEnvironmentHealth.md) |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewAppEnvironment

`func NewAppEnvironment(id int32, name string, title string, environmentType string, status string, outdated bool, mainRouteCert NullableAppEnvironmentMainRouteCert, appId int32, clusterId int32, stackId int32, stackRevId int32, stackName string, stackTitle string, stackIcon string, stackRevNumber int32, stackVersion string, routingMode string, routingPending bool, maintenanceMode bool, maintenanceModeActive bool, configurationReady bool, configurationIssues []AppServiceConfigurationIssue, health AppEnvironmentHealth, createdAt time.Time, updatedAt time.Time, ) *AppEnvironment`

NewAppEnvironment instantiates a new AppEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppEnvironmentWithDefaults

`func NewAppEnvironmentWithDefaults() *AppEnvironment`

NewAppEnvironmentWithDefaults instantiates a new AppEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppEnvironment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppEnvironment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppEnvironment) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *AppEnvironment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppEnvironment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppEnvironment) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *AppEnvironment) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AppEnvironment) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AppEnvironment) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetEnvironmentType

`func (o *AppEnvironment) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *AppEnvironment) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *AppEnvironment) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.


### GetStatus

`func (o *AppEnvironment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppEnvironment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppEnvironment) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetOutdated

`func (o *AppEnvironment) GetOutdated() bool`

GetOutdated returns the Outdated field if non-nil, zero value otherwise.

### GetOutdatedOk

`func (o *AppEnvironment) GetOutdatedOk() (*bool, bool)`

GetOutdatedOk returns a tuple with the Outdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutdated

`func (o *AppEnvironment) SetOutdated(v bool)`

SetOutdated sets Outdated field to given value.


### GetMainDomain

`func (o *AppEnvironment) GetMainDomain() string`

GetMainDomain returns the MainDomain field if non-nil, zero value otherwise.

### GetMainDomainOk

`func (o *AppEnvironment) GetMainDomainOk() (*string, bool)`

GetMainDomainOk returns a tuple with the MainDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainDomain

`func (o *AppEnvironment) SetMainDomain(v string)`

SetMainDomain sets MainDomain field to given value.

### HasMainDomain

`func (o *AppEnvironment) HasMainDomain() bool`

HasMainDomain returns a boolean if a field has been set.

### SetMainDomainNil

`func (o *AppEnvironment) SetMainDomainNil(b bool)`

 SetMainDomainNil sets the value for MainDomain to be an explicit nil

### UnsetMainDomain
`func (o *AppEnvironment) UnsetMainDomain()`

UnsetMainDomain ensures that no value is present for MainDomain, not even an explicit nil
### GetMainRouteCert

`func (o *AppEnvironment) GetMainRouteCert() AppEnvironmentMainRouteCert`

GetMainRouteCert returns the MainRouteCert field if non-nil, zero value otherwise.

### GetMainRouteCertOk

`func (o *AppEnvironment) GetMainRouteCertOk() (*AppEnvironmentMainRouteCert, bool)`

GetMainRouteCertOk returns a tuple with the MainRouteCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainRouteCert

`func (o *AppEnvironment) SetMainRouteCert(v AppEnvironmentMainRouteCert)`

SetMainRouteCert sets MainRouteCert field to given value.


### SetMainRouteCertNil

`func (o *AppEnvironment) SetMainRouteCertNil(b bool)`

 SetMainRouteCertNil sets the value for MainRouteCert to be an explicit nil

### UnsetMainRouteCert
`func (o *AppEnvironment) UnsetMainRouteCert()`

UnsetMainRouteCert ensures that no value is present for MainRouteCert, not even an explicit nil
### GetAppId

`func (o *AppEnvironment) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *AppEnvironment) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *AppEnvironment) SetAppId(v int32)`

SetAppId sets AppId field to given value.


### GetClusterId

`func (o *AppEnvironment) GetClusterId() int32`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *AppEnvironment) GetClusterIdOk() (*int32, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *AppEnvironment) SetClusterId(v int32)`

SetClusterId sets ClusterId field to given value.


### GetStackId

`func (o *AppEnvironment) GetStackId() int32`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *AppEnvironment) GetStackIdOk() (*int32, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *AppEnvironment) SetStackId(v int32)`

SetStackId sets StackId field to given value.


### GetStackRevId

`func (o *AppEnvironment) GetStackRevId() int32`

GetStackRevId returns the StackRevId field if non-nil, zero value otherwise.

### GetStackRevIdOk

`func (o *AppEnvironment) GetStackRevIdOk() (*int32, bool)`

GetStackRevIdOk returns a tuple with the StackRevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackRevId

`func (o *AppEnvironment) SetStackRevId(v int32)`

SetStackRevId sets StackRevId field to given value.


### GetStackName

`func (o *AppEnvironment) GetStackName() string`

GetStackName returns the StackName field if non-nil, zero value otherwise.

### GetStackNameOk

`func (o *AppEnvironment) GetStackNameOk() (*string, bool)`

GetStackNameOk returns a tuple with the StackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackName

`func (o *AppEnvironment) SetStackName(v string)`

SetStackName sets StackName field to given value.


### GetStackTitle

`func (o *AppEnvironment) GetStackTitle() string`

GetStackTitle returns the StackTitle field if non-nil, zero value otherwise.

### GetStackTitleOk

`func (o *AppEnvironment) GetStackTitleOk() (*string, bool)`

GetStackTitleOk returns a tuple with the StackTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackTitle

`func (o *AppEnvironment) SetStackTitle(v string)`

SetStackTitle sets StackTitle field to given value.


### GetStackIcon

`func (o *AppEnvironment) GetStackIcon() string`

GetStackIcon returns the StackIcon field if non-nil, zero value otherwise.

### GetStackIconOk

`func (o *AppEnvironment) GetStackIconOk() (*string, bool)`

GetStackIconOk returns a tuple with the StackIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackIcon

`func (o *AppEnvironment) SetStackIcon(v string)`

SetStackIcon sets StackIcon field to given value.


### GetStackRevNumber

`func (o *AppEnvironment) GetStackRevNumber() int32`

GetStackRevNumber returns the StackRevNumber field if non-nil, zero value otherwise.

### GetStackRevNumberOk

`func (o *AppEnvironment) GetStackRevNumberOk() (*int32, bool)`

GetStackRevNumberOk returns a tuple with the StackRevNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackRevNumber

`func (o *AppEnvironment) SetStackRevNumber(v int32)`

SetStackRevNumber sets StackRevNumber field to given value.


### GetStackVersion

`func (o *AppEnvironment) GetStackVersion() string`

GetStackVersion returns the StackVersion field if non-nil, zero value otherwise.

### GetStackVersionOk

`func (o *AppEnvironment) GetStackVersionOk() (*string, bool)`

GetStackVersionOk returns a tuple with the StackVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackVersion

`func (o *AppEnvironment) SetStackVersion(v string)`

SetStackVersion sets StackVersion field to given value.


### GetAccess

`func (o *AppEnvironment) GetAccess() AppAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *AppEnvironment) GetAccessOk() (*AppAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *AppEnvironment) SetAccess(v AppAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *AppEnvironment) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### SetAccessNil

`func (o *AppEnvironment) SetAccessNil(b bool)`

 SetAccessNil sets the value for Access to be an explicit nil

### UnsetAccess
`func (o *AppEnvironment) UnsetAccess()`

UnsetAccess ensures that no value is present for Access, not even an explicit nil
### GetRoutingMode

`func (o *AppEnvironment) GetRoutingMode() string`

GetRoutingMode returns the RoutingMode field if non-nil, zero value otherwise.

### GetRoutingModeOk

`func (o *AppEnvironment) GetRoutingModeOk() (*string, bool)`

GetRoutingModeOk returns a tuple with the RoutingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingMode

`func (o *AppEnvironment) SetRoutingMode(v string)`

SetRoutingMode sets RoutingMode field to given value.


### GetRoutingPending

`func (o *AppEnvironment) GetRoutingPending() bool`

GetRoutingPending returns the RoutingPending field if non-nil, zero value otherwise.

### GetRoutingPendingOk

`func (o *AppEnvironment) GetRoutingPendingOk() (*bool, bool)`

GetRoutingPendingOk returns a tuple with the RoutingPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingPending

`func (o *AppEnvironment) SetRoutingPending(v bool)`

SetRoutingPending sets RoutingPending field to given value.


### GetMaintenanceMode

`func (o *AppEnvironment) GetMaintenanceMode() bool`

GetMaintenanceMode returns the MaintenanceMode field if non-nil, zero value otherwise.

### GetMaintenanceModeOk

`func (o *AppEnvironment) GetMaintenanceModeOk() (*bool, bool)`

GetMaintenanceModeOk returns a tuple with the MaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceMode

`func (o *AppEnvironment) SetMaintenanceMode(v bool)`

SetMaintenanceMode sets MaintenanceMode field to given value.


### GetMaintenanceModeActive

`func (o *AppEnvironment) GetMaintenanceModeActive() bool`

GetMaintenanceModeActive returns the MaintenanceModeActive field if non-nil, zero value otherwise.

### GetMaintenanceModeActiveOk

`func (o *AppEnvironment) GetMaintenanceModeActiveOk() (*bool, bool)`

GetMaintenanceModeActiveOk returns a tuple with the MaintenanceModeActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceModeActive

`func (o *AppEnvironment) SetMaintenanceModeActive(v bool)`

SetMaintenanceModeActive sets MaintenanceModeActive field to given value.


### GetConfigurationReady

`func (o *AppEnvironment) GetConfigurationReady() bool`

GetConfigurationReady returns the ConfigurationReady field if non-nil, zero value otherwise.

### GetConfigurationReadyOk

`func (o *AppEnvironment) GetConfigurationReadyOk() (*bool, bool)`

GetConfigurationReadyOk returns a tuple with the ConfigurationReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationReady

`func (o *AppEnvironment) SetConfigurationReady(v bool)`

SetConfigurationReady sets ConfigurationReady field to given value.


### GetConfigurationIssues

`func (o *AppEnvironment) GetConfigurationIssues() []AppServiceConfigurationIssue`

GetConfigurationIssues returns the ConfigurationIssues field if non-nil, zero value otherwise.

### GetConfigurationIssuesOk

`func (o *AppEnvironment) GetConfigurationIssuesOk() (*[]AppServiceConfigurationIssue, bool)`

GetConfigurationIssuesOk returns a tuple with the ConfigurationIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationIssues

`func (o *AppEnvironment) SetConfigurationIssues(v []AppServiceConfigurationIssue)`

SetConfigurationIssues sets ConfigurationIssues field to given value.


### GetSettings

`func (o *AppEnvironment) GetSettings() AppEnvironmentSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *AppEnvironment) GetSettingsOk() (*AppEnvironmentSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *AppEnvironment) SetSettings(v AppEnvironmentSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *AppEnvironment) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetHealth

`func (o *AppEnvironment) GetHealth() AppEnvironmentHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *AppEnvironment) GetHealthOk() (*AppEnvironmentHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *AppEnvironment) SetHealth(v AppEnvironmentHealth)`

SetHealth sets Health field to given value.


### GetCreatedAt

`func (o *AppEnvironment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppEnvironment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppEnvironment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AppEnvironment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AppEnvironment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AppEnvironment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


