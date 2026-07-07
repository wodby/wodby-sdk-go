# HelmChartStackScaffoldResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Analysis** | [**HelmChartAnalysis**](HelmChartAnalysis.md) |  | 
**ServiceManifestYaml** | **string** |  | 
**ServiceManifest** | **map[string]interface{}** |  | 
**StackManifestYaml** | **string** |  | 
**StackManifest** | **map[string]interface{}** |  | 
**Warnings** | Pointer to **[]string** |  | [optional] 

## Methods

### NewHelmChartStackScaffoldResponse

`func NewHelmChartStackScaffoldResponse(analysis HelmChartAnalysis, serviceManifestYaml string, serviceManifest map[string]interface{}, stackManifestYaml string, stackManifest map[string]interface{}, ) *HelmChartStackScaffoldResponse`

NewHelmChartStackScaffoldResponse instantiates a new HelmChartStackScaffoldResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmChartStackScaffoldResponseWithDefaults

`func NewHelmChartStackScaffoldResponseWithDefaults() *HelmChartStackScaffoldResponse`

NewHelmChartStackScaffoldResponseWithDefaults instantiates a new HelmChartStackScaffoldResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysis

`func (o *HelmChartStackScaffoldResponse) GetAnalysis() HelmChartAnalysis`

GetAnalysis returns the Analysis field if non-nil, zero value otherwise.

### GetAnalysisOk

`func (o *HelmChartStackScaffoldResponse) GetAnalysisOk() (*HelmChartAnalysis, bool)`

GetAnalysisOk returns a tuple with the Analysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysis

`func (o *HelmChartStackScaffoldResponse) SetAnalysis(v HelmChartAnalysis)`

SetAnalysis sets Analysis field to given value.


### GetServiceManifestYaml

`func (o *HelmChartStackScaffoldResponse) GetServiceManifestYaml() string`

GetServiceManifestYaml returns the ServiceManifestYaml field if non-nil, zero value otherwise.

### GetServiceManifestYamlOk

`func (o *HelmChartStackScaffoldResponse) GetServiceManifestYamlOk() (*string, bool)`

GetServiceManifestYamlOk returns a tuple with the ServiceManifestYaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceManifestYaml

`func (o *HelmChartStackScaffoldResponse) SetServiceManifestYaml(v string)`

SetServiceManifestYaml sets ServiceManifestYaml field to given value.


### GetServiceManifest

`func (o *HelmChartStackScaffoldResponse) GetServiceManifest() map[string]interface{}`

GetServiceManifest returns the ServiceManifest field if non-nil, zero value otherwise.

### GetServiceManifestOk

`func (o *HelmChartStackScaffoldResponse) GetServiceManifestOk() (*map[string]interface{}, bool)`

GetServiceManifestOk returns a tuple with the ServiceManifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceManifest

`func (o *HelmChartStackScaffoldResponse) SetServiceManifest(v map[string]interface{})`

SetServiceManifest sets ServiceManifest field to given value.


### GetStackManifestYaml

`func (o *HelmChartStackScaffoldResponse) GetStackManifestYaml() string`

GetStackManifestYaml returns the StackManifestYaml field if non-nil, zero value otherwise.

### GetStackManifestYamlOk

`func (o *HelmChartStackScaffoldResponse) GetStackManifestYamlOk() (*string, bool)`

GetStackManifestYamlOk returns a tuple with the StackManifestYaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackManifestYaml

`func (o *HelmChartStackScaffoldResponse) SetStackManifestYaml(v string)`

SetStackManifestYaml sets StackManifestYaml field to given value.


### GetStackManifest

`func (o *HelmChartStackScaffoldResponse) GetStackManifest() map[string]interface{}`

GetStackManifest returns the StackManifest field if non-nil, zero value otherwise.

### GetStackManifestOk

`func (o *HelmChartStackScaffoldResponse) GetStackManifestOk() (*map[string]interface{}, bool)`

GetStackManifestOk returns a tuple with the StackManifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackManifest

`func (o *HelmChartStackScaffoldResponse) SetStackManifest(v map[string]interface{})`

SetStackManifest sets StackManifest field to given value.


### GetWarnings

`func (o *HelmChartStackScaffoldResponse) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *HelmChartStackScaffoldResponse) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *HelmChartStackScaffoldResponse) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *HelmChartStackScaffoldResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


