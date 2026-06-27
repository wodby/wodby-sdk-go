# AppServiceAnnotation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**AppServiceId** | **int32** |  | 
**Name** | **string** |  | 
**Value** | **string** |  | 
**EnvType** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to [**NullableAppServiceAnnotationSource**](AppServiceAnnotationSource.md) |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewAppServiceAnnotation

`func NewAppServiceAnnotation(id int32, appServiceId int32, name string, value string, ) *AppServiceAnnotation`

NewAppServiceAnnotation instantiates a new AppServiceAnnotation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceAnnotationWithDefaults

`func NewAppServiceAnnotationWithDefaults() *AppServiceAnnotation`

NewAppServiceAnnotationWithDefaults instantiates a new AppServiceAnnotation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppServiceAnnotation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppServiceAnnotation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppServiceAnnotation) SetId(v int32)`

SetId sets Id field to given value.


### GetAppServiceId

`func (o *AppServiceAnnotation) GetAppServiceId() int32`

GetAppServiceId returns the AppServiceId field if non-nil, zero value otherwise.

### GetAppServiceIdOk

`func (o *AppServiceAnnotation) GetAppServiceIdOk() (*int32, bool)`

GetAppServiceIdOk returns a tuple with the AppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceId

`func (o *AppServiceAnnotation) SetAppServiceId(v int32)`

SetAppServiceId sets AppServiceId field to given value.


### GetName

`func (o *AppServiceAnnotation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppServiceAnnotation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppServiceAnnotation) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *AppServiceAnnotation) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AppServiceAnnotation) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AppServiceAnnotation) SetValue(v string)`

SetValue sets Value field to given value.


### GetEnvType

`func (o *AppServiceAnnotation) GetEnvType() string`

GetEnvType returns the EnvType field if non-nil, zero value otherwise.

### GetEnvTypeOk

`func (o *AppServiceAnnotation) GetEnvTypeOk() (*string, bool)`

GetEnvTypeOk returns a tuple with the EnvType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvType

`func (o *AppServiceAnnotation) SetEnvType(v string)`

SetEnvType sets EnvType field to given value.

### HasEnvType

`func (o *AppServiceAnnotation) HasEnvType() bool`

HasEnvType returns a boolean if a field has been set.

### SetEnvTypeNil

`func (o *AppServiceAnnotation) SetEnvTypeNil(b bool)`

 SetEnvTypeNil sets the value for EnvType to be an explicit nil

### UnsetEnvType
`func (o *AppServiceAnnotation) UnsetEnvType()`

UnsetEnvType ensures that no value is present for EnvType, not even an explicit nil
### GetSource

`func (o *AppServiceAnnotation) GetSource() AppServiceAnnotationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AppServiceAnnotation) GetSourceOk() (*AppServiceAnnotationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AppServiceAnnotation) SetSource(v AppServiceAnnotationSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *AppServiceAnnotation) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *AppServiceAnnotation) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *AppServiceAnnotation) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetCreatedAt

`func (o *AppServiceAnnotation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppServiceAnnotation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppServiceAnnotation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AppServiceAnnotation) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *AppServiceAnnotation) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *AppServiceAnnotation) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


