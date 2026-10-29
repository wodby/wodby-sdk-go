# AppPreviewPullRequestList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configured** | **bool** | True when preview environments are enabled and the base environment exists. | 
**PullRequests** | [**[]AppPreviewPullRequest**](AppPreviewPullRequest.md) |  | 

## Methods

### NewAppPreviewPullRequestList

`func NewAppPreviewPullRequestList(configured bool, pullRequests []AppPreviewPullRequest, ) *AppPreviewPullRequestList`

NewAppPreviewPullRequestList instantiates a new AppPreviewPullRequestList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPreviewPullRequestListWithDefaults

`func NewAppPreviewPullRequestListWithDefaults() *AppPreviewPullRequestList`

NewAppPreviewPullRequestListWithDefaults instantiates a new AppPreviewPullRequestList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigured

`func (o *AppPreviewPullRequestList) GetConfigured() bool`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *AppPreviewPullRequestList) GetConfiguredOk() (*bool, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *AppPreviewPullRequestList) SetConfigured(v bool)`

SetConfigured sets Configured field to given value.


### GetPullRequests

`func (o *AppPreviewPullRequestList) GetPullRequests() []AppPreviewPullRequest`

GetPullRequests returns the PullRequests field if non-nil, zero value otherwise.

### GetPullRequestsOk

`func (o *AppPreviewPullRequestList) GetPullRequestsOk() (*[]AppPreviewPullRequest, bool)`

GetPullRequestsOk returns a tuple with the PullRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequests

`func (o *AppPreviewPullRequestList) SetPullRequests(v []AppPreviewPullRequest)`

SetPullRequests sets PullRequests field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


