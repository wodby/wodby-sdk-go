# ClusterVulnerabilitiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ClusterVulnerability**](ClusterVulnerability.md) |  | 
**TotalCount** | **int32** |  | 
**NextPage** | Pointer to **int32** |  | [optional] 
**Counts** | [**VulnerabilityCounts**](VulnerabilityCounts.md) |  | 
**Sources** | [**[]ClusterVulnerabilitySource**](ClusterVulnerabilitySource.md) |  | 
**Scanning** | **bool** | Whether a vulnerability check of the cluster is waiting or running. Until it ends, the list shows the results of the checks that finished before it. | 
**RestartRequired** | **int32** | The cluster&#39;s critical and high vulnerabilities, not ignored, that a restart of its K3s servers fixes: the fix is installed and runs after the restart. Covers the whole cluster, whatever the filter. | 
**WaitingForUpstream** | **int32** | The cluster&#39;s critical and high vulnerabilities, not ignored, that wait for a fix from an infrastructure image&#39;s maintainers. They aren&#39;t in counts. Covers the whole cluster, whatever the filter. | 

## Methods

### NewClusterVulnerabilitiesResponse

`func NewClusterVulnerabilitiesResponse(items []ClusterVulnerability, totalCount int32, counts VulnerabilityCounts, sources []ClusterVulnerabilitySource, scanning bool, restartRequired int32, waitingForUpstream int32, ) *ClusterVulnerabilitiesResponse`

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


### GetScanning

`func (o *ClusterVulnerabilitiesResponse) GetScanning() bool`

GetScanning returns the Scanning field if non-nil, zero value otherwise.

### GetScanningOk

`func (o *ClusterVulnerabilitiesResponse) GetScanningOk() (*bool, bool)`

GetScanningOk returns a tuple with the Scanning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanning

`func (o *ClusterVulnerabilitiesResponse) SetScanning(v bool)`

SetScanning sets Scanning field to given value.


### GetRestartRequired

`func (o *ClusterVulnerabilitiesResponse) GetRestartRequired() int32`

GetRestartRequired returns the RestartRequired field if non-nil, zero value otherwise.

### GetRestartRequiredOk

`func (o *ClusterVulnerabilitiesResponse) GetRestartRequiredOk() (*int32, bool)`

GetRestartRequiredOk returns a tuple with the RestartRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartRequired

`func (o *ClusterVulnerabilitiesResponse) SetRestartRequired(v int32)`

SetRestartRequired sets RestartRequired field to given value.


### GetWaitingForUpstream

`func (o *ClusterVulnerabilitiesResponse) GetWaitingForUpstream() int32`

GetWaitingForUpstream returns the WaitingForUpstream field if non-nil, zero value otherwise.

### GetWaitingForUpstreamOk

`func (o *ClusterVulnerabilitiesResponse) GetWaitingForUpstreamOk() (*int32, bool)`

GetWaitingForUpstreamOk returns a tuple with the WaitingForUpstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitingForUpstream

`func (o *ClusterVulnerabilitiesResponse) SetWaitingForUpstream(v int32)`

SetWaitingForUpstream sets WaitingForUpstream field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


