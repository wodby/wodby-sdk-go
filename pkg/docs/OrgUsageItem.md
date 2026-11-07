# OrgUsageItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Item** | **string** |  | 
**Title** | **string** |  | 
**Quantity** | **float64** | Usage in unit over the period. | 
**Unit** | **string** | service-hour, minute, GB-hour, or USD for Wodby Cloud usage. | 
**Amount** | **float64** | Quantity at the public list price in USD, before included usage and discounts. | 
**InvoiceShare** | **NullableFloat64** | This entry&#39;s share of what the invoices charged for the item&#39;s usage in the period, in USD before tax: the invoiced amount after included usage, discounts, and prorations, split between projects in proportion to amount. An item invoiced without recorded usage is shared to the organization row. Null while the invoiced amount is unknown. | 

## Methods

### NewOrgUsageItem

`func NewOrgUsageItem(item string, title string, quantity float64, unit string, amount float64, invoiceShare NullableFloat64, ) *OrgUsageItem`

NewOrgUsageItem instantiates a new OrgUsageItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgUsageItemWithDefaults

`func NewOrgUsageItemWithDefaults() *OrgUsageItem`

NewOrgUsageItemWithDefaults instantiates a new OrgUsageItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItem

`func (o *OrgUsageItem) GetItem() string`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *OrgUsageItem) GetItemOk() (*string, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *OrgUsageItem) SetItem(v string)`

SetItem sets Item field to given value.


### GetTitle

`func (o *OrgUsageItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OrgUsageItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OrgUsageItem) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetQuantity

`func (o *OrgUsageItem) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrgUsageItem) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrgUsageItem) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.


### GetUnit

`func (o *OrgUsageItem) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *OrgUsageItem) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *OrgUsageItem) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetAmount

`func (o *OrgUsageItem) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *OrgUsageItem) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *OrgUsageItem) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetInvoiceShare

`func (o *OrgUsageItem) GetInvoiceShare() float64`

GetInvoiceShare returns the InvoiceShare field if non-nil, zero value otherwise.

### GetInvoiceShareOk

`func (o *OrgUsageItem) GetInvoiceShareOk() (*float64, bool)`

GetInvoiceShareOk returns a tuple with the InvoiceShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceShare

`func (o *OrgUsageItem) SetInvoiceShare(v float64)`

SetInvoiceShare sets InvoiceShare field to given value.


### SetInvoiceShareNil

`func (o *OrgUsageItem) SetInvoiceShareNil(b bool)`

 SetInvoiceShareNil sets the value for InvoiceShare to be an explicit nil

### UnsetInvoiceShare
`func (o *OrgUsageItem) UnsetInvoiceShare()`

UnsetInvoiceShare ensures that no value is present for InvoiceShare, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


