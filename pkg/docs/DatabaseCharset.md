# DatabaseCharset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Title** | **string** |  | 
**DefaultCollation** | **string** |  | 
**Default** | **bool** |  | 

## Methods

### NewDatabaseCharset

`func NewDatabaseCharset(name string, title string, defaultCollation string, default_ bool, ) *DatabaseCharset`

NewDatabaseCharset instantiates a new DatabaseCharset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseCharsetWithDefaults

`func NewDatabaseCharsetWithDefaults() *DatabaseCharset`

NewDatabaseCharsetWithDefaults instantiates a new DatabaseCharset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DatabaseCharset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatabaseCharset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatabaseCharset) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *DatabaseCharset) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DatabaseCharset) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DatabaseCharset) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDefaultCollation

`func (o *DatabaseCharset) GetDefaultCollation() string`

GetDefaultCollation returns the DefaultCollation field if non-nil, zero value otherwise.

### GetDefaultCollationOk

`func (o *DatabaseCharset) GetDefaultCollationOk() (*string, bool)`

GetDefaultCollationOk returns a tuple with the DefaultCollation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCollation

`func (o *DatabaseCharset) SetDefaultCollation(v string)`

SetDefaultCollation sets DefaultCollation field to given value.


### GetDefault

`func (o *DatabaseCharset) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *DatabaseCharset) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *DatabaseCharset) SetDefault(v bool)`

SetDefault sets Default field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


