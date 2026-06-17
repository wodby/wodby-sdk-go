# UpdateIntegrationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Name** | **string** |  | 
**Kinds** | **[]string** |  | 
**Scope** | Pointer to **NullableString** |  | [optional] 
**FieldsInput** | Pointer to [**[]FieldInput**](FieldInput.md) |  | [optional] 

## Methods

### NewUpdateIntegrationInput

`func NewUpdateIntegrationInput(title string, name string, kinds []string, ) *UpdateIntegrationInput`

NewUpdateIntegrationInput instantiates a new UpdateIntegrationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIntegrationInputWithDefaults

`func NewUpdateIntegrationInputWithDefaults() *UpdateIntegrationInput`

NewUpdateIntegrationInputWithDefaults instantiates a new UpdateIntegrationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *UpdateIntegrationInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateIntegrationInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateIntegrationInput) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetName

`func (o *UpdateIntegrationInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateIntegrationInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateIntegrationInput) SetName(v string)`

SetName sets Name field to given value.


### GetKinds

`func (o *UpdateIntegrationInput) GetKinds() []string`

GetKinds returns the Kinds field if non-nil, zero value otherwise.

### GetKindsOk

`func (o *UpdateIntegrationInput) GetKindsOk() (*[]string, bool)`

GetKindsOk returns a tuple with the Kinds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKinds

`func (o *UpdateIntegrationInput) SetKinds(v []string)`

SetKinds sets Kinds field to given value.


### GetScope

`func (o *UpdateIntegrationInput) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *UpdateIntegrationInput) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *UpdateIntegrationInput) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *UpdateIntegrationInput) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *UpdateIntegrationInput) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *UpdateIntegrationInput) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetFieldsInput

`func (o *UpdateIntegrationInput) GetFieldsInput() []FieldInput`

GetFieldsInput returns the FieldsInput field if non-nil, zero value otherwise.

### GetFieldsInputOk

`func (o *UpdateIntegrationInput) GetFieldsInputOk() (*[]FieldInput, bool)`

GetFieldsInputOk returns a tuple with the FieldsInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsInput

`func (o *UpdateIntegrationInput) SetFieldsInput(v []FieldInput)`

SetFieldsInput sets FieldsInput field to given value.

### HasFieldsInput

`func (o *UpdateIntegrationInput) HasFieldsInput() bool`

HasFieldsInput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


