# DatabaseDB

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**DatabaseId** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**Collation** | **string** |  | 
**Charset** | **string** |  | 
**Status** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewDatabaseDB

`func NewDatabaseDB(id int32, name string, collation string, charset string, status string, updatedAt time.Time, createdAt time.Time, ) *DatabaseDB`

NewDatabaseDB instantiates a new DatabaseDB object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseDBWithDefaults

`func NewDatabaseDBWithDefaults() *DatabaseDB`

NewDatabaseDBWithDefaults instantiates a new DatabaseDB object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DatabaseDB) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatabaseDB) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatabaseDB) SetId(v int32)`

SetId sets Id field to given value.


### GetDatabaseId

`func (o *DatabaseDB) GetDatabaseId() int32`

GetDatabaseId returns the DatabaseId field if non-nil, zero value otherwise.

### GetDatabaseIdOk

`func (o *DatabaseDB) GetDatabaseIdOk() (*int32, bool)`

GetDatabaseIdOk returns a tuple with the DatabaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseId

`func (o *DatabaseDB) SetDatabaseId(v int32)`

SetDatabaseId sets DatabaseId field to given value.

### HasDatabaseId

`func (o *DatabaseDB) HasDatabaseId() bool`

HasDatabaseId returns a boolean if a field has been set.

### GetName

`func (o *DatabaseDB) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatabaseDB) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatabaseDB) SetName(v string)`

SetName sets Name field to given value.


### GetCollation

`func (o *DatabaseDB) GetCollation() string`

GetCollation returns the Collation field if non-nil, zero value otherwise.

### GetCollationOk

`func (o *DatabaseDB) GetCollationOk() (*string, bool)`

GetCollationOk returns a tuple with the Collation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollation

`func (o *DatabaseDB) SetCollation(v string)`

SetCollation sets Collation field to given value.


### GetCharset

`func (o *DatabaseDB) GetCharset() string`

GetCharset returns the Charset field if non-nil, zero value otherwise.

### GetCharsetOk

`func (o *DatabaseDB) GetCharsetOk() (*string, bool)`

GetCharsetOk returns a tuple with the Charset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharset

`func (o *DatabaseDB) SetCharset(v string)`

SetCharset sets Charset field to given value.


### GetStatus

`func (o *DatabaseDB) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DatabaseDB) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DatabaseDB) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *DatabaseDB) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DatabaseDB) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DatabaseDB) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *DatabaseDB) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DatabaseDB) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DatabaseDB) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


