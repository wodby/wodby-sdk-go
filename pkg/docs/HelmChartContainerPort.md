# HelmChartContainerPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Number** | **int32** |  | 
**Protocol** | **string** |  | 

## Methods

### NewHelmChartContainerPort

`func NewHelmChartContainerPort(number int32, protocol string, ) *HelmChartContainerPort`

NewHelmChartContainerPort instantiates a new HelmChartContainerPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmChartContainerPortWithDefaults

`func NewHelmChartContainerPortWithDefaults() *HelmChartContainerPort`

NewHelmChartContainerPortWithDefaults instantiates a new HelmChartContainerPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HelmChartContainerPort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HelmChartContainerPort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HelmChartContainerPort) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HelmChartContainerPort) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumber

`func (o *HelmChartContainerPort) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *HelmChartContainerPort) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *HelmChartContainerPort) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetProtocol

`func (o *HelmChartContainerPort) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *HelmChartContainerPort) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *HelmChartContainerPort) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


