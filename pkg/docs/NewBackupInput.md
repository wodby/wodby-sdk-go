# NewBackupInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppServiceID** | Pointer to **NullableInt32** |  | [optional] 
**DatabaseDBID** | Pointer to **NullableInt32** |  | [optional] 
**BackupName** | Pointer to **NullableString** |  | [optional] 
**IntegrationID** | **int32** |  | 
**Bucket** | **string** |  | 
**StorageClass** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNewBackupInput

`func NewNewBackupInput(integrationID int32, bucket string, ) *NewBackupInput`

NewNewBackupInput instantiates a new NewBackupInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewBackupInputWithDefaults

`func NewNewBackupInputWithDefaults() *NewBackupInput`

NewNewBackupInputWithDefaults instantiates a new NewBackupInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppServiceID

`func (o *NewBackupInput) GetAppServiceID() int32`

GetAppServiceID returns the AppServiceID field if non-nil, zero value otherwise.

### GetAppServiceIDOk

`func (o *NewBackupInput) GetAppServiceIDOk() (*int32, bool)`

GetAppServiceIDOk returns a tuple with the AppServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceID

`func (o *NewBackupInput) SetAppServiceID(v int32)`

SetAppServiceID sets AppServiceID field to given value.

### HasAppServiceID

`func (o *NewBackupInput) HasAppServiceID() bool`

HasAppServiceID returns a boolean if a field has been set.

### SetAppServiceIDNil

`func (o *NewBackupInput) SetAppServiceIDNil(b bool)`

 SetAppServiceIDNil sets the value for AppServiceID to be an explicit nil

### UnsetAppServiceID
`func (o *NewBackupInput) UnsetAppServiceID()`

UnsetAppServiceID ensures that no value is present for AppServiceID, not even an explicit nil
### GetDatabaseDBID

`func (o *NewBackupInput) GetDatabaseDBID() int32`

GetDatabaseDBID returns the DatabaseDBID field if non-nil, zero value otherwise.

### GetDatabaseDBIDOk

`func (o *NewBackupInput) GetDatabaseDBIDOk() (*int32, bool)`

GetDatabaseDBIDOk returns a tuple with the DatabaseDBID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseDBID

`func (o *NewBackupInput) SetDatabaseDBID(v int32)`

SetDatabaseDBID sets DatabaseDBID field to given value.

### HasDatabaseDBID

`func (o *NewBackupInput) HasDatabaseDBID() bool`

HasDatabaseDBID returns a boolean if a field has been set.

### SetDatabaseDBIDNil

`func (o *NewBackupInput) SetDatabaseDBIDNil(b bool)`

 SetDatabaseDBIDNil sets the value for DatabaseDBID to be an explicit nil

### UnsetDatabaseDBID
`func (o *NewBackupInput) UnsetDatabaseDBID()`

UnsetDatabaseDBID ensures that no value is present for DatabaseDBID, not even an explicit nil
### GetBackupName

`func (o *NewBackupInput) GetBackupName() string`

GetBackupName returns the BackupName field if non-nil, zero value otherwise.

### GetBackupNameOk

`func (o *NewBackupInput) GetBackupNameOk() (*string, bool)`

GetBackupNameOk returns a tuple with the BackupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupName

`func (o *NewBackupInput) SetBackupName(v string)`

SetBackupName sets BackupName field to given value.

### HasBackupName

`func (o *NewBackupInput) HasBackupName() bool`

HasBackupName returns a boolean if a field has been set.

### SetBackupNameNil

`func (o *NewBackupInput) SetBackupNameNil(b bool)`

 SetBackupNameNil sets the value for BackupName to be an explicit nil

### UnsetBackupName
`func (o *NewBackupInput) UnsetBackupName()`

UnsetBackupName ensures that no value is present for BackupName, not even an explicit nil
### GetIntegrationID

`func (o *NewBackupInput) GetIntegrationID() int32`

GetIntegrationID returns the IntegrationID field if non-nil, zero value otherwise.

### GetIntegrationIDOk

`func (o *NewBackupInput) GetIntegrationIDOk() (*int32, bool)`

GetIntegrationIDOk returns a tuple with the IntegrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationID

`func (o *NewBackupInput) SetIntegrationID(v int32)`

SetIntegrationID sets IntegrationID field to given value.


### GetBucket

`func (o *NewBackupInput) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *NewBackupInput) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *NewBackupInput) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetStorageClass

`func (o *NewBackupInput) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *NewBackupInput) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *NewBackupInput) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *NewBackupInput) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.

### SetStorageClassNil

`func (o *NewBackupInput) SetStorageClassNil(b bool)`

 SetStorageClassNil sets the value for StorageClass to be an explicit nil

### UnsetStorageClass
`func (o *NewBackupInput) UnsetStorageClass()`

UnsetStorageClass ensures that no value is present for StorageClass, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


