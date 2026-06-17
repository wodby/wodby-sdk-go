# AppServiceDatabaseInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseID** | **int32** |  | 
**DatabaseDBID** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewAppServiceDatabaseInput

`func NewAppServiceDatabaseInput(databaseID int32, ) *AppServiceDatabaseInput`

NewAppServiceDatabaseInput instantiates a new AppServiceDatabaseInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceDatabaseInputWithDefaults

`func NewAppServiceDatabaseInputWithDefaults() *AppServiceDatabaseInput`

NewAppServiceDatabaseInputWithDefaults instantiates a new AppServiceDatabaseInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseID

`func (o *AppServiceDatabaseInput) GetDatabaseID() int32`

GetDatabaseID returns the DatabaseID field if non-nil, zero value otherwise.

### GetDatabaseIDOk

`func (o *AppServiceDatabaseInput) GetDatabaseIDOk() (*int32, bool)`

GetDatabaseIDOk returns a tuple with the DatabaseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseID

`func (o *AppServiceDatabaseInput) SetDatabaseID(v int32)`

SetDatabaseID sets DatabaseID field to given value.


### GetDatabaseDBID

`func (o *AppServiceDatabaseInput) GetDatabaseDBID() int32`

GetDatabaseDBID returns the DatabaseDBID field if non-nil, zero value otherwise.

### GetDatabaseDBIDOk

`func (o *AppServiceDatabaseInput) GetDatabaseDBIDOk() (*int32, bool)`

GetDatabaseDBIDOk returns a tuple with the DatabaseDBID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseDBID

`func (o *AppServiceDatabaseInput) SetDatabaseDBID(v int32)`

SetDatabaseDBID sets DatabaseDBID field to given value.

### HasDatabaseDBID

`func (o *AppServiceDatabaseInput) HasDatabaseDBID() bool`

HasDatabaseDBID returns a boolean if a field has been set.

### SetDatabaseDBIDNil

`func (o *AppServiceDatabaseInput) SetDatabaseDBIDNil(b bool)`

 SetDatabaseDBIDNil sets the value for DatabaseDBID to be an explicit nil

### UnsetDatabaseDBID
`func (o *AppServiceDatabaseInput) UnsetDatabaseDBID()`

UnsetDatabaseDBID ensures that no value is present for DatabaseDBID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


