# OrgSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Plan** | Pointer to [**OrgSubscriptionPlan**](OrgSubscriptionPlan.md) |  | [optional] 

## Methods

### NewOrgSubscription

`func NewOrgSubscription(status string, ) *OrgSubscription`

NewOrgSubscription instantiates a new OrgSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgSubscriptionWithDefaults

`func NewOrgSubscriptionWithDefaults() *OrgSubscription`

NewOrgSubscriptionWithDefaults instantiates a new OrgSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *OrgSubscription) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrgSubscription) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrgSubscription) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPlan

`func (o *OrgSubscription) GetPlan() OrgSubscriptionPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *OrgSubscription) GetPlanOk() (*OrgSubscriptionPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *OrgSubscription) SetPlan(v OrgSubscriptionPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *OrgSubscription) HasPlan() bool`

HasPlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


