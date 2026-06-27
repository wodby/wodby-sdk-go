# StackServiceAnnotation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**StackServiceId** | **int32** |  | 
**Name** | **string** |  | 
**Value** | Pointer to **NullableString** |  | [optional] 
**EnvType** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewStackServiceAnnotation

`func NewStackServiceAnnotation(id int32, stackServiceId int32, name string, createdAt time.Time, ) *StackServiceAnnotation`

NewStackServiceAnnotation instantiates a new StackServiceAnnotation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackServiceAnnotationWithDefaults

`func NewStackServiceAnnotationWithDefaults() *StackServiceAnnotation`

NewStackServiceAnnotationWithDefaults instantiates a new StackServiceAnnotation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StackServiceAnnotation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StackServiceAnnotation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StackServiceAnnotation) SetId(v int32)`

SetId sets Id field to given value.


### GetStackServiceId

`func (o *StackServiceAnnotation) GetStackServiceId() int32`

GetStackServiceId returns the StackServiceId field if non-nil, zero value otherwise.

### GetStackServiceIdOk

`func (o *StackServiceAnnotation) GetStackServiceIdOk() (*int32, bool)`

GetStackServiceIdOk returns a tuple with the StackServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackServiceId

`func (o *StackServiceAnnotation) SetStackServiceId(v int32)`

SetStackServiceId sets StackServiceId field to given value.


### GetName

`func (o *StackServiceAnnotation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StackServiceAnnotation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StackServiceAnnotation) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *StackServiceAnnotation) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StackServiceAnnotation) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StackServiceAnnotation) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *StackServiceAnnotation) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *StackServiceAnnotation) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *StackServiceAnnotation) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetEnvType

`func (o *StackServiceAnnotation) GetEnvType() string`

GetEnvType returns the EnvType field if non-nil, zero value otherwise.

### GetEnvTypeOk

`func (o *StackServiceAnnotation) GetEnvTypeOk() (*string, bool)`

GetEnvTypeOk returns a tuple with the EnvType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvType

`func (o *StackServiceAnnotation) SetEnvType(v string)`

SetEnvType sets EnvType field to given value.

### HasEnvType

`func (o *StackServiceAnnotation) HasEnvType() bool`

HasEnvType returns a boolean if a field has been set.

### SetEnvTypeNil

`func (o *StackServiceAnnotation) SetEnvTypeNil(b bool)`

 SetEnvTypeNil sets the value for EnvType to be an explicit nil

### UnsetEnvType
`func (o *StackServiceAnnotation) UnsetEnvType()`

UnsetEnvType ensures that no value is present for EnvType, not even an explicit nil
### GetCreatedAt

`func (o *StackServiceAnnotation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StackServiceAnnotation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StackServiceAnnotation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


