# ProviderManifestInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. | [optional] 
**ProjectId** | Pointer to **NullableInt32** |  | [optional] 
**ManifestYaml** | **string** | Complete variable provider manifest YAML. | 

## Methods

### NewProviderManifestInput

`func NewProviderManifestInput(manifestYaml string, ) *ProviderManifestInput`

NewProviderManifestInput instantiates a new ProviderManifestInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderManifestInputWithDefaults

`func NewProviderManifestInputWithDefaults() *ProviderManifestInput`

NewProviderManifestInputWithDefaults instantiates a new ProviderManifestInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *ProviderManifestInput) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ProviderManifestInput) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ProviderManifestInput) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ProviderManifestInput) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProjectId

`func (o *ProviderManifestInput) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ProviderManifestInput) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ProviderManifestInput) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ProviderManifestInput) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *ProviderManifestInput) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *ProviderManifestInput) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetManifestYaml

`func (o *ProviderManifestInput) GetManifestYaml() string`

GetManifestYaml returns the ManifestYaml field if non-nil, zero value otherwise.

### GetManifestYamlOk

`func (o *ProviderManifestInput) GetManifestYamlOk() (*string, bool)`

GetManifestYamlOk returns a tuple with the ManifestYaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestYaml

`func (o *ProviderManifestInput) SetManifestYaml(v string)`

SetManifestYaml sets ManifestYaml field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


