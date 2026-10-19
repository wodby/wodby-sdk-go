# NewWodbyCloudClusterInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. | [optional] 
**ProjectId** | Pointer to **NullableInt32** |  | [optional] 
**AvailableToAllProjects** | Pointer to **NullableBool** | Lets every project in the organization use the new resource without a project share. Only an organization-owned resource, created without projectId, can be available to all projects. | [optional] 
**Demo** | **bool** | Create a free demo cluster with limited resources. | 
**SingleNode** | Pointer to **NullableBool** |  | [optional] 
**Region** | Pointer to **NullableString** |  | [optional] 
**MachineType** | Pointer to **NullableString** | Machine type name from GET /wodby-cloud/machine-types. | [optional] 
**MinNodeCount** | Pointer to **NullableInt32** |  | [optional] 
**MaxNodeCount** | Pointer to **NullableInt32** |  | [optional] 
**EnvironmentPolicy** | Pointer to [**ClusterEnvironmentPolicyInput**](ClusterEnvironmentPolicyInput.md) |  | [optional] 

## Methods

### NewNewWodbyCloudClusterInput

`func NewNewWodbyCloudClusterInput(demo bool, ) *NewWodbyCloudClusterInput`

NewNewWodbyCloudClusterInput instantiates a new NewWodbyCloudClusterInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewWodbyCloudClusterInputWithDefaults

`func NewNewWodbyCloudClusterInputWithDefaults() *NewWodbyCloudClusterInput`

NewNewWodbyCloudClusterInputWithDefaults instantiates a new NewWodbyCloudClusterInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *NewWodbyCloudClusterInput) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *NewWodbyCloudClusterInput) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *NewWodbyCloudClusterInput) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *NewWodbyCloudClusterInput) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProjectId

`func (o *NewWodbyCloudClusterInput) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *NewWodbyCloudClusterInput) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *NewWodbyCloudClusterInput) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *NewWodbyCloudClusterInput) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *NewWodbyCloudClusterInput) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *NewWodbyCloudClusterInput) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetAvailableToAllProjects

`func (o *NewWodbyCloudClusterInput) GetAvailableToAllProjects() bool`

GetAvailableToAllProjects returns the AvailableToAllProjects field if non-nil, zero value otherwise.

### GetAvailableToAllProjectsOk

`func (o *NewWodbyCloudClusterInput) GetAvailableToAllProjectsOk() (*bool, bool)`

GetAvailableToAllProjectsOk returns a tuple with the AvailableToAllProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableToAllProjects

`func (o *NewWodbyCloudClusterInput) SetAvailableToAllProjects(v bool)`

SetAvailableToAllProjects sets AvailableToAllProjects field to given value.

### HasAvailableToAllProjects

`func (o *NewWodbyCloudClusterInput) HasAvailableToAllProjects() bool`

HasAvailableToAllProjects returns a boolean if a field has been set.

### SetAvailableToAllProjectsNil

`func (o *NewWodbyCloudClusterInput) SetAvailableToAllProjectsNil(b bool)`

 SetAvailableToAllProjectsNil sets the value for AvailableToAllProjects to be an explicit nil

### UnsetAvailableToAllProjects
`func (o *NewWodbyCloudClusterInput) UnsetAvailableToAllProjects()`

UnsetAvailableToAllProjects ensures that no value is present for AvailableToAllProjects, not even an explicit nil
### GetDemo

`func (o *NewWodbyCloudClusterInput) GetDemo() bool`

GetDemo returns the Demo field if non-nil, zero value otherwise.

### GetDemoOk

`func (o *NewWodbyCloudClusterInput) GetDemoOk() (*bool, bool)`

GetDemoOk returns a tuple with the Demo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemo

`func (o *NewWodbyCloudClusterInput) SetDemo(v bool)`

SetDemo sets Demo field to given value.


### GetSingleNode

`func (o *NewWodbyCloudClusterInput) GetSingleNode() bool`

GetSingleNode returns the SingleNode field if non-nil, zero value otherwise.

### GetSingleNodeOk

`func (o *NewWodbyCloudClusterInput) GetSingleNodeOk() (*bool, bool)`

GetSingleNodeOk returns a tuple with the SingleNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleNode

`func (o *NewWodbyCloudClusterInput) SetSingleNode(v bool)`

SetSingleNode sets SingleNode field to given value.

### HasSingleNode

`func (o *NewWodbyCloudClusterInput) HasSingleNode() bool`

HasSingleNode returns a boolean if a field has been set.

### SetSingleNodeNil

`func (o *NewWodbyCloudClusterInput) SetSingleNodeNil(b bool)`

 SetSingleNodeNil sets the value for SingleNode to be an explicit nil

### UnsetSingleNode
`func (o *NewWodbyCloudClusterInput) UnsetSingleNode()`

UnsetSingleNode ensures that no value is present for SingleNode, not even an explicit nil
### GetRegion

`func (o *NewWodbyCloudClusterInput) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NewWodbyCloudClusterInput) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NewWodbyCloudClusterInput) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *NewWodbyCloudClusterInput) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *NewWodbyCloudClusterInput) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *NewWodbyCloudClusterInput) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetMachineType

`func (o *NewWodbyCloudClusterInput) GetMachineType() string`

GetMachineType returns the MachineType field if non-nil, zero value otherwise.

### GetMachineTypeOk

`func (o *NewWodbyCloudClusterInput) GetMachineTypeOk() (*string, bool)`

GetMachineTypeOk returns a tuple with the MachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineType

`func (o *NewWodbyCloudClusterInput) SetMachineType(v string)`

SetMachineType sets MachineType field to given value.

### HasMachineType

`func (o *NewWodbyCloudClusterInput) HasMachineType() bool`

HasMachineType returns a boolean if a field has been set.

### SetMachineTypeNil

`func (o *NewWodbyCloudClusterInput) SetMachineTypeNil(b bool)`

 SetMachineTypeNil sets the value for MachineType to be an explicit nil

### UnsetMachineType
`func (o *NewWodbyCloudClusterInput) UnsetMachineType()`

UnsetMachineType ensures that no value is present for MachineType, not even an explicit nil
### GetMinNodeCount

`func (o *NewWodbyCloudClusterInput) GetMinNodeCount() int32`

GetMinNodeCount returns the MinNodeCount field if non-nil, zero value otherwise.

### GetMinNodeCountOk

`func (o *NewWodbyCloudClusterInput) GetMinNodeCountOk() (*int32, bool)`

GetMinNodeCountOk returns a tuple with the MinNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNodeCount

`func (o *NewWodbyCloudClusterInput) SetMinNodeCount(v int32)`

SetMinNodeCount sets MinNodeCount field to given value.

### HasMinNodeCount

`func (o *NewWodbyCloudClusterInput) HasMinNodeCount() bool`

HasMinNodeCount returns a boolean if a field has been set.

### SetMinNodeCountNil

`func (o *NewWodbyCloudClusterInput) SetMinNodeCountNil(b bool)`

 SetMinNodeCountNil sets the value for MinNodeCount to be an explicit nil

### UnsetMinNodeCount
`func (o *NewWodbyCloudClusterInput) UnsetMinNodeCount()`

UnsetMinNodeCount ensures that no value is present for MinNodeCount, not even an explicit nil
### GetMaxNodeCount

`func (o *NewWodbyCloudClusterInput) GetMaxNodeCount() int32`

GetMaxNodeCount returns the MaxNodeCount field if non-nil, zero value otherwise.

### GetMaxNodeCountOk

`func (o *NewWodbyCloudClusterInput) GetMaxNodeCountOk() (*int32, bool)`

GetMaxNodeCountOk returns a tuple with the MaxNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNodeCount

`func (o *NewWodbyCloudClusterInput) SetMaxNodeCount(v int32)`

SetMaxNodeCount sets MaxNodeCount field to given value.

### HasMaxNodeCount

`func (o *NewWodbyCloudClusterInput) HasMaxNodeCount() bool`

HasMaxNodeCount returns a boolean if a field has been set.

### SetMaxNodeCountNil

`func (o *NewWodbyCloudClusterInput) SetMaxNodeCountNil(b bool)`

 SetMaxNodeCountNil sets the value for MaxNodeCount to be an explicit nil

### UnsetMaxNodeCount
`func (o *NewWodbyCloudClusterInput) UnsetMaxNodeCount()`

UnsetMaxNodeCount ensures that no value is present for MaxNodeCount, not even an explicit nil
### GetEnvironmentPolicy

`func (o *NewWodbyCloudClusterInput) GetEnvironmentPolicy() ClusterEnvironmentPolicyInput`

GetEnvironmentPolicy returns the EnvironmentPolicy field if non-nil, zero value otherwise.

### GetEnvironmentPolicyOk

`func (o *NewWodbyCloudClusterInput) GetEnvironmentPolicyOk() (*ClusterEnvironmentPolicyInput, bool)`

GetEnvironmentPolicyOk returns a tuple with the EnvironmentPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentPolicy

`func (o *NewWodbyCloudClusterInput) SetEnvironmentPolicy(v ClusterEnvironmentPolicyInput)`

SetEnvironmentPolicy sets EnvironmentPolicy field to given value.

### HasEnvironmentPolicy

`func (o *NewWodbyCloudClusterInput) HasEnvironmentPolicy() bool`

HasEnvironmentPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


