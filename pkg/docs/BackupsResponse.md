# BackupsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]Backup**](Backup.md) |  | 
**TotalCount** | **int32** |  | 
**NextPage** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewBackupsResponse

`func NewBackupsResponse(items []Backup, totalCount int32, ) *BackupsResponse`

NewBackupsResponse instantiates a new BackupsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupsResponseWithDefaults

`func NewBackupsResponseWithDefaults() *BackupsResponse`

NewBackupsResponseWithDefaults instantiates a new BackupsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *BackupsResponse) GetItems() []Backup`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BackupsResponse) GetItemsOk() (*[]Backup, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BackupsResponse) SetItems(v []Backup)`

SetItems sets Items field to given value.


### GetTotalCount

`func (o *BackupsResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *BackupsResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *BackupsResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetNextPage

`func (o *BackupsResponse) GetNextPage() int32`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *BackupsResponse) GetNextPageOk() (*int32, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *BackupsResponse) SetNextPage(v int32)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *BackupsResponse) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### SetNextPageNil

`func (o *BackupsResponse) SetNextPageNil(b bool)`

 SetNextPageNil sets the value for NextPage to be an explicit nil

### UnsetNextPage
`func (o *BackupsResponse) UnsetNextPage()`

UnsetNextPage ensures that no value is present for NextPage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


