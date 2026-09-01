# ServiceWorkloadDeploymentConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Kind** | **string** |  | 
**Override** | [**ServiceDeploymentPolicy**](ServiceDeploymentPolicy.md) |  | 
**Effective** | [**ServiceDeploymentPolicy**](ServiceDeploymentPolicy.md) |  | 

## Methods

### NewServiceWorkloadDeploymentConfiguration

`func NewServiceWorkloadDeploymentConfiguration(name string, kind string, override ServiceDeploymentPolicy, effective ServiceDeploymentPolicy, ) *ServiceWorkloadDeploymentConfiguration`

NewServiceWorkloadDeploymentConfiguration instantiates a new ServiceWorkloadDeploymentConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWorkloadDeploymentConfigurationWithDefaults

`func NewServiceWorkloadDeploymentConfigurationWithDefaults() *ServiceWorkloadDeploymentConfiguration`

NewServiceWorkloadDeploymentConfigurationWithDefaults instantiates a new ServiceWorkloadDeploymentConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServiceWorkloadDeploymentConfiguration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceWorkloadDeploymentConfiguration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceWorkloadDeploymentConfiguration) SetName(v string)`

SetName sets Name field to given value.


### GetKind

`func (o *ServiceWorkloadDeploymentConfiguration) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ServiceWorkloadDeploymentConfiguration) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ServiceWorkloadDeploymentConfiguration) SetKind(v string)`

SetKind sets Kind field to given value.


### GetOverride

`func (o *ServiceWorkloadDeploymentConfiguration) GetOverride() ServiceDeploymentPolicy`

GetOverride returns the Override field if non-nil, zero value otherwise.

### GetOverrideOk

`func (o *ServiceWorkloadDeploymentConfiguration) GetOverrideOk() (*ServiceDeploymentPolicy, bool)`

GetOverrideOk returns a tuple with the Override field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverride

`func (o *ServiceWorkloadDeploymentConfiguration) SetOverride(v ServiceDeploymentPolicy)`

SetOverride sets Override field to given value.


### GetEffective

`func (o *ServiceWorkloadDeploymentConfiguration) GetEffective() ServiceDeploymentPolicy`

GetEffective returns the Effective field if non-nil, zero value otherwise.

### GetEffectiveOk

`func (o *ServiceWorkloadDeploymentConfiguration) GetEffectiveOk() (*ServiceDeploymentPolicy, bool)`

GetEffectiveOk returns a tuple with the Effective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffective

`func (o *ServiceWorkloadDeploymentConfiguration) SetEffective(v ServiceDeploymentPolicy)`

SetEffective sets Effective field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


