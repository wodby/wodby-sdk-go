# NewImportInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppServiceID** | Pointer to **NullableInt32** |  | [optional] 
**DatabaseDBID** | Pointer to **NullableInt32** |  | [optional] 
**Import** | [**ImportInput**](ImportInput.md) |  | 

## Methods

### NewNewImportInput

`func NewNewImportInput(import_ ImportInput, ) *NewImportInput`

NewNewImportInput instantiates a new NewImportInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewImportInputWithDefaults

`func NewNewImportInputWithDefaults() *NewImportInput`

NewNewImportInputWithDefaults instantiates a new NewImportInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppServiceID

`func (o *NewImportInput) GetAppServiceID() int32`

GetAppServiceID returns the AppServiceID field if non-nil, zero value otherwise.

### GetAppServiceIDOk

`func (o *NewImportInput) GetAppServiceIDOk() (*int32, bool)`

GetAppServiceIDOk returns a tuple with the AppServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceID

`func (o *NewImportInput) SetAppServiceID(v int32)`

SetAppServiceID sets AppServiceID field to given value.

### HasAppServiceID

`func (o *NewImportInput) HasAppServiceID() bool`

HasAppServiceID returns a boolean if a field has been set.

### SetAppServiceIDNil

`func (o *NewImportInput) SetAppServiceIDNil(b bool)`

 SetAppServiceIDNil sets the value for AppServiceID to be an explicit nil

### UnsetAppServiceID
`func (o *NewImportInput) UnsetAppServiceID()`

UnsetAppServiceID ensures that no value is present for AppServiceID, not even an explicit nil
### GetDatabaseDBID

`func (o *NewImportInput) GetDatabaseDBID() int32`

GetDatabaseDBID returns the DatabaseDBID field if non-nil, zero value otherwise.

### GetDatabaseDBIDOk

`func (o *NewImportInput) GetDatabaseDBIDOk() (*int32, bool)`

GetDatabaseDBIDOk returns a tuple with the DatabaseDBID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseDBID

`func (o *NewImportInput) SetDatabaseDBID(v int32)`

SetDatabaseDBID sets DatabaseDBID field to given value.

### HasDatabaseDBID

`func (o *NewImportInput) HasDatabaseDBID() bool`

HasDatabaseDBID returns a boolean if a field has been set.

### SetDatabaseDBIDNil

`func (o *NewImportInput) SetDatabaseDBIDNil(b bool)`

 SetDatabaseDBIDNil sets the value for DatabaseDBID to be an explicit nil

### UnsetDatabaseDBID
`func (o *NewImportInput) UnsetDatabaseDBID()`

UnsetDatabaseDBID ensures that no value is present for DatabaseDBID, not even an explicit nil
### GetImport

`func (o *NewImportInput) GetImport() ImportInput`

GetImport returns the Import field if non-nil, zero value otherwise.

### GetImportOk

`func (o *NewImportInput) GetImportOk() (*ImportInput, bool)`

GetImportOk returns a tuple with the Import field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImport

`func (o *NewImportInput) SetImport(v ImportInput)`

SetImport sets Import field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


