# AppEnvironmentMigrationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetClusterId** | **int32** | Cluster of the same organization that the environment moves to. | 
**Volumes** | Pointer to [**[]AppEnvironmentMigrationVolumeInput**](AppEnvironmentMigrationVolumeInput.md) | Target cluster storage classes for the environment&#39;s volumes. A volume left out uses its current storage class, which must exist on the target cluster. | [optional] 
**ExistingBackupIds** | Pointer to **[]int32** | Completed backups of this environment to restore instead of creating final backups. Data written after a backup was taken does not reach the target. Components without a backup listed here get a final backup. | [optional] 

## Methods

### NewAppEnvironmentMigrationInput

`func NewAppEnvironmentMigrationInput(targetClusterId int32, ) *AppEnvironmentMigrationInput`

NewAppEnvironmentMigrationInput instantiates a new AppEnvironmentMigrationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppEnvironmentMigrationInputWithDefaults

`func NewAppEnvironmentMigrationInputWithDefaults() *AppEnvironmentMigrationInput`

NewAppEnvironmentMigrationInputWithDefaults instantiates a new AppEnvironmentMigrationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetClusterId

`func (o *AppEnvironmentMigrationInput) GetTargetClusterId() int32`

GetTargetClusterId returns the TargetClusterId field if non-nil, zero value otherwise.

### GetTargetClusterIdOk

`func (o *AppEnvironmentMigrationInput) GetTargetClusterIdOk() (*int32, bool)`

GetTargetClusterIdOk returns a tuple with the TargetClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetClusterId

`func (o *AppEnvironmentMigrationInput) SetTargetClusterId(v int32)`

SetTargetClusterId sets TargetClusterId field to given value.


### GetVolumes

`func (o *AppEnvironmentMigrationInput) GetVolumes() []AppEnvironmentMigrationVolumeInput`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *AppEnvironmentMigrationInput) GetVolumesOk() (*[]AppEnvironmentMigrationVolumeInput, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *AppEnvironmentMigrationInput) SetVolumes(v []AppEnvironmentMigrationVolumeInput)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *AppEnvironmentMigrationInput) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetExistingBackupIds

`func (o *AppEnvironmentMigrationInput) GetExistingBackupIds() []int32`

GetExistingBackupIds returns the ExistingBackupIds field if non-nil, zero value otherwise.

### GetExistingBackupIdsOk

`func (o *AppEnvironmentMigrationInput) GetExistingBackupIdsOk() (*[]int32, bool)`

GetExistingBackupIdsOk returns a tuple with the ExistingBackupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingBackupIds

`func (o *AppEnvironmentMigrationInput) SetExistingBackupIds(v []int32)`

SetExistingBackupIds sets ExistingBackupIds field to given value.

### HasExistingBackupIds

`func (o *AppEnvironmentMigrationInput) HasExistingBackupIds() bool`

HasExistingBackupIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


