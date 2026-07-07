# HelmChartInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceName** | Pointer to **string** | Optional chart source name to use in generated Wodby manifests. | [optional] 
**Source** | Pointer to **string** | Optional Helm repository or OCI source URL. | [optional] 
**Chart** | **string** | Helm chart reference, such as bitnami/redis, oci://registry.example.com/chart, a chart archive URL, or a server-local chart path. | 
**Version** | Pointer to **string** | Optional Helm chart version. | [optional] 
**Release** | Pointer to **string** | Optional Helm release name used for rendering analysis. | [optional] 
**Namespace** | Pointer to **string** | Optional Kubernetes namespace used for rendering analysis. | [optional] 
**Values** | Pointer to **map[string]interface{}** |  | [optional] 
**ValuesYaml** | Pointer to **string** | Optional Helm values YAML. Use either values or valuesYaml, not both. | [optional] 

## Methods

### NewHelmChartInput

`func NewHelmChartInput(chart string, ) *HelmChartInput`

NewHelmChartInput instantiates a new HelmChartInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmChartInputWithDefaults

`func NewHelmChartInputWithDefaults() *HelmChartInput`

NewHelmChartInputWithDefaults instantiates a new HelmChartInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceName

`func (o *HelmChartInput) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *HelmChartInput) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *HelmChartInput) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *HelmChartInput) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### GetSource

`func (o *HelmChartInput) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *HelmChartInput) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *HelmChartInput) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *HelmChartInput) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetChart

`func (o *HelmChartInput) GetChart() string`

GetChart returns the Chart field if non-nil, zero value otherwise.

### GetChartOk

`func (o *HelmChartInput) GetChartOk() (*string, bool)`

GetChartOk returns a tuple with the Chart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChart

`func (o *HelmChartInput) SetChart(v string)`

SetChart sets Chart field to given value.


### GetVersion

`func (o *HelmChartInput) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HelmChartInput) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HelmChartInput) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HelmChartInput) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRelease

`func (o *HelmChartInput) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *HelmChartInput) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *HelmChartInput) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *HelmChartInput) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetNamespace

`func (o *HelmChartInput) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *HelmChartInput) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *HelmChartInput) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *HelmChartInput) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetValues

`func (o *HelmChartInput) GetValues() map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *HelmChartInput) GetValuesOk() (*map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *HelmChartInput) SetValues(v map[string]interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *HelmChartInput) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetValuesYaml

`func (o *HelmChartInput) GetValuesYaml() string`

GetValuesYaml returns the ValuesYaml field if non-nil, zero value otherwise.

### GetValuesYamlOk

`func (o *HelmChartInput) GetValuesYamlOk() (*string, bool)`

GetValuesYamlOk returns a tuple with the ValuesYaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesYaml

`func (o *HelmChartInput) SetValuesYaml(v string)`

SetValuesYaml sets ValuesYaml field to given value.

### HasValuesYaml

`func (o *HelmChartInput) HasValuesYaml() bool`

HasValuesYaml returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


