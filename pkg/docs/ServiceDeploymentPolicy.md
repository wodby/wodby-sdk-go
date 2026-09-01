# ServiceDeploymentPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Strategy** | Pointer to [**ServiceDeploymentStrategy**](ServiceDeploymentStrategy.md) |  | [optional] 
**MaxUnavailable** | Pointer to **string** | Absolute number or percentage, such as &#x60;0&#x60; or &#x60;25%&#x60;. | [optional] 
**MaxSurge** | Pointer to **string** | Absolute number or percentage, such as &#x60;1&#x60; or &#x60;25%&#x60;. | [optional] 
**MinReady** | Pointer to **string** | Go duration, such as &#x60;10s&#x60;. | [optional] 
**ProgressDeadline** | Pointer to **string** | Go duration, such as &#x60;15m&#x60;. | [optional] 
**ShutdownGracePeriod** | Pointer to **string** | Go duration, such as &#x60;11m&#x60;. | [optional] 

## Methods

### NewServiceDeploymentPolicy

`func NewServiceDeploymentPolicy() *ServiceDeploymentPolicy`

NewServiceDeploymentPolicy instantiates a new ServiceDeploymentPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDeploymentPolicyWithDefaults

`func NewServiceDeploymentPolicyWithDefaults() *ServiceDeploymentPolicy`

NewServiceDeploymentPolicyWithDefaults instantiates a new ServiceDeploymentPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStrategy

`func (o *ServiceDeploymentPolicy) GetStrategy() ServiceDeploymentStrategy`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *ServiceDeploymentPolicy) GetStrategyOk() (*ServiceDeploymentStrategy, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *ServiceDeploymentPolicy) SetStrategy(v ServiceDeploymentStrategy)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *ServiceDeploymentPolicy) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetMaxUnavailable

`func (o *ServiceDeploymentPolicy) GetMaxUnavailable() string`

GetMaxUnavailable returns the MaxUnavailable field if non-nil, zero value otherwise.

### GetMaxUnavailableOk

`func (o *ServiceDeploymentPolicy) GetMaxUnavailableOk() (*string, bool)`

GetMaxUnavailableOk returns a tuple with the MaxUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUnavailable

`func (o *ServiceDeploymentPolicy) SetMaxUnavailable(v string)`

SetMaxUnavailable sets MaxUnavailable field to given value.

### HasMaxUnavailable

`func (o *ServiceDeploymentPolicy) HasMaxUnavailable() bool`

HasMaxUnavailable returns a boolean if a field has been set.

### GetMaxSurge

`func (o *ServiceDeploymentPolicy) GetMaxSurge() string`

GetMaxSurge returns the MaxSurge field if non-nil, zero value otherwise.

### GetMaxSurgeOk

`func (o *ServiceDeploymentPolicy) GetMaxSurgeOk() (*string, bool)`

GetMaxSurgeOk returns a tuple with the MaxSurge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSurge

`func (o *ServiceDeploymentPolicy) SetMaxSurge(v string)`

SetMaxSurge sets MaxSurge field to given value.

### HasMaxSurge

`func (o *ServiceDeploymentPolicy) HasMaxSurge() bool`

HasMaxSurge returns a boolean if a field has been set.

### GetMinReady

`func (o *ServiceDeploymentPolicy) GetMinReady() string`

GetMinReady returns the MinReady field if non-nil, zero value otherwise.

### GetMinReadyOk

`func (o *ServiceDeploymentPolicy) GetMinReadyOk() (*string, bool)`

GetMinReadyOk returns a tuple with the MinReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReady

`func (o *ServiceDeploymentPolicy) SetMinReady(v string)`

SetMinReady sets MinReady field to given value.

### HasMinReady

`func (o *ServiceDeploymentPolicy) HasMinReady() bool`

HasMinReady returns a boolean if a field has been set.

### GetProgressDeadline

`func (o *ServiceDeploymentPolicy) GetProgressDeadline() string`

GetProgressDeadline returns the ProgressDeadline field if non-nil, zero value otherwise.

### GetProgressDeadlineOk

`func (o *ServiceDeploymentPolicy) GetProgressDeadlineOk() (*string, bool)`

GetProgressDeadlineOk returns a tuple with the ProgressDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressDeadline

`func (o *ServiceDeploymentPolicy) SetProgressDeadline(v string)`

SetProgressDeadline sets ProgressDeadline field to given value.

### HasProgressDeadline

`func (o *ServiceDeploymentPolicy) HasProgressDeadline() bool`

HasProgressDeadline returns a boolean if a field has been set.

### GetShutdownGracePeriod

`func (o *ServiceDeploymentPolicy) GetShutdownGracePeriod() string`

GetShutdownGracePeriod returns the ShutdownGracePeriod field if non-nil, zero value otherwise.

### GetShutdownGracePeriodOk

`func (o *ServiceDeploymentPolicy) GetShutdownGracePeriodOk() (*string, bool)`

GetShutdownGracePeriodOk returns a tuple with the ShutdownGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownGracePeriod

`func (o *ServiceDeploymentPolicy) SetShutdownGracePeriod(v string)`

SetShutdownGracePeriod sets ShutdownGracePeriod field to given value.

### HasShutdownGracePeriod

`func (o *ServiceDeploymentPolicy) HasShutdownGracePeriod() bool`

HasShutdownGracePeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


