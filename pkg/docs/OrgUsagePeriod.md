# OrgUsagePeriod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | **time.Time** |  | 
**End** | **time.Time** | Where the period ends. A past period ends where the next one starts; a plan change starts a new period early. | 
**Current** | **bool** | The subscription&#39;s current period, whose usage is still growing. | 

## Methods

### NewOrgUsagePeriod

`func NewOrgUsagePeriod(start time.Time, end time.Time, current bool, ) *OrgUsagePeriod`

NewOrgUsagePeriod instantiates a new OrgUsagePeriod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgUsagePeriodWithDefaults

`func NewOrgUsagePeriodWithDefaults() *OrgUsagePeriod`

NewOrgUsagePeriodWithDefaults instantiates a new OrgUsagePeriod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *OrgUsagePeriod) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *OrgUsagePeriod) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *OrgUsagePeriod) SetStart(v time.Time)`

SetStart sets Start field to given value.


### GetEnd

`func (o *OrgUsagePeriod) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *OrgUsagePeriod) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *OrgUsagePeriod) SetEnd(v time.Time)`

SetEnd sets End field to given value.


### GetCurrent

`func (o *OrgUsagePeriod) GetCurrent() bool`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *OrgUsagePeriod) GetCurrentOk() (*bool, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *OrgUsagePeriod) SetCurrent(v bool)`

SetCurrent sets Current field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


