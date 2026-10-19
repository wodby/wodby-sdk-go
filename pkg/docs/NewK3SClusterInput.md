# NewK3SClusterInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. | [optional] 
**ProjectId** | Pointer to **NullableInt32** |  | [optional] 
**AvailableToAllProjects** | Pointer to **NullableBool** | Lets every project in the organization use the new resource without a project share. Only an organization-owned resource, created without projectId, can be available to all projects. | [optional] 
**Name** | **string** |  | 
**Title** | **string** |  | 
**DisableMonitoring** | **bool** |  | 
**AutoInfrastructureUpgrade** | Pointer to **NullableBool** |  | [optional] 
**EnvironmentPolicy** | Pointer to [**ClusterEnvironmentPolicyInput**](ClusterEnvironmentPolicyInput.md) |  | [optional] 

## Methods

### NewNewK3SClusterInput

`func NewNewK3SClusterInput(name string, title string, disableMonitoring bool, ) *NewK3SClusterInput`

NewNewK3SClusterInput instantiates a new NewK3SClusterInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewK3SClusterInputWithDefaults

`func NewNewK3SClusterInputWithDefaults() *NewK3SClusterInput`

NewNewK3SClusterInputWithDefaults instantiates a new NewK3SClusterInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *NewK3SClusterInput) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *NewK3SClusterInput) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *NewK3SClusterInput) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *NewK3SClusterInput) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProjectId

`func (o *NewK3SClusterInput) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *NewK3SClusterInput) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *NewK3SClusterInput) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *NewK3SClusterInput) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *NewK3SClusterInput) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *NewK3SClusterInput) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetAvailableToAllProjects

`func (o *NewK3SClusterInput) GetAvailableToAllProjects() bool`

GetAvailableToAllProjects returns the AvailableToAllProjects field if non-nil, zero value otherwise.

### GetAvailableToAllProjectsOk

`func (o *NewK3SClusterInput) GetAvailableToAllProjectsOk() (*bool, bool)`

GetAvailableToAllProjectsOk returns a tuple with the AvailableToAllProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableToAllProjects

`func (o *NewK3SClusterInput) SetAvailableToAllProjects(v bool)`

SetAvailableToAllProjects sets AvailableToAllProjects field to given value.

### HasAvailableToAllProjects

`func (o *NewK3SClusterInput) HasAvailableToAllProjects() bool`

HasAvailableToAllProjects returns a boolean if a field has been set.

### SetAvailableToAllProjectsNil

`func (o *NewK3SClusterInput) SetAvailableToAllProjectsNil(b bool)`

 SetAvailableToAllProjectsNil sets the value for AvailableToAllProjects to be an explicit nil

### UnsetAvailableToAllProjects
`func (o *NewK3SClusterInput) UnsetAvailableToAllProjects()`

UnsetAvailableToAllProjects ensures that no value is present for AvailableToAllProjects, not even an explicit nil
### GetName

`func (o *NewK3SClusterInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewK3SClusterInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewK3SClusterInput) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *NewK3SClusterInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewK3SClusterInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NewK3SClusterInput) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDisableMonitoring

`func (o *NewK3SClusterInput) GetDisableMonitoring() bool`

GetDisableMonitoring returns the DisableMonitoring field if non-nil, zero value otherwise.

### GetDisableMonitoringOk

`func (o *NewK3SClusterInput) GetDisableMonitoringOk() (*bool, bool)`

GetDisableMonitoringOk returns a tuple with the DisableMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableMonitoring

`func (o *NewK3SClusterInput) SetDisableMonitoring(v bool)`

SetDisableMonitoring sets DisableMonitoring field to given value.


### GetAutoInfrastructureUpgrade

`func (o *NewK3SClusterInput) GetAutoInfrastructureUpgrade() bool`

GetAutoInfrastructureUpgrade returns the AutoInfrastructureUpgrade field if non-nil, zero value otherwise.

### GetAutoInfrastructureUpgradeOk

`func (o *NewK3SClusterInput) GetAutoInfrastructureUpgradeOk() (*bool, bool)`

GetAutoInfrastructureUpgradeOk returns a tuple with the AutoInfrastructureUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoInfrastructureUpgrade

`func (o *NewK3SClusterInput) SetAutoInfrastructureUpgrade(v bool)`

SetAutoInfrastructureUpgrade sets AutoInfrastructureUpgrade field to given value.

### HasAutoInfrastructureUpgrade

`func (o *NewK3SClusterInput) HasAutoInfrastructureUpgrade() bool`

HasAutoInfrastructureUpgrade returns a boolean if a field has been set.

### SetAutoInfrastructureUpgradeNil

`func (o *NewK3SClusterInput) SetAutoInfrastructureUpgradeNil(b bool)`

 SetAutoInfrastructureUpgradeNil sets the value for AutoInfrastructureUpgrade to be an explicit nil

### UnsetAutoInfrastructureUpgrade
`func (o *NewK3SClusterInput) UnsetAutoInfrastructureUpgrade()`

UnsetAutoInfrastructureUpgrade ensures that no value is present for AutoInfrastructureUpgrade, not even an explicit nil
### GetEnvironmentPolicy

`func (o *NewK3SClusterInput) GetEnvironmentPolicy() ClusterEnvironmentPolicyInput`

GetEnvironmentPolicy returns the EnvironmentPolicy field if non-nil, zero value otherwise.

### GetEnvironmentPolicyOk

`func (o *NewK3SClusterInput) GetEnvironmentPolicyOk() (*ClusterEnvironmentPolicyInput, bool)`

GetEnvironmentPolicyOk returns a tuple with the EnvironmentPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentPolicy

`func (o *NewK3SClusterInput) SetEnvironmentPolicy(v ClusterEnvironmentPolicyInput)`

SetEnvironmentPolicy sets EnvironmentPolicy field to given value.

### HasEnvironmentPolicy

`func (o *NewK3SClusterInput) HasEnvironmentPolicy() bool`

HasEnvironmentPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


