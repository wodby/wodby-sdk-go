# StackServiceToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**StackServiceId** | **int32** |  | 
**Name** | **string** |  | 
**Value** | Pointer to **NullableString** |  | [optional] 
**Regex** | Pointer to **NullableString** |  | [optional] 
**ValueSecretId** | Pointer to **NullableInt32** |  | [optional] 
**EnvType** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewStackServiceToken

`func NewStackServiceToken(id int32, stackServiceId int32, name string, createdAt time.Time, ) *StackServiceToken`

NewStackServiceToken instantiates a new StackServiceToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackServiceTokenWithDefaults

`func NewStackServiceTokenWithDefaults() *StackServiceToken`

NewStackServiceTokenWithDefaults instantiates a new StackServiceToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StackServiceToken) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StackServiceToken) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StackServiceToken) SetId(v int32)`

SetId sets Id field to given value.


### GetStackServiceId

`func (o *StackServiceToken) GetStackServiceId() int32`

GetStackServiceId returns the StackServiceId field if non-nil, zero value otherwise.

### GetStackServiceIdOk

`func (o *StackServiceToken) GetStackServiceIdOk() (*int32, bool)`

GetStackServiceIdOk returns a tuple with the StackServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackServiceId

`func (o *StackServiceToken) SetStackServiceId(v int32)`

SetStackServiceId sets StackServiceId field to given value.


### GetName

`func (o *StackServiceToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StackServiceToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StackServiceToken) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *StackServiceToken) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StackServiceToken) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StackServiceToken) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *StackServiceToken) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *StackServiceToken) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *StackServiceToken) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetRegex

`func (o *StackServiceToken) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *StackServiceToken) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *StackServiceToken) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *StackServiceToken) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### SetRegexNil

`func (o *StackServiceToken) SetRegexNil(b bool)`

 SetRegexNil sets the value for Regex to be an explicit nil

### UnsetRegex
`func (o *StackServiceToken) UnsetRegex()`

UnsetRegex ensures that no value is present for Regex, not even an explicit nil
### GetValueSecretId

`func (o *StackServiceToken) GetValueSecretId() int32`

GetValueSecretId returns the ValueSecretId field if non-nil, zero value otherwise.

### GetValueSecretIdOk

`func (o *StackServiceToken) GetValueSecretIdOk() (*int32, bool)`

GetValueSecretIdOk returns a tuple with the ValueSecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueSecretId

`func (o *StackServiceToken) SetValueSecretId(v int32)`

SetValueSecretId sets ValueSecretId field to given value.

### HasValueSecretId

`func (o *StackServiceToken) HasValueSecretId() bool`

HasValueSecretId returns a boolean if a field has been set.

### SetValueSecretIdNil

`func (o *StackServiceToken) SetValueSecretIdNil(b bool)`

 SetValueSecretIdNil sets the value for ValueSecretId to be an explicit nil

### UnsetValueSecretId
`func (o *StackServiceToken) UnsetValueSecretId()`

UnsetValueSecretId ensures that no value is present for ValueSecretId, not even an explicit nil
### GetEnvType

`func (o *StackServiceToken) GetEnvType() string`

GetEnvType returns the EnvType field if non-nil, zero value otherwise.

### GetEnvTypeOk

`func (o *StackServiceToken) GetEnvTypeOk() (*string, bool)`

GetEnvTypeOk returns a tuple with the EnvType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvType

`func (o *StackServiceToken) SetEnvType(v string)`

SetEnvType sets EnvType field to given value.

### HasEnvType

`func (o *StackServiceToken) HasEnvType() bool`

HasEnvType returns a boolean if a field has been set.

### SetEnvTypeNil

`func (o *StackServiceToken) SetEnvTypeNil(b bool)`

 SetEnvTypeNil sets the value for EnvType to be an explicit nil

### UnsetEnvType
`func (o *StackServiceToken) UnsetEnvType()`

UnsetEnvType ensures that no value is present for EnvType, not even an explicit nil
### GetCreatedAt

`func (o *StackServiceToken) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StackServiceToken) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StackServiceToken) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


