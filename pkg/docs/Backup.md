# Backup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Status** | **string** |  | 
**AppInstanceId** | Pointer to **NullableInt32** |  | [optional] 
**AppServiceId** | Pointer to **NullableInt32** |  | [optional] 
**DatabaseId** | Pointer to **NullableInt32** |  | [optional] 
**DatabaseDbId** | Pointer to **NullableInt32** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewBackup

`func NewBackup(id int32, name string, status string, createdAt time.Time, updatedAt time.Time, ) *Backup`

NewBackup instantiates a new Backup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupWithDefaults

`func NewBackupWithDefaults() *Backup`

NewBackupWithDefaults instantiates a new Backup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Backup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Backup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Backup) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Backup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Backup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Backup) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *Backup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Backup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Backup) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAppInstanceId

`func (o *Backup) GetAppInstanceId() int32`

GetAppInstanceId returns the AppInstanceId field if non-nil, zero value otherwise.

### GetAppInstanceIdOk

`func (o *Backup) GetAppInstanceIdOk() (*int32, bool)`

GetAppInstanceIdOk returns a tuple with the AppInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceId

`func (o *Backup) SetAppInstanceId(v int32)`

SetAppInstanceId sets AppInstanceId field to given value.

### HasAppInstanceId

`func (o *Backup) HasAppInstanceId() bool`

HasAppInstanceId returns a boolean if a field has been set.

### SetAppInstanceIdNil

`func (o *Backup) SetAppInstanceIdNil(b bool)`

 SetAppInstanceIdNil sets the value for AppInstanceId to be an explicit nil

### UnsetAppInstanceId
`func (o *Backup) UnsetAppInstanceId()`

UnsetAppInstanceId ensures that no value is present for AppInstanceId, not even an explicit nil
### GetAppServiceId

`func (o *Backup) GetAppServiceId() int32`

GetAppServiceId returns the AppServiceId field if non-nil, zero value otherwise.

### GetAppServiceIdOk

`func (o *Backup) GetAppServiceIdOk() (*int32, bool)`

GetAppServiceIdOk returns a tuple with the AppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceId

`func (o *Backup) SetAppServiceId(v int32)`

SetAppServiceId sets AppServiceId field to given value.

### HasAppServiceId

`func (o *Backup) HasAppServiceId() bool`

HasAppServiceId returns a boolean if a field has been set.

### SetAppServiceIdNil

`func (o *Backup) SetAppServiceIdNil(b bool)`

 SetAppServiceIdNil sets the value for AppServiceId to be an explicit nil

### UnsetAppServiceId
`func (o *Backup) UnsetAppServiceId()`

UnsetAppServiceId ensures that no value is present for AppServiceId, not even an explicit nil
### GetDatabaseId

`func (o *Backup) GetDatabaseId() int32`

GetDatabaseId returns the DatabaseId field if non-nil, zero value otherwise.

### GetDatabaseIdOk

`func (o *Backup) GetDatabaseIdOk() (*int32, bool)`

GetDatabaseIdOk returns a tuple with the DatabaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseId

`func (o *Backup) SetDatabaseId(v int32)`

SetDatabaseId sets DatabaseId field to given value.

### HasDatabaseId

`func (o *Backup) HasDatabaseId() bool`

HasDatabaseId returns a boolean if a field has been set.

### SetDatabaseIdNil

`func (o *Backup) SetDatabaseIdNil(b bool)`

 SetDatabaseIdNil sets the value for DatabaseId to be an explicit nil

### UnsetDatabaseId
`func (o *Backup) UnsetDatabaseId()`

UnsetDatabaseId ensures that no value is present for DatabaseId, not even an explicit nil
### GetDatabaseDbId

`func (o *Backup) GetDatabaseDbId() int32`

GetDatabaseDbId returns the DatabaseDbId field if non-nil, zero value otherwise.

### GetDatabaseDbIdOk

`func (o *Backup) GetDatabaseDbIdOk() (*int32, bool)`

GetDatabaseDbIdOk returns a tuple with the DatabaseDbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseDbId

`func (o *Backup) SetDatabaseDbId(v int32)`

SetDatabaseDbId sets DatabaseDbId field to given value.

### HasDatabaseDbId

`func (o *Backup) HasDatabaseDbId() bool`

HasDatabaseDbId returns a boolean if a field has been set.

### SetDatabaseDbIdNil

`func (o *Backup) SetDatabaseDbIdNil(b bool)`

 SetDatabaseDbIdNil sets the value for DatabaseDbId to be an explicit nil

### UnsetDatabaseDbId
`func (o *Backup) UnsetDatabaseDbId()`

UnsetDatabaseDbId ensures that no value is present for DatabaseDbId, not even an explicit nil
### GetCreatedAt

`func (o *Backup) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Backup) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Backup) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Backup) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Backup) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Backup) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


