# KubernetesVersionUpgradePlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentVersion** | **string** |  | 
**ProviderVersion** | **string** |  | 
**Provider** | **string** |  | 
**Supported** | **bool** |  | 
**Targets** | [**[]KubernetesVersionUpgradeTarget**](KubernetesVersionUpgradeTarget.md) |  | 
**Blockers** | **[]string** |  | 
**Warnings** | **[]string** |  | 
**ObservedAt** | **time.Time** |  | 

## Methods

### NewKubernetesVersionUpgradePlan

`func NewKubernetesVersionUpgradePlan(currentVersion string, providerVersion string, provider string, supported bool, targets []KubernetesVersionUpgradeTarget, blockers []string, warnings []string, observedAt time.Time, ) *KubernetesVersionUpgradePlan`

NewKubernetesVersionUpgradePlan instantiates a new KubernetesVersionUpgradePlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesVersionUpgradePlanWithDefaults

`func NewKubernetesVersionUpgradePlanWithDefaults() *KubernetesVersionUpgradePlan`

NewKubernetesVersionUpgradePlanWithDefaults instantiates a new KubernetesVersionUpgradePlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentVersion

`func (o *KubernetesVersionUpgradePlan) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *KubernetesVersionUpgradePlan) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *KubernetesVersionUpgradePlan) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.


### GetProviderVersion

`func (o *KubernetesVersionUpgradePlan) GetProviderVersion() string`

GetProviderVersion returns the ProviderVersion field if non-nil, zero value otherwise.

### GetProviderVersionOk

`func (o *KubernetesVersionUpgradePlan) GetProviderVersionOk() (*string, bool)`

GetProviderVersionOk returns a tuple with the ProviderVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderVersion

`func (o *KubernetesVersionUpgradePlan) SetProviderVersion(v string)`

SetProviderVersion sets ProviderVersion field to given value.


### GetProvider

`func (o *KubernetesVersionUpgradePlan) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *KubernetesVersionUpgradePlan) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *KubernetesVersionUpgradePlan) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetSupported

`func (o *KubernetesVersionUpgradePlan) GetSupported() bool`

GetSupported returns the Supported field if non-nil, zero value otherwise.

### GetSupportedOk

`func (o *KubernetesVersionUpgradePlan) GetSupportedOk() (*bool, bool)`

GetSupportedOk returns a tuple with the Supported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupported

`func (o *KubernetesVersionUpgradePlan) SetSupported(v bool)`

SetSupported sets Supported field to given value.


### GetTargets

`func (o *KubernetesVersionUpgradePlan) GetTargets() []KubernetesVersionUpgradeTarget`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *KubernetesVersionUpgradePlan) GetTargetsOk() (*[]KubernetesVersionUpgradeTarget, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *KubernetesVersionUpgradePlan) SetTargets(v []KubernetesVersionUpgradeTarget)`

SetTargets sets Targets field to given value.


### GetBlockers

`func (o *KubernetesVersionUpgradePlan) GetBlockers() []string`

GetBlockers returns the Blockers field if non-nil, zero value otherwise.

### GetBlockersOk

`func (o *KubernetesVersionUpgradePlan) GetBlockersOk() (*[]string, bool)`

GetBlockersOk returns a tuple with the Blockers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockers

`func (o *KubernetesVersionUpgradePlan) SetBlockers(v []string)`

SetBlockers sets Blockers field to given value.


### GetWarnings

`func (o *KubernetesVersionUpgradePlan) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *KubernetesVersionUpgradePlan) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *KubernetesVersionUpgradePlan) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.


### GetObservedAt

`func (o *KubernetesVersionUpgradePlan) GetObservedAt() time.Time`

GetObservedAt returns the ObservedAt field if non-nil, zero value otherwise.

### GetObservedAtOk

`func (o *KubernetesVersionUpgradePlan) GetObservedAtOk() (*time.Time, bool)`

GetObservedAtOk returns a tuple with the ObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAt

`func (o *KubernetesVersionUpgradePlan) SetObservedAt(v time.Time)`

SetObservedAt sets ObservedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


