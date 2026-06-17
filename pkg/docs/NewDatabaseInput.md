# NewDatabaseInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgID** | **int32** |  | 
**ProjectID** | Pointer to **NullableInt32** |  | [optional] 
**EnvID** | **int32** |  | 
**Name** | **string** |  | 
**Title** | **string** |  | 
**IntegrationKindID** | **int32** |  | 
**Type** | **string** |  | 
**Version** | **string** |  | 
**MachineType** | **string** |  | 
**StorageSize** | Pointer to **NullableInt32** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**StorageAutoscaling** | Pointer to **NullableBool** |  | [optional] 
**HighAvailability** | Pointer to **NullableBool** |  | [optional] 
**Region** | Pointer to **NullableString** |  | [optional] 
**Zone** | Pointer to **NullableString** |  | [optional] 
**ResidedClusterID** | Pointer to **NullableInt32** |  | [optional] 
**Iops** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewNewDatabaseInput

`func NewNewDatabaseInput(orgID int32, envID int32, name string, title string, integrationKindID int32, type_ string, version string, machineType string, ) *NewDatabaseInput`

NewNewDatabaseInput instantiates a new NewDatabaseInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewDatabaseInputWithDefaults

`func NewNewDatabaseInputWithDefaults() *NewDatabaseInput`

NewNewDatabaseInputWithDefaults instantiates a new NewDatabaseInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgID

`func (o *NewDatabaseInput) GetOrgID() int32`

GetOrgID returns the OrgID field if non-nil, zero value otherwise.

### GetOrgIDOk

`func (o *NewDatabaseInput) GetOrgIDOk() (*int32, bool)`

GetOrgIDOk returns a tuple with the OrgID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgID

`func (o *NewDatabaseInput) SetOrgID(v int32)`

SetOrgID sets OrgID field to given value.


### GetProjectID

`func (o *NewDatabaseInput) GetProjectID() int32`

GetProjectID returns the ProjectID field if non-nil, zero value otherwise.

### GetProjectIDOk

`func (o *NewDatabaseInput) GetProjectIDOk() (*int32, bool)`

GetProjectIDOk returns a tuple with the ProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectID

`func (o *NewDatabaseInput) SetProjectID(v int32)`

SetProjectID sets ProjectID field to given value.

### HasProjectID

`func (o *NewDatabaseInput) HasProjectID() bool`

HasProjectID returns a boolean if a field has been set.

### SetProjectIDNil

`func (o *NewDatabaseInput) SetProjectIDNil(b bool)`

 SetProjectIDNil sets the value for ProjectID to be an explicit nil

### UnsetProjectID
`func (o *NewDatabaseInput) UnsetProjectID()`

UnsetProjectID ensures that no value is present for ProjectID, not even an explicit nil
### GetEnvID

`func (o *NewDatabaseInput) GetEnvID() int32`

GetEnvID returns the EnvID field if non-nil, zero value otherwise.

### GetEnvIDOk

`func (o *NewDatabaseInput) GetEnvIDOk() (*int32, bool)`

GetEnvIDOk returns a tuple with the EnvID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvID

`func (o *NewDatabaseInput) SetEnvID(v int32)`

SetEnvID sets EnvID field to given value.


### GetName

`func (o *NewDatabaseInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewDatabaseInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewDatabaseInput) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *NewDatabaseInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewDatabaseInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NewDatabaseInput) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetIntegrationKindID

`func (o *NewDatabaseInput) GetIntegrationKindID() int32`

GetIntegrationKindID returns the IntegrationKindID field if non-nil, zero value otherwise.

### GetIntegrationKindIDOk

`func (o *NewDatabaseInput) GetIntegrationKindIDOk() (*int32, bool)`

GetIntegrationKindIDOk returns a tuple with the IntegrationKindID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationKindID

`func (o *NewDatabaseInput) SetIntegrationKindID(v int32)`

SetIntegrationKindID sets IntegrationKindID field to given value.


### GetType

`func (o *NewDatabaseInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NewDatabaseInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NewDatabaseInput) SetType(v string)`

SetType sets Type field to given value.


### GetVersion

`func (o *NewDatabaseInput) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NewDatabaseInput) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NewDatabaseInput) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetMachineType

`func (o *NewDatabaseInput) GetMachineType() string`

GetMachineType returns the MachineType field if non-nil, zero value otherwise.

### GetMachineTypeOk

`func (o *NewDatabaseInput) GetMachineTypeOk() (*string, bool)`

GetMachineTypeOk returns a tuple with the MachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineType

`func (o *NewDatabaseInput) SetMachineType(v string)`

SetMachineType sets MachineType field to given value.


### GetStorageSize

`func (o *NewDatabaseInput) GetStorageSize() int32`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *NewDatabaseInput) GetStorageSizeOk() (*int32, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *NewDatabaseInput) SetStorageSize(v int32)`

SetStorageSize sets StorageSize field to given value.

### HasStorageSize

`func (o *NewDatabaseInput) HasStorageSize() bool`

HasStorageSize returns a boolean if a field has been set.

### SetStorageSizeNil

`func (o *NewDatabaseInput) SetStorageSizeNil(b bool)`

 SetStorageSizeNil sets the value for StorageSize to be an explicit nil

### UnsetStorageSize
`func (o *NewDatabaseInput) UnsetStorageSize()`

UnsetStorageSize ensures that no value is present for StorageSize, not even an explicit nil
### GetPassword

`func (o *NewDatabaseInput) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *NewDatabaseInput) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *NewDatabaseInput) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *NewDatabaseInput) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *NewDatabaseInput) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *NewDatabaseInput) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetStorageAutoscaling

`func (o *NewDatabaseInput) GetStorageAutoscaling() bool`

GetStorageAutoscaling returns the StorageAutoscaling field if non-nil, zero value otherwise.

### GetStorageAutoscalingOk

`func (o *NewDatabaseInput) GetStorageAutoscalingOk() (*bool, bool)`

GetStorageAutoscalingOk returns a tuple with the StorageAutoscaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAutoscaling

`func (o *NewDatabaseInput) SetStorageAutoscaling(v bool)`

SetStorageAutoscaling sets StorageAutoscaling field to given value.

### HasStorageAutoscaling

`func (o *NewDatabaseInput) HasStorageAutoscaling() bool`

HasStorageAutoscaling returns a boolean if a field has been set.

### SetStorageAutoscalingNil

`func (o *NewDatabaseInput) SetStorageAutoscalingNil(b bool)`

 SetStorageAutoscalingNil sets the value for StorageAutoscaling to be an explicit nil

### UnsetStorageAutoscaling
`func (o *NewDatabaseInput) UnsetStorageAutoscaling()`

UnsetStorageAutoscaling ensures that no value is present for StorageAutoscaling, not even an explicit nil
### GetHighAvailability

`func (o *NewDatabaseInput) GetHighAvailability() bool`

GetHighAvailability returns the HighAvailability field if non-nil, zero value otherwise.

### GetHighAvailabilityOk

`func (o *NewDatabaseInput) GetHighAvailabilityOk() (*bool, bool)`

GetHighAvailabilityOk returns a tuple with the HighAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighAvailability

`func (o *NewDatabaseInput) SetHighAvailability(v bool)`

SetHighAvailability sets HighAvailability field to given value.

### HasHighAvailability

`func (o *NewDatabaseInput) HasHighAvailability() bool`

HasHighAvailability returns a boolean if a field has been set.

### SetHighAvailabilityNil

`func (o *NewDatabaseInput) SetHighAvailabilityNil(b bool)`

 SetHighAvailabilityNil sets the value for HighAvailability to be an explicit nil

### UnsetHighAvailability
`func (o *NewDatabaseInput) UnsetHighAvailability()`

UnsetHighAvailability ensures that no value is present for HighAvailability, not even an explicit nil
### GetRegion

`func (o *NewDatabaseInput) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NewDatabaseInput) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NewDatabaseInput) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *NewDatabaseInput) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *NewDatabaseInput) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *NewDatabaseInput) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetZone

`func (o *NewDatabaseInput) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *NewDatabaseInput) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *NewDatabaseInput) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *NewDatabaseInput) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *NewDatabaseInput) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *NewDatabaseInput) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil
### GetResidedClusterID

`func (o *NewDatabaseInput) GetResidedClusterID() int32`

GetResidedClusterID returns the ResidedClusterID field if non-nil, zero value otherwise.

### GetResidedClusterIDOk

`func (o *NewDatabaseInput) GetResidedClusterIDOk() (*int32, bool)`

GetResidedClusterIDOk returns a tuple with the ResidedClusterID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidedClusterID

`func (o *NewDatabaseInput) SetResidedClusterID(v int32)`

SetResidedClusterID sets ResidedClusterID field to given value.

### HasResidedClusterID

`func (o *NewDatabaseInput) HasResidedClusterID() bool`

HasResidedClusterID returns a boolean if a field has been set.

### SetResidedClusterIDNil

`func (o *NewDatabaseInput) SetResidedClusterIDNil(b bool)`

 SetResidedClusterIDNil sets the value for ResidedClusterID to be an explicit nil

### UnsetResidedClusterID
`func (o *NewDatabaseInput) UnsetResidedClusterID()`

UnsetResidedClusterID ensures that no value is present for ResidedClusterID, not even an explicit nil
### GetIops

`func (o *NewDatabaseInput) GetIops() int32`

GetIops returns the Iops field if non-nil, zero value otherwise.

### GetIopsOk

`func (o *NewDatabaseInput) GetIopsOk() (*int32, bool)`

GetIopsOk returns a tuple with the Iops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIops

`func (o *NewDatabaseInput) SetIops(v int32)`

SetIops sets Iops field to given value.

### HasIops

`func (o *NewDatabaseInput) HasIops() bool`

HasIops returns a boolean if a field has been set.

### SetIopsNil

`func (o *NewDatabaseInput) SetIopsNil(b bool)`

 SetIopsNil sets the value for Iops to be an explicit nil

### UnsetIops
`func (o *NewDatabaseInput) UnsetIops()`

UnsetIops ensures that no value is present for Iops, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


