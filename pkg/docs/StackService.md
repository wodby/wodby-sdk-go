# StackService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Title** | **string** |  | 
**Type** | **string** |  | 
**Main** | **bool** |  | 
**Disabled** | **bool** |  | 
**Required** | **bool** |  | 
**BasicAuthOverride** | **string** | Whether the stack overrides the service&#39;s basic auth default. | 
**BasicAuthDefault** | **bool** | Whether new app services of this stack service get their own basic auth entry by default. Always false without a public HTTP endpoint. | 
**Replicas** | **int32** |  | 
**ServiceRevPinned** | **bool** |  | 
**Outdated** | **bool** |  | 
**Icon** | Pointer to **string** |  | [optional] 
**Eol** | Pointer to **bool** | True when the version of the stack service&#39;s default option has reached its end of life. | [optional] 
**Derivative** | Pointer to **NullableString** | The service derivative the stack service runs. Null for the service itself. | [optional] 
**ParentStackServiceId** | Pointer to **NullableInt32** | The stack service this derivative stack service belongs to. Null for the service itself. | [optional] 
**ServiceRevId** | **int32** |  | 
**ServiceRevName** | **string** |  | 
**ServiceRevTitle** | **string** |  | 
**ServiceRevVersion** | **string** |  | 
**BuildSourceIntegrationId** | Pointer to **NullableInt32** |  | [optional] 
**BuildSourceRemoteRepoId** | Pointer to **NullableString** |  | [optional] 
**Options** | Pointer to [**[]StackServiceOption**](StackServiceOption.md) |  | [optional] 
**Settings** | Pointer to [**[]StackServiceSetting**](StackServiceSetting.md) |  | [optional] 
**Containers** | Pointer to [**[]StackServiceContainer**](StackServiceContainer.md) |  | [optional] 
**DeploymentConfiguration** | [**ServiceDeploymentConfiguration**](ServiceDeploymentConfiguration.md) |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewStackService

`func NewStackService(id int32, name string, title string, type_ string, main bool, disabled bool, required bool, basicAuthOverride string, basicAuthDefault bool, replicas int32, serviceRevPinned bool, outdated bool, serviceRevId int32, serviceRevName string, serviceRevTitle string, serviceRevVersion string, deploymentConfiguration ServiceDeploymentConfiguration, createdAt time.Time, updatedAt time.Time, ) *StackService`

NewStackService instantiates a new StackService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackServiceWithDefaults

`func NewStackServiceWithDefaults() *StackService`

NewStackServiceWithDefaults instantiates a new StackService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StackService) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StackService) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StackService) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *StackService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StackService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StackService) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *StackService) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StackService) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StackService) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *StackService) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StackService) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StackService) SetType(v string)`

SetType sets Type field to given value.


### GetMain

`func (o *StackService) GetMain() bool`

GetMain returns the Main field if non-nil, zero value otherwise.

### GetMainOk

`func (o *StackService) GetMainOk() (*bool, bool)`

GetMainOk returns a tuple with the Main field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMain

`func (o *StackService) SetMain(v bool)`

SetMain sets Main field to given value.


### GetDisabled

`func (o *StackService) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *StackService) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *StackService) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetRequired

`func (o *StackService) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *StackService) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *StackService) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetBasicAuthOverride

`func (o *StackService) GetBasicAuthOverride() string`

GetBasicAuthOverride returns the BasicAuthOverride field if non-nil, zero value otherwise.

### GetBasicAuthOverrideOk

`func (o *StackService) GetBasicAuthOverrideOk() (*string, bool)`

GetBasicAuthOverrideOk returns a tuple with the BasicAuthOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthOverride

`func (o *StackService) SetBasicAuthOverride(v string)`

SetBasicAuthOverride sets BasicAuthOverride field to given value.


### GetBasicAuthDefault

`func (o *StackService) GetBasicAuthDefault() bool`

GetBasicAuthDefault returns the BasicAuthDefault field if non-nil, zero value otherwise.

### GetBasicAuthDefaultOk

`func (o *StackService) GetBasicAuthDefaultOk() (*bool, bool)`

GetBasicAuthDefaultOk returns a tuple with the BasicAuthDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthDefault

`func (o *StackService) SetBasicAuthDefault(v bool)`

SetBasicAuthDefault sets BasicAuthDefault field to given value.


### GetReplicas

`func (o *StackService) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *StackService) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *StackService) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.


### GetServiceRevPinned

`func (o *StackService) GetServiceRevPinned() bool`

GetServiceRevPinned returns the ServiceRevPinned field if non-nil, zero value otherwise.

### GetServiceRevPinnedOk

`func (o *StackService) GetServiceRevPinnedOk() (*bool, bool)`

GetServiceRevPinnedOk returns a tuple with the ServiceRevPinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRevPinned

`func (o *StackService) SetServiceRevPinned(v bool)`

SetServiceRevPinned sets ServiceRevPinned field to given value.


### GetOutdated

`func (o *StackService) GetOutdated() bool`

GetOutdated returns the Outdated field if non-nil, zero value otherwise.

### GetOutdatedOk

`func (o *StackService) GetOutdatedOk() (*bool, bool)`

GetOutdatedOk returns a tuple with the Outdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutdated

`func (o *StackService) SetOutdated(v bool)`

SetOutdated sets Outdated field to given value.


### GetIcon

`func (o *StackService) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *StackService) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *StackService) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *StackService) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetEol

`func (o *StackService) GetEol() bool`

GetEol returns the Eol field if non-nil, zero value otherwise.

### GetEolOk

`func (o *StackService) GetEolOk() (*bool, bool)`

GetEolOk returns a tuple with the Eol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEol

`func (o *StackService) SetEol(v bool)`

SetEol sets Eol field to given value.

### HasEol

`func (o *StackService) HasEol() bool`

HasEol returns a boolean if a field has been set.

### GetDerivative

`func (o *StackService) GetDerivative() string`

GetDerivative returns the Derivative field if non-nil, zero value otherwise.

### GetDerivativeOk

`func (o *StackService) GetDerivativeOk() (*string, bool)`

GetDerivativeOk returns a tuple with the Derivative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivative

`func (o *StackService) SetDerivative(v string)`

SetDerivative sets Derivative field to given value.

### HasDerivative

`func (o *StackService) HasDerivative() bool`

HasDerivative returns a boolean if a field has been set.

### SetDerivativeNil

`func (o *StackService) SetDerivativeNil(b bool)`

 SetDerivativeNil sets the value for Derivative to be an explicit nil

### UnsetDerivative
`func (o *StackService) UnsetDerivative()`

UnsetDerivative ensures that no value is present for Derivative, not even an explicit nil
### GetParentStackServiceId

`func (o *StackService) GetParentStackServiceId() int32`

GetParentStackServiceId returns the ParentStackServiceId field if non-nil, zero value otherwise.

### GetParentStackServiceIdOk

`func (o *StackService) GetParentStackServiceIdOk() (*int32, bool)`

GetParentStackServiceIdOk returns a tuple with the ParentStackServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentStackServiceId

`func (o *StackService) SetParentStackServiceId(v int32)`

SetParentStackServiceId sets ParentStackServiceId field to given value.

### HasParentStackServiceId

`func (o *StackService) HasParentStackServiceId() bool`

HasParentStackServiceId returns a boolean if a field has been set.

### SetParentStackServiceIdNil

`func (o *StackService) SetParentStackServiceIdNil(b bool)`

 SetParentStackServiceIdNil sets the value for ParentStackServiceId to be an explicit nil

### UnsetParentStackServiceId
`func (o *StackService) UnsetParentStackServiceId()`

UnsetParentStackServiceId ensures that no value is present for ParentStackServiceId, not even an explicit nil
### GetServiceRevId

`func (o *StackService) GetServiceRevId() int32`

GetServiceRevId returns the ServiceRevId field if non-nil, zero value otherwise.

### GetServiceRevIdOk

`func (o *StackService) GetServiceRevIdOk() (*int32, bool)`

GetServiceRevIdOk returns a tuple with the ServiceRevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRevId

`func (o *StackService) SetServiceRevId(v int32)`

SetServiceRevId sets ServiceRevId field to given value.


### GetServiceRevName

`func (o *StackService) GetServiceRevName() string`

GetServiceRevName returns the ServiceRevName field if non-nil, zero value otherwise.

### GetServiceRevNameOk

`func (o *StackService) GetServiceRevNameOk() (*string, bool)`

GetServiceRevNameOk returns a tuple with the ServiceRevName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRevName

`func (o *StackService) SetServiceRevName(v string)`

SetServiceRevName sets ServiceRevName field to given value.


### GetServiceRevTitle

`func (o *StackService) GetServiceRevTitle() string`

GetServiceRevTitle returns the ServiceRevTitle field if non-nil, zero value otherwise.

### GetServiceRevTitleOk

`func (o *StackService) GetServiceRevTitleOk() (*string, bool)`

GetServiceRevTitleOk returns a tuple with the ServiceRevTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRevTitle

`func (o *StackService) SetServiceRevTitle(v string)`

SetServiceRevTitle sets ServiceRevTitle field to given value.


### GetServiceRevVersion

`func (o *StackService) GetServiceRevVersion() string`

GetServiceRevVersion returns the ServiceRevVersion field if non-nil, zero value otherwise.

### GetServiceRevVersionOk

`func (o *StackService) GetServiceRevVersionOk() (*string, bool)`

GetServiceRevVersionOk returns a tuple with the ServiceRevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRevVersion

`func (o *StackService) SetServiceRevVersion(v string)`

SetServiceRevVersion sets ServiceRevVersion field to given value.


### GetBuildSourceIntegrationId

`func (o *StackService) GetBuildSourceIntegrationId() int32`

GetBuildSourceIntegrationId returns the BuildSourceIntegrationId field if non-nil, zero value otherwise.

### GetBuildSourceIntegrationIdOk

`func (o *StackService) GetBuildSourceIntegrationIdOk() (*int32, bool)`

GetBuildSourceIntegrationIdOk returns a tuple with the BuildSourceIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildSourceIntegrationId

`func (o *StackService) SetBuildSourceIntegrationId(v int32)`

SetBuildSourceIntegrationId sets BuildSourceIntegrationId field to given value.

### HasBuildSourceIntegrationId

`func (o *StackService) HasBuildSourceIntegrationId() bool`

HasBuildSourceIntegrationId returns a boolean if a field has been set.

### SetBuildSourceIntegrationIdNil

`func (o *StackService) SetBuildSourceIntegrationIdNil(b bool)`

 SetBuildSourceIntegrationIdNil sets the value for BuildSourceIntegrationId to be an explicit nil

### UnsetBuildSourceIntegrationId
`func (o *StackService) UnsetBuildSourceIntegrationId()`

UnsetBuildSourceIntegrationId ensures that no value is present for BuildSourceIntegrationId, not even an explicit nil
### GetBuildSourceRemoteRepoId

`func (o *StackService) GetBuildSourceRemoteRepoId() string`

GetBuildSourceRemoteRepoId returns the BuildSourceRemoteRepoId field if non-nil, zero value otherwise.

### GetBuildSourceRemoteRepoIdOk

`func (o *StackService) GetBuildSourceRemoteRepoIdOk() (*string, bool)`

GetBuildSourceRemoteRepoIdOk returns a tuple with the BuildSourceRemoteRepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildSourceRemoteRepoId

`func (o *StackService) SetBuildSourceRemoteRepoId(v string)`

SetBuildSourceRemoteRepoId sets BuildSourceRemoteRepoId field to given value.

### HasBuildSourceRemoteRepoId

`func (o *StackService) HasBuildSourceRemoteRepoId() bool`

HasBuildSourceRemoteRepoId returns a boolean if a field has been set.

### SetBuildSourceRemoteRepoIdNil

`func (o *StackService) SetBuildSourceRemoteRepoIdNil(b bool)`

 SetBuildSourceRemoteRepoIdNil sets the value for BuildSourceRemoteRepoId to be an explicit nil

### UnsetBuildSourceRemoteRepoId
`func (o *StackService) UnsetBuildSourceRemoteRepoId()`

UnsetBuildSourceRemoteRepoId ensures that no value is present for BuildSourceRemoteRepoId, not even an explicit nil
### GetOptions

`func (o *StackService) GetOptions() []StackServiceOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *StackService) GetOptionsOk() (*[]StackServiceOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *StackService) SetOptions(v []StackServiceOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *StackService) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetSettings

`func (o *StackService) GetSettings() []StackServiceSetting`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *StackService) GetSettingsOk() (*[]StackServiceSetting, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *StackService) SetSettings(v []StackServiceSetting)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *StackService) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetContainers

`func (o *StackService) GetContainers() []StackServiceContainer`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *StackService) GetContainersOk() (*[]StackServiceContainer, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *StackService) SetContainers(v []StackServiceContainer)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *StackService) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetDeploymentConfiguration

`func (o *StackService) GetDeploymentConfiguration() ServiceDeploymentConfiguration`

GetDeploymentConfiguration returns the DeploymentConfiguration field if non-nil, zero value otherwise.

### GetDeploymentConfigurationOk

`func (o *StackService) GetDeploymentConfigurationOk() (*ServiceDeploymentConfiguration, bool)`

GetDeploymentConfigurationOk returns a tuple with the DeploymentConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentConfiguration

`func (o *StackService) SetDeploymentConfiguration(v ServiceDeploymentConfiguration)`

SetDeploymentConfiguration sets DeploymentConfiguration field to given value.


### GetCreatedAt

`func (o *StackService) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StackService) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StackService) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *StackService) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *StackService) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *StackService) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


