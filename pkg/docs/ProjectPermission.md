# ProjectPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Title** | **string** |  | 
**Description** | **string** |  | 
**Scope** | **string** |  | 
**AlwaysGranted** | **bool** | Every project role grants it, so it cannot be removed from a custom role. | 

## Methods

### NewProjectPermission

`func NewProjectPermission(name string, title string, description string, scope string, alwaysGranted bool, ) *ProjectPermission`

NewProjectPermission instantiates a new ProjectPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectPermissionWithDefaults

`func NewProjectPermissionWithDefaults() *ProjectPermission`

NewProjectPermissionWithDefaults instantiates a new ProjectPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProjectPermission) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectPermission) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectPermission) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *ProjectPermission) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProjectPermission) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProjectPermission) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *ProjectPermission) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectPermission) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectPermission) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetScope

`func (o *ProjectPermission) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ProjectPermission) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ProjectPermission) SetScope(v string)`

SetScope sets Scope field to given value.


### GetAlwaysGranted

`func (o *ProjectPermission) GetAlwaysGranted() bool`

GetAlwaysGranted returns the AlwaysGranted field if non-nil, zero value otherwise.

### GetAlwaysGrantedOk

`func (o *ProjectPermission) GetAlwaysGrantedOk() (*bool, bool)`

GetAlwaysGrantedOk returns a tuple with the AlwaysGranted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysGranted

`func (o *ProjectPermission) SetAlwaysGranted(v bool)`

SetAlwaysGranted sets AlwaysGranted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


