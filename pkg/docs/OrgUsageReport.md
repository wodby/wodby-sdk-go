# OrgUsageReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Period** | [**OrgUsagePeriod**](OrgUsagePeriod.md) |  | 
**Amount** | **float64** | Total of all projects at the public list price in USD. | 
**InvoiceAmount** | **NullableFloat64** | What the invoices charged for the report&#39;s items in the period, in USD before tax. Null while it is unknown: for the current period, a period that was not invoiced, or while an invoice is a draft. | 
**Projects** | [**[]OrgUsageProject**](OrgUsageProject.md) | Projects ordered by title, with the organization row last. | 

## Methods

### NewOrgUsageReport

`func NewOrgUsageReport(period OrgUsagePeriod, amount float64, invoiceAmount NullableFloat64, projects []OrgUsageProject, ) *OrgUsageReport`

NewOrgUsageReport instantiates a new OrgUsageReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgUsageReportWithDefaults

`func NewOrgUsageReportWithDefaults() *OrgUsageReport`

NewOrgUsageReportWithDefaults instantiates a new OrgUsageReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriod

`func (o *OrgUsageReport) GetPeriod() OrgUsagePeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *OrgUsageReport) GetPeriodOk() (*OrgUsagePeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *OrgUsageReport) SetPeriod(v OrgUsagePeriod)`

SetPeriod sets Period field to given value.


### GetAmount

`func (o *OrgUsageReport) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *OrgUsageReport) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *OrgUsageReport) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetInvoiceAmount

`func (o *OrgUsageReport) GetInvoiceAmount() float64`

GetInvoiceAmount returns the InvoiceAmount field if non-nil, zero value otherwise.

### GetInvoiceAmountOk

`func (o *OrgUsageReport) GetInvoiceAmountOk() (*float64, bool)`

GetInvoiceAmountOk returns a tuple with the InvoiceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceAmount

`func (o *OrgUsageReport) SetInvoiceAmount(v float64)`

SetInvoiceAmount sets InvoiceAmount field to given value.


### SetInvoiceAmountNil

`func (o *OrgUsageReport) SetInvoiceAmountNil(b bool)`

 SetInvoiceAmountNil sets the value for InvoiceAmount to be an explicit nil

### UnsetInvoiceAmount
`func (o *OrgUsageReport) UnsetInvoiceAmount()`

UnsetInvoiceAmount ensures that no value is present for InvoiceAmount, not even an explicit nil
### GetProjects

`func (o *OrgUsageReport) GetProjects() []OrgUsageProject`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *OrgUsageReport) GetProjectsOk() (*[]OrgUsageProject, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *OrgUsageReport) SetProjects(v []OrgUsageProject)`

SetProjects sets Projects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


