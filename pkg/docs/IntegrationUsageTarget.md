# IntegrationUsageTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Type** | **string** |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Title** | **string** |  | 
**Icon** | Pointer to **NullableString** |  | [optional] 
**Parent** | Pointer to [**IntegrationUsageTarget**](IntegrationUsageTarget.md) | Resource that contains the target, such as the app environment of an app service. | [optional] 

## Methods

### NewIntegrationUsageTarget

`func NewIntegrationUsageTarget(id int32, type_ string, title string, ) *IntegrationUsageTarget`

NewIntegrationUsageTarget instantiates a new IntegrationUsageTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationUsageTargetWithDefaults

`func NewIntegrationUsageTargetWithDefaults() *IntegrationUsageTarget`

NewIntegrationUsageTargetWithDefaults instantiates a new IntegrationUsageTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntegrationUsageTarget) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationUsageTarget) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationUsageTarget) SetId(v int32)`

SetId sets Id field to given value.


### GetType

`func (o *IntegrationUsageTarget) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegrationUsageTarget) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntegrationUsageTarget) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *IntegrationUsageTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationUsageTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationUsageTarget) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IntegrationUsageTarget) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IntegrationUsageTarget) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IntegrationUsageTarget) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTitle

`func (o *IntegrationUsageTarget) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IntegrationUsageTarget) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IntegrationUsageTarget) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetIcon

`func (o *IntegrationUsageTarget) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *IntegrationUsageTarget) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *IntegrationUsageTarget) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *IntegrationUsageTarget) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *IntegrationUsageTarget) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *IntegrationUsageTarget) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetParent

`func (o *IntegrationUsageTarget) GetParent() IntegrationUsageTarget`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *IntegrationUsageTarget) GetParentOk() (*IntegrationUsageTarget, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *IntegrationUsageTarget) SetParent(v IntegrationUsageTarget)`

SetParent sets Parent field to given value.

### HasParent

`func (o *IntegrationUsageTarget) HasParent() bool`

HasParent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


