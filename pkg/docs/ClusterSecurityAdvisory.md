# ClusterSecurityAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Advisory** | **string** | The advisory, such as CVE-2026-98432. | 
**Title** | **string** |  | 
**State** | **string** | mitigated or not_mitigated for a K3S cluster, by the protections Wodby keeps on its server. fixed or not_fixed for a managed cluster, by the node images its provider ships. | 
**Action** | Pointer to **string** | The operation that fixes or mitigates the advisory. Absent when nothing is needed, or when the provider has no fixed node image yet. | [optional] 
**Message** | **string** | Why the cluster is in this state and what to do about it. | 

## Methods

### NewClusterSecurityAdvisory

`func NewClusterSecurityAdvisory(advisory string, title string, state string, message string, ) *ClusterSecurityAdvisory`

NewClusterSecurityAdvisory instantiates a new ClusterSecurityAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterSecurityAdvisoryWithDefaults

`func NewClusterSecurityAdvisoryWithDefaults() *ClusterSecurityAdvisory`

NewClusterSecurityAdvisoryWithDefaults instantiates a new ClusterSecurityAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvisory

`func (o *ClusterSecurityAdvisory) GetAdvisory() string`

GetAdvisory returns the Advisory field if non-nil, zero value otherwise.

### GetAdvisoryOk

`func (o *ClusterSecurityAdvisory) GetAdvisoryOk() (*string, bool)`

GetAdvisoryOk returns a tuple with the Advisory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisory

`func (o *ClusterSecurityAdvisory) SetAdvisory(v string)`

SetAdvisory sets Advisory field to given value.


### GetTitle

`func (o *ClusterSecurityAdvisory) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ClusterSecurityAdvisory) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ClusterSecurityAdvisory) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetState

`func (o *ClusterSecurityAdvisory) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ClusterSecurityAdvisory) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ClusterSecurityAdvisory) SetState(v string)`

SetState sets State field to given value.


### GetAction

`func (o *ClusterSecurityAdvisory) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ClusterSecurityAdvisory) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ClusterSecurityAdvisory) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ClusterSecurityAdvisory) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetMessage

`func (o *ClusterSecurityAdvisory) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ClusterSecurityAdvisory) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ClusterSecurityAdvisory) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


