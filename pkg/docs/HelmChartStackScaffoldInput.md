# HelmChartStackScaffoldInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chart** | [**HelmChartInput**](HelmChartInput.md) |  | 
**ServiceName** | Pointer to **string** | Optional generated Wodby service name. | [optional] 
**ServiceTitle** | Pointer to **string** | Optional generated Wodby service title. | [optional] 
**ServiceType** | Pointer to **string** | Optional generated Wodby service type. Defaults to service. | [optional] 
**StackName** | Pointer to **string** | Optional generated Wodby stack name. | [optional] 
**StackTitle** | Pointer to **string** | Optional generated Wodby stack title. | [optional] 
**Icon** | Pointer to **string** | Optional generated service and stack icon. | [optional] 

## Methods

### NewHelmChartStackScaffoldInput

`func NewHelmChartStackScaffoldInput(chart HelmChartInput, ) *HelmChartStackScaffoldInput`

NewHelmChartStackScaffoldInput instantiates a new HelmChartStackScaffoldInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmChartStackScaffoldInputWithDefaults

`func NewHelmChartStackScaffoldInputWithDefaults() *HelmChartStackScaffoldInput`

NewHelmChartStackScaffoldInputWithDefaults instantiates a new HelmChartStackScaffoldInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChart

`func (o *HelmChartStackScaffoldInput) GetChart() HelmChartInput`

GetChart returns the Chart field if non-nil, zero value otherwise.

### GetChartOk

`func (o *HelmChartStackScaffoldInput) GetChartOk() (*HelmChartInput, bool)`

GetChartOk returns a tuple with the Chart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChart

`func (o *HelmChartStackScaffoldInput) SetChart(v HelmChartInput)`

SetChart sets Chart field to given value.


### GetServiceName

`func (o *HelmChartStackScaffoldInput) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *HelmChartStackScaffoldInput) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *HelmChartStackScaffoldInput) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *HelmChartStackScaffoldInput) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetServiceTitle

`func (o *HelmChartStackScaffoldInput) GetServiceTitle() string`

GetServiceTitle returns the ServiceTitle field if non-nil, zero value otherwise.

### GetServiceTitleOk

`func (o *HelmChartStackScaffoldInput) GetServiceTitleOk() (*string, bool)`

GetServiceTitleOk returns a tuple with the ServiceTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTitle

`func (o *HelmChartStackScaffoldInput) SetServiceTitle(v string)`

SetServiceTitle sets ServiceTitle field to given value.

### HasServiceTitle

`func (o *HelmChartStackScaffoldInput) HasServiceTitle() bool`

HasServiceTitle returns a boolean if a field has been set.

### GetServiceType

`func (o *HelmChartStackScaffoldInput) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *HelmChartStackScaffoldInput) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *HelmChartStackScaffoldInput) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *HelmChartStackScaffoldInput) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetStackName

`func (o *HelmChartStackScaffoldInput) GetStackName() string`

GetStackName returns the StackName field if non-nil, zero value otherwise.

### GetStackNameOk

`func (o *HelmChartStackScaffoldInput) GetStackNameOk() (*string, bool)`

GetStackNameOk returns a tuple with the StackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackName

`func (o *HelmChartStackScaffoldInput) SetStackName(v string)`

SetStackName sets StackName field to given value.

### HasStackName

`func (o *HelmChartStackScaffoldInput) HasStackName() bool`

HasStackName returns a boolean if a field has been set.

### GetStackTitle

`func (o *HelmChartStackScaffoldInput) GetStackTitle() string`

GetStackTitle returns the StackTitle field if non-nil, zero value otherwise.

### GetStackTitleOk

`func (o *HelmChartStackScaffoldInput) GetStackTitleOk() (*string, bool)`

GetStackTitleOk returns a tuple with the StackTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackTitle

`func (o *HelmChartStackScaffoldInput) SetStackTitle(v string)`

SetStackTitle sets StackTitle field to given value.

### HasStackTitle

`func (o *HelmChartStackScaffoldInput) HasStackTitle() bool`

HasStackTitle returns a boolean if a field has been set.

### GetIcon

`func (o *HelmChartStackScaffoldInput) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *HelmChartStackScaffoldInput) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *HelmChartStackScaffoldInput) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *HelmChartStackScaffoldInput) HasIcon() bool`

HasIcon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


