# AppServiceCapacityPreflight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowed** | **bool** | Whether the proposed usage increase is permitted by the effective backend policy. | 
**Enforced** | **bool** | Whether subscription capacity limits apply to this organization. | 
**Usage** | **float64** |  | 
**UsageIncluded** | **float64** |  | 
**ProjectedUsage** | **float64** |  | 

## Methods

### NewAppServiceCapacityPreflight

`func NewAppServiceCapacityPreflight(allowed bool, enforced bool, usage float64, usageIncluded float64, projectedUsage float64, ) *AppServiceCapacityPreflight`

NewAppServiceCapacityPreflight instantiates a new AppServiceCapacityPreflight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceCapacityPreflightWithDefaults

`func NewAppServiceCapacityPreflightWithDefaults() *AppServiceCapacityPreflight`

NewAppServiceCapacityPreflightWithDefaults instantiates a new AppServiceCapacityPreflight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowed

`func (o *AppServiceCapacityPreflight) GetAllowed() bool`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *AppServiceCapacityPreflight) GetAllowedOk() (*bool, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *AppServiceCapacityPreflight) SetAllowed(v bool)`

SetAllowed sets Allowed field to given value.


### GetEnforced

`func (o *AppServiceCapacityPreflight) GetEnforced() bool`

GetEnforced returns the Enforced field if non-nil, zero value otherwise.

### GetEnforcedOk

`func (o *AppServiceCapacityPreflight) GetEnforcedOk() (*bool, bool)`

GetEnforcedOk returns a tuple with the Enforced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforced

`func (o *AppServiceCapacityPreflight) SetEnforced(v bool)`

SetEnforced sets Enforced field to given value.


### GetUsage

`func (o *AppServiceCapacityPreflight) GetUsage() float64`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *AppServiceCapacityPreflight) GetUsageOk() (*float64, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *AppServiceCapacityPreflight) SetUsage(v float64)`

SetUsage sets Usage field to given value.


### GetUsageIncluded

`func (o *AppServiceCapacityPreflight) GetUsageIncluded() float64`

GetUsageIncluded returns the UsageIncluded field if non-nil, zero value otherwise.

### GetUsageIncludedOk

`func (o *AppServiceCapacityPreflight) GetUsageIncludedOk() (*float64, bool)`

GetUsageIncludedOk returns a tuple with the UsageIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageIncluded

`func (o *AppServiceCapacityPreflight) SetUsageIncluded(v float64)`

SetUsageIncluded sets UsageIncluded field to given value.


### GetProjectedUsage

`func (o *AppServiceCapacityPreflight) GetProjectedUsage() float64`

GetProjectedUsage returns the ProjectedUsage field if non-nil, zero value otherwise.

### GetProjectedUsageOk

`func (o *AppServiceCapacityPreflight) GetProjectedUsageOk() (*float64, bool)`

GetProjectedUsageOk returns a tuple with the ProjectedUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectedUsage

`func (o *AppServiceCapacityPreflight) SetProjectedUsage(v float64)`

SetProjectedUsage sets ProjectedUsage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


