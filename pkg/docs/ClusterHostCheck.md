# ClusterHostCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Identifies the check, such as ipv4_options_filter or mptcp_disabled. More checks may be added. | 
**Title** | **string** |  | 
**Advisory** | **string** | The security advisory the check protects against, such as CVE-2026-98432. Group checks by this field. | 
**AdvisoryTitle** | **string** | Describes the advisory, such as Linux kernel TCP bug. | 
**Node** | **string** | The node that reported the check; empty when no node has reported. | 
**State** | **string** | Whether the protection is in place. unknown means the server hasn&#39;t reported the check, or its last report is more than two days old. | 
**Message** | **string** | What the state means and what to do about it. | 
**Kernel** | **string** | The kernel release the server ran at the check; empty when it hasn&#39;t reported. | 
**CheckedAt** | Pointer to **time.Time** | When the server ran the check. Absent when it hasn&#39;t reported. | [optional] 
**ObservedAt** | Pointer to **time.Time** | When Wodby read the report. Absent when the server hasn&#39;t reported. | [optional] 

## Methods

### NewClusterHostCheck

`func NewClusterHostCheck(name string, title string, advisory string, advisoryTitle string, node string, state string, message string, kernel string, ) *ClusterHostCheck`

NewClusterHostCheck instantiates a new ClusterHostCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterHostCheckWithDefaults

`func NewClusterHostCheckWithDefaults() *ClusterHostCheck`

NewClusterHostCheckWithDefaults instantiates a new ClusterHostCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClusterHostCheck) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterHostCheck) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterHostCheck) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *ClusterHostCheck) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ClusterHostCheck) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ClusterHostCheck) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetAdvisory

`func (o *ClusterHostCheck) GetAdvisory() string`

GetAdvisory returns the Advisory field if non-nil, zero value otherwise.

### GetAdvisoryOk

`func (o *ClusterHostCheck) GetAdvisoryOk() (*string, bool)`

GetAdvisoryOk returns a tuple with the Advisory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisory

`func (o *ClusterHostCheck) SetAdvisory(v string)`

SetAdvisory sets Advisory field to given value.


### GetAdvisoryTitle

`func (o *ClusterHostCheck) GetAdvisoryTitle() string`

GetAdvisoryTitle returns the AdvisoryTitle field if non-nil, zero value otherwise.

### GetAdvisoryTitleOk

`func (o *ClusterHostCheck) GetAdvisoryTitleOk() (*string, bool)`

GetAdvisoryTitleOk returns a tuple with the AdvisoryTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisoryTitle

`func (o *ClusterHostCheck) SetAdvisoryTitle(v string)`

SetAdvisoryTitle sets AdvisoryTitle field to given value.


### GetNode

`func (o *ClusterHostCheck) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *ClusterHostCheck) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *ClusterHostCheck) SetNode(v string)`

SetNode sets Node field to given value.


### GetState

`func (o *ClusterHostCheck) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ClusterHostCheck) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ClusterHostCheck) SetState(v string)`

SetState sets State field to given value.


### GetMessage

`func (o *ClusterHostCheck) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ClusterHostCheck) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ClusterHostCheck) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetKernel

`func (o *ClusterHostCheck) GetKernel() string`

GetKernel returns the Kernel field if non-nil, zero value otherwise.

### GetKernelOk

`func (o *ClusterHostCheck) GetKernelOk() (*string, bool)`

GetKernelOk returns a tuple with the Kernel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernel

`func (o *ClusterHostCheck) SetKernel(v string)`

SetKernel sets Kernel field to given value.


### GetCheckedAt

`func (o *ClusterHostCheck) GetCheckedAt() time.Time`

GetCheckedAt returns the CheckedAt field if non-nil, zero value otherwise.

### GetCheckedAtOk

`func (o *ClusterHostCheck) GetCheckedAtOk() (*time.Time, bool)`

GetCheckedAtOk returns a tuple with the CheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckedAt

`func (o *ClusterHostCheck) SetCheckedAt(v time.Time)`

SetCheckedAt sets CheckedAt field to given value.

### HasCheckedAt

`func (o *ClusterHostCheck) HasCheckedAt() bool`

HasCheckedAt returns a boolean if a field has been set.

### GetObservedAt

`func (o *ClusterHostCheck) GetObservedAt() time.Time`

GetObservedAt returns the ObservedAt field if non-nil, zero value otherwise.

### GetObservedAtOk

`func (o *ClusterHostCheck) GetObservedAtOk() (*time.Time, bool)`

GetObservedAtOk returns a tuple with the ObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAt

`func (o *ClusterHostCheck) SetObservedAt(v time.Time)`

SetObservedAt sets ObservedAt field to given value.

### HasObservedAt

`func (o *ClusterHostCheck) HasObservedAt() bool`

HasObservedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


