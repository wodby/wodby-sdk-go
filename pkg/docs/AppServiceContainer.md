# AppServiceContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**AppServiceId** | **int32** |  | 
**Workload** | **string** |  | 
**Name** | **string** |  | 
**RequestCPU** | Pointer to **NullableInt32** |  | [optional] 
**RequestMem** | Pointer to **NullableInt32** |  | [optional] 
**LimitCPU** | Pointer to **NullableInt32** |  | [optional] 
**LimitMem** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewAppServiceContainer

`func NewAppServiceContainer(appServiceId int32, workload string, name string, ) *AppServiceContainer`

NewAppServiceContainer instantiates a new AppServiceContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceContainerWithDefaults

`func NewAppServiceContainerWithDefaults() *AppServiceContainer`

NewAppServiceContainerWithDefaults instantiates a new AppServiceContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppServiceContainer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppServiceContainer) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppServiceContainer) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AppServiceContainer) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AppServiceContainer) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AppServiceContainer) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAppServiceId

`func (o *AppServiceContainer) GetAppServiceId() int32`

GetAppServiceId returns the AppServiceId field if non-nil, zero value otherwise.

### GetAppServiceIdOk

`func (o *AppServiceContainer) GetAppServiceIdOk() (*int32, bool)`

GetAppServiceIdOk returns a tuple with the AppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceId

`func (o *AppServiceContainer) SetAppServiceId(v int32)`

SetAppServiceId sets AppServiceId field to given value.


### GetWorkload

`func (o *AppServiceContainer) GetWorkload() string`

GetWorkload returns the Workload field if non-nil, zero value otherwise.

### GetWorkloadOk

`func (o *AppServiceContainer) GetWorkloadOk() (*string, bool)`

GetWorkloadOk returns a tuple with the Workload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkload

`func (o *AppServiceContainer) SetWorkload(v string)`

SetWorkload sets Workload field to given value.


### GetName

`func (o *AppServiceContainer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppServiceContainer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppServiceContainer) SetName(v string)`

SetName sets Name field to given value.


### GetRequestCPU

`func (o *AppServiceContainer) GetRequestCPU() int32`

GetRequestCPU returns the RequestCPU field if non-nil, zero value otherwise.

### GetRequestCPUOk

`func (o *AppServiceContainer) GetRequestCPUOk() (*int32, bool)`

GetRequestCPUOk returns a tuple with the RequestCPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCPU

`func (o *AppServiceContainer) SetRequestCPU(v int32)`

SetRequestCPU sets RequestCPU field to given value.

### HasRequestCPU

`func (o *AppServiceContainer) HasRequestCPU() bool`

HasRequestCPU returns a boolean if a field has been set.

### SetRequestCPUNil

`func (o *AppServiceContainer) SetRequestCPUNil(b bool)`

 SetRequestCPUNil sets the value for RequestCPU to be an explicit nil

### UnsetRequestCPU
`func (o *AppServiceContainer) UnsetRequestCPU()`

UnsetRequestCPU ensures that no value is present for RequestCPU, not even an explicit nil
### GetRequestMem

`func (o *AppServiceContainer) GetRequestMem() int32`

GetRequestMem returns the RequestMem field if non-nil, zero value otherwise.

### GetRequestMemOk

`func (o *AppServiceContainer) GetRequestMemOk() (*int32, bool)`

GetRequestMemOk returns a tuple with the RequestMem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMem

`func (o *AppServiceContainer) SetRequestMem(v int32)`

SetRequestMem sets RequestMem field to given value.

### HasRequestMem

`func (o *AppServiceContainer) HasRequestMem() bool`

HasRequestMem returns a boolean if a field has been set.

### SetRequestMemNil

`func (o *AppServiceContainer) SetRequestMemNil(b bool)`

 SetRequestMemNil sets the value for RequestMem to be an explicit nil

### UnsetRequestMem
`func (o *AppServiceContainer) UnsetRequestMem()`

UnsetRequestMem ensures that no value is present for RequestMem, not even an explicit nil
### GetLimitCPU

`func (o *AppServiceContainer) GetLimitCPU() int32`

GetLimitCPU returns the LimitCPU field if non-nil, zero value otherwise.

### GetLimitCPUOk

`func (o *AppServiceContainer) GetLimitCPUOk() (*int32, bool)`

GetLimitCPUOk returns a tuple with the LimitCPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitCPU

`func (o *AppServiceContainer) SetLimitCPU(v int32)`

SetLimitCPU sets LimitCPU field to given value.

### HasLimitCPU

`func (o *AppServiceContainer) HasLimitCPU() bool`

HasLimitCPU returns a boolean if a field has been set.

### SetLimitCPUNil

`func (o *AppServiceContainer) SetLimitCPUNil(b bool)`

 SetLimitCPUNil sets the value for LimitCPU to be an explicit nil

### UnsetLimitCPU
`func (o *AppServiceContainer) UnsetLimitCPU()`

UnsetLimitCPU ensures that no value is present for LimitCPU, not even an explicit nil
### GetLimitMem

`func (o *AppServiceContainer) GetLimitMem() int32`

GetLimitMem returns the LimitMem field if non-nil, zero value otherwise.

### GetLimitMemOk

`func (o *AppServiceContainer) GetLimitMemOk() (*int32, bool)`

GetLimitMemOk returns a tuple with the LimitMem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitMem

`func (o *AppServiceContainer) SetLimitMem(v int32)`

SetLimitMem sets LimitMem field to given value.

### HasLimitMem

`func (o *AppServiceContainer) HasLimitMem() bool`

HasLimitMem returns a boolean if a field has been set.

### SetLimitMemNil

`func (o *AppServiceContainer) SetLimitMemNil(b bool)`

 SetLimitMemNil sets the value for LimitMem to be an explicit nil

### UnsetLimitMem
`func (o *AppServiceContainer) UnsetLimitMem()`

UnsetLimitMem ensures that no value is present for LimitMem, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


