# CreateImportInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppServiceId** | Pointer to **NullableInt32** |  | [optional] 
**DatabaseDbId** | Pointer to **NullableInt32** |  | [optional] 
**Import** | [**ImportInput**](ImportInput.md) |  | 

## Methods

### NewCreateImportInput

`func NewCreateImportInput(import_ ImportInput, ) *CreateImportInput`

NewCreateImportInput instantiates a new CreateImportInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateImportInputWithDefaults

`func NewCreateImportInputWithDefaults() *CreateImportInput`

NewCreateImportInputWithDefaults instantiates a new CreateImportInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppServiceId

`func (o *CreateImportInput) GetAppServiceId() int32`

GetAppServiceId returns the AppServiceId field if non-nil, zero value otherwise.

### GetAppServiceIdOk

`func (o *CreateImportInput) GetAppServiceIdOk() (*int32, bool)`

GetAppServiceIdOk returns a tuple with the AppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceId

`func (o *CreateImportInput) SetAppServiceId(v int32)`

SetAppServiceId sets AppServiceId field to given value.

### HasAppServiceId

`func (o *CreateImportInput) HasAppServiceId() bool`

HasAppServiceId returns a boolean if a field has been set.

### SetAppServiceIdNil

`func (o *CreateImportInput) SetAppServiceIdNil(b bool)`

 SetAppServiceIdNil sets the value for AppServiceId to be an explicit nil

### UnsetAppServiceId
`func (o *CreateImportInput) UnsetAppServiceId()`

UnsetAppServiceId ensures that no value is present for AppServiceId, not even an explicit nil
### GetDatabaseDbId

`func (o *CreateImportInput) GetDatabaseDbId() int32`

GetDatabaseDbId returns the DatabaseDbId field if non-nil, zero value otherwise.

### GetDatabaseDbIdOk

`func (o *CreateImportInput) GetDatabaseDbIdOk() (*int32, bool)`

GetDatabaseDbIdOk returns a tuple with the DatabaseDbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseDbId

`func (o *CreateImportInput) SetDatabaseDbId(v int32)`

SetDatabaseDbId sets DatabaseDbId field to given value.

### HasDatabaseDbId

`func (o *CreateImportInput) HasDatabaseDbId() bool`

HasDatabaseDbId returns a boolean if a field has been set.

### SetDatabaseDbIdNil

`func (o *CreateImportInput) SetDatabaseDbIdNil(b bool)`

 SetDatabaseDbIdNil sets the value for DatabaseDbId to be an explicit nil

### UnsetDatabaseDbId
`func (o *CreateImportInput) UnsetDatabaseDbId()`

UnsetDatabaseDbId ensures that no value is present for DatabaseDbId, not even an explicit nil
### GetImport

`func (o *CreateImportInput) GetImport() ImportInput`

GetImport returns the Import field if non-nil, zero value otherwise.

### GetImportOk

`func (o *CreateImportInput) GetImportOk() (*ImportInput, bool)`

GetImportOk returns a tuple with the Import field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImport

`func (o *CreateImportInput) SetImport(v ImportInput)`

SetImport sets Import field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


