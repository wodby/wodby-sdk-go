# AppServiceCustomLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppServiceId** | **int32** |  | 
**Name** | **string** |  | 
**Values** | [**[]AppServiceCustomLinkValue**](AppServiceCustomLinkValue.md) | The values the link declares, in its order. | 

## Methods

### NewAppServiceCustomLink

`func NewAppServiceCustomLink(appServiceId int32, name string, values []AppServiceCustomLinkValue, ) *AppServiceCustomLink`

NewAppServiceCustomLink instantiates a new AppServiceCustomLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceCustomLinkWithDefaults

`func NewAppServiceCustomLinkWithDefaults() *AppServiceCustomLink`

NewAppServiceCustomLinkWithDefaults instantiates a new AppServiceCustomLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppServiceId

`func (o *AppServiceCustomLink) GetAppServiceId() int32`

GetAppServiceId returns the AppServiceId field if non-nil, zero value otherwise.

### GetAppServiceIdOk

`func (o *AppServiceCustomLink) GetAppServiceIdOk() (*int32, bool)`

GetAppServiceIdOk returns a tuple with the AppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceId

`func (o *AppServiceCustomLink) SetAppServiceId(v int32)`

SetAppServiceId sets AppServiceId field to given value.


### GetName

`func (o *AppServiceCustomLink) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppServiceCustomLink) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppServiceCustomLink) SetName(v string)`

SetName sets Name field to given value.


### GetValues

`func (o *AppServiceCustomLink) GetValues() []AppServiceCustomLinkValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *AppServiceCustomLink) GetValuesOk() (*[]AppServiceCustomLinkValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *AppServiceCustomLink) SetValues(v []AppServiceCustomLinkValue)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


