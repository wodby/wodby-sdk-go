# StackServiceHelmValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**StackServiceId** | **int32** |  | 
**Name** | **string** |  | 
**Value** | Pointer to **NullableString** |  | [optional] 
**ValueSecretId** | Pointer to **NullableInt32** |  | [optional] 
**EnvType** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewStackServiceHelmValue

`func NewStackServiceHelmValue(id int32, stackServiceId int32, name string, createdAt time.Time, ) *StackServiceHelmValue`

NewStackServiceHelmValue instantiates a new StackServiceHelmValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackServiceHelmValueWithDefaults

`func NewStackServiceHelmValueWithDefaults() *StackServiceHelmValue`

NewStackServiceHelmValueWithDefaults instantiates a new StackServiceHelmValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StackServiceHelmValue) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StackServiceHelmValue) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StackServiceHelmValue) SetId(v int32)`

SetId sets Id field to given value.


### GetStackServiceId

`func (o *StackServiceHelmValue) GetStackServiceId() int32`

GetStackServiceId returns the StackServiceId field if non-nil, zero value otherwise.

### GetStackServiceIdOk

`func (o *StackServiceHelmValue) GetStackServiceIdOk() (*int32, bool)`

GetStackServiceIdOk returns a tuple with the StackServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackServiceId

`func (o *StackServiceHelmValue) SetStackServiceId(v int32)`

SetStackServiceId sets StackServiceId field to given value.


### GetName

`func (o *StackServiceHelmValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StackServiceHelmValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StackServiceHelmValue) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *StackServiceHelmValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StackServiceHelmValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StackServiceHelmValue) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *StackServiceHelmValue) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *StackServiceHelmValue) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *StackServiceHelmValue) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetValueSecretId

`func (o *StackServiceHelmValue) GetValueSecretId() int32`

GetValueSecretId returns the ValueSecretId field if non-nil, zero value otherwise.

### GetValueSecretIdOk

`func (o *StackServiceHelmValue) GetValueSecretIdOk() (*int32, bool)`

GetValueSecretIdOk returns a tuple with the ValueSecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueSecretId

`func (o *StackServiceHelmValue) SetValueSecretId(v int32)`

SetValueSecretId sets ValueSecretId field to given value.

### HasValueSecretId

`func (o *StackServiceHelmValue) HasValueSecretId() bool`

HasValueSecretId returns a boolean if a field has been set.

### SetValueSecretIdNil

`func (o *StackServiceHelmValue) SetValueSecretIdNil(b bool)`

 SetValueSecretIdNil sets the value for ValueSecretId to be an explicit nil

### UnsetValueSecretId
`func (o *StackServiceHelmValue) UnsetValueSecretId()`

UnsetValueSecretId ensures that no value is present for ValueSecretId, not even an explicit nil
### GetEnvType

`func (o *StackServiceHelmValue) GetEnvType() string`

GetEnvType returns the EnvType field if non-nil, zero value otherwise.

### GetEnvTypeOk

`func (o *StackServiceHelmValue) GetEnvTypeOk() (*string, bool)`

GetEnvTypeOk returns a tuple with the EnvType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvType

`func (o *StackServiceHelmValue) SetEnvType(v string)`

SetEnvType sets EnvType field to given value.

### HasEnvType

`func (o *StackServiceHelmValue) HasEnvType() bool`

HasEnvType returns a boolean if a field has been set.

### SetEnvTypeNil

`func (o *StackServiceHelmValue) SetEnvTypeNil(b bool)`

 SetEnvTypeNil sets the value for EnvType to be an explicit nil

### UnsetEnvType
`func (o *StackServiceHelmValue) UnsetEnvType()`

UnsetEnvType ensures that no value is present for EnvType, not even an explicit nil
### GetCreatedAt

`func (o *StackServiceHelmValue) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StackServiceHelmValue) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StackServiceHelmValue) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


