# AppServiceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**AppServiceId** | **int32** |  | 
**Name** | **string** |  | 
**Title** | **string** |  | 
**Config** | Pointer to **NullableString** |  | [optional] 
**Disabled** | **bool** |  | 

## Methods

### NewAppServiceConfig

`func NewAppServiceConfig(id int32, appServiceId int32, name string, title string, disabled bool, ) *AppServiceConfig`

NewAppServiceConfig instantiates a new AppServiceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceConfigWithDefaults

`func NewAppServiceConfigWithDefaults() *AppServiceConfig`

NewAppServiceConfigWithDefaults instantiates a new AppServiceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppServiceConfig) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppServiceConfig) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppServiceConfig) SetId(v int32)`

SetId sets Id field to given value.


### GetAppServiceId

`func (o *AppServiceConfig) GetAppServiceId() int32`

GetAppServiceId returns the AppServiceId field if non-nil, zero value otherwise.

### GetAppServiceIdOk

`func (o *AppServiceConfig) GetAppServiceIdOk() (*int32, bool)`

GetAppServiceIdOk returns a tuple with the AppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceId

`func (o *AppServiceConfig) SetAppServiceId(v int32)`

SetAppServiceId sets AppServiceId field to given value.


### GetName

`func (o *AppServiceConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppServiceConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppServiceConfig) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *AppServiceConfig) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AppServiceConfig) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AppServiceConfig) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetConfig

`func (o *AppServiceConfig) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AppServiceConfig) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AppServiceConfig) SetConfig(v string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AppServiceConfig) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *AppServiceConfig) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *AppServiceConfig) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetDisabled

`func (o *AppServiceConfig) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *AppServiceConfig) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *AppServiceConfig) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


