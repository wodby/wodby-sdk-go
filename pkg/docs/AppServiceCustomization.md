# AppServiceCustomization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Section** | **string** | The part of the app service where the customization is: general, resources, variables, helm, annotations, settings, links, integrations, volumes or configs. | 
**Name** | **string** |  | 
**Title** | **string** |  | 

## Methods

### NewAppServiceCustomization

`func NewAppServiceCustomization(section string, name string, title string, ) *AppServiceCustomization`

NewAppServiceCustomization instantiates a new AppServiceCustomization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceCustomizationWithDefaults

`func NewAppServiceCustomizationWithDefaults() *AppServiceCustomization`

NewAppServiceCustomizationWithDefaults instantiates a new AppServiceCustomization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSection

`func (o *AppServiceCustomization) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *AppServiceCustomization) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *AppServiceCustomization) SetSection(v string)`

SetSection sets Section field to given value.


### GetName

`func (o *AppServiceCustomization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppServiceCustomization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppServiceCustomization) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *AppServiceCustomization) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AppServiceCustomization) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AppServiceCustomization) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


