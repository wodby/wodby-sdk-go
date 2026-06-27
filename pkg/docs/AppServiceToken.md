# AppServiceToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**AppServiceId** | **int32** |  | 
**Name** | **string** |  | 
**Value** | **string** |  | 
**ValueSecretId** | Pointer to **NullableInt32** |  | [optional] 
**EnvType** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewAppServiceToken

`func NewAppServiceToken(id int32, appServiceId int32, name string, value string, createdAt time.Time, ) *AppServiceToken`

NewAppServiceToken instantiates a new AppServiceToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceTokenWithDefaults

`func NewAppServiceTokenWithDefaults() *AppServiceToken`

NewAppServiceTokenWithDefaults instantiates a new AppServiceToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppServiceToken) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppServiceToken) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppServiceToken) SetId(v int32)`

SetId sets Id field to given value.


### GetAppServiceId

`func (o *AppServiceToken) GetAppServiceId() int32`

GetAppServiceId returns the AppServiceId field if non-nil, zero value otherwise.

### GetAppServiceIdOk

`func (o *AppServiceToken) GetAppServiceIdOk() (*int32, bool)`

GetAppServiceIdOk returns a tuple with the AppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceId

`func (o *AppServiceToken) SetAppServiceId(v int32)`

SetAppServiceId sets AppServiceId field to given value.


### GetName

`func (o *AppServiceToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppServiceToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppServiceToken) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *AppServiceToken) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AppServiceToken) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AppServiceToken) SetValue(v string)`

SetValue sets Value field to given value.


### GetValueSecretId

`func (o *AppServiceToken) GetValueSecretId() int32`

GetValueSecretId returns the ValueSecretId field if non-nil, zero value otherwise.

### GetValueSecretIdOk

`func (o *AppServiceToken) GetValueSecretIdOk() (*int32, bool)`

GetValueSecretIdOk returns a tuple with the ValueSecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueSecretId

`func (o *AppServiceToken) SetValueSecretId(v int32)`

SetValueSecretId sets ValueSecretId field to given value.

### HasValueSecretId

`func (o *AppServiceToken) HasValueSecretId() bool`

HasValueSecretId returns a boolean if a field has been set.

### SetValueSecretIdNil

`func (o *AppServiceToken) SetValueSecretIdNil(b bool)`

 SetValueSecretIdNil sets the value for ValueSecretId to be an explicit nil

### UnsetValueSecretId
`func (o *AppServiceToken) UnsetValueSecretId()`

UnsetValueSecretId ensures that no value is present for ValueSecretId, not even an explicit nil
### GetEnvType

`func (o *AppServiceToken) GetEnvType() string`

GetEnvType returns the EnvType field if non-nil, zero value otherwise.

### GetEnvTypeOk

`func (o *AppServiceToken) GetEnvTypeOk() (*string, bool)`

GetEnvTypeOk returns a tuple with the EnvType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvType

`func (o *AppServiceToken) SetEnvType(v string)`

SetEnvType sets EnvType field to given value.

### HasEnvType

`func (o *AppServiceToken) HasEnvType() bool`

HasEnvType returns a boolean if a field has been set.

### SetEnvTypeNil

`func (o *AppServiceToken) SetEnvTypeNil(b bool)`

 SetEnvTypeNil sets the value for EnvType to be an explicit nil

### UnsetEnvType
`func (o *AppServiceToken) UnsetEnvType()`

UnsetEnvType ensures that no value is present for EnvType, not even an explicit nil
### GetCreatedAt

`func (o *AppServiceToken) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppServiceToken) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppServiceToken) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


