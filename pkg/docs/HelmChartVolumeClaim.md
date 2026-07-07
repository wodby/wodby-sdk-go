# HelmChartVolumeClaim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Size** | Pointer to **string** |  | [optional] 
**StorageClassName** | Pointer to **string** |  | [optional] 
**AccessModes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewHelmChartVolumeClaim

`func NewHelmChartVolumeClaim(name string, ) *HelmChartVolumeClaim`

NewHelmChartVolumeClaim instantiates a new HelmChartVolumeClaim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmChartVolumeClaimWithDefaults

`func NewHelmChartVolumeClaimWithDefaults() *HelmChartVolumeClaim`

NewHelmChartVolumeClaimWithDefaults instantiates a new HelmChartVolumeClaim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HelmChartVolumeClaim) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HelmChartVolumeClaim) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HelmChartVolumeClaim) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *HelmChartVolumeClaim) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *HelmChartVolumeClaim) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *HelmChartVolumeClaim) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *HelmChartVolumeClaim) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStorageClassName

`func (o *HelmChartVolumeClaim) GetStorageClassName() string`

GetStorageClassName returns the StorageClassName field if non-nil, zero value otherwise.

### GetStorageClassNameOk

`func (o *HelmChartVolumeClaim) GetStorageClassNameOk() (*string, bool)`

GetStorageClassNameOk returns a tuple with the StorageClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClassName

`func (o *HelmChartVolumeClaim) SetStorageClassName(v string)`

SetStorageClassName sets StorageClassName field to given value.

### HasStorageClassName

`func (o *HelmChartVolumeClaim) HasStorageClassName() bool`

HasStorageClassName returns a boolean if a field has been set.

### GetAccessModes

`func (o *HelmChartVolumeClaim) GetAccessModes() []string`

GetAccessModes returns the AccessModes field if non-nil, zero value otherwise.

### GetAccessModesOk

`func (o *HelmChartVolumeClaim) GetAccessModesOk() (*[]string, bool)`

GetAccessModesOk returns a tuple with the AccessModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessModes

`func (o *HelmChartVolumeClaim) SetAccessModes(v []string)`

SetAccessModes sets AccessModes field to given value.

### HasAccessModes

`func (o *HelmChartVolumeClaim) HasAccessModes() bool`

HasAccessModes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


