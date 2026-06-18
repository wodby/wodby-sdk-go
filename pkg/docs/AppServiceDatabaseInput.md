# AppServiceDatabaseInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseId** | **int32** |  | 
**DatabaseDbId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewAppServiceDatabaseInput

`func NewAppServiceDatabaseInput(databaseId int32, ) *AppServiceDatabaseInput`

NewAppServiceDatabaseInput instantiates a new AppServiceDatabaseInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceDatabaseInputWithDefaults

`func NewAppServiceDatabaseInputWithDefaults() *AppServiceDatabaseInput`

NewAppServiceDatabaseInputWithDefaults instantiates a new AppServiceDatabaseInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseId

`func (o *AppServiceDatabaseInput) GetDatabaseId() int32`

GetDatabaseId returns the DatabaseId field if non-nil, zero value otherwise.

### GetDatabaseIdOk

`func (o *AppServiceDatabaseInput) GetDatabaseIdOk() (*int32, bool)`

GetDatabaseIdOk returns a tuple with the DatabaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseId

`func (o *AppServiceDatabaseInput) SetDatabaseId(v int32)`

SetDatabaseId sets DatabaseId field to given value.


### GetDatabaseDbId

`func (o *AppServiceDatabaseInput) GetDatabaseDbId() int32`

GetDatabaseDbId returns the DatabaseDbId field if non-nil, zero value otherwise.

### GetDatabaseDbIdOk

`func (o *AppServiceDatabaseInput) GetDatabaseDbIdOk() (*int32, bool)`

GetDatabaseDbIdOk returns a tuple with the DatabaseDbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseDbId

`func (o *AppServiceDatabaseInput) SetDatabaseDbId(v int32)`

SetDatabaseDbId sets DatabaseDbId field to given value.

### HasDatabaseDbId

`func (o *AppServiceDatabaseInput) HasDatabaseDbId() bool`

HasDatabaseDbId returns a boolean if a field has been set.

### SetDatabaseDbIdNil

`func (o *AppServiceDatabaseInput) SetDatabaseDbIdNil(b bool)`

 SetDatabaseDbIdNil sets the value for DatabaseDbId to be an explicit nil

### UnsetDatabaseDbId
`func (o *AppServiceDatabaseInput) UnsetDatabaseDbId()`

UnsetDatabaseDbId ensures that no value is present for DatabaseDbId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


