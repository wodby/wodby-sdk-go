# OrgUsageProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **NullableInt32** | Null for usage of organization-owned resources and usage without a project. | 
**Name** | **string** | The project name, empty for the organization row and deleted projects. | 
**Title** | **string** |  | 
**Deleted** | **bool** | The project no longer exists. Its recorded usage stays in the report. | 
**Amount** | **float64** | Total of the items at the public list price in USD. | 
**InvoiceShare** | **NullableFloat64** | Total of the items&#39; invoice shares in USD before tax, null while the invoiced amount is unknown. | 
**Items** | [**[]OrgUsageItem**](OrgUsageItem.md) |  | 

## Methods

### NewOrgUsageProject

`func NewOrgUsageProject(projectId NullableInt32, name string, title string, deleted bool, amount float64, invoiceShare NullableFloat64, items []OrgUsageItem, ) *OrgUsageProject`

NewOrgUsageProject instantiates a new OrgUsageProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgUsageProjectWithDefaults

`func NewOrgUsageProjectWithDefaults() *OrgUsageProject`

NewOrgUsageProjectWithDefaults instantiates a new OrgUsageProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *OrgUsageProject) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *OrgUsageProject) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *OrgUsageProject) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### SetProjectIdNil

`func (o *OrgUsageProject) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *OrgUsageProject) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetName

`func (o *OrgUsageProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrgUsageProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrgUsageProject) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *OrgUsageProject) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OrgUsageProject) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OrgUsageProject) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDeleted

`func (o *OrgUsageProject) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *OrgUsageProject) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *OrgUsageProject) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.


### GetAmount

`func (o *OrgUsageProject) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *OrgUsageProject) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *OrgUsageProject) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetInvoiceShare

`func (o *OrgUsageProject) GetInvoiceShare() float64`

GetInvoiceShare returns the InvoiceShare field if non-nil, zero value otherwise.

### GetInvoiceShareOk

`func (o *OrgUsageProject) GetInvoiceShareOk() (*float64, bool)`

GetInvoiceShareOk returns a tuple with the InvoiceShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceShare

`func (o *OrgUsageProject) SetInvoiceShare(v float64)`

SetInvoiceShare sets InvoiceShare field to given value.


### SetInvoiceShareNil

`func (o *OrgUsageProject) SetInvoiceShareNil(b bool)`

 SetInvoiceShareNil sets the value for InvoiceShare to be an explicit nil

### UnsetInvoiceShare
`func (o *OrgUsageProject) UnsetInvoiceShare()`

UnsetInvoiceShare ensures that no value is present for InvoiceShare, not even an explicit nil
### GetItems

`func (o *OrgUsageProject) GetItems() []OrgUsageItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *OrgUsageProject) GetItemsOk() (*[]OrgUsageItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *OrgUsageProject) SetItems(v []OrgUsageItem)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


