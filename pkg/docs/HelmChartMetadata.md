# HelmChartMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Version** | Pointer to **string** |  | [optional] 
**AppVersion** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Chart** | **string** |  | 

## Methods

### NewHelmChartMetadata

`func NewHelmChartMetadata(name string, chart string, ) *HelmChartMetadata`

NewHelmChartMetadata instantiates a new HelmChartMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmChartMetadataWithDefaults

`func NewHelmChartMetadataWithDefaults() *HelmChartMetadata`

NewHelmChartMetadataWithDefaults instantiates a new HelmChartMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HelmChartMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HelmChartMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HelmChartMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *HelmChartMetadata) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HelmChartMetadata) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HelmChartMetadata) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HelmChartMetadata) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAppVersion

`func (o *HelmChartMetadata) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *HelmChartMetadata) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *HelmChartMetadata) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *HelmChartMetadata) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### GetDescription

`func (o *HelmChartMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HelmChartMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HelmChartMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HelmChartMetadata) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSource

`func (o *HelmChartMetadata) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *HelmChartMetadata) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *HelmChartMetadata) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *HelmChartMetadata) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetChart

`func (o *HelmChartMetadata) GetChart() string`

GetChart returns the Chart field if non-nil, zero value otherwise.

### GetChartOk

`func (o *HelmChartMetadata) GetChartOk() (*string, bool)`

GetChartOk returns a tuple with the Chart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChart

`func (o *HelmChartMetadata) SetChart(v string)`

SetChart sets Chart field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


