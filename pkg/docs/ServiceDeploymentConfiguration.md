# ServiceDeploymentConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Override** | [**ServiceDeploymentPolicy**](ServiceDeploymentPolicy.md) |  | 
**Effective** | [**ServiceDeploymentPolicy**](ServiceDeploymentPolicy.md) |  | 
**Workloads** | [**[]ServiceWorkloadDeploymentConfiguration**](ServiceWorkloadDeploymentConfiguration.md) |  | 

## Methods

### NewServiceDeploymentConfiguration

`func NewServiceDeploymentConfiguration(override ServiceDeploymentPolicy, effective ServiceDeploymentPolicy, workloads []ServiceWorkloadDeploymentConfiguration, ) *ServiceDeploymentConfiguration`

NewServiceDeploymentConfiguration instantiates a new ServiceDeploymentConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDeploymentConfigurationWithDefaults

`func NewServiceDeploymentConfigurationWithDefaults() *ServiceDeploymentConfiguration`

NewServiceDeploymentConfigurationWithDefaults instantiates a new ServiceDeploymentConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverride

`func (o *ServiceDeploymentConfiguration) GetOverride() ServiceDeploymentPolicy`

GetOverride returns the Override field if non-nil, zero value otherwise.

### GetOverrideOk

`func (o *ServiceDeploymentConfiguration) GetOverrideOk() (*ServiceDeploymentPolicy, bool)`

GetOverrideOk returns a tuple with the Override field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverride

`func (o *ServiceDeploymentConfiguration) SetOverride(v ServiceDeploymentPolicy)`

SetOverride sets Override field to given value.


### GetEffective

`func (o *ServiceDeploymentConfiguration) GetEffective() ServiceDeploymentPolicy`

GetEffective returns the Effective field if non-nil, zero value otherwise.

### GetEffectiveOk

`func (o *ServiceDeploymentConfiguration) GetEffectiveOk() (*ServiceDeploymentPolicy, bool)`

GetEffectiveOk returns a tuple with the Effective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffective

`func (o *ServiceDeploymentConfiguration) SetEffective(v ServiceDeploymentPolicy)`

SetEffective sets Effective field to given value.


### GetWorkloads

`func (o *ServiceDeploymentConfiguration) GetWorkloads() []ServiceWorkloadDeploymentConfiguration`

GetWorkloads returns the Workloads field if non-nil, zero value otherwise.

### GetWorkloadsOk

`func (o *ServiceDeploymentConfiguration) GetWorkloadsOk() (*[]ServiceWorkloadDeploymentConfiguration, bool)`

GetWorkloadsOk returns a tuple with the Workloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloads

`func (o *ServiceDeploymentConfiguration) SetWorkloads(v []ServiceWorkloadDeploymentConfiguration)`

SetWorkloads sets Workloads field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


