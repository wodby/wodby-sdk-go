# HelmChartWorkload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Name** | **string** |  | 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**Selector** | Pointer to **map[string]string** |  | [optional] 
**PodLabels** | Pointer to **map[string]string** |  | [optional] 
**Containers** | Pointer to [**[]HelmChartContainer**](HelmChartContainer.md) |  | [optional] 
**InitContainers** | Pointer to [**[]HelmChartContainer**](HelmChartContainer.md) |  | [optional] 
**Volumes** | Pointer to [**[]HelmChartVolumeClaim**](HelmChartVolumeClaim.md) |  | [optional] 

## Methods

### NewHelmChartWorkload

`func NewHelmChartWorkload(kind string, name string, ) *HelmChartWorkload`

NewHelmChartWorkload instantiates a new HelmChartWorkload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmChartWorkloadWithDefaults

`func NewHelmChartWorkloadWithDefaults() *HelmChartWorkload`

NewHelmChartWorkloadWithDefaults instantiates a new HelmChartWorkload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *HelmChartWorkload) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *HelmChartWorkload) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *HelmChartWorkload) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *HelmChartWorkload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HelmChartWorkload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HelmChartWorkload) SetName(v string)`

SetName sets Name field to given value.


### GetLabels

`func (o *HelmChartWorkload) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *HelmChartWorkload) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *HelmChartWorkload) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *HelmChartWorkload) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetSelector

`func (o *HelmChartWorkload) GetSelector() map[string]string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *HelmChartWorkload) GetSelectorOk() (*map[string]string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *HelmChartWorkload) SetSelector(v map[string]string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *HelmChartWorkload) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetPodLabels

`func (o *HelmChartWorkload) GetPodLabels() map[string]string`

GetPodLabels returns the PodLabels field if non-nil, zero value otherwise.

### GetPodLabelsOk

`func (o *HelmChartWorkload) GetPodLabelsOk() (*map[string]string, bool)`

GetPodLabelsOk returns a tuple with the PodLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodLabels

`func (o *HelmChartWorkload) SetPodLabels(v map[string]string)`

SetPodLabels sets PodLabels field to given value.

### HasPodLabels

`func (o *HelmChartWorkload) HasPodLabels() bool`

HasPodLabels returns a boolean if a field has been set.

### GetContainers

`func (o *HelmChartWorkload) GetContainers() []HelmChartContainer`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *HelmChartWorkload) GetContainersOk() (*[]HelmChartContainer, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *HelmChartWorkload) SetContainers(v []HelmChartContainer)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *HelmChartWorkload) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetInitContainers

`func (o *HelmChartWorkload) GetInitContainers() []HelmChartContainer`

GetInitContainers returns the InitContainers field if non-nil, zero value otherwise.

### GetInitContainersOk

`func (o *HelmChartWorkload) GetInitContainersOk() (*[]HelmChartContainer, bool)`

GetInitContainersOk returns a tuple with the InitContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitContainers

`func (o *HelmChartWorkload) SetInitContainers(v []HelmChartContainer)`

SetInitContainers sets InitContainers field to given value.

### HasInitContainers

`func (o *HelmChartWorkload) HasInitContainers() bool`

HasInitContainers returns a boolean if a field has been set.

### GetVolumes

`func (o *HelmChartWorkload) GetVolumes() []HelmChartVolumeClaim`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *HelmChartWorkload) GetVolumesOk() (*[]HelmChartVolumeClaim, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *HelmChartWorkload) SetVolumes(v []HelmChartVolumeClaim)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *HelmChartWorkload) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


