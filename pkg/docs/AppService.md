# AppService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Title** | **string** |  | 
**Type** | **string** |  | 
**Status** | **string** |  | 
**Replicas** | **int32** |  | 
**Scalability** | Pointer to [**NullableAppServiceScalability**](AppServiceScalability.md) |  | [optional] 
**Version** | **string** |  | 
**Main** | **bool** |  | 
**Disabled** | **bool** |  | 
**External** | **bool** |  | 
**Required** | **bool** |  | 
**NeedsRebuild** | **bool** |  | 
**NeedsRedeploy** | **bool** |  | 
**StackState** | **string** |  | 
**ConfigurationReady** | **bool** |  | 
**BuildSourceBoilerplate** | Pointer to **NullableString** |  | [optional] 
**CiPolicy** | **string** |  | 
**EffectiveCiIntegrationId** | Pointer to **NullableInt32** |  | [optional] 
**ConfigurationIssues** | [**[]AppServiceConfigurationIssue**](AppServiceConfigurationIssue.md) |  | 
**AppInstanceId** | **int32** |  | 
**ServiceRevId** | **int32** |  | 
**ParentAppServiceId** | Pointer to **NullableInt32** |  | [optional] 
**DeploymentConfiguration** | [**ServiceDeploymentConfiguration**](ServiceDeploymentConfiguration.md) |  | 
**Icon** | Pointer to **string** |  | [optional] 
**Eol** | Pointer to **bool** | True when the app service&#39;s version has reached its end of life. | [optional] 
**BuildSourceType** | Pointer to **NullableString** | How the app service gets its source code to build, public, clone or connect. Null when it is not built from source. | [optional] 
**Derivative** | Pointer to **NullableString** | The service derivative the app service runs. Null for the service itself. | [optional] 
**GitRepoId** | Pointer to **NullableInt32** |  | [optional] 
**WorkspaceMounts** | Pointer to [**[]AppServiceWorkspaceMount**](AppServiceWorkspaceMount.md) | Workspace code mounts of the service, with its settings resolved in the paths. Empty without a workspace contract. | [optional] 
**DatabaseDbId** | Pointer to **NullableInt32** | The database the app service uses on a provider-managed database server. | [optional] 
**DatabaseUserId** | Pointer to **NullableInt32** | The database user the app service connects as on a provider-managed database server. | [optional] 
**Customizations** | Pointer to [**[]AppServiceCustomization**](AppServiceCustomization.md) | What this app environment set differently from the stack service the app service was created from, such as a setting value, an extra variable or another link target. Only names are returned, never values. Returned only by GET /app-services/{id}. | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewAppService

`func NewAppService(id int32, name string, title string, type_ string, status string, replicas int32, version string, main bool, disabled bool, external bool, required bool, needsRebuild bool, needsRedeploy bool, stackState string, configurationReady bool, ciPolicy string, configurationIssues []AppServiceConfigurationIssue, appInstanceId int32, serviceRevId int32, deploymentConfiguration ServiceDeploymentConfiguration, createdAt time.Time, updatedAt time.Time, ) *AppService`

NewAppService instantiates a new AppService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceWithDefaults

`func NewAppServiceWithDefaults() *AppService`

NewAppServiceWithDefaults instantiates a new AppService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppService) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppService) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppService) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *AppService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppService) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *AppService) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AppService) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AppService) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *AppService) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppService) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppService) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *AppService) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppService) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppService) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetReplicas

`func (o *AppService) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *AppService) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *AppService) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.


### GetScalability

`func (o *AppService) GetScalability() AppServiceScalability`

GetScalability returns the Scalability field if non-nil, zero value otherwise.

### GetScalabilityOk

`func (o *AppService) GetScalabilityOk() (*AppServiceScalability, bool)`

GetScalabilityOk returns a tuple with the Scalability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalability

`func (o *AppService) SetScalability(v AppServiceScalability)`

SetScalability sets Scalability field to given value.

### HasScalability

`func (o *AppService) HasScalability() bool`

HasScalability returns a boolean if a field has been set.

### SetScalabilityNil

`func (o *AppService) SetScalabilityNil(b bool)`

 SetScalabilityNil sets the value for Scalability to be an explicit nil

### UnsetScalability
`func (o *AppService) UnsetScalability()`

UnsetScalability ensures that no value is present for Scalability, not even an explicit nil
### GetVersion

`func (o *AppService) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AppService) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AppService) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetMain

`func (o *AppService) GetMain() bool`

GetMain returns the Main field if non-nil, zero value otherwise.

### GetMainOk

`func (o *AppService) GetMainOk() (*bool, bool)`

GetMainOk returns a tuple with the Main field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMain

`func (o *AppService) SetMain(v bool)`

SetMain sets Main field to given value.


### GetDisabled

`func (o *AppService) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *AppService) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *AppService) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetExternal

`func (o *AppService) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *AppService) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *AppService) SetExternal(v bool)`

SetExternal sets External field to given value.


### GetRequired

`func (o *AppService) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *AppService) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *AppService) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetNeedsRebuild

`func (o *AppService) GetNeedsRebuild() bool`

GetNeedsRebuild returns the NeedsRebuild field if non-nil, zero value otherwise.

### GetNeedsRebuildOk

`func (o *AppService) GetNeedsRebuildOk() (*bool, bool)`

GetNeedsRebuildOk returns a tuple with the NeedsRebuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsRebuild

`func (o *AppService) SetNeedsRebuild(v bool)`

SetNeedsRebuild sets NeedsRebuild field to given value.


### GetNeedsRedeploy

`func (o *AppService) GetNeedsRedeploy() bool`

GetNeedsRedeploy returns the NeedsRedeploy field if non-nil, zero value otherwise.

### GetNeedsRedeployOk

`func (o *AppService) GetNeedsRedeployOk() (*bool, bool)`

GetNeedsRedeployOk returns a tuple with the NeedsRedeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsRedeploy

`func (o *AppService) SetNeedsRedeploy(v bool)`

SetNeedsRedeploy sets NeedsRedeploy field to given value.


### GetStackState

`func (o *AppService) GetStackState() string`

GetStackState returns the StackState field if non-nil, zero value otherwise.

### GetStackStateOk

`func (o *AppService) GetStackStateOk() (*string, bool)`

GetStackStateOk returns a tuple with the StackState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackState

`func (o *AppService) SetStackState(v string)`

SetStackState sets StackState field to given value.


### GetConfigurationReady

`func (o *AppService) GetConfigurationReady() bool`

GetConfigurationReady returns the ConfigurationReady field if non-nil, zero value otherwise.

### GetConfigurationReadyOk

`func (o *AppService) GetConfigurationReadyOk() (*bool, bool)`

GetConfigurationReadyOk returns a tuple with the ConfigurationReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationReady

`func (o *AppService) SetConfigurationReady(v bool)`

SetConfigurationReady sets ConfigurationReady field to given value.


### GetBuildSourceBoilerplate

`func (o *AppService) GetBuildSourceBoilerplate() string`

GetBuildSourceBoilerplate returns the BuildSourceBoilerplate field if non-nil, zero value otherwise.

### GetBuildSourceBoilerplateOk

`func (o *AppService) GetBuildSourceBoilerplateOk() (*string, bool)`

GetBuildSourceBoilerplateOk returns a tuple with the BuildSourceBoilerplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildSourceBoilerplate

`func (o *AppService) SetBuildSourceBoilerplate(v string)`

SetBuildSourceBoilerplate sets BuildSourceBoilerplate field to given value.

### HasBuildSourceBoilerplate

`func (o *AppService) HasBuildSourceBoilerplate() bool`

HasBuildSourceBoilerplate returns a boolean if a field has been set.

### SetBuildSourceBoilerplateNil

`func (o *AppService) SetBuildSourceBoilerplateNil(b bool)`

 SetBuildSourceBoilerplateNil sets the value for BuildSourceBoilerplate to be an explicit nil

### UnsetBuildSourceBoilerplate
`func (o *AppService) UnsetBuildSourceBoilerplate()`

UnsetBuildSourceBoilerplate ensures that no value is present for BuildSourceBoilerplate, not even an explicit nil
### GetCiPolicy

`func (o *AppService) GetCiPolicy() string`

GetCiPolicy returns the CiPolicy field if non-nil, zero value otherwise.

### GetCiPolicyOk

`func (o *AppService) GetCiPolicyOk() (*string, bool)`

GetCiPolicyOk returns a tuple with the CiPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiPolicy

`func (o *AppService) SetCiPolicy(v string)`

SetCiPolicy sets CiPolicy field to given value.


### GetEffectiveCiIntegrationId

`func (o *AppService) GetEffectiveCiIntegrationId() int32`

GetEffectiveCiIntegrationId returns the EffectiveCiIntegrationId field if non-nil, zero value otherwise.

### GetEffectiveCiIntegrationIdOk

`func (o *AppService) GetEffectiveCiIntegrationIdOk() (*int32, bool)`

GetEffectiveCiIntegrationIdOk returns a tuple with the EffectiveCiIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveCiIntegrationId

`func (o *AppService) SetEffectiveCiIntegrationId(v int32)`

SetEffectiveCiIntegrationId sets EffectiveCiIntegrationId field to given value.

### HasEffectiveCiIntegrationId

`func (o *AppService) HasEffectiveCiIntegrationId() bool`

HasEffectiveCiIntegrationId returns a boolean if a field has been set.

### SetEffectiveCiIntegrationIdNil

`func (o *AppService) SetEffectiveCiIntegrationIdNil(b bool)`

 SetEffectiveCiIntegrationIdNil sets the value for EffectiveCiIntegrationId to be an explicit nil

### UnsetEffectiveCiIntegrationId
`func (o *AppService) UnsetEffectiveCiIntegrationId()`

UnsetEffectiveCiIntegrationId ensures that no value is present for EffectiveCiIntegrationId, not even an explicit nil
### GetConfigurationIssues

`func (o *AppService) GetConfigurationIssues() []AppServiceConfigurationIssue`

GetConfigurationIssues returns the ConfigurationIssues field if non-nil, zero value otherwise.

### GetConfigurationIssuesOk

`func (o *AppService) GetConfigurationIssuesOk() (*[]AppServiceConfigurationIssue, bool)`

GetConfigurationIssuesOk returns a tuple with the ConfigurationIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationIssues

`func (o *AppService) SetConfigurationIssues(v []AppServiceConfigurationIssue)`

SetConfigurationIssues sets ConfigurationIssues field to given value.


### GetAppInstanceId

`func (o *AppService) GetAppInstanceId() int32`

GetAppInstanceId returns the AppInstanceId field if non-nil, zero value otherwise.

### GetAppInstanceIdOk

`func (o *AppService) GetAppInstanceIdOk() (*int32, bool)`

GetAppInstanceIdOk returns a tuple with the AppInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceId

`func (o *AppService) SetAppInstanceId(v int32)`

SetAppInstanceId sets AppInstanceId field to given value.


### GetServiceRevId

`func (o *AppService) GetServiceRevId() int32`

GetServiceRevId returns the ServiceRevId field if non-nil, zero value otherwise.

### GetServiceRevIdOk

`func (o *AppService) GetServiceRevIdOk() (*int32, bool)`

GetServiceRevIdOk returns a tuple with the ServiceRevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRevId

`func (o *AppService) SetServiceRevId(v int32)`

SetServiceRevId sets ServiceRevId field to given value.


### GetParentAppServiceId

`func (o *AppService) GetParentAppServiceId() int32`

GetParentAppServiceId returns the ParentAppServiceId field if non-nil, zero value otherwise.

### GetParentAppServiceIdOk

`func (o *AppService) GetParentAppServiceIdOk() (*int32, bool)`

GetParentAppServiceIdOk returns a tuple with the ParentAppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAppServiceId

`func (o *AppService) SetParentAppServiceId(v int32)`

SetParentAppServiceId sets ParentAppServiceId field to given value.

### HasParentAppServiceId

`func (o *AppService) HasParentAppServiceId() bool`

HasParentAppServiceId returns a boolean if a field has been set.

### SetParentAppServiceIdNil

`func (o *AppService) SetParentAppServiceIdNil(b bool)`

 SetParentAppServiceIdNil sets the value for ParentAppServiceId to be an explicit nil

### UnsetParentAppServiceId
`func (o *AppService) UnsetParentAppServiceId()`

UnsetParentAppServiceId ensures that no value is present for ParentAppServiceId, not even an explicit nil
### GetDeploymentConfiguration

`func (o *AppService) GetDeploymentConfiguration() ServiceDeploymentConfiguration`

GetDeploymentConfiguration returns the DeploymentConfiguration field if non-nil, zero value otherwise.

### GetDeploymentConfigurationOk

`func (o *AppService) GetDeploymentConfigurationOk() (*ServiceDeploymentConfiguration, bool)`

GetDeploymentConfigurationOk returns a tuple with the DeploymentConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentConfiguration

`func (o *AppService) SetDeploymentConfiguration(v ServiceDeploymentConfiguration)`

SetDeploymentConfiguration sets DeploymentConfiguration field to given value.


### GetIcon

`func (o *AppService) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *AppService) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *AppService) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *AppService) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetEol

`func (o *AppService) GetEol() bool`

GetEol returns the Eol field if non-nil, zero value otherwise.

### GetEolOk

`func (o *AppService) GetEolOk() (*bool, bool)`

GetEolOk returns a tuple with the Eol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEol

`func (o *AppService) SetEol(v bool)`

SetEol sets Eol field to given value.

### HasEol

`func (o *AppService) HasEol() bool`

HasEol returns a boolean if a field has been set.

### GetBuildSourceType

`func (o *AppService) GetBuildSourceType() string`

GetBuildSourceType returns the BuildSourceType field if non-nil, zero value otherwise.

### GetBuildSourceTypeOk

`func (o *AppService) GetBuildSourceTypeOk() (*string, bool)`

GetBuildSourceTypeOk returns a tuple with the BuildSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildSourceType

`func (o *AppService) SetBuildSourceType(v string)`

SetBuildSourceType sets BuildSourceType field to given value.

### HasBuildSourceType

`func (o *AppService) HasBuildSourceType() bool`

HasBuildSourceType returns a boolean if a field has been set.

### SetBuildSourceTypeNil

`func (o *AppService) SetBuildSourceTypeNil(b bool)`

 SetBuildSourceTypeNil sets the value for BuildSourceType to be an explicit nil

### UnsetBuildSourceType
`func (o *AppService) UnsetBuildSourceType()`

UnsetBuildSourceType ensures that no value is present for BuildSourceType, not even an explicit nil
### GetDerivative

`func (o *AppService) GetDerivative() string`

GetDerivative returns the Derivative field if non-nil, zero value otherwise.

### GetDerivativeOk

`func (o *AppService) GetDerivativeOk() (*string, bool)`

GetDerivativeOk returns a tuple with the Derivative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivative

`func (o *AppService) SetDerivative(v string)`

SetDerivative sets Derivative field to given value.

### HasDerivative

`func (o *AppService) HasDerivative() bool`

HasDerivative returns a boolean if a field has been set.

### SetDerivativeNil

`func (o *AppService) SetDerivativeNil(b bool)`

 SetDerivativeNil sets the value for Derivative to be an explicit nil

### UnsetDerivative
`func (o *AppService) UnsetDerivative()`

UnsetDerivative ensures that no value is present for Derivative, not even an explicit nil
### GetGitRepoId

`func (o *AppService) GetGitRepoId() int32`

GetGitRepoId returns the GitRepoId field if non-nil, zero value otherwise.

### GetGitRepoIdOk

`func (o *AppService) GetGitRepoIdOk() (*int32, bool)`

GetGitRepoIdOk returns a tuple with the GitRepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepoId

`func (o *AppService) SetGitRepoId(v int32)`

SetGitRepoId sets GitRepoId field to given value.

### HasGitRepoId

`func (o *AppService) HasGitRepoId() bool`

HasGitRepoId returns a boolean if a field has been set.

### SetGitRepoIdNil

`func (o *AppService) SetGitRepoIdNil(b bool)`

 SetGitRepoIdNil sets the value for GitRepoId to be an explicit nil

### UnsetGitRepoId
`func (o *AppService) UnsetGitRepoId()`

UnsetGitRepoId ensures that no value is present for GitRepoId, not even an explicit nil
### GetWorkspaceMounts

`func (o *AppService) GetWorkspaceMounts() []AppServiceWorkspaceMount`

GetWorkspaceMounts returns the WorkspaceMounts field if non-nil, zero value otherwise.

### GetWorkspaceMountsOk

`func (o *AppService) GetWorkspaceMountsOk() (*[]AppServiceWorkspaceMount, bool)`

GetWorkspaceMountsOk returns a tuple with the WorkspaceMounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceMounts

`func (o *AppService) SetWorkspaceMounts(v []AppServiceWorkspaceMount)`

SetWorkspaceMounts sets WorkspaceMounts field to given value.

### HasWorkspaceMounts

`func (o *AppService) HasWorkspaceMounts() bool`

HasWorkspaceMounts returns a boolean if a field has been set.

### GetDatabaseDbId

`func (o *AppService) GetDatabaseDbId() int32`

GetDatabaseDbId returns the DatabaseDbId field if non-nil, zero value otherwise.

### GetDatabaseDbIdOk

`func (o *AppService) GetDatabaseDbIdOk() (*int32, bool)`

GetDatabaseDbIdOk returns a tuple with the DatabaseDbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseDbId

`func (o *AppService) SetDatabaseDbId(v int32)`

SetDatabaseDbId sets DatabaseDbId field to given value.

### HasDatabaseDbId

`func (o *AppService) HasDatabaseDbId() bool`

HasDatabaseDbId returns a boolean if a field has been set.

### SetDatabaseDbIdNil

`func (o *AppService) SetDatabaseDbIdNil(b bool)`

 SetDatabaseDbIdNil sets the value for DatabaseDbId to be an explicit nil

### UnsetDatabaseDbId
`func (o *AppService) UnsetDatabaseDbId()`

UnsetDatabaseDbId ensures that no value is present for DatabaseDbId, not even an explicit nil
### GetDatabaseUserId

`func (o *AppService) GetDatabaseUserId() int32`

GetDatabaseUserId returns the DatabaseUserId field if non-nil, zero value otherwise.

### GetDatabaseUserIdOk

`func (o *AppService) GetDatabaseUserIdOk() (*int32, bool)`

GetDatabaseUserIdOk returns a tuple with the DatabaseUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUserId

`func (o *AppService) SetDatabaseUserId(v int32)`

SetDatabaseUserId sets DatabaseUserId field to given value.

### HasDatabaseUserId

`func (o *AppService) HasDatabaseUserId() bool`

HasDatabaseUserId returns a boolean if a field has been set.

### SetDatabaseUserIdNil

`func (o *AppService) SetDatabaseUserIdNil(b bool)`

 SetDatabaseUserIdNil sets the value for DatabaseUserId to be an explicit nil

### UnsetDatabaseUserId
`func (o *AppService) UnsetDatabaseUserId()`

UnsetDatabaseUserId ensures that no value is present for DatabaseUserId, not even an explicit nil
### GetCustomizations

`func (o *AppService) GetCustomizations() []AppServiceCustomization`

GetCustomizations returns the Customizations field if non-nil, zero value otherwise.

### GetCustomizationsOk

`func (o *AppService) GetCustomizationsOk() (*[]AppServiceCustomization, bool)`

GetCustomizationsOk returns a tuple with the Customizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizations

`func (o *AppService) SetCustomizations(v []AppServiceCustomization)`

SetCustomizations sets Customizations field to given value.

### HasCustomizations

`func (o *AppService) HasCustomizations() bool`

HasCustomizations returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AppService) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppService) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppService) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AppService) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AppService) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AppService) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


