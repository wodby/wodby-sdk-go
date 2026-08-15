# AppInstanceStackUpgradeChangelog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreviousStackVersion** | **string** |  | 
**StackVersion** | **string** |  | 
**PreviousStackRevNumber** | **int32** |  | 
**StackRevNumber** | **int32** |  | 
**ServiceChanges** | [**[]ServiceRevisionChange**](ServiceRevisionChange.md) |  | 

## Methods

### NewAppInstanceStackUpgradeChangelog

`func NewAppInstanceStackUpgradeChangelog(previousStackVersion string, stackVersion string, previousStackRevNumber int32, stackRevNumber int32, serviceChanges []ServiceRevisionChange, ) *AppInstanceStackUpgradeChangelog`

NewAppInstanceStackUpgradeChangelog instantiates a new AppInstanceStackUpgradeChangelog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppInstanceStackUpgradeChangelogWithDefaults

`func NewAppInstanceStackUpgradeChangelogWithDefaults() *AppInstanceStackUpgradeChangelog`

NewAppInstanceStackUpgradeChangelogWithDefaults instantiates a new AppInstanceStackUpgradeChangelog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreviousStackVersion

`func (o *AppInstanceStackUpgradeChangelog) GetPreviousStackVersion() string`

GetPreviousStackVersion returns the PreviousStackVersion field if non-nil, zero value otherwise.

### GetPreviousStackVersionOk

`func (o *AppInstanceStackUpgradeChangelog) GetPreviousStackVersionOk() (*string, bool)`

GetPreviousStackVersionOk returns a tuple with the PreviousStackVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousStackVersion

`func (o *AppInstanceStackUpgradeChangelog) SetPreviousStackVersion(v string)`

SetPreviousStackVersion sets PreviousStackVersion field to given value.


### GetStackVersion

`func (o *AppInstanceStackUpgradeChangelog) GetStackVersion() string`

GetStackVersion returns the StackVersion field if non-nil, zero value otherwise.

### GetStackVersionOk

`func (o *AppInstanceStackUpgradeChangelog) GetStackVersionOk() (*string, bool)`

GetStackVersionOk returns a tuple with the StackVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackVersion

`func (o *AppInstanceStackUpgradeChangelog) SetStackVersion(v string)`

SetStackVersion sets StackVersion field to given value.


### GetPreviousStackRevNumber

`func (o *AppInstanceStackUpgradeChangelog) GetPreviousStackRevNumber() int32`

GetPreviousStackRevNumber returns the PreviousStackRevNumber field if non-nil, zero value otherwise.

### GetPreviousStackRevNumberOk

`func (o *AppInstanceStackUpgradeChangelog) GetPreviousStackRevNumberOk() (*int32, bool)`

GetPreviousStackRevNumberOk returns a tuple with the PreviousStackRevNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousStackRevNumber

`func (o *AppInstanceStackUpgradeChangelog) SetPreviousStackRevNumber(v int32)`

SetPreviousStackRevNumber sets PreviousStackRevNumber field to given value.


### GetStackRevNumber

`func (o *AppInstanceStackUpgradeChangelog) GetStackRevNumber() int32`

GetStackRevNumber returns the StackRevNumber field if non-nil, zero value otherwise.

### GetStackRevNumberOk

`func (o *AppInstanceStackUpgradeChangelog) GetStackRevNumberOk() (*int32, bool)`

GetStackRevNumberOk returns a tuple with the StackRevNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackRevNumber

`func (o *AppInstanceStackUpgradeChangelog) SetStackRevNumber(v int32)`

SetStackRevNumber sets StackRevNumber field to given value.


### GetServiceChanges

`func (o *AppInstanceStackUpgradeChangelog) GetServiceChanges() []ServiceRevisionChange`

GetServiceChanges returns the ServiceChanges field if non-nil, zero value otherwise.

### GetServiceChangesOk

`func (o *AppInstanceStackUpgradeChangelog) GetServiceChangesOk() (*[]ServiceRevisionChange, bool)`

GetServiceChangesOk returns a tuple with the ServiceChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceChanges

`func (o *AppInstanceStackUpgradeChangelog) SetServiceChanges(v []ServiceRevisionChange)`

SetServiceChanges sets ServiceChanges field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


