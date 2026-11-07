# ResourcePresetInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **int32** | Used when creating a preset. Optional for API-key requests; defaults to the API key&#39;s organization. | [optional] 
**Name** | **string** | Lowercase letters, digits and dashes; unique in the organization and not the name of a Wodby preset. | 
**Title** | **string** |  | 
**RequestCPU** | Pointer to **NullableInt32** |  | [optional] 
**RequestMem** | Pointer to **NullableInt32** |  | [optional] 
**LimitCPU** | Pointer to **NullableInt32** |  | [optional] 
**LimitMem** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewResourcePresetInput

`func NewResourcePresetInput(name string, title string, ) *ResourcePresetInput`

NewResourcePresetInput instantiates a new ResourcePresetInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcePresetInputWithDefaults

`func NewResourcePresetInputWithDefaults() *ResourcePresetInput`

NewResourcePresetInputWithDefaults instantiates a new ResourcePresetInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *ResourcePresetInput) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ResourcePresetInput) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ResourcePresetInput) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ResourcePresetInput) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetName

`func (o *ResourcePresetInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourcePresetInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourcePresetInput) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *ResourcePresetInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ResourcePresetInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ResourcePresetInput) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetRequestCPU

`func (o *ResourcePresetInput) GetRequestCPU() int32`

GetRequestCPU returns the RequestCPU field if non-nil, zero value otherwise.

### GetRequestCPUOk

`func (o *ResourcePresetInput) GetRequestCPUOk() (*int32, bool)`

GetRequestCPUOk returns a tuple with the RequestCPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCPU

`func (o *ResourcePresetInput) SetRequestCPU(v int32)`

SetRequestCPU sets RequestCPU field to given value.

### HasRequestCPU

`func (o *ResourcePresetInput) HasRequestCPU() bool`

HasRequestCPU returns a boolean if a field has been set.

### SetRequestCPUNil

`func (o *ResourcePresetInput) SetRequestCPUNil(b bool)`

 SetRequestCPUNil sets the value for RequestCPU to be an explicit nil

### UnsetRequestCPU
`func (o *ResourcePresetInput) UnsetRequestCPU()`

UnsetRequestCPU ensures that no value is present for RequestCPU, not even an explicit nil
### GetRequestMem

`func (o *ResourcePresetInput) GetRequestMem() int32`

GetRequestMem returns the RequestMem field if non-nil, zero value otherwise.

### GetRequestMemOk

`func (o *ResourcePresetInput) GetRequestMemOk() (*int32, bool)`

GetRequestMemOk returns a tuple with the RequestMem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMem

`func (o *ResourcePresetInput) SetRequestMem(v int32)`

SetRequestMem sets RequestMem field to given value.

### HasRequestMem

`func (o *ResourcePresetInput) HasRequestMem() bool`

HasRequestMem returns a boolean if a field has been set.

### SetRequestMemNil

`func (o *ResourcePresetInput) SetRequestMemNil(b bool)`

 SetRequestMemNil sets the value for RequestMem to be an explicit nil

### UnsetRequestMem
`func (o *ResourcePresetInput) UnsetRequestMem()`

UnsetRequestMem ensures that no value is present for RequestMem, not even an explicit nil
### GetLimitCPU

`func (o *ResourcePresetInput) GetLimitCPU() int32`

GetLimitCPU returns the LimitCPU field if non-nil, zero value otherwise.

### GetLimitCPUOk

`func (o *ResourcePresetInput) GetLimitCPUOk() (*int32, bool)`

GetLimitCPUOk returns a tuple with the LimitCPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitCPU

`func (o *ResourcePresetInput) SetLimitCPU(v int32)`

SetLimitCPU sets LimitCPU field to given value.

### HasLimitCPU

`func (o *ResourcePresetInput) HasLimitCPU() bool`

HasLimitCPU returns a boolean if a field has been set.

### SetLimitCPUNil

`func (o *ResourcePresetInput) SetLimitCPUNil(b bool)`

 SetLimitCPUNil sets the value for LimitCPU to be an explicit nil

### UnsetLimitCPU
`func (o *ResourcePresetInput) UnsetLimitCPU()`

UnsetLimitCPU ensures that no value is present for LimitCPU, not even an explicit nil
### GetLimitMem

`func (o *ResourcePresetInput) GetLimitMem() int32`

GetLimitMem returns the LimitMem field if non-nil, zero value otherwise.

### GetLimitMemOk

`func (o *ResourcePresetInput) GetLimitMemOk() (*int32, bool)`

GetLimitMemOk returns a tuple with the LimitMem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitMem

`func (o *ResourcePresetInput) SetLimitMem(v int32)`

SetLimitMem sets LimitMem field to given value.

### HasLimitMem

`func (o *ResourcePresetInput) HasLimitMem() bool`

HasLimitMem returns a boolean if a field has been set.

### SetLimitMemNil

`func (o *ResourcePresetInput) SetLimitMemNil(b bool)`

 SetLimitMemNil sets the value for LimitMem to be an explicit nil

### UnsetLimitMem
`func (o *ResourcePresetInput) UnsetLimitMem()`

UnsetLimitMem ensures that no value is present for LimitMem, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


