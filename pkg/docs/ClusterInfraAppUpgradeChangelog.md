# ClusterInfraAppUpgradeChangelog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppInstanceId** | **int32** |  | 
**AppName** | **string** |  | 
**AppTitle** | **string** |  | 
**PreviousStackVersion** | **string** |  | 
**StackVersion** | **string** |  | 
**PreviousStackRevNumber** | **int32** |  | 
**StackRevNumber** | **int32** |  | 
**ServiceChanges** | [**[]ServiceRevisionChange**](ServiceRevisionChange.md) |  | 

## Methods

### NewClusterInfraAppUpgradeChangelog

`func NewClusterInfraAppUpgradeChangelog(appInstanceId int32, appName string, appTitle string, previousStackVersion string, stackVersion string, previousStackRevNumber int32, stackRevNumber int32, serviceChanges []ServiceRevisionChange, ) *ClusterInfraAppUpgradeChangelog`

NewClusterInfraAppUpgradeChangelog instantiates a new ClusterInfraAppUpgradeChangelog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterInfraAppUpgradeChangelogWithDefaults

`func NewClusterInfraAppUpgradeChangelogWithDefaults() *ClusterInfraAppUpgradeChangelog`

NewClusterInfraAppUpgradeChangelogWithDefaults instantiates a new ClusterInfraAppUpgradeChangelog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppInstanceId

`func (o *ClusterInfraAppUpgradeChangelog) GetAppInstanceId() int32`

GetAppInstanceId returns the AppInstanceId field if non-nil, zero value otherwise.

### GetAppInstanceIdOk

`func (o *ClusterInfraAppUpgradeChangelog) GetAppInstanceIdOk() (*int32, bool)`

GetAppInstanceIdOk returns a tuple with the AppInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceId

`func (o *ClusterInfraAppUpgradeChangelog) SetAppInstanceId(v int32)`

SetAppInstanceId sets AppInstanceId field to given value.


### GetAppName

`func (o *ClusterInfraAppUpgradeChangelog) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *ClusterInfraAppUpgradeChangelog) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *ClusterInfraAppUpgradeChangelog) SetAppName(v string)`

SetAppName sets AppName field to given value.


### GetAppTitle

`func (o *ClusterInfraAppUpgradeChangelog) GetAppTitle() string`

GetAppTitle returns the AppTitle field if non-nil, zero value otherwise.

### GetAppTitleOk

`func (o *ClusterInfraAppUpgradeChangelog) GetAppTitleOk() (*string, bool)`

GetAppTitleOk returns a tuple with the AppTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppTitle

`func (o *ClusterInfraAppUpgradeChangelog) SetAppTitle(v string)`

SetAppTitle sets AppTitle field to given value.


### GetPreviousStackVersion

`func (o *ClusterInfraAppUpgradeChangelog) GetPreviousStackVersion() string`

GetPreviousStackVersion returns the PreviousStackVersion field if non-nil, zero value otherwise.

### GetPreviousStackVersionOk

`func (o *ClusterInfraAppUpgradeChangelog) GetPreviousStackVersionOk() (*string, bool)`

GetPreviousStackVersionOk returns a tuple with the PreviousStackVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousStackVersion

`func (o *ClusterInfraAppUpgradeChangelog) SetPreviousStackVersion(v string)`

SetPreviousStackVersion sets PreviousStackVersion field to given value.


### GetStackVersion

`func (o *ClusterInfraAppUpgradeChangelog) GetStackVersion() string`

GetStackVersion returns the StackVersion field if non-nil, zero value otherwise.

### GetStackVersionOk

`func (o *ClusterInfraAppUpgradeChangelog) GetStackVersionOk() (*string, bool)`

GetStackVersionOk returns a tuple with the StackVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackVersion

`func (o *ClusterInfraAppUpgradeChangelog) SetStackVersion(v string)`

SetStackVersion sets StackVersion field to given value.


### GetPreviousStackRevNumber

`func (o *ClusterInfraAppUpgradeChangelog) GetPreviousStackRevNumber() int32`

GetPreviousStackRevNumber returns the PreviousStackRevNumber field if non-nil, zero value otherwise.

### GetPreviousStackRevNumberOk

`func (o *ClusterInfraAppUpgradeChangelog) GetPreviousStackRevNumberOk() (*int32, bool)`

GetPreviousStackRevNumberOk returns a tuple with the PreviousStackRevNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousStackRevNumber

`func (o *ClusterInfraAppUpgradeChangelog) SetPreviousStackRevNumber(v int32)`

SetPreviousStackRevNumber sets PreviousStackRevNumber field to given value.


### GetStackRevNumber

`func (o *ClusterInfraAppUpgradeChangelog) GetStackRevNumber() int32`

GetStackRevNumber returns the StackRevNumber field if non-nil, zero value otherwise.

### GetStackRevNumberOk

`func (o *ClusterInfraAppUpgradeChangelog) GetStackRevNumberOk() (*int32, bool)`

GetStackRevNumberOk returns a tuple with the StackRevNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackRevNumber

`func (o *ClusterInfraAppUpgradeChangelog) SetStackRevNumber(v int32)`

SetStackRevNumber sets StackRevNumber field to given value.


### GetServiceChanges

`func (o *ClusterInfraAppUpgradeChangelog) GetServiceChanges() []ServiceRevisionChange`

GetServiceChanges returns the ServiceChanges field if non-nil, zero value otherwise.

### GetServiceChangesOk

`func (o *ClusterInfraAppUpgradeChangelog) GetServiceChangesOk() (*[]ServiceRevisionChange, bool)`

GetServiceChangesOk returns a tuple with the ServiceChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceChanges

`func (o *ClusterInfraAppUpgradeChangelog) SetServiceChanges(v []ServiceRevisionChange)`

SetServiceChanges sets ServiceChanges field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


