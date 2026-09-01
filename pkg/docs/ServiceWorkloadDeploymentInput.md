# ServiceWorkloadDeploymentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Strategy** | Pointer to [**ServiceDeploymentStrategy**](ServiceDeploymentStrategy.md) |  | [optional] 
**MaxUnavailable** | Pointer to **string** | Absolute number or percentage, such as &#x60;0&#x60; or &#x60;25%&#x60;. | [optional] 
**MaxSurge** | Pointer to **string** | Absolute number or percentage, such as &#x60;1&#x60; or &#x60;25%&#x60;. | [optional] 
**MinReady** | Pointer to **string** | Go duration, such as &#x60;10s&#x60;. | [optional] 
**ProgressDeadline** | Pointer to **string** | Go duration, such as &#x60;15m&#x60;. | [optional] 
**ShutdownGracePeriod** | Pointer to **string** | Go duration, such as &#x60;11m&#x60;. | [optional] 
**Name** | **string** |  | 

## Methods

### NewServiceWorkloadDeploymentInput

`func NewServiceWorkloadDeploymentInput(name string, ) *ServiceWorkloadDeploymentInput`

NewServiceWorkloadDeploymentInput instantiates a new ServiceWorkloadDeploymentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWorkloadDeploymentInputWithDefaults

`func NewServiceWorkloadDeploymentInputWithDefaults() *ServiceWorkloadDeploymentInput`

NewServiceWorkloadDeploymentInputWithDefaults instantiates a new ServiceWorkloadDeploymentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStrategy

`func (o *ServiceWorkloadDeploymentInput) GetStrategy() ServiceDeploymentStrategy`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *ServiceWorkloadDeploymentInput) GetStrategyOk() (*ServiceDeploymentStrategy, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *ServiceWorkloadDeploymentInput) SetStrategy(v ServiceDeploymentStrategy)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *ServiceWorkloadDeploymentInput) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetMaxUnavailable

`func (o *ServiceWorkloadDeploymentInput) GetMaxUnavailable() string`

GetMaxUnavailable returns the MaxUnavailable field if non-nil, zero value otherwise.

### GetMaxUnavailableOk

`func (o *ServiceWorkloadDeploymentInput) GetMaxUnavailableOk() (*string, bool)`

GetMaxUnavailableOk returns a tuple with the MaxUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUnavailable

`func (o *ServiceWorkloadDeploymentInput) SetMaxUnavailable(v string)`

SetMaxUnavailable sets MaxUnavailable field to given value.

### HasMaxUnavailable

`func (o *ServiceWorkloadDeploymentInput) HasMaxUnavailable() bool`

HasMaxUnavailable returns a boolean if a field has been set.

### GetMaxSurge

`func (o *ServiceWorkloadDeploymentInput) GetMaxSurge() string`

GetMaxSurge returns the MaxSurge field if non-nil, zero value otherwise.

### GetMaxSurgeOk

`func (o *ServiceWorkloadDeploymentInput) GetMaxSurgeOk() (*string, bool)`

GetMaxSurgeOk returns a tuple with the MaxSurge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSurge

`func (o *ServiceWorkloadDeploymentInput) SetMaxSurge(v string)`

SetMaxSurge sets MaxSurge field to given value.

### HasMaxSurge

`func (o *ServiceWorkloadDeploymentInput) HasMaxSurge() bool`

HasMaxSurge returns a boolean if a field has been set.

### GetMinReady

`func (o *ServiceWorkloadDeploymentInput) GetMinReady() string`

GetMinReady returns the MinReady field if non-nil, zero value otherwise.

### GetMinReadyOk

`func (o *ServiceWorkloadDeploymentInput) GetMinReadyOk() (*string, bool)`

GetMinReadyOk returns a tuple with the MinReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReady

`func (o *ServiceWorkloadDeploymentInput) SetMinReady(v string)`

SetMinReady sets MinReady field to given value.

### HasMinReady

`func (o *ServiceWorkloadDeploymentInput) HasMinReady() bool`

HasMinReady returns a boolean if a field has been set.

### GetProgressDeadline

`func (o *ServiceWorkloadDeploymentInput) GetProgressDeadline() string`

GetProgressDeadline returns the ProgressDeadline field if non-nil, zero value otherwise.

### GetProgressDeadlineOk

`func (o *ServiceWorkloadDeploymentInput) GetProgressDeadlineOk() (*string, bool)`

GetProgressDeadlineOk returns a tuple with the ProgressDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressDeadline

`func (o *ServiceWorkloadDeploymentInput) SetProgressDeadline(v string)`

SetProgressDeadline sets ProgressDeadline field to given value.

### HasProgressDeadline

`func (o *ServiceWorkloadDeploymentInput) HasProgressDeadline() bool`

HasProgressDeadline returns a boolean if a field has been set.

### GetShutdownGracePeriod

`func (o *ServiceWorkloadDeploymentInput) GetShutdownGracePeriod() string`

GetShutdownGracePeriod returns the ShutdownGracePeriod field if non-nil, zero value otherwise.

### GetShutdownGracePeriodOk

`func (o *ServiceWorkloadDeploymentInput) GetShutdownGracePeriodOk() (*string, bool)`

GetShutdownGracePeriodOk returns a tuple with the ShutdownGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownGracePeriod

`func (o *ServiceWorkloadDeploymentInput) SetShutdownGracePeriod(v string)`

SetShutdownGracePeriod sets ShutdownGracePeriod field to given value.

### HasShutdownGracePeriod

`func (o *ServiceWorkloadDeploymentInput) HasShutdownGracePeriod() bool`

HasShutdownGracePeriod returns a boolean if a field has been set.

### GetName

`func (o *ServiceWorkloadDeploymentInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceWorkloadDeploymentInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceWorkloadDeploymentInput) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


