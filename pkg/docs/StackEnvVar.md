# StackEnvVar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Value** | Pointer to **NullableString** |  | [optional] 
**ValueSecretId** | Pointer to **NullableInt32** |  | [optional] 
**EnvType** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewStackEnvVar

`func NewStackEnvVar(id int32, name string, createdAt time.Time, ) *StackEnvVar`

NewStackEnvVar instantiates a new StackEnvVar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackEnvVarWithDefaults

`func NewStackEnvVarWithDefaults() *StackEnvVar`

NewStackEnvVarWithDefaults instantiates a new StackEnvVar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StackEnvVar) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StackEnvVar) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StackEnvVar) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *StackEnvVar) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StackEnvVar) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StackEnvVar) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *StackEnvVar) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StackEnvVar) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StackEnvVar) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *StackEnvVar) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *StackEnvVar) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *StackEnvVar) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetValueSecretId

`func (o *StackEnvVar) GetValueSecretId() int32`

GetValueSecretId returns the ValueSecretId field if non-nil, zero value otherwise.

### GetValueSecretIdOk

`func (o *StackEnvVar) GetValueSecretIdOk() (*int32, bool)`

GetValueSecretIdOk returns a tuple with the ValueSecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueSecretId

`func (o *StackEnvVar) SetValueSecretId(v int32)`

SetValueSecretId sets ValueSecretId field to given value.

### HasValueSecretId

`func (o *StackEnvVar) HasValueSecretId() bool`

HasValueSecretId returns a boolean if a field has been set.

### SetValueSecretIdNil

`func (o *StackEnvVar) SetValueSecretIdNil(b bool)`

 SetValueSecretIdNil sets the value for ValueSecretId to be an explicit nil

### UnsetValueSecretId
`func (o *StackEnvVar) UnsetValueSecretId()`

UnsetValueSecretId ensures that no value is present for ValueSecretId, not even an explicit nil
### GetEnvType

`func (o *StackEnvVar) GetEnvType() string`

GetEnvType returns the EnvType field if non-nil, zero value otherwise.

### GetEnvTypeOk

`func (o *StackEnvVar) GetEnvTypeOk() (*string, bool)`

GetEnvTypeOk returns a tuple with the EnvType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvType

`func (o *StackEnvVar) SetEnvType(v string)`

SetEnvType sets EnvType field to given value.

### HasEnvType

`func (o *StackEnvVar) HasEnvType() bool`

HasEnvType returns a boolean if a field has been set.

### SetEnvTypeNil

`func (o *StackEnvVar) SetEnvTypeNil(b bool)`

 SetEnvTypeNil sets the value for EnvType to be an explicit nil

### UnsetEnvType
`func (o *StackEnvVar) UnsetEnvType()`

UnsetEnvType ensures that no value is present for EnvType, not even an explicit nil
### GetCreatedAt

`func (o *StackEnvVar) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StackEnvVar) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StackEnvVar) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


