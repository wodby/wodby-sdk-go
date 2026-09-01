# ServiceDeploymentConfigurationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Strategy** | Pointer to [**ServiceDeploymentStrategy**](ServiceDeploymentStrategy.md) |  | [optional] 
**MaxUnavailable** | Pointer to **string** | Absolute number or percentage, such as &#x60;0&#x60; or &#x60;25%&#x60;. | [optional] 
**MaxSurge** | Pointer to **string** | Absolute number or percentage, such as &#x60;1&#x60; or &#x60;25%&#x60;. | [optional] 
**MinReady** | Pointer to **string** | Go duration, such as &#x60;10s&#x60;. | [optional] 
**ProgressDeadline** | Pointer to **string** | Go duration, such as &#x60;15m&#x60;. | [optional] 
**ShutdownGracePeriod** | Pointer to **string** | Go duration, such as &#x60;11m&#x60;. | [optional] 
**Workloads** | Pointer to [**[]ServiceWorkloadDeploymentInput**](ServiceWorkloadDeploymentInput.md) |  | [optional] 

## Methods

### NewServiceDeploymentConfigurationInput

`func NewServiceDeploymentConfigurationInput() *ServiceDeploymentConfigurationInput`

NewServiceDeploymentConfigurationInput instantiates a new ServiceDeploymentConfigurationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDeploymentConfigurationInputWithDefaults

`func NewServiceDeploymentConfigurationInputWithDefaults() *ServiceDeploymentConfigurationInput`

NewServiceDeploymentConfigurationInputWithDefaults instantiates a new ServiceDeploymentConfigurationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStrategy

`func (o *ServiceDeploymentConfigurationInput) GetStrategy() ServiceDeploymentStrategy`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *ServiceDeploymentConfigurationInput) GetStrategyOk() (*ServiceDeploymentStrategy, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *ServiceDeploymentConfigurationInput) SetStrategy(v ServiceDeploymentStrategy)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *ServiceDeploymentConfigurationInput) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetMaxUnavailable

`func (o *ServiceDeploymentConfigurationInput) GetMaxUnavailable() string`

GetMaxUnavailable returns the MaxUnavailable field if non-nil, zero value otherwise.

### GetMaxUnavailableOk

`func (o *ServiceDeploymentConfigurationInput) GetMaxUnavailableOk() (*string, bool)`

GetMaxUnavailableOk returns a tuple with the MaxUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUnavailable

`func (o *ServiceDeploymentConfigurationInput) SetMaxUnavailable(v string)`

SetMaxUnavailable sets MaxUnavailable field to given value.

### HasMaxUnavailable

`func (o *ServiceDeploymentConfigurationInput) HasMaxUnavailable() bool`

HasMaxUnavailable returns a boolean if a field has been set.

### GetMaxSurge

`func (o *ServiceDeploymentConfigurationInput) GetMaxSurge() string`

GetMaxSurge returns the MaxSurge field if non-nil, zero value otherwise.

### GetMaxSurgeOk

`func (o *ServiceDeploymentConfigurationInput) GetMaxSurgeOk() (*string, bool)`

GetMaxSurgeOk returns a tuple with the MaxSurge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSurge

`func (o *ServiceDeploymentConfigurationInput) SetMaxSurge(v string)`

SetMaxSurge sets MaxSurge field to given value.

### HasMaxSurge

`func (o *ServiceDeploymentConfigurationInput) HasMaxSurge() bool`

HasMaxSurge returns a boolean if a field has been set.

### GetMinReady

`func (o *ServiceDeploymentConfigurationInput) GetMinReady() string`

GetMinReady returns the MinReady field if non-nil, zero value otherwise.

### GetMinReadyOk

`func (o *ServiceDeploymentConfigurationInput) GetMinReadyOk() (*string, bool)`

GetMinReadyOk returns a tuple with the MinReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReady

`func (o *ServiceDeploymentConfigurationInput) SetMinReady(v string)`

SetMinReady sets MinReady field to given value.

### HasMinReady

`func (o *ServiceDeploymentConfigurationInput) HasMinReady() bool`

HasMinReady returns a boolean if a field has been set.

### GetProgressDeadline

`func (o *ServiceDeploymentConfigurationInput) GetProgressDeadline() string`

GetProgressDeadline returns the ProgressDeadline field if non-nil, zero value otherwise.

### GetProgressDeadlineOk

`func (o *ServiceDeploymentConfigurationInput) GetProgressDeadlineOk() (*string, bool)`

GetProgressDeadlineOk returns a tuple with the ProgressDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressDeadline

`func (o *ServiceDeploymentConfigurationInput) SetProgressDeadline(v string)`

SetProgressDeadline sets ProgressDeadline field to given value.

### HasProgressDeadline

`func (o *ServiceDeploymentConfigurationInput) HasProgressDeadline() bool`

HasProgressDeadline returns a boolean if a field has been set.

### GetShutdownGracePeriod

`func (o *ServiceDeploymentConfigurationInput) GetShutdownGracePeriod() string`

GetShutdownGracePeriod returns the ShutdownGracePeriod field if non-nil, zero value otherwise.

### GetShutdownGracePeriodOk

`func (o *ServiceDeploymentConfigurationInput) GetShutdownGracePeriodOk() (*string, bool)`

GetShutdownGracePeriodOk returns a tuple with the ShutdownGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownGracePeriod

`func (o *ServiceDeploymentConfigurationInput) SetShutdownGracePeriod(v string)`

SetShutdownGracePeriod sets ShutdownGracePeriod field to given value.

### HasShutdownGracePeriod

`func (o *ServiceDeploymentConfigurationInput) HasShutdownGracePeriod() bool`

HasShutdownGracePeriod returns a boolean if a field has been set.

### GetWorkloads

`func (o *ServiceDeploymentConfigurationInput) GetWorkloads() []ServiceWorkloadDeploymentInput`

GetWorkloads returns the Workloads field if non-nil, zero value otherwise.

### GetWorkloadsOk

`func (o *ServiceDeploymentConfigurationInput) GetWorkloadsOk() (*[]ServiceWorkloadDeploymentInput, bool)`

GetWorkloadsOk returns a tuple with the Workloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloads

`func (o *ServiceDeploymentConfigurationInput) SetWorkloads(v []ServiceWorkloadDeploymentInput)`

SetWorkloads sets Workloads field to given value.

### HasWorkloads

`func (o *ServiceDeploymentConfigurationInput) HasWorkloads() bool`

HasWorkloads returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


