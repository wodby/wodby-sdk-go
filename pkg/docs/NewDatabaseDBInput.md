# NewDatabaseDBInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseId** | **int32** |  | 
**Name** | **string** |  | 
**Charset** | Pointer to **NullableString** |  | [optional] 
**Collation** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNewDatabaseDBInput

`func NewNewDatabaseDBInput(databaseId int32, name string, ) *NewDatabaseDBInput`

NewNewDatabaseDBInput instantiates a new NewDatabaseDBInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewDatabaseDBInputWithDefaults

`func NewNewDatabaseDBInputWithDefaults() *NewDatabaseDBInput`

NewNewDatabaseDBInputWithDefaults instantiates a new NewDatabaseDBInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseId

`func (o *NewDatabaseDBInput) GetDatabaseId() int32`

GetDatabaseId returns the DatabaseId field if non-nil, zero value otherwise.

### GetDatabaseIdOk

`func (o *NewDatabaseDBInput) GetDatabaseIdOk() (*int32, bool)`

GetDatabaseIdOk returns a tuple with the DatabaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseId

`func (o *NewDatabaseDBInput) SetDatabaseId(v int32)`

SetDatabaseId sets DatabaseId field to given value.


### GetName

`func (o *NewDatabaseDBInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewDatabaseDBInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewDatabaseDBInput) SetName(v string)`

SetName sets Name field to given value.


### GetCharset

`func (o *NewDatabaseDBInput) GetCharset() string`

GetCharset returns the Charset field if non-nil, zero value otherwise.

### GetCharsetOk

`func (o *NewDatabaseDBInput) GetCharsetOk() (*string, bool)`

GetCharsetOk returns a tuple with the Charset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharset

`func (o *NewDatabaseDBInput) SetCharset(v string)`

SetCharset sets Charset field to given value.

### HasCharset

`func (o *NewDatabaseDBInput) HasCharset() bool`

HasCharset returns a boolean if a field has been set.

### SetCharsetNil

`func (o *NewDatabaseDBInput) SetCharsetNil(b bool)`

 SetCharsetNil sets the value for Charset to be an explicit nil

### UnsetCharset
`func (o *NewDatabaseDBInput) UnsetCharset()`

UnsetCharset ensures that no value is present for Charset, not even an explicit nil
### GetCollation

`func (o *NewDatabaseDBInput) GetCollation() string`

GetCollation returns the Collation field if non-nil, zero value otherwise.

### GetCollationOk

`func (o *NewDatabaseDBInput) GetCollationOk() (*string, bool)`

GetCollationOk returns a tuple with the Collation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollation

`func (o *NewDatabaseDBInput) SetCollation(v string)`

SetCollation sets Collation field to given value.

### HasCollation

`func (o *NewDatabaseDBInput) HasCollation() bool`

HasCollation returns a boolean if a field has been set.

### SetCollationNil

`func (o *NewDatabaseDBInput) SetCollationNil(b bool)`

 SetCollationNil sets the value for Collation to be an explicit nil

### UnsetCollation
`func (o *NewDatabaseDBInput) UnsetCollation()`

UnsetCollation ensures that no value is present for Collation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


