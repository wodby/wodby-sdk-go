# StackUpgradeVolumeDeletion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppServiceName** | **string** |  | 
**AppServiceTitle** | **string** |  | 
**Volume** | **string** |  | 
**ServiceRemoved** | **bool** | True when the target stack revision no longer has the app service&#39;s stack service, so the upgrade deletes the whole app service. False when the service no longer uses the volume. | 

## Methods

### NewStackUpgradeVolumeDeletion

`func NewStackUpgradeVolumeDeletion(appServiceName string, appServiceTitle string, volume string, serviceRemoved bool, ) *StackUpgradeVolumeDeletion`

NewStackUpgradeVolumeDeletion instantiates a new StackUpgradeVolumeDeletion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackUpgradeVolumeDeletionWithDefaults

`func NewStackUpgradeVolumeDeletionWithDefaults() *StackUpgradeVolumeDeletion`

NewStackUpgradeVolumeDeletionWithDefaults instantiates a new StackUpgradeVolumeDeletion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppServiceName

`func (o *StackUpgradeVolumeDeletion) GetAppServiceName() string`

GetAppServiceName returns the AppServiceName field if non-nil, zero value otherwise.

### GetAppServiceNameOk

`func (o *StackUpgradeVolumeDeletion) GetAppServiceNameOk() (*string, bool)`

GetAppServiceNameOk returns a tuple with the AppServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceName

`func (o *StackUpgradeVolumeDeletion) SetAppServiceName(v string)`

SetAppServiceName sets AppServiceName field to given value.


### GetAppServiceTitle

`func (o *StackUpgradeVolumeDeletion) GetAppServiceTitle() string`

GetAppServiceTitle returns the AppServiceTitle field if non-nil, zero value otherwise.

### GetAppServiceTitleOk

`func (o *StackUpgradeVolumeDeletion) GetAppServiceTitleOk() (*string, bool)`

GetAppServiceTitleOk returns a tuple with the AppServiceTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceTitle

`func (o *StackUpgradeVolumeDeletion) SetAppServiceTitle(v string)`

SetAppServiceTitle sets AppServiceTitle field to given value.


### GetVolume

`func (o *StackUpgradeVolumeDeletion) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *StackUpgradeVolumeDeletion) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *StackUpgradeVolumeDeletion) SetVolume(v string)`

SetVolume sets Volume field to given value.


### GetServiceRemoved

`func (o *StackUpgradeVolumeDeletion) GetServiceRemoved() bool`

GetServiceRemoved returns the ServiceRemoved field if non-nil, zero value otherwise.

### GetServiceRemovedOk

`func (o *StackUpgradeVolumeDeletion) GetServiceRemovedOk() (*bool, bool)`

GetServiceRemovedOk returns a tuple with the ServiceRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRemoved

`func (o *StackUpgradeVolumeDeletion) SetServiceRemoved(v bool)`

SetServiceRemoved sets ServiceRemoved field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


