# NewDatabaseUserInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseId** | **int32** |  | 
**Name** | **string** |  | 
**Password** | **string** |  | 
**Hostname** | Pointer to **NullableString** |  | [optional] 
**DatabaseDbIds** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewNewDatabaseUserInput

`func NewNewDatabaseUserInput(databaseId int32, name string, password string, ) *NewDatabaseUserInput`

NewNewDatabaseUserInput instantiates a new NewDatabaseUserInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewDatabaseUserInputWithDefaults

`func NewNewDatabaseUserInputWithDefaults() *NewDatabaseUserInput`

NewNewDatabaseUserInputWithDefaults instantiates a new NewDatabaseUserInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseId

`func (o *NewDatabaseUserInput) GetDatabaseId() int32`

GetDatabaseId returns the DatabaseId field if non-nil, zero value otherwise.

### GetDatabaseIdOk

`func (o *NewDatabaseUserInput) GetDatabaseIdOk() (*int32, bool)`

GetDatabaseIdOk returns a tuple with the DatabaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseId

`func (o *NewDatabaseUserInput) SetDatabaseId(v int32)`

SetDatabaseId sets DatabaseId field to given value.


### GetName

`func (o *NewDatabaseUserInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewDatabaseUserInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewDatabaseUserInput) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *NewDatabaseUserInput) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *NewDatabaseUserInput) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *NewDatabaseUserInput) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetHostname

`func (o *NewDatabaseUserInput) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *NewDatabaseUserInput) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *NewDatabaseUserInput) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *NewDatabaseUserInput) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### SetHostnameNil

`func (o *NewDatabaseUserInput) SetHostnameNil(b bool)`

 SetHostnameNil sets the value for Hostname to be an explicit nil

### UnsetHostname
`func (o *NewDatabaseUserInput) UnsetHostname()`

UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
### GetDatabaseDbIds

`func (o *NewDatabaseUserInput) GetDatabaseDbIds() []int32`

GetDatabaseDbIds returns the DatabaseDbIds field if non-nil, zero value otherwise.

### GetDatabaseDbIdsOk

`func (o *NewDatabaseUserInput) GetDatabaseDbIdsOk() (*[]int32, bool)`

GetDatabaseDbIdsOk returns a tuple with the DatabaseDbIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseDbIds

`func (o *NewDatabaseUserInput) SetDatabaseDbIds(v []int32)`

SetDatabaseDbIds sets DatabaseDbIds field to given value.

### HasDatabaseDbIds

`func (o *NewDatabaseUserInput) HasDatabaseDbIds() bool`

HasDatabaseDbIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


