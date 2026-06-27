# AppServiceCronJobsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]AppServiceCronJob**](AppServiceCronJob.md) |  | 
**TotalCount** | **int32** |  | 
**NextPage** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewAppServiceCronJobsResponse

`func NewAppServiceCronJobsResponse(items []AppServiceCronJob, totalCount int32, ) *AppServiceCronJobsResponse`

NewAppServiceCronJobsResponse instantiates a new AppServiceCronJobsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceCronJobsResponseWithDefaults

`func NewAppServiceCronJobsResponseWithDefaults() *AppServiceCronJobsResponse`

NewAppServiceCronJobsResponseWithDefaults instantiates a new AppServiceCronJobsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *AppServiceCronJobsResponse) GetItems() []AppServiceCronJob`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AppServiceCronJobsResponse) GetItemsOk() (*[]AppServiceCronJob, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AppServiceCronJobsResponse) SetItems(v []AppServiceCronJob)`

SetItems sets Items field to given value.


### GetTotalCount

`func (o *AppServiceCronJobsResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AppServiceCronJobsResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AppServiceCronJobsResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetNextPage

`func (o *AppServiceCronJobsResponse) GetNextPage() int32`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *AppServiceCronJobsResponse) GetNextPageOk() (*int32, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *AppServiceCronJobsResponse) SetNextPage(v int32)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *AppServiceCronJobsResponse) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### SetNextPageNil

`func (o *AppServiceCronJobsResponse) SetNextPageNil(b bool)`

 SetNextPageNil sets the value for NextPage to be an explicit nil

### UnsetNextPage
`func (o *AppServiceCronJobsResponse) UnsetNextPage()`

UnsetNextPage ensures that no value is present for NextPage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


