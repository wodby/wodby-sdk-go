# AppRouteSettingInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | [**AppRouteSettingName**](AppRouteSettingName.md) |  | 
**Value** | **string** |  | 

## Methods

### NewAppRouteSettingInput

`func NewAppRouteSettingInput(name AppRouteSettingName, value string, ) *AppRouteSettingInput`

NewAppRouteSettingInput instantiates a new AppRouteSettingInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRouteSettingInputWithDefaults

`func NewAppRouteSettingInputWithDefaults() *AppRouteSettingInput`

NewAppRouteSettingInputWithDefaults instantiates a new AppRouteSettingInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AppRouteSettingInput) GetName() AppRouteSettingName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppRouteSettingInput) GetNameOk() (*AppRouteSettingName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppRouteSettingInput) SetName(v AppRouteSettingName)`

SetName sets Name field to given value.


### GetValue

`func (o *AppRouteSettingInput) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AppRouteSettingInput) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AppRouteSettingInput) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


