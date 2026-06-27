# AppServiceSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**AppServiceId** | **int32** |  | 
**Name** | **string** |  | 
**Value** | **string** |  | 
**Var** | **string** |  | 
**Runtime** | **bool** |  | 
**Build** | **bool** |  | 
**FromSettingId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewAppServiceSetting

`func NewAppServiceSetting(id int32, appServiceId int32, name string, value string, var_ string, runtime bool, build bool, ) *AppServiceSetting`

NewAppServiceSetting instantiates a new AppServiceSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceSettingWithDefaults

`func NewAppServiceSettingWithDefaults() *AppServiceSetting`

NewAppServiceSettingWithDefaults instantiates a new AppServiceSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppServiceSetting) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppServiceSetting) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppServiceSetting) SetId(v int32)`

SetId sets Id field to given value.


### GetAppServiceId

`func (o *AppServiceSetting) GetAppServiceId() int32`

GetAppServiceId returns the AppServiceId field if non-nil, zero value otherwise.

### GetAppServiceIdOk

`func (o *AppServiceSetting) GetAppServiceIdOk() (*int32, bool)`

GetAppServiceIdOk returns a tuple with the AppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceId

`func (o *AppServiceSetting) SetAppServiceId(v int32)`

SetAppServiceId sets AppServiceId field to given value.


### GetName

`func (o *AppServiceSetting) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppServiceSetting) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppServiceSetting) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *AppServiceSetting) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AppServiceSetting) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AppServiceSetting) SetValue(v string)`

SetValue sets Value field to given value.


### GetVar

`func (o *AppServiceSetting) GetVar() string`

GetVar returns the Var field if non-nil, zero value otherwise.

### GetVarOk

`func (o *AppServiceSetting) GetVarOk() (*string, bool)`

GetVarOk returns a tuple with the Var field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar

`func (o *AppServiceSetting) SetVar(v string)`

SetVar sets Var field to given value.


### GetRuntime

`func (o *AppServiceSetting) GetRuntime() bool`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *AppServiceSetting) GetRuntimeOk() (*bool, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *AppServiceSetting) SetRuntime(v bool)`

SetRuntime sets Runtime field to given value.


### GetBuild

`func (o *AppServiceSetting) GetBuild() bool`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *AppServiceSetting) GetBuildOk() (*bool, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *AppServiceSetting) SetBuild(v bool)`

SetBuild sets Build field to given value.


### GetFromSettingId

`func (o *AppServiceSetting) GetFromSettingId() int32`

GetFromSettingId returns the FromSettingId field if non-nil, zero value otherwise.

### GetFromSettingIdOk

`func (o *AppServiceSetting) GetFromSettingIdOk() (*int32, bool)`

GetFromSettingIdOk returns a tuple with the FromSettingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromSettingId

`func (o *AppServiceSetting) SetFromSettingId(v int32)`

SetFromSettingId sets FromSettingId field to given value.

### HasFromSettingId

`func (o *AppServiceSetting) HasFromSettingId() bool`

HasFromSettingId returns a boolean if a field has been set.

### SetFromSettingIdNil

`func (o *AppServiceSetting) SetFromSettingIdNil(b bool)`

 SetFromSettingIdNil sets the value for FromSettingId to be an explicit nil

### UnsetFromSettingId
`func (o *AppServiceSetting) UnsetFromSettingId()`

UnsetFromSettingId ensures that no value is present for FromSettingId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


