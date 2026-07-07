# HelmChartServiceScaffoldResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Analysis** | [**HelmChartAnalysis**](HelmChartAnalysis.md) |  | 
**ManifestYaml** | **string** |  | 
**Manifest** | **map[string]interface{}** |  | 
**Warnings** | Pointer to **[]string** |  | [optional] 

## Methods

### NewHelmChartServiceScaffoldResponse

`func NewHelmChartServiceScaffoldResponse(analysis HelmChartAnalysis, manifestYaml string, manifest map[string]interface{}, ) *HelmChartServiceScaffoldResponse`

NewHelmChartServiceScaffoldResponse instantiates a new HelmChartServiceScaffoldResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmChartServiceScaffoldResponseWithDefaults

`func NewHelmChartServiceScaffoldResponseWithDefaults() *HelmChartServiceScaffoldResponse`

NewHelmChartServiceScaffoldResponseWithDefaults instantiates a new HelmChartServiceScaffoldResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysis

`func (o *HelmChartServiceScaffoldResponse) GetAnalysis() HelmChartAnalysis`

GetAnalysis returns the Analysis field if non-nil, zero value otherwise.

### GetAnalysisOk

`func (o *HelmChartServiceScaffoldResponse) GetAnalysisOk() (*HelmChartAnalysis, bool)`

GetAnalysisOk returns a tuple with the Analysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysis

`func (o *HelmChartServiceScaffoldResponse) SetAnalysis(v HelmChartAnalysis)`

SetAnalysis sets Analysis field to given value.


### GetManifestYaml

`func (o *HelmChartServiceScaffoldResponse) GetManifestYaml() string`

GetManifestYaml returns the ManifestYaml field if non-nil, zero value otherwise.

### GetManifestYamlOk

`func (o *HelmChartServiceScaffoldResponse) GetManifestYamlOk() (*string, bool)`

GetManifestYamlOk returns a tuple with the ManifestYaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestYaml

`func (o *HelmChartServiceScaffoldResponse) SetManifestYaml(v string)`

SetManifestYaml sets ManifestYaml field to given value.


### GetManifest

`func (o *HelmChartServiceScaffoldResponse) GetManifest() map[string]interface{}`

GetManifest returns the Manifest field if non-nil, zero value otherwise.

### GetManifestOk

`func (o *HelmChartServiceScaffoldResponse) GetManifestOk() (*map[string]interface{}, bool)`

GetManifestOk returns a tuple with the Manifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifest

`func (o *HelmChartServiceScaffoldResponse) SetManifest(v map[string]interface{})`

SetManifest sets Manifest field to given value.


### GetWarnings

`func (o *HelmChartServiceScaffoldResponse) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *HelmChartServiceScaffoldResponse) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *HelmChartServiceScaffoldResponse) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *HelmChartServiceScaffoldResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


