# ServiceManifestUpdateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | Optional service revision version; defaults to the current service version. | [optional] 
**ManifestYaml** | **string** | Complete Wodby service.yml manifest content. | 
**Files** | Pointer to **map[string]string** | Optional referenced file contents keyed by manifest-relative path, for example Dockerfile or configs/app.conf. | [optional] 

## Methods

### NewServiceManifestUpdateInput

`func NewServiceManifestUpdateInput(manifestYaml string, ) *ServiceManifestUpdateInput`

NewServiceManifestUpdateInput instantiates a new ServiceManifestUpdateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceManifestUpdateInputWithDefaults

`func NewServiceManifestUpdateInputWithDefaults() *ServiceManifestUpdateInput`

NewServiceManifestUpdateInputWithDefaults instantiates a new ServiceManifestUpdateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ServiceManifestUpdateInput) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ServiceManifestUpdateInput) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ServiceManifestUpdateInput) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ServiceManifestUpdateInput) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetManifestYaml

`func (o *ServiceManifestUpdateInput) GetManifestYaml() string`

GetManifestYaml returns the ManifestYaml field if non-nil, zero value otherwise.

### GetManifestYamlOk

`func (o *ServiceManifestUpdateInput) GetManifestYamlOk() (*string, bool)`

GetManifestYamlOk returns a tuple with the ManifestYaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestYaml

`func (o *ServiceManifestUpdateInput) SetManifestYaml(v string)`

SetManifestYaml sets ManifestYaml field to given value.


### GetFiles

`func (o *ServiceManifestUpdateInput) GetFiles() map[string]string`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ServiceManifestUpdateInput) GetFilesOk() (*map[string]string, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ServiceManifestUpdateInput) SetFiles(v map[string]string)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ServiceManifestUpdateInput) HasFiles() bool`

HasFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


