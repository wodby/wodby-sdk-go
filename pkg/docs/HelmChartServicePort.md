# HelmChartServicePort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Number** | **int32** |  | 
**TargetPort** | Pointer to [**HelmChartServicePortTargetPort**](HelmChartServicePortTargetPort.md) |  | [optional] 
**Protocol** | **string** |  | 

## Methods

### NewHelmChartServicePort

`func NewHelmChartServicePort(number int32, protocol string, ) *HelmChartServicePort`

NewHelmChartServicePort instantiates a new HelmChartServicePort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmChartServicePortWithDefaults

`func NewHelmChartServicePortWithDefaults() *HelmChartServicePort`

NewHelmChartServicePortWithDefaults instantiates a new HelmChartServicePort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HelmChartServicePort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HelmChartServicePort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HelmChartServicePort) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HelmChartServicePort) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumber

`func (o *HelmChartServicePort) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *HelmChartServicePort) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *HelmChartServicePort) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetTargetPort

`func (o *HelmChartServicePort) GetTargetPort() HelmChartServicePortTargetPort`

GetTargetPort returns the TargetPort field if non-nil, zero value otherwise.

### GetTargetPortOk

`func (o *HelmChartServicePort) GetTargetPortOk() (*HelmChartServicePortTargetPort, bool)`

GetTargetPortOk returns a tuple with the TargetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPort

`func (o *HelmChartServicePort) SetTargetPort(v HelmChartServicePortTargetPort)`

SetTargetPort sets TargetPort field to given value.

### HasTargetPort

`func (o *HelmChartServicePort) HasTargetPort() bool`

HasTargetPort returns a boolean if a field has been set.

### GetProtocol

`func (o *HelmChartServicePort) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *HelmChartServicePort) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *HelmChartServicePort) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


