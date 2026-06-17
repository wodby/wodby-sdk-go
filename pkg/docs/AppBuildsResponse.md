# AppBuildsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]AppBuild**](AppBuild.md) |  | 
**TotalCount** | **int32** |  | 
**NextPage** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewAppBuildsResponse

`func NewAppBuildsResponse(items []AppBuild, totalCount int32, ) *AppBuildsResponse`

NewAppBuildsResponse instantiates a new AppBuildsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppBuildsResponseWithDefaults

`func NewAppBuildsResponseWithDefaults() *AppBuildsResponse`

NewAppBuildsResponseWithDefaults instantiates a new AppBuildsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *AppBuildsResponse) GetItems() []AppBuild`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AppBuildsResponse) GetItemsOk() (*[]AppBuild, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AppBuildsResponse) SetItems(v []AppBuild)`

SetItems sets Items field to given value.


### GetTotalCount

`func (o *AppBuildsResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AppBuildsResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AppBuildsResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetNextPage

`func (o *AppBuildsResponse) GetNextPage() int32`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *AppBuildsResponse) GetNextPageOk() (*int32, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *AppBuildsResponse) SetNextPage(v int32)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *AppBuildsResponse) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### SetNextPageNil

`func (o *AppBuildsResponse) SetNextPageNil(b bool)`

 SetNextPageNil sets the value for NextPage to be an explicit nil

### UnsetNextPage
`func (o *AppBuildsResponse) UnsetNextPage()`

UnsetNextPage ensures that no value is present for NextPage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


