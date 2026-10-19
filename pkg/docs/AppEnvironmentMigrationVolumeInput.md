# AppEnvironmentMigrationVolumeInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppServiceVolumeId** | **int32** |  | 
**StorageClassName** | Pointer to **NullableString** | Target cluster storage class for the volume. | [optional] 

## Methods

### NewAppEnvironmentMigrationVolumeInput

`func NewAppEnvironmentMigrationVolumeInput(appServiceVolumeId int32, ) *AppEnvironmentMigrationVolumeInput`

NewAppEnvironmentMigrationVolumeInput instantiates a new AppEnvironmentMigrationVolumeInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppEnvironmentMigrationVolumeInputWithDefaults

`func NewAppEnvironmentMigrationVolumeInputWithDefaults() *AppEnvironmentMigrationVolumeInput`

NewAppEnvironmentMigrationVolumeInputWithDefaults instantiates a new AppEnvironmentMigrationVolumeInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppServiceVolumeId

`func (o *AppEnvironmentMigrationVolumeInput) GetAppServiceVolumeId() int32`

GetAppServiceVolumeId returns the AppServiceVolumeId field if non-nil, zero value otherwise.

### GetAppServiceVolumeIdOk

`func (o *AppEnvironmentMigrationVolumeInput) GetAppServiceVolumeIdOk() (*int32, bool)`

GetAppServiceVolumeIdOk returns a tuple with the AppServiceVolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceVolumeId

`func (o *AppEnvironmentMigrationVolumeInput) SetAppServiceVolumeId(v int32)`

SetAppServiceVolumeId sets AppServiceVolumeId field to given value.


### GetStorageClassName

`func (o *AppEnvironmentMigrationVolumeInput) GetStorageClassName() string`

GetStorageClassName returns the StorageClassName field if non-nil, zero value otherwise.

### GetStorageClassNameOk

`func (o *AppEnvironmentMigrationVolumeInput) GetStorageClassNameOk() (*string, bool)`

GetStorageClassNameOk returns a tuple with the StorageClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClassName

`func (o *AppEnvironmentMigrationVolumeInput) SetStorageClassName(v string)`

SetStorageClassName sets StorageClassName field to given value.

### HasStorageClassName

`func (o *AppEnvironmentMigrationVolumeInput) HasStorageClassName() bool`

HasStorageClassName returns a boolean if a field has been set.

### SetStorageClassNameNil

`func (o *AppEnvironmentMigrationVolumeInput) SetStorageClassNameNil(b bool)`

 SetStorageClassNameNil sets the value for StorageClassName to be an explicit nil

### UnsetStorageClassName
`func (o *AppEnvironmentMigrationVolumeInput) UnsetStorageClassName()`

UnsetStorageClassName ensures that no value is present for StorageClassName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


