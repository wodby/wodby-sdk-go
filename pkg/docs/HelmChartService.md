# HelmChartService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**Selector** | Pointer to **map[string]string** |  | [optional] 
**Ports** | Pointer to [**[]HelmChartServicePort**](HelmChartServicePort.md) |  | [optional] 
**Headless** | **bool** |  | 

## Methods

### NewHelmChartService

`func NewHelmChartService(name string, headless bool, ) *HelmChartService`

NewHelmChartService instantiates a new HelmChartService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmChartServiceWithDefaults

`func NewHelmChartServiceWithDefaults() *HelmChartService`

NewHelmChartServiceWithDefaults instantiates a new HelmChartService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HelmChartService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HelmChartService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HelmChartService) SetName(v string)`

SetName sets Name field to given value.


### GetLabels

`func (o *HelmChartService) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *HelmChartService) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *HelmChartService) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *HelmChartService) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetSelector

`func (o *HelmChartService) GetSelector() map[string]string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *HelmChartService) GetSelectorOk() (*map[string]string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *HelmChartService) SetSelector(v map[string]string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *HelmChartService) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetPorts

`func (o *HelmChartService) GetPorts() []HelmChartServicePort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *HelmChartService) GetPortsOk() (*[]HelmChartServicePort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *HelmChartService) SetPorts(v []HelmChartServicePort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *HelmChartService) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetHeadless

`func (o *HelmChartService) GetHeadless() bool`

GetHeadless returns the Headless field if non-nil, zero value otherwise.

### GetHeadlessOk

`func (o *HelmChartService) GetHeadlessOk() (*bool, bool)`

GetHeadlessOk returns a tuple with the Headless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadless

`func (o *HelmChartService) SetHeadless(v bool)`

SetHeadless sets Headless field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


