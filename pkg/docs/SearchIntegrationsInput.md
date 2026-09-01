# SearchIntegrationsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. | [optional] 
**ProjectIds** | Pointer to **[]int32** |  | [optional] 
**Types** | Pointer to **[]string** |  | [optional] 
**Statuses** | Pointer to **[]string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Variables** | Pointer to [**[]IntegrationVariableRequirementInput**](IntegrationVariableRequirementInput.md) |  | [optional] 
**EnvId** | Pointer to **NullableInt32** |  | [optional] 
**EnvType** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSearchIntegrationsInput

`func NewSearchIntegrationsInput() *SearchIntegrationsInput`

NewSearchIntegrationsInput instantiates a new SearchIntegrationsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchIntegrationsInputWithDefaults

`func NewSearchIntegrationsInputWithDefaults() *SearchIntegrationsInput`

NewSearchIntegrationsInputWithDefaults instantiates a new SearchIntegrationsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *SearchIntegrationsInput) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *SearchIntegrationsInput) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *SearchIntegrationsInput) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *SearchIntegrationsInput) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProjectIds

`func (o *SearchIntegrationsInput) GetProjectIds() []int32`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *SearchIntegrationsInput) GetProjectIdsOk() (*[]int32, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *SearchIntegrationsInput) SetProjectIds(v []int32)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *SearchIntegrationsInput) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### GetTypes

`func (o *SearchIntegrationsInput) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *SearchIntegrationsInput) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *SearchIntegrationsInput) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *SearchIntegrationsInput) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetStatuses

`func (o *SearchIntegrationsInput) GetStatuses() []string`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *SearchIntegrationsInput) GetStatusesOk() (*[]string, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *SearchIntegrationsInput) SetStatuses(v []string)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *SearchIntegrationsInput) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetLabels

`func (o *SearchIntegrationsInput) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *SearchIntegrationsInput) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *SearchIntegrationsInput) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *SearchIntegrationsInput) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetVariables

`func (o *SearchIntegrationsInput) GetVariables() []IntegrationVariableRequirementInput`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *SearchIntegrationsInput) GetVariablesOk() (*[]IntegrationVariableRequirementInput, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *SearchIntegrationsInput) SetVariables(v []IntegrationVariableRequirementInput)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *SearchIntegrationsInput) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetEnvId

`func (o *SearchIntegrationsInput) GetEnvId() int32`

GetEnvId returns the EnvId field if non-nil, zero value otherwise.

### GetEnvIdOk

`func (o *SearchIntegrationsInput) GetEnvIdOk() (*int32, bool)`

GetEnvIdOk returns a tuple with the EnvId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvId

`func (o *SearchIntegrationsInput) SetEnvId(v int32)`

SetEnvId sets EnvId field to given value.

### HasEnvId

`func (o *SearchIntegrationsInput) HasEnvId() bool`

HasEnvId returns a boolean if a field has been set.

### SetEnvIdNil

`func (o *SearchIntegrationsInput) SetEnvIdNil(b bool)`

 SetEnvIdNil sets the value for EnvId to be an explicit nil

### UnsetEnvId
`func (o *SearchIntegrationsInput) UnsetEnvId()`

UnsetEnvId ensures that no value is present for EnvId, not even an explicit nil
### GetEnvType

`func (o *SearchIntegrationsInput) GetEnvType() string`

GetEnvType returns the EnvType field if non-nil, zero value otherwise.

### GetEnvTypeOk

`func (o *SearchIntegrationsInput) GetEnvTypeOk() (*string, bool)`

GetEnvTypeOk returns a tuple with the EnvType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvType

`func (o *SearchIntegrationsInput) SetEnvType(v string)`

SetEnvType sets EnvType field to given value.

### HasEnvType

`func (o *SearchIntegrationsInput) HasEnvType() bool`

HasEnvType returns a boolean if a field has been set.

### SetEnvTypeNil

`func (o *SearchIntegrationsInput) SetEnvTypeNil(b bool)`

 SetEnvTypeNil sets the value for EnvType to be an explicit nil

### UnsetEnvType
`func (o *SearchIntegrationsInput) UnsetEnvType()`

UnsetEnvType ensures that no value is present for EnvType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


