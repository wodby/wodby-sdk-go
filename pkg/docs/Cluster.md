# Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Title** | **string** |  | 
**Status** | **string** |  | 
**Serverless** | **bool** |  | 
**Demo** | **bool** |  | 
**Wodby** | **bool** |  | 
**K3s** | **bool** |  | 
**SingleNode** | **bool** |  | 
**Version** | Pointer to **NullableString** |  | [optional] 
**InfraVersion** | **string** |  | 
**MinNodeCount** | Pointer to **NullableInt32** |  | [optional] 
**MaxNodeCount** | Pointer to **NullableInt32** |  | [optional] 
**LastNodesReady** | Pointer to **NullableInt32** |  | [optional] 
**LastNodesTotal** | Pointer to **NullableInt32** |  | [optional] 
**Region** | Pointer to **NullableString** |  | [optional] 
**Zone** | Pointer to **NullableString** |  | [optional] 
**Ips** | Pointer to **[]string** |  | [optional] 
**Hostname** | Pointer to **NullableString** |  | [optional] 
**IntegrationId** | Pointer to **NullableInt32** |  | [optional] 
**EnvId** | **int32** | Legacy internal environment entity ID. Use envType. | 
**EnvType** | **string** |  | 
**EnvScope** | **string** |  | 
**AllowedEnvIds** | **[]int32** | Legacy internal environment entity IDs. Use allowedEnvTypes. | 
**AllowedEnvTypes** | **[]string** |  | 
**OrgId** | **int32** |  | 
**OwnershipScope** | **string** |  | 
**OwnerProjectId** | Pointer to **NullableInt32** |  | [optional] 
**Capabilities** | [**ClusterCapabilities**](ClusterCapabilities.md) |  | 
**Settings** | Pointer to [**ClusterSettings**](ClusterSettings.md) |  | [optional] 
**StorageClasses** | Pointer to [**[]StorageClass**](StorageClass.md) |  | [optional] 
**StorageClassesObservedAt** | Pointer to **NullableTime** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewCluster

`func NewCluster(id int32, name string, title string, status string, serverless bool, demo bool, wodby bool, k3s bool, singleNode bool, infraVersion string, envId int32, envType string, envScope string, allowedEnvIds []int32, allowedEnvTypes []string, orgId int32, ownershipScope string, capabilities ClusterCapabilities, createdAt time.Time, updatedAt time.Time, ) *Cluster`

NewCluster instantiates a new Cluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterWithDefaults

`func NewClusterWithDefaults() *Cluster`

NewClusterWithDefaults instantiates a new Cluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Cluster) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cluster) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cluster) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Cluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cluster) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *Cluster) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Cluster) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Cluster) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetStatus

`func (o *Cluster) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Cluster) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Cluster) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetServerless

`func (o *Cluster) GetServerless() bool`

GetServerless returns the Serverless field if non-nil, zero value otherwise.

### GetServerlessOk

`func (o *Cluster) GetServerlessOk() (*bool, bool)`

GetServerlessOk returns a tuple with the Serverless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerless

`func (o *Cluster) SetServerless(v bool)`

SetServerless sets Serverless field to given value.


### GetDemo

`func (o *Cluster) GetDemo() bool`

GetDemo returns the Demo field if non-nil, zero value otherwise.

### GetDemoOk

`func (o *Cluster) GetDemoOk() (*bool, bool)`

GetDemoOk returns a tuple with the Demo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemo

`func (o *Cluster) SetDemo(v bool)`

SetDemo sets Demo field to given value.


### GetWodby

`func (o *Cluster) GetWodby() bool`

GetWodby returns the Wodby field if non-nil, zero value otherwise.

### GetWodbyOk

`func (o *Cluster) GetWodbyOk() (*bool, bool)`

GetWodbyOk returns a tuple with the Wodby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWodby

`func (o *Cluster) SetWodby(v bool)`

SetWodby sets Wodby field to given value.


### GetK3s

`func (o *Cluster) GetK3s() bool`

GetK3s returns the K3s field if non-nil, zero value otherwise.

### GetK3sOk

`func (o *Cluster) GetK3sOk() (*bool, bool)`

GetK3sOk returns a tuple with the K3s field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK3s

`func (o *Cluster) SetK3s(v bool)`

SetK3s sets K3s field to given value.


### GetSingleNode

`func (o *Cluster) GetSingleNode() bool`

GetSingleNode returns the SingleNode field if non-nil, zero value otherwise.

### GetSingleNodeOk

`func (o *Cluster) GetSingleNodeOk() (*bool, bool)`

GetSingleNodeOk returns a tuple with the SingleNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleNode

`func (o *Cluster) SetSingleNode(v bool)`

SetSingleNode sets SingleNode field to given value.


### GetVersion

`func (o *Cluster) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Cluster) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Cluster) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Cluster) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *Cluster) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *Cluster) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetInfraVersion

`func (o *Cluster) GetInfraVersion() string`

GetInfraVersion returns the InfraVersion field if non-nil, zero value otherwise.

### GetInfraVersionOk

`func (o *Cluster) GetInfraVersionOk() (*string, bool)`

GetInfraVersionOk returns a tuple with the InfraVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraVersion

`func (o *Cluster) SetInfraVersion(v string)`

SetInfraVersion sets InfraVersion field to given value.


### GetMinNodeCount

`func (o *Cluster) GetMinNodeCount() int32`

GetMinNodeCount returns the MinNodeCount field if non-nil, zero value otherwise.

### GetMinNodeCountOk

`func (o *Cluster) GetMinNodeCountOk() (*int32, bool)`

GetMinNodeCountOk returns a tuple with the MinNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNodeCount

`func (o *Cluster) SetMinNodeCount(v int32)`

SetMinNodeCount sets MinNodeCount field to given value.

### HasMinNodeCount

`func (o *Cluster) HasMinNodeCount() bool`

HasMinNodeCount returns a boolean if a field has been set.

### SetMinNodeCountNil

`func (o *Cluster) SetMinNodeCountNil(b bool)`

 SetMinNodeCountNil sets the value for MinNodeCount to be an explicit nil

### UnsetMinNodeCount
`func (o *Cluster) UnsetMinNodeCount()`

UnsetMinNodeCount ensures that no value is present for MinNodeCount, not even an explicit nil
### GetMaxNodeCount

`func (o *Cluster) GetMaxNodeCount() int32`

GetMaxNodeCount returns the MaxNodeCount field if non-nil, zero value otherwise.

### GetMaxNodeCountOk

`func (o *Cluster) GetMaxNodeCountOk() (*int32, bool)`

GetMaxNodeCountOk returns a tuple with the MaxNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNodeCount

`func (o *Cluster) SetMaxNodeCount(v int32)`

SetMaxNodeCount sets MaxNodeCount field to given value.

### HasMaxNodeCount

`func (o *Cluster) HasMaxNodeCount() bool`

HasMaxNodeCount returns a boolean if a field has been set.

### SetMaxNodeCountNil

`func (o *Cluster) SetMaxNodeCountNil(b bool)`

 SetMaxNodeCountNil sets the value for MaxNodeCount to be an explicit nil

### UnsetMaxNodeCount
`func (o *Cluster) UnsetMaxNodeCount()`

UnsetMaxNodeCount ensures that no value is present for MaxNodeCount, not even an explicit nil
### GetLastNodesReady

`func (o *Cluster) GetLastNodesReady() int32`

GetLastNodesReady returns the LastNodesReady field if non-nil, zero value otherwise.

### GetLastNodesReadyOk

`func (o *Cluster) GetLastNodesReadyOk() (*int32, bool)`

GetLastNodesReadyOk returns a tuple with the LastNodesReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNodesReady

`func (o *Cluster) SetLastNodesReady(v int32)`

SetLastNodesReady sets LastNodesReady field to given value.

### HasLastNodesReady

`func (o *Cluster) HasLastNodesReady() bool`

HasLastNodesReady returns a boolean if a field has been set.

### SetLastNodesReadyNil

`func (o *Cluster) SetLastNodesReadyNil(b bool)`

 SetLastNodesReadyNil sets the value for LastNodesReady to be an explicit nil

### UnsetLastNodesReady
`func (o *Cluster) UnsetLastNodesReady()`

UnsetLastNodesReady ensures that no value is present for LastNodesReady, not even an explicit nil
### GetLastNodesTotal

`func (o *Cluster) GetLastNodesTotal() int32`

GetLastNodesTotal returns the LastNodesTotal field if non-nil, zero value otherwise.

### GetLastNodesTotalOk

`func (o *Cluster) GetLastNodesTotalOk() (*int32, bool)`

GetLastNodesTotalOk returns a tuple with the LastNodesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNodesTotal

`func (o *Cluster) SetLastNodesTotal(v int32)`

SetLastNodesTotal sets LastNodesTotal field to given value.

### HasLastNodesTotal

`func (o *Cluster) HasLastNodesTotal() bool`

HasLastNodesTotal returns a boolean if a field has been set.

### SetLastNodesTotalNil

`func (o *Cluster) SetLastNodesTotalNil(b bool)`

 SetLastNodesTotalNil sets the value for LastNodesTotal to be an explicit nil

### UnsetLastNodesTotal
`func (o *Cluster) UnsetLastNodesTotal()`

UnsetLastNodesTotal ensures that no value is present for LastNodesTotal, not even an explicit nil
### GetRegion

`func (o *Cluster) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Cluster) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Cluster) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Cluster) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *Cluster) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *Cluster) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetZone

`func (o *Cluster) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *Cluster) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *Cluster) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *Cluster) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *Cluster) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *Cluster) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil
### GetIps

`func (o *Cluster) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *Cluster) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *Cluster) SetIps(v []string)`

SetIps sets Ips field to given value.

### HasIps

`func (o *Cluster) HasIps() bool`

HasIps returns a boolean if a field has been set.

### SetIpsNil

`func (o *Cluster) SetIpsNil(b bool)`

 SetIpsNil sets the value for Ips to be an explicit nil

### UnsetIps
`func (o *Cluster) UnsetIps()`

UnsetIps ensures that no value is present for Ips, not even an explicit nil
### GetHostname

`func (o *Cluster) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *Cluster) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *Cluster) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *Cluster) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### SetHostnameNil

`func (o *Cluster) SetHostnameNil(b bool)`

 SetHostnameNil sets the value for Hostname to be an explicit nil

### UnsetHostname
`func (o *Cluster) UnsetHostname()`

UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
### GetIntegrationId

`func (o *Cluster) GetIntegrationId() int32`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *Cluster) GetIntegrationIdOk() (*int32, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *Cluster) SetIntegrationId(v int32)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *Cluster) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### SetIntegrationIdNil

`func (o *Cluster) SetIntegrationIdNil(b bool)`

 SetIntegrationIdNil sets the value for IntegrationId to be an explicit nil

### UnsetIntegrationId
`func (o *Cluster) UnsetIntegrationId()`

UnsetIntegrationId ensures that no value is present for IntegrationId, not even an explicit nil
### GetEnvId

`func (o *Cluster) GetEnvId() int32`

GetEnvId returns the EnvId field if non-nil, zero value otherwise.

### GetEnvIdOk

`func (o *Cluster) GetEnvIdOk() (*int32, bool)`

GetEnvIdOk returns a tuple with the EnvId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvId

`func (o *Cluster) SetEnvId(v int32)`

SetEnvId sets EnvId field to given value.


### GetEnvType

`func (o *Cluster) GetEnvType() string`

GetEnvType returns the EnvType field if non-nil, zero value otherwise.

### GetEnvTypeOk

`func (o *Cluster) GetEnvTypeOk() (*string, bool)`

GetEnvTypeOk returns a tuple with the EnvType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvType

`func (o *Cluster) SetEnvType(v string)`

SetEnvType sets EnvType field to given value.


### GetEnvScope

`func (o *Cluster) GetEnvScope() string`

GetEnvScope returns the EnvScope field if non-nil, zero value otherwise.

### GetEnvScopeOk

`func (o *Cluster) GetEnvScopeOk() (*string, bool)`

GetEnvScopeOk returns a tuple with the EnvScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvScope

`func (o *Cluster) SetEnvScope(v string)`

SetEnvScope sets EnvScope field to given value.


### GetAllowedEnvIds

`func (o *Cluster) GetAllowedEnvIds() []int32`

GetAllowedEnvIds returns the AllowedEnvIds field if non-nil, zero value otherwise.

### GetAllowedEnvIdsOk

`func (o *Cluster) GetAllowedEnvIdsOk() (*[]int32, bool)`

GetAllowedEnvIdsOk returns a tuple with the AllowedEnvIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedEnvIds

`func (o *Cluster) SetAllowedEnvIds(v []int32)`

SetAllowedEnvIds sets AllowedEnvIds field to given value.


### GetAllowedEnvTypes

`func (o *Cluster) GetAllowedEnvTypes() []string`

GetAllowedEnvTypes returns the AllowedEnvTypes field if non-nil, zero value otherwise.

### GetAllowedEnvTypesOk

`func (o *Cluster) GetAllowedEnvTypesOk() (*[]string, bool)`

GetAllowedEnvTypesOk returns a tuple with the AllowedEnvTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedEnvTypes

`func (o *Cluster) SetAllowedEnvTypes(v []string)`

SetAllowedEnvTypes sets AllowedEnvTypes field to given value.


### GetOrgId

`func (o *Cluster) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Cluster) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Cluster) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.


### GetOwnershipScope

`func (o *Cluster) GetOwnershipScope() string`

GetOwnershipScope returns the OwnershipScope field if non-nil, zero value otherwise.

### GetOwnershipScopeOk

`func (o *Cluster) GetOwnershipScopeOk() (*string, bool)`

GetOwnershipScopeOk returns a tuple with the OwnershipScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipScope

`func (o *Cluster) SetOwnershipScope(v string)`

SetOwnershipScope sets OwnershipScope field to given value.


### GetOwnerProjectId

`func (o *Cluster) GetOwnerProjectId() int32`

GetOwnerProjectId returns the OwnerProjectId field if non-nil, zero value otherwise.

### GetOwnerProjectIdOk

`func (o *Cluster) GetOwnerProjectIdOk() (*int32, bool)`

GetOwnerProjectIdOk returns a tuple with the OwnerProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerProjectId

`func (o *Cluster) SetOwnerProjectId(v int32)`

SetOwnerProjectId sets OwnerProjectId field to given value.

### HasOwnerProjectId

`func (o *Cluster) HasOwnerProjectId() bool`

HasOwnerProjectId returns a boolean if a field has been set.

### SetOwnerProjectIdNil

`func (o *Cluster) SetOwnerProjectIdNil(b bool)`

 SetOwnerProjectIdNil sets the value for OwnerProjectId to be an explicit nil

### UnsetOwnerProjectId
`func (o *Cluster) UnsetOwnerProjectId()`

UnsetOwnerProjectId ensures that no value is present for OwnerProjectId, not even an explicit nil
### GetCapabilities

`func (o *Cluster) GetCapabilities() ClusterCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *Cluster) GetCapabilitiesOk() (*ClusterCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *Cluster) SetCapabilities(v ClusterCapabilities)`

SetCapabilities sets Capabilities field to given value.


### GetSettings

`func (o *Cluster) GetSettings() ClusterSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *Cluster) GetSettingsOk() (*ClusterSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *Cluster) SetSettings(v ClusterSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *Cluster) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetStorageClasses

`func (o *Cluster) GetStorageClasses() []StorageClass`

GetStorageClasses returns the StorageClasses field if non-nil, zero value otherwise.

### GetStorageClassesOk

`func (o *Cluster) GetStorageClassesOk() (*[]StorageClass, bool)`

GetStorageClassesOk returns a tuple with the StorageClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClasses

`func (o *Cluster) SetStorageClasses(v []StorageClass)`

SetStorageClasses sets StorageClasses field to given value.

### HasStorageClasses

`func (o *Cluster) HasStorageClasses() bool`

HasStorageClasses returns a boolean if a field has been set.

### SetStorageClassesNil

`func (o *Cluster) SetStorageClassesNil(b bool)`

 SetStorageClassesNil sets the value for StorageClasses to be an explicit nil

### UnsetStorageClasses
`func (o *Cluster) UnsetStorageClasses()`

UnsetStorageClasses ensures that no value is present for StorageClasses, not even an explicit nil
### GetStorageClassesObservedAt

`func (o *Cluster) GetStorageClassesObservedAt() time.Time`

GetStorageClassesObservedAt returns the StorageClassesObservedAt field if non-nil, zero value otherwise.

### GetStorageClassesObservedAtOk

`func (o *Cluster) GetStorageClassesObservedAtOk() (*time.Time, bool)`

GetStorageClassesObservedAtOk returns a tuple with the StorageClassesObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClassesObservedAt

`func (o *Cluster) SetStorageClassesObservedAt(v time.Time)`

SetStorageClassesObservedAt sets StorageClassesObservedAt field to given value.

### HasStorageClassesObservedAt

`func (o *Cluster) HasStorageClassesObservedAt() bool`

HasStorageClassesObservedAt returns a boolean if a field has been set.

### SetStorageClassesObservedAtNil

`func (o *Cluster) SetStorageClassesObservedAtNil(b bool)`

 SetStorageClassesObservedAtNil sets the value for StorageClassesObservedAt to be an explicit nil

### UnsetStorageClassesObservedAt
`func (o *Cluster) UnsetStorageClassesObservedAt()`

UnsetStorageClassesObservedAt ensures that no value is present for StorageClassesObservedAt, not even an explicit nil
### GetCreatedAt

`func (o *Cluster) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Cluster) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Cluster) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Cluster) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Cluster) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Cluster) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


