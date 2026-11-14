# Database

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Title** | **string** |  | 
**Type** | **string** |  | 
**Kind** | **string** |  | 
**Status** | **string** |  | 
**Version** | **string** |  | 
**Region** | Pointer to **NullableString** |  | [optional] 
**Zone** | Pointer to **NullableString** |  | [optional] 
**Replicas** | Pointer to **int32** |  | [optional] 
**StorageSize** | Pointer to **NullableInt32** | Storage size in GB. Null when the database does not set one. | [optional] 
**StorageIOPS** | Pointer to **NullableInt32** | Provisioned storage IOPS. Null when the database does not set them. | [optional] 
**MachineType** | Pointer to **NullableString** |  | [optional] 
**HighAvailability** | Pointer to **bool** |  | [optional] 
**External** | Pointer to **bool** | True for a database that runs outside Wodby, such as a provider-managed database server. | [optional] 
**ExternalId** | Pointer to **NullableString** | The database&#39;s ID at its provider. | [optional] 
**Ssl** | Pointer to **bool** | True when connections to the database use TLS. | [optional] 
**PublicHost** | Pointer to **NullableString** | Host name for connections from outside the database&#39;s network. Null when the database has none. | [optional] 
**PublicPort** | Pointer to **NullableInt32** |  | [optional] 
**PrivateHost** | Pointer to **NullableString** | Host name for connections from inside the database&#39;s network. | [optional] 
**ClusterId** | Pointer to **NullableInt32** | The cluster that runs the database. | [optional] 
**ResidedClusterId** | Pointer to **NullableInt32** | The cluster whose network a provider-managed database is placed in. | [optional] 
**IntegrationId** | Pointer to **NullableInt32** |  | [optional] 
**AppServiceId** | Pointer to **NullableInt32** |  | [optional] 
**OwnershipScope** | **string** | org when the organization owns the resource, project when one project owns it. | 
**OwnerProjectId** | Pointer to **NullableInt32** |  | [optional] 
**EnvId** | **int32** | Legacy internal environment entity ID. Use envType. | 
**EnvType** | **string** |  | 
**OrgId** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewDatabase

`func NewDatabase(id int32, name string, title string, type_ string, kind string, status string, version string, ownershipScope string, envId int32, envType string, orgId int32, createdAt time.Time, updatedAt time.Time, ) *Database`

NewDatabase instantiates a new Database object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseWithDefaults

`func NewDatabaseWithDefaults() *Database`

NewDatabaseWithDefaults instantiates a new Database object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Database) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Database) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Database) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Database) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Database) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Database) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *Database) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Database) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Database) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *Database) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Database) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Database) SetType(v string)`

SetType sets Type field to given value.


### GetKind

`func (o *Database) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Database) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Database) SetKind(v string)`

SetKind sets Kind field to given value.


### GetStatus

`func (o *Database) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Database) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Database) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetVersion

`func (o *Database) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Database) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Database) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetRegion

`func (o *Database) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Database) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Database) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Database) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *Database) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *Database) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetZone

`func (o *Database) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *Database) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *Database) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *Database) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *Database) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *Database) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil
### GetReplicas

`func (o *Database) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *Database) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *Database) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *Database) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetStorageSize

`func (o *Database) GetStorageSize() int32`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *Database) GetStorageSizeOk() (*int32, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *Database) SetStorageSize(v int32)`

SetStorageSize sets StorageSize field to given value.

### HasStorageSize

`func (o *Database) HasStorageSize() bool`

HasStorageSize returns a boolean if a field has been set.

### SetStorageSizeNil

`func (o *Database) SetStorageSizeNil(b bool)`

 SetStorageSizeNil sets the value for StorageSize to be an explicit nil

### UnsetStorageSize
`func (o *Database) UnsetStorageSize()`

UnsetStorageSize ensures that no value is present for StorageSize, not even an explicit nil
### GetStorageIOPS

`func (o *Database) GetStorageIOPS() int32`

GetStorageIOPS returns the StorageIOPS field if non-nil, zero value otherwise.

### GetStorageIOPSOk

`func (o *Database) GetStorageIOPSOk() (*int32, bool)`

GetStorageIOPSOk returns a tuple with the StorageIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageIOPS

`func (o *Database) SetStorageIOPS(v int32)`

SetStorageIOPS sets StorageIOPS field to given value.

### HasStorageIOPS

`func (o *Database) HasStorageIOPS() bool`

HasStorageIOPS returns a boolean if a field has been set.

### SetStorageIOPSNil

`func (o *Database) SetStorageIOPSNil(b bool)`

 SetStorageIOPSNil sets the value for StorageIOPS to be an explicit nil

### UnsetStorageIOPS
`func (o *Database) UnsetStorageIOPS()`

UnsetStorageIOPS ensures that no value is present for StorageIOPS, not even an explicit nil
### GetMachineType

`func (o *Database) GetMachineType() string`

GetMachineType returns the MachineType field if non-nil, zero value otherwise.

### GetMachineTypeOk

`func (o *Database) GetMachineTypeOk() (*string, bool)`

GetMachineTypeOk returns a tuple with the MachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineType

`func (o *Database) SetMachineType(v string)`

SetMachineType sets MachineType field to given value.

### HasMachineType

`func (o *Database) HasMachineType() bool`

HasMachineType returns a boolean if a field has been set.

### SetMachineTypeNil

`func (o *Database) SetMachineTypeNil(b bool)`

 SetMachineTypeNil sets the value for MachineType to be an explicit nil

### UnsetMachineType
`func (o *Database) UnsetMachineType()`

UnsetMachineType ensures that no value is present for MachineType, not even an explicit nil
### GetHighAvailability

`func (o *Database) GetHighAvailability() bool`

GetHighAvailability returns the HighAvailability field if non-nil, zero value otherwise.

### GetHighAvailabilityOk

`func (o *Database) GetHighAvailabilityOk() (*bool, bool)`

GetHighAvailabilityOk returns a tuple with the HighAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighAvailability

`func (o *Database) SetHighAvailability(v bool)`

SetHighAvailability sets HighAvailability field to given value.

### HasHighAvailability

`func (o *Database) HasHighAvailability() bool`

HasHighAvailability returns a boolean if a field has been set.

### GetExternal

`func (o *Database) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *Database) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *Database) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *Database) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetExternalId

`func (o *Database) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Database) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Database) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Database) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *Database) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *Database) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetSsl

`func (o *Database) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *Database) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *Database) SetSsl(v bool)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *Database) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetPublicHost

`func (o *Database) GetPublicHost() string`

GetPublicHost returns the PublicHost field if non-nil, zero value otherwise.

### GetPublicHostOk

`func (o *Database) GetPublicHostOk() (*string, bool)`

GetPublicHostOk returns a tuple with the PublicHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicHost

`func (o *Database) SetPublicHost(v string)`

SetPublicHost sets PublicHost field to given value.

### HasPublicHost

`func (o *Database) HasPublicHost() bool`

HasPublicHost returns a boolean if a field has been set.

### SetPublicHostNil

`func (o *Database) SetPublicHostNil(b bool)`

 SetPublicHostNil sets the value for PublicHost to be an explicit nil

### UnsetPublicHost
`func (o *Database) UnsetPublicHost()`

UnsetPublicHost ensures that no value is present for PublicHost, not even an explicit nil
### GetPublicPort

`func (o *Database) GetPublicPort() int32`

GetPublicPort returns the PublicPort field if non-nil, zero value otherwise.

### GetPublicPortOk

`func (o *Database) GetPublicPortOk() (*int32, bool)`

GetPublicPortOk returns a tuple with the PublicPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicPort

`func (o *Database) SetPublicPort(v int32)`

SetPublicPort sets PublicPort field to given value.

### HasPublicPort

`func (o *Database) HasPublicPort() bool`

HasPublicPort returns a boolean if a field has been set.

### SetPublicPortNil

`func (o *Database) SetPublicPortNil(b bool)`

 SetPublicPortNil sets the value for PublicPort to be an explicit nil

### UnsetPublicPort
`func (o *Database) UnsetPublicPort()`

UnsetPublicPort ensures that no value is present for PublicPort, not even an explicit nil
### GetPrivateHost

`func (o *Database) GetPrivateHost() string`

GetPrivateHost returns the PrivateHost field if non-nil, zero value otherwise.

### GetPrivateHostOk

`func (o *Database) GetPrivateHostOk() (*string, bool)`

GetPrivateHostOk returns a tuple with the PrivateHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateHost

`func (o *Database) SetPrivateHost(v string)`

SetPrivateHost sets PrivateHost field to given value.

### HasPrivateHost

`func (o *Database) HasPrivateHost() bool`

HasPrivateHost returns a boolean if a field has been set.

### SetPrivateHostNil

`func (o *Database) SetPrivateHostNil(b bool)`

 SetPrivateHostNil sets the value for PrivateHost to be an explicit nil

### UnsetPrivateHost
`func (o *Database) UnsetPrivateHost()`

UnsetPrivateHost ensures that no value is present for PrivateHost, not even an explicit nil
### GetClusterId

`func (o *Database) GetClusterId() int32`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *Database) GetClusterIdOk() (*int32, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *Database) SetClusterId(v int32)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *Database) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *Database) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *Database) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetResidedClusterId

`func (o *Database) GetResidedClusterId() int32`

GetResidedClusterId returns the ResidedClusterId field if non-nil, zero value otherwise.

### GetResidedClusterIdOk

`func (o *Database) GetResidedClusterIdOk() (*int32, bool)`

GetResidedClusterIdOk returns a tuple with the ResidedClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidedClusterId

`func (o *Database) SetResidedClusterId(v int32)`

SetResidedClusterId sets ResidedClusterId field to given value.

### HasResidedClusterId

`func (o *Database) HasResidedClusterId() bool`

HasResidedClusterId returns a boolean if a field has been set.

### SetResidedClusterIdNil

`func (o *Database) SetResidedClusterIdNil(b bool)`

 SetResidedClusterIdNil sets the value for ResidedClusterId to be an explicit nil

### UnsetResidedClusterId
`func (o *Database) UnsetResidedClusterId()`

UnsetResidedClusterId ensures that no value is present for ResidedClusterId, not even an explicit nil
### GetIntegrationId

`func (o *Database) GetIntegrationId() int32`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *Database) GetIntegrationIdOk() (*int32, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *Database) SetIntegrationId(v int32)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *Database) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### SetIntegrationIdNil

`func (o *Database) SetIntegrationIdNil(b bool)`

 SetIntegrationIdNil sets the value for IntegrationId to be an explicit nil

### UnsetIntegrationId
`func (o *Database) UnsetIntegrationId()`

UnsetIntegrationId ensures that no value is present for IntegrationId, not even an explicit nil
### GetAppServiceId

`func (o *Database) GetAppServiceId() int32`

GetAppServiceId returns the AppServiceId field if non-nil, zero value otherwise.

### GetAppServiceIdOk

`func (o *Database) GetAppServiceIdOk() (*int32, bool)`

GetAppServiceIdOk returns a tuple with the AppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceId

`func (o *Database) SetAppServiceId(v int32)`

SetAppServiceId sets AppServiceId field to given value.

### HasAppServiceId

`func (o *Database) HasAppServiceId() bool`

HasAppServiceId returns a boolean if a field has been set.

### SetAppServiceIdNil

`func (o *Database) SetAppServiceIdNil(b bool)`

 SetAppServiceIdNil sets the value for AppServiceId to be an explicit nil

### UnsetAppServiceId
`func (o *Database) UnsetAppServiceId()`

UnsetAppServiceId ensures that no value is present for AppServiceId, not even an explicit nil
### GetOwnershipScope

`func (o *Database) GetOwnershipScope() string`

GetOwnershipScope returns the OwnershipScope field if non-nil, zero value otherwise.

### GetOwnershipScopeOk

`func (o *Database) GetOwnershipScopeOk() (*string, bool)`

GetOwnershipScopeOk returns a tuple with the OwnershipScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipScope

`func (o *Database) SetOwnershipScope(v string)`

SetOwnershipScope sets OwnershipScope field to given value.


### GetOwnerProjectId

`func (o *Database) GetOwnerProjectId() int32`

GetOwnerProjectId returns the OwnerProjectId field if non-nil, zero value otherwise.

### GetOwnerProjectIdOk

`func (o *Database) GetOwnerProjectIdOk() (*int32, bool)`

GetOwnerProjectIdOk returns a tuple with the OwnerProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerProjectId

`func (o *Database) SetOwnerProjectId(v int32)`

SetOwnerProjectId sets OwnerProjectId field to given value.

### HasOwnerProjectId

`func (o *Database) HasOwnerProjectId() bool`

HasOwnerProjectId returns a boolean if a field has been set.

### SetOwnerProjectIdNil

`func (o *Database) SetOwnerProjectIdNil(b bool)`

 SetOwnerProjectIdNil sets the value for OwnerProjectId to be an explicit nil

### UnsetOwnerProjectId
`func (o *Database) UnsetOwnerProjectId()`

UnsetOwnerProjectId ensures that no value is present for OwnerProjectId, not even an explicit nil
### GetEnvId

`func (o *Database) GetEnvId() int32`

GetEnvId returns the EnvId field if non-nil, zero value otherwise.

### GetEnvIdOk

`func (o *Database) GetEnvIdOk() (*int32, bool)`

GetEnvIdOk returns a tuple with the EnvId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvId

`func (o *Database) SetEnvId(v int32)`

SetEnvId sets EnvId field to given value.


### GetEnvType

`func (o *Database) GetEnvType() string`

GetEnvType returns the EnvType field if non-nil, zero value otherwise.

### GetEnvTypeOk

`func (o *Database) GetEnvTypeOk() (*string, bool)`

GetEnvTypeOk returns a tuple with the EnvType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvType

`func (o *Database) SetEnvType(v string)`

SetEnvType sets EnvType field to given value.


### GetOrgId

`func (o *Database) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Database) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Database) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.


### GetCreatedAt

`func (o *Database) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Database) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Database) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Database) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Database) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Database) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


