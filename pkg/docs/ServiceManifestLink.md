# ServiceManifestLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Title** | **string** |  | 
**Required** | **bool** |  | 
**Selectors** | Pointer to [**[]ServiceManifestLinkSelector**](ServiceManifestLinkSelector.md) | The services the link may point at. A target must match one selector. | [optional] 
**Custom** | **bool** | Whether the link accepts a custom target, whose values a user enters instead of choosing a target service. | 
**CrossApp** | **bool** | Whether the link may point at a service of another app environment that accepts it. Across environments the link reads only the target&#39;s exports. | 
**Values** | [**[]ServiceManifestLinkValue**](ServiceManifestLinkValue.md) | The values the link hands to the service, whatever it points at. | 

## Methods

### NewServiceManifestLink

`func NewServiceManifestLink(name string, title string, required bool, custom bool, crossApp bool, values []ServiceManifestLinkValue, ) *ServiceManifestLink`

NewServiceManifestLink instantiates a new ServiceManifestLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceManifestLinkWithDefaults

`func NewServiceManifestLinkWithDefaults() *ServiceManifestLink`

NewServiceManifestLinkWithDefaults instantiates a new ServiceManifestLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServiceManifestLink) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceManifestLink) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceManifestLink) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *ServiceManifestLink) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ServiceManifestLink) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ServiceManifestLink) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetRequired

`func (o *ServiceManifestLink) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ServiceManifestLink) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ServiceManifestLink) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetSelectors

`func (o *ServiceManifestLink) GetSelectors() []ServiceManifestLinkSelector`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *ServiceManifestLink) GetSelectorsOk() (*[]ServiceManifestLinkSelector, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *ServiceManifestLink) SetSelectors(v []ServiceManifestLinkSelector)`

SetSelectors sets Selectors field to given value.

### HasSelectors

`func (o *ServiceManifestLink) HasSelectors() bool`

HasSelectors returns a boolean if a field has been set.

### GetCustom

`func (o *ServiceManifestLink) GetCustom() bool`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *ServiceManifestLink) GetCustomOk() (*bool, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *ServiceManifestLink) SetCustom(v bool)`

SetCustom sets Custom field to given value.


### GetCrossApp

`func (o *ServiceManifestLink) GetCrossApp() bool`

GetCrossApp returns the CrossApp field if non-nil, zero value otherwise.

### GetCrossAppOk

`func (o *ServiceManifestLink) GetCrossAppOk() (*bool, bool)`

GetCrossAppOk returns a tuple with the CrossApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossApp

`func (o *ServiceManifestLink) SetCrossApp(v bool)`

SetCrossApp sets CrossApp field to given value.


### GetValues

`func (o *ServiceManifestLink) GetValues() []ServiceManifestLinkValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ServiceManifestLink) GetValuesOk() (*[]ServiceManifestLinkValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ServiceManifestLink) SetValues(v []ServiceManifestLinkValue)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


