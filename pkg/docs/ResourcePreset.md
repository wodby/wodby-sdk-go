# ResourcePreset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Set only for an organization&#39;s custom presets. | [optional] 
**Name** | **string** |  | 
**Title** | **string** |  | 
**RequestCPU** | Pointer to **NullableInt32** | Requested CPU in millicores. | [optional] 
**RequestMem** | Pointer to **NullableInt32** | Requested memory in MiB. | [optional] 
**LimitCPU** | Pointer to **NullableInt32** | CPU limit in millicores. | [optional] 
**LimitMem** | Pointer to **NullableInt32** | Memory limit in MiB. | [optional] 
**System** | **bool** | Wodby&#39;s own presets, whose values never change. | 

## Methods

### NewResourcePreset

`func NewResourcePreset(name string, title string, system bool, ) *ResourcePreset`

NewResourcePreset instantiates a new ResourcePreset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcePresetWithDefaults

`func NewResourcePresetWithDefaults() *ResourcePreset`

NewResourcePresetWithDefaults instantiates a new ResourcePreset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourcePreset) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourcePreset) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourcePreset) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ResourcePreset) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ResourcePreset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourcePreset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourcePreset) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *ResourcePreset) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ResourcePreset) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ResourcePreset) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetRequestCPU

`func (o *ResourcePreset) GetRequestCPU() int32`

GetRequestCPU returns the RequestCPU field if non-nil, zero value otherwise.

### GetRequestCPUOk

`func (o *ResourcePreset) GetRequestCPUOk() (*int32, bool)`

GetRequestCPUOk returns a tuple with the RequestCPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCPU

`func (o *ResourcePreset) SetRequestCPU(v int32)`

SetRequestCPU sets RequestCPU field to given value.

### HasRequestCPU

`func (o *ResourcePreset) HasRequestCPU() bool`

HasRequestCPU returns a boolean if a field has been set.

### SetRequestCPUNil

`func (o *ResourcePreset) SetRequestCPUNil(b bool)`

 SetRequestCPUNil sets the value for RequestCPU to be an explicit nil

### UnsetRequestCPU
`func (o *ResourcePreset) UnsetRequestCPU()`

UnsetRequestCPU ensures that no value is present for RequestCPU, not even an explicit nil
### GetRequestMem

`func (o *ResourcePreset) GetRequestMem() int32`

GetRequestMem returns the RequestMem field if non-nil, zero value otherwise.

### GetRequestMemOk

`func (o *ResourcePreset) GetRequestMemOk() (*int32, bool)`

GetRequestMemOk returns a tuple with the RequestMem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMem

`func (o *ResourcePreset) SetRequestMem(v int32)`

SetRequestMem sets RequestMem field to given value.

### HasRequestMem

`func (o *ResourcePreset) HasRequestMem() bool`

HasRequestMem returns a boolean if a field has been set.

### SetRequestMemNil

`func (o *ResourcePreset) SetRequestMemNil(b bool)`

 SetRequestMemNil sets the value for RequestMem to be an explicit nil

### UnsetRequestMem
`func (o *ResourcePreset) UnsetRequestMem()`

UnsetRequestMem ensures that no value is present for RequestMem, not even an explicit nil
### GetLimitCPU

`func (o *ResourcePreset) GetLimitCPU() int32`

GetLimitCPU returns the LimitCPU field if non-nil, zero value otherwise.

### GetLimitCPUOk

`func (o *ResourcePreset) GetLimitCPUOk() (*int32, bool)`

GetLimitCPUOk returns a tuple with the LimitCPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitCPU

`func (o *ResourcePreset) SetLimitCPU(v int32)`

SetLimitCPU sets LimitCPU field to given value.

### HasLimitCPU

`func (o *ResourcePreset) HasLimitCPU() bool`

HasLimitCPU returns a boolean if a field has been set.

### SetLimitCPUNil

`func (o *ResourcePreset) SetLimitCPUNil(b bool)`

 SetLimitCPUNil sets the value for LimitCPU to be an explicit nil

### UnsetLimitCPU
`func (o *ResourcePreset) UnsetLimitCPU()`

UnsetLimitCPU ensures that no value is present for LimitCPU, not even an explicit nil
### GetLimitMem

`func (o *ResourcePreset) GetLimitMem() int32`

GetLimitMem returns the LimitMem field if non-nil, zero value otherwise.

### GetLimitMemOk

`func (o *ResourcePreset) GetLimitMemOk() (*int32, bool)`

GetLimitMemOk returns a tuple with the LimitMem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitMem

`func (o *ResourcePreset) SetLimitMem(v int32)`

SetLimitMem sets LimitMem field to given value.

### HasLimitMem

`func (o *ResourcePreset) HasLimitMem() bool`

HasLimitMem returns a boolean if a field has been set.

### SetLimitMemNil

`func (o *ResourcePreset) SetLimitMemNil(b bool)`

 SetLimitMemNil sets the value for LimitMem to be an explicit nil

### UnsetLimitMem
`func (o *ResourcePreset) UnsetLimitMem()`

UnsetLimitMem ensures that no value is present for LimitMem, not even an explicit nil
### GetSystem

`func (o *ResourcePreset) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *ResourcePreset) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *ResourcePreset) SetSystem(v bool)`

SetSystem sets System field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


