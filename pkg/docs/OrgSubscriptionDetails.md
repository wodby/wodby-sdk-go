# OrgSubscriptionDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**BillingSubscriptionStatus**](BillingSubscriptionStatus.md) |  | 
**Plan** | [**OrgSubscriptionPlanDetails**](OrgSubscriptionPlanDetails.md) |  | 

## Methods

### NewOrgSubscriptionDetails

`func NewOrgSubscriptionDetails(status BillingSubscriptionStatus, plan OrgSubscriptionPlanDetails, ) *OrgSubscriptionDetails`

NewOrgSubscriptionDetails instantiates a new OrgSubscriptionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgSubscriptionDetailsWithDefaults

`func NewOrgSubscriptionDetailsWithDefaults() *OrgSubscriptionDetails`

NewOrgSubscriptionDetailsWithDefaults instantiates a new OrgSubscriptionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *OrgSubscriptionDetails) GetStatus() BillingSubscriptionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrgSubscriptionDetails) GetStatusOk() (*BillingSubscriptionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrgSubscriptionDetails) SetStatus(v BillingSubscriptionStatus)`

SetStatus sets Status field to given value.


### GetPlan

`func (o *OrgSubscriptionDetails) GetPlan() OrgSubscriptionPlanDetails`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *OrgSubscriptionDetails) GetPlanOk() (*OrgSubscriptionPlanDetails, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *OrgSubscriptionDetails) SetPlan(v OrgSubscriptionPlanDetails)`

SetPlan sets Plan field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


