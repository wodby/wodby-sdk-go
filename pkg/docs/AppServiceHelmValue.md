# AppServiceHelmValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**AppServiceId** | **int32** |  | 
**Name** | **string** |  | 
**Value** | **string** |  | 
**ValueSecretId** | Pointer to **NullableInt32** |  | [optional] 
**Source** | Pointer to [**NullableAppServiceHelmValueSource**](AppServiceHelmValueSource.md) |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewAppServiceHelmValue

`func NewAppServiceHelmValue(id int32, appServiceId int32, name string, value string, ) *AppServiceHelmValue`

NewAppServiceHelmValue instantiates a new AppServiceHelmValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceHelmValueWithDefaults

`func NewAppServiceHelmValueWithDefaults() *AppServiceHelmValue`

NewAppServiceHelmValueWithDefaults instantiates a new AppServiceHelmValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppServiceHelmValue) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppServiceHelmValue) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppServiceHelmValue) SetId(v int32)`

SetId sets Id field to given value.


### GetAppServiceId

`func (o *AppServiceHelmValue) GetAppServiceId() int32`

GetAppServiceId returns the AppServiceId field if non-nil, zero value otherwise.

### GetAppServiceIdOk

`func (o *AppServiceHelmValue) GetAppServiceIdOk() (*int32, bool)`

GetAppServiceIdOk returns a tuple with the AppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceId

`func (o *AppServiceHelmValue) SetAppServiceId(v int32)`

SetAppServiceId sets AppServiceId field to given value.


### GetName

`func (o *AppServiceHelmValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppServiceHelmValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppServiceHelmValue) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *AppServiceHelmValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AppServiceHelmValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AppServiceHelmValue) SetValue(v string)`

SetValue sets Value field to given value.


### GetValueSecretId

`func (o *AppServiceHelmValue) GetValueSecretId() int32`

GetValueSecretId returns the ValueSecretId field if non-nil, zero value otherwise.

### GetValueSecretIdOk

`func (o *AppServiceHelmValue) GetValueSecretIdOk() (*int32, bool)`

GetValueSecretIdOk returns a tuple with the ValueSecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueSecretId

`func (o *AppServiceHelmValue) SetValueSecretId(v int32)`

SetValueSecretId sets ValueSecretId field to given value.

### HasValueSecretId

`func (o *AppServiceHelmValue) HasValueSecretId() bool`

HasValueSecretId returns a boolean if a field has been set.

### SetValueSecretIdNil

`func (o *AppServiceHelmValue) SetValueSecretIdNil(b bool)`

 SetValueSecretIdNil sets the value for ValueSecretId to be an explicit nil

### UnsetValueSecretId
`func (o *AppServiceHelmValue) UnsetValueSecretId()`

UnsetValueSecretId ensures that no value is present for ValueSecretId, not even an explicit nil
### GetSource

`func (o *AppServiceHelmValue) GetSource() AppServiceHelmValueSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AppServiceHelmValue) GetSourceOk() (*AppServiceHelmValueSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AppServiceHelmValue) SetSource(v AppServiceHelmValueSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *AppServiceHelmValue) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *AppServiceHelmValue) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *AppServiceHelmValue) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetCreatedAt

`func (o *AppServiceHelmValue) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppServiceHelmValue) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppServiceHelmValue) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AppServiceHelmValue) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *AppServiceHelmValue) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *AppServiceHelmValue) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


