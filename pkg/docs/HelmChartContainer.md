# HelmChartContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Image** | Pointer to **string** |  | [optional] 
**Command** | Pointer to **[]string** |  | [optional] 
**Args** | Pointer to **[]string** |  | [optional] 
**Ports** | Pointer to [**[]HelmChartContainerPort**](HelmChartContainerPort.md) |  | [optional] 
**Env** | Pointer to **[]string** |  | [optional] 

## Methods

### NewHelmChartContainer

`func NewHelmChartContainer(name string, ) *HelmChartContainer`

NewHelmChartContainer instantiates a new HelmChartContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmChartContainerWithDefaults

`func NewHelmChartContainerWithDefaults() *HelmChartContainer`

NewHelmChartContainerWithDefaults instantiates a new HelmChartContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HelmChartContainer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HelmChartContainer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HelmChartContainer) SetName(v string)`

SetName sets Name field to given value.


### GetImage

`func (o *HelmChartContainer) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *HelmChartContainer) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *HelmChartContainer) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *HelmChartContainer) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetCommand

`func (o *HelmChartContainer) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *HelmChartContainer) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *HelmChartContainer) SetCommand(v []string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *HelmChartContainer) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetArgs

`func (o *HelmChartContainer) GetArgs() []string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *HelmChartContainer) GetArgsOk() (*[]string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *HelmChartContainer) SetArgs(v []string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *HelmChartContainer) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetPorts

`func (o *HelmChartContainer) GetPorts() []HelmChartContainerPort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *HelmChartContainer) GetPortsOk() (*[]HelmChartContainerPort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *HelmChartContainer) SetPorts(v []HelmChartContainerPort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *HelmChartContainer) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetEnv

`func (o *HelmChartContainer) GetEnv() []string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *HelmChartContainer) GetEnvOk() (*[]string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *HelmChartContainer) SetEnv(v []string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *HelmChartContainer) HasEnv() bool`

HasEnv returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


