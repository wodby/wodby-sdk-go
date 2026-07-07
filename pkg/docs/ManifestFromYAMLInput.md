# ManifestFromYAMLInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. | [optional] 
**ProjectId** | Pointer to **NullableInt32** |  | [optional] 
**Version** | Pointer to **string** | Optional revision version label for the generated non-Git resource. | [optional] 
**ManifestYaml** | **string** | Complete Wodby service.yml or stack.yml manifest content. | 
**Files** | Pointer to **map[string]string** | Optional referenced file contents keyed by manifest-relative path, for example Dockerfile or configs/app.conf. | [optional] 

## Methods

### NewManifestFromYAMLInput

`func NewManifestFromYAMLInput(manifestYaml string, ) *ManifestFromYAMLInput`

NewManifestFromYAMLInput instantiates a new ManifestFromYAMLInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManifestFromYAMLInputWithDefaults

`func NewManifestFromYAMLInputWithDefaults() *ManifestFromYAMLInput`

NewManifestFromYAMLInputWithDefaults instantiates a new ManifestFromYAMLInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *ManifestFromYAMLInput) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ManifestFromYAMLInput) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ManifestFromYAMLInput) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ManifestFromYAMLInput) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProjectId

`func (o *ManifestFromYAMLInput) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ManifestFromYAMLInput) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ManifestFromYAMLInput) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ManifestFromYAMLInput) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *ManifestFromYAMLInput) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *ManifestFromYAMLInput) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetVersion

`func (o *ManifestFromYAMLInput) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ManifestFromYAMLInput) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ManifestFromYAMLInput) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ManifestFromYAMLInput) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetManifestYaml

`func (o *ManifestFromYAMLInput) GetManifestYaml() string`

GetManifestYaml returns the ManifestYaml field if non-nil, zero value otherwise.

### GetManifestYamlOk

`func (o *ManifestFromYAMLInput) GetManifestYamlOk() (*string, bool)`

GetManifestYamlOk returns a tuple with the ManifestYaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestYaml

`func (o *ManifestFromYAMLInput) SetManifestYaml(v string)`

SetManifestYaml sets ManifestYaml field to given value.


### GetFiles

`func (o *ManifestFromYAMLInput) GetFiles() map[string]string`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ManifestFromYAMLInput) GetFilesOk() (*map[string]string, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ManifestFromYAMLInput) SetFiles(v map[string]string)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ManifestFromYAMLInput) HasFiles() bool`

HasFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


