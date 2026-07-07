# HelmChartServiceScaffoldInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chart** | [**HelmChartInput**](HelmChartInput.md) |  | 
**ServiceName** | Pointer to **string** | Optional generated Wodby service name. | [optional] 
**ServiceTitle** | Pointer to **string** | Optional generated Wodby service title. | [optional] 
**ServiceType** | Pointer to **string** | Optional generated Wodby service type. Defaults to service. | [optional] 
**Icon** | Pointer to **string** | Optional generated service icon. | [optional] 

## Methods

### NewHelmChartServiceScaffoldInput

`func NewHelmChartServiceScaffoldInput(chart HelmChartInput, ) *HelmChartServiceScaffoldInput`

NewHelmChartServiceScaffoldInput instantiates a new HelmChartServiceScaffoldInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmChartServiceScaffoldInputWithDefaults

`func NewHelmChartServiceScaffoldInputWithDefaults() *HelmChartServiceScaffoldInput`

NewHelmChartServiceScaffoldInputWithDefaults instantiates a new HelmChartServiceScaffoldInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChart

`func (o *HelmChartServiceScaffoldInput) GetChart() HelmChartInput`

GetChart returns the Chart field if non-nil, zero value otherwise.

### GetChartOk

`func (o *HelmChartServiceScaffoldInput) GetChartOk() (*HelmChartInput, bool)`

GetChartOk returns a tuple with the Chart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChart

`func (o *HelmChartServiceScaffoldInput) SetChart(v HelmChartInput)`

SetChart sets Chart field to given value.


### GetServiceName

`func (o *HelmChartServiceScaffoldInput) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *HelmChartServiceScaffoldInput) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *HelmChartServiceScaffoldInput) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *HelmChartServiceScaffoldInput) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetServiceTitle

`func (o *HelmChartServiceScaffoldInput) GetServiceTitle() string`

GetServiceTitle returns the ServiceTitle field if non-nil, zero value otherwise.

### GetServiceTitleOk

`func (o *HelmChartServiceScaffoldInput) GetServiceTitleOk() (*string, bool)`

GetServiceTitleOk returns a tuple with the ServiceTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTitle

`func (o *HelmChartServiceScaffoldInput) SetServiceTitle(v string)`

SetServiceTitle sets ServiceTitle field to given value.

### HasServiceTitle

`func (o *HelmChartServiceScaffoldInput) HasServiceTitle() bool`

HasServiceTitle returns a boolean if a field has been set.

### GetServiceType

`func (o *HelmChartServiceScaffoldInput) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *HelmChartServiceScaffoldInput) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *HelmChartServiceScaffoldInput) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *HelmChartServiceScaffoldInput) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetIcon

`func (o *HelmChartServiceScaffoldInput) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *HelmChartServiceScaffoldInput) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *HelmChartServiceScaffoldInput) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *HelmChartServiceScaffoldInput) HasIcon() bool`

HasIcon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


