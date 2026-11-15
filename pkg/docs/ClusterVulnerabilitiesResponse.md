# ClusterVulnerabilitiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ClusterVulnerability**](ClusterVulnerability.md) |  | 
**TotalCount** | **int32** |  | 
**NextPage** | Pointer to **int32** |  | [optional] 
**Counts** | [**VulnerabilityCounts**](VulnerabilityCounts.md) |  | 
**Sources** | [**[]ClusterVulnerabilitySource**](ClusterVulnerabilitySource.md) |  | 

## Methods

### NewClusterVulnerabilitiesResponse

`func NewClusterVulnerabilitiesResponse(items []ClusterVulnerability, totalCount int32, counts VulnerabilityCounts, sources []ClusterVulnerabilitySource, ) *ClusterVulnerabilitiesResponse`

NewClusterVulnerabilitiesResponse instantiates a new ClusterVulnerabilitiesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterVulnerabilitiesResponseWithDefaults

`func NewClusterVulnerabilitiesResponseWithDefaults() *ClusterVulnerabilitiesResponse`

NewClusterVulnerabilitiesResponseWithDefaults instantiates a new ClusterVulnerabilitiesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ClusterVulnerabilitiesResponse) GetItems() []ClusterVulnerability`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ClusterVulnerabilitiesResponse) GetItemsOk() (*[]ClusterVulnerability, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ClusterVulnerabilitiesResponse) SetItems(v []ClusterVulnerability)`

SetItems sets Items field to given value.


### GetTotalCount

`func (o *ClusterVulnerabilitiesResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ClusterVulnerabilitiesResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ClusterVulnerabilitiesResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetNextPage

`func (o *ClusterVulnerabilitiesResponse) GetNextPage() int32`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *ClusterVulnerabilitiesResponse) GetNextPageOk() (*int32, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *ClusterVulnerabilitiesResponse) SetNextPage(v int32)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *ClusterVulnerabilitiesResponse) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### GetCounts

`func (o *ClusterVulnerabilitiesResponse) GetCounts() VulnerabilityCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *ClusterVulnerabilitiesResponse) GetCountsOk() (*VulnerabilityCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *ClusterVulnerabilitiesResponse) SetCounts(v VulnerabilityCounts)`

SetCounts sets Counts field to given value.


### GetSources

`func (o *ClusterVulnerabilitiesResponse) GetSources() []ClusterVulnerabilitySource`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *ClusterVulnerabilitiesResponse) GetSourcesOk() (*[]ClusterVulnerabilitySource, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *ClusterVulnerabilitiesResponse) SetSources(v []ClusterVulnerabilitySource)`

SetSources sets Sources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


