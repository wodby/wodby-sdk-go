# NewClusterInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. | [optional] 
**ProjectId** | Pointer to **NullableInt32** |  | [optional] 
**IntegrationId** | **int32** |  | 
**Name** | **string** |  | 
**Title** | **string** |  | 
**Serverless** | **bool** |  | 
**SingleNode** | Pointer to **NullableBool** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**MachineType** | Pointer to **NullableString** |  | [optional] 
**MinNodeCount** | Pointer to **NullableInt32** |  | [optional] 
**MaxNodeCount** | Pointer to **NullableInt32** |  | [optional] 
**NodeDiskSize** | Pointer to **NullableInt32** |  | [optional] 
**Zone** | Pointer to **NullableString** |  | [optional] 
**Region** | Pointer to **NullableString** |  | [optional] 
**BillingOption** | Pointer to **NullableString** |  | [optional] 
**DisableMonitoring** | **bool** |  | 
**AutoInfrastructureUpgrade** | Pointer to **NullableBool** |  | [optional] 
**EnvironmentPolicy** | Pointer to [**ClusterEnvironmentPolicyInput**](ClusterEnvironmentPolicyInput.md) |  | [optional] 

## Methods

### NewNewClusterInput

`func NewNewClusterInput(integrationId int32, name string, title string, serverless bool, disableMonitoring bool, ) *NewClusterInput`

NewNewClusterInput instantiates a new NewClusterInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewClusterInputWithDefaults

`func NewNewClusterInputWithDefaults() *NewClusterInput`

NewNewClusterInputWithDefaults instantiates a new NewClusterInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *NewClusterInput) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *NewClusterInput) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *NewClusterInput) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *NewClusterInput) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProjectId

`func (o *NewClusterInput) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *NewClusterInput) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *NewClusterInput) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *NewClusterInput) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *NewClusterInput) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *NewClusterInput) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetIntegrationId

`func (o *NewClusterInput) GetIntegrationId() int32`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *NewClusterInput) GetIntegrationIdOk() (*int32, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *NewClusterInput) SetIntegrationId(v int32)`

SetIntegrationId sets IntegrationId field to given value.


### GetName

`func (o *NewClusterInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewClusterInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewClusterInput) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *NewClusterInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewClusterInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NewClusterInput) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetServerless

`func (o *NewClusterInput) GetServerless() bool`

GetServerless returns the Serverless field if non-nil, zero value otherwise.

### GetServerlessOk

`func (o *NewClusterInput) GetServerlessOk() (*bool, bool)`

GetServerlessOk returns a tuple with the Serverless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerless

`func (o *NewClusterInput) SetServerless(v bool)`

SetServerless sets Serverless field to given value.


### GetSingleNode

`func (o *NewClusterInput) GetSingleNode() bool`

GetSingleNode returns the SingleNode field if non-nil, zero value otherwise.

### GetSingleNodeOk

`func (o *NewClusterInput) GetSingleNodeOk() (*bool, bool)`

GetSingleNodeOk returns a tuple with the SingleNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleNode

`func (o *NewClusterInput) SetSingleNode(v bool)`

SetSingleNode sets SingleNode field to given value.

### HasSingleNode

`func (o *NewClusterInput) HasSingleNode() bool`

HasSingleNode returns a boolean if a field has been set.

### SetSingleNodeNil

`func (o *NewClusterInput) SetSingleNodeNil(b bool)`

 SetSingleNodeNil sets the value for SingleNode to be an explicit nil

### UnsetSingleNode
`func (o *NewClusterInput) UnsetSingleNode()`

UnsetSingleNode ensures that no value is present for SingleNode, not even an explicit nil
### GetVersion

`func (o *NewClusterInput) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NewClusterInput) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NewClusterInput) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NewClusterInput) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *NewClusterInput) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *NewClusterInput) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetMachineType

`func (o *NewClusterInput) GetMachineType() string`

GetMachineType returns the MachineType field if non-nil, zero value otherwise.

### GetMachineTypeOk

`func (o *NewClusterInput) GetMachineTypeOk() (*string, bool)`

GetMachineTypeOk returns a tuple with the MachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineType

`func (o *NewClusterInput) SetMachineType(v string)`

SetMachineType sets MachineType field to given value.

### HasMachineType

`func (o *NewClusterInput) HasMachineType() bool`

HasMachineType returns a boolean if a field has been set.

### SetMachineTypeNil

`func (o *NewClusterInput) SetMachineTypeNil(b bool)`

 SetMachineTypeNil sets the value for MachineType to be an explicit nil

### UnsetMachineType
`func (o *NewClusterInput) UnsetMachineType()`

UnsetMachineType ensures that no value is present for MachineType, not even an explicit nil
### GetMinNodeCount

`func (o *NewClusterInput) GetMinNodeCount() int32`

GetMinNodeCount returns the MinNodeCount field if non-nil, zero value otherwise.

### GetMinNodeCountOk

`func (o *NewClusterInput) GetMinNodeCountOk() (*int32, bool)`

GetMinNodeCountOk returns a tuple with the MinNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNodeCount

`func (o *NewClusterInput) SetMinNodeCount(v int32)`

SetMinNodeCount sets MinNodeCount field to given value.

### HasMinNodeCount

`func (o *NewClusterInput) HasMinNodeCount() bool`

HasMinNodeCount returns a boolean if a field has been set.

### SetMinNodeCountNil

`func (o *NewClusterInput) SetMinNodeCountNil(b bool)`

 SetMinNodeCountNil sets the value for MinNodeCount to be an explicit nil

### UnsetMinNodeCount
`func (o *NewClusterInput) UnsetMinNodeCount()`

UnsetMinNodeCount ensures that no value is present for MinNodeCount, not even an explicit nil
### GetMaxNodeCount

`func (o *NewClusterInput) GetMaxNodeCount() int32`

GetMaxNodeCount returns the MaxNodeCount field if non-nil, zero value otherwise.

### GetMaxNodeCountOk

`func (o *NewClusterInput) GetMaxNodeCountOk() (*int32, bool)`

GetMaxNodeCountOk returns a tuple with the MaxNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNodeCount

`func (o *NewClusterInput) SetMaxNodeCount(v int32)`

SetMaxNodeCount sets MaxNodeCount field to given value.

### HasMaxNodeCount

`func (o *NewClusterInput) HasMaxNodeCount() bool`

HasMaxNodeCount returns a boolean if a field has been set.

### SetMaxNodeCountNil

`func (o *NewClusterInput) SetMaxNodeCountNil(b bool)`

 SetMaxNodeCountNil sets the value for MaxNodeCount to be an explicit nil

### UnsetMaxNodeCount
`func (o *NewClusterInput) UnsetMaxNodeCount()`

UnsetMaxNodeCount ensures that no value is present for MaxNodeCount, not even an explicit nil
### GetNodeDiskSize

`func (o *NewClusterInput) GetNodeDiskSize() int32`

GetNodeDiskSize returns the NodeDiskSize field if non-nil, zero value otherwise.

### GetNodeDiskSizeOk

`func (o *NewClusterInput) GetNodeDiskSizeOk() (*int32, bool)`

GetNodeDiskSizeOk returns a tuple with the NodeDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeDiskSize

`func (o *NewClusterInput) SetNodeDiskSize(v int32)`

SetNodeDiskSize sets NodeDiskSize field to given value.

### HasNodeDiskSize

`func (o *NewClusterInput) HasNodeDiskSize() bool`

HasNodeDiskSize returns a boolean if a field has been set.

### SetNodeDiskSizeNil

`func (o *NewClusterInput) SetNodeDiskSizeNil(b bool)`

 SetNodeDiskSizeNil sets the value for NodeDiskSize to be an explicit nil

### UnsetNodeDiskSize
`func (o *NewClusterInput) UnsetNodeDiskSize()`

UnsetNodeDiskSize ensures that no value is present for NodeDiskSize, not even an explicit nil
### GetZone

`func (o *NewClusterInput) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *NewClusterInput) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *NewClusterInput) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *NewClusterInput) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *NewClusterInput) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *NewClusterInput) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil
### GetRegion

`func (o *NewClusterInput) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NewClusterInput) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NewClusterInput) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *NewClusterInput) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *NewClusterInput) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *NewClusterInput) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetBillingOption

`func (o *NewClusterInput) GetBillingOption() string`

GetBillingOption returns the BillingOption field if non-nil, zero value otherwise.

### GetBillingOptionOk

`func (o *NewClusterInput) GetBillingOptionOk() (*string, bool)`

GetBillingOptionOk returns a tuple with the BillingOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingOption

`func (o *NewClusterInput) SetBillingOption(v string)`

SetBillingOption sets BillingOption field to given value.

### HasBillingOption

`func (o *NewClusterInput) HasBillingOption() bool`

HasBillingOption returns a boolean if a field has been set.

### SetBillingOptionNil

`func (o *NewClusterInput) SetBillingOptionNil(b bool)`

 SetBillingOptionNil sets the value for BillingOption to be an explicit nil

### UnsetBillingOption
`func (o *NewClusterInput) UnsetBillingOption()`

UnsetBillingOption ensures that no value is present for BillingOption, not even an explicit nil
### GetDisableMonitoring

`func (o *NewClusterInput) GetDisableMonitoring() bool`

GetDisableMonitoring returns the DisableMonitoring field if non-nil, zero value otherwise.

### GetDisableMonitoringOk

`func (o *NewClusterInput) GetDisableMonitoringOk() (*bool, bool)`

GetDisableMonitoringOk returns a tuple with the DisableMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableMonitoring

`func (o *NewClusterInput) SetDisableMonitoring(v bool)`

SetDisableMonitoring sets DisableMonitoring field to given value.


### GetAutoInfrastructureUpgrade

`func (o *NewClusterInput) GetAutoInfrastructureUpgrade() bool`

GetAutoInfrastructureUpgrade returns the AutoInfrastructureUpgrade field if non-nil, zero value otherwise.

### GetAutoInfrastructureUpgradeOk

`func (o *NewClusterInput) GetAutoInfrastructureUpgradeOk() (*bool, bool)`

GetAutoInfrastructureUpgradeOk returns a tuple with the AutoInfrastructureUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoInfrastructureUpgrade

`func (o *NewClusterInput) SetAutoInfrastructureUpgrade(v bool)`

SetAutoInfrastructureUpgrade sets AutoInfrastructureUpgrade field to given value.

### HasAutoInfrastructureUpgrade

`func (o *NewClusterInput) HasAutoInfrastructureUpgrade() bool`

HasAutoInfrastructureUpgrade returns a boolean if a field has been set.

### SetAutoInfrastructureUpgradeNil

`func (o *NewClusterInput) SetAutoInfrastructureUpgradeNil(b bool)`

 SetAutoInfrastructureUpgradeNil sets the value for AutoInfrastructureUpgrade to be an explicit nil

### UnsetAutoInfrastructureUpgrade
`func (o *NewClusterInput) UnsetAutoInfrastructureUpgrade()`

UnsetAutoInfrastructureUpgrade ensures that no value is present for AutoInfrastructureUpgrade, not even an explicit nil
### GetEnvironmentPolicy

`func (o *NewClusterInput) GetEnvironmentPolicy() ClusterEnvironmentPolicyInput`

GetEnvironmentPolicy returns the EnvironmentPolicy field if non-nil, zero value otherwise.

### GetEnvironmentPolicyOk

`func (o *NewClusterInput) GetEnvironmentPolicyOk() (*ClusterEnvironmentPolicyInput, bool)`

GetEnvironmentPolicyOk returns a tuple with the EnvironmentPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentPolicy

`func (o *NewClusterInput) SetEnvironmentPolicy(v ClusterEnvironmentPolicyInput)`

SetEnvironmentPolicy sets EnvironmentPolicy field to given value.

### HasEnvironmentPolicy

`func (o *NewClusterInput) HasEnvironmentPolicy() bool`

HasEnvironmentPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


