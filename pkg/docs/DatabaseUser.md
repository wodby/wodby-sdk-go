# DatabaseUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**DatabaseId** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**PasswordSecretId** | **int32** |  | 
**Hostname** | Pointer to **NullableString** |  | [optional] 
**Status** | **string** |  | 
**Dbs** | Pointer to [**[]DatabaseDB**](DatabaseDB.md) |  | [optional] 
**UpdatedAt** | **time.Time** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewDatabaseUser

`func NewDatabaseUser(id int32, name string, passwordSecretId int32, status string, updatedAt time.Time, createdAt time.Time, ) *DatabaseUser`

NewDatabaseUser instantiates a new DatabaseUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseUserWithDefaults

`func NewDatabaseUserWithDefaults() *DatabaseUser`

NewDatabaseUserWithDefaults instantiates a new DatabaseUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DatabaseUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatabaseUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatabaseUser) SetId(v int32)`

SetId sets Id field to given value.


### GetDatabaseId

`func (o *DatabaseUser) GetDatabaseId() int32`

GetDatabaseId returns the DatabaseId field if non-nil, zero value otherwise.

### GetDatabaseIdOk

`func (o *DatabaseUser) GetDatabaseIdOk() (*int32, bool)`

GetDatabaseIdOk returns a tuple with the DatabaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseId

`func (o *DatabaseUser) SetDatabaseId(v int32)`

SetDatabaseId sets DatabaseId field to given value.

### HasDatabaseId

`func (o *DatabaseUser) HasDatabaseId() bool`

HasDatabaseId returns a boolean if a field has been set.

### GetName

`func (o *DatabaseUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatabaseUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatabaseUser) SetName(v string)`

SetName sets Name field to given value.


### GetPasswordSecretId

`func (o *DatabaseUser) GetPasswordSecretId() int32`

GetPasswordSecretId returns the PasswordSecretId field if non-nil, zero value otherwise.

### GetPasswordSecretIdOk

`func (o *DatabaseUser) GetPasswordSecretIdOk() (*int32, bool)`

GetPasswordSecretIdOk returns a tuple with the PasswordSecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordSecretId

`func (o *DatabaseUser) SetPasswordSecretId(v int32)`

SetPasswordSecretId sets PasswordSecretId field to given value.


### GetHostname

`func (o *DatabaseUser) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *DatabaseUser) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *DatabaseUser) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *DatabaseUser) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### SetHostnameNil

`func (o *DatabaseUser) SetHostnameNil(b bool)`

 SetHostnameNil sets the value for Hostname to be an explicit nil

### UnsetHostname
`func (o *DatabaseUser) UnsetHostname()`

UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
### GetStatus

`func (o *DatabaseUser) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DatabaseUser) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DatabaseUser) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDbs

`func (o *DatabaseUser) GetDbs() []DatabaseDB`

GetDbs returns the Dbs field if non-nil, zero value otherwise.

### GetDbsOk

`func (o *DatabaseUser) GetDbsOk() (*[]DatabaseDB, bool)`

GetDbsOk returns a tuple with the Dbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbs

`func (o *DatabaseUser) SetDbs(v []DatabaseDB)`

SetDbs sets Dbs field to given value.

### HasDbs

`func (o *DatabaseUser) HasDbs() bool`

HasDbs returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DatabaseUser) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DatabaseUser) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DatabaseUser) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *DatabaseUser) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DatabaseUser) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DatabaseUser) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


