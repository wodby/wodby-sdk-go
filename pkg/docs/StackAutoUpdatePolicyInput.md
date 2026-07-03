# StackAutoUpdatePolicyInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | Pointer to **NullableString** |  | [optional] 
**IncludeDisabled** | Pointer to **NullableBool** |  | [optional] 
**VersionPolicy** | Pointer to [**StackAutoUpdateVersionPolicyInput**](StackAutoUpdateVersionPolicyInput.md) |  | [optional] 

## Methods

### NewStackAutoUpdatePolicyInput

`func NewStackAutoUpdatePolicyInput() *StackAutoUpdatePolicyInput`

NewStackAutoUpdatePolicyInput instantiates a new StackAutoUpdatePolicyInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackAutoUpdatePolicyInputWithDefaults

`func NewStackAutoUpdatePolicyInputWithDefaults() *StackAutoUpdatePolicyInput`

NewStackAutoUpdatePolicyInputWithDefaults instantiates a new StackAutoUpdatePolicyInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *StackAutoUpdatePolicyInput) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *StackAutoUpdatePolicyInput) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *StackAutoUpdatePolicyInput) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *StackAutoUpdatePolicyInput) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *StackAutoUpdatePolicyInput) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *StackAutoUpdatePolicyInput) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetIncludeDisabled

`func (o *StackAutoUpdatePolicyInput) GetIncludeDisabled() bool`

GetIncludeDisabled returns the IncludeDisabled field if non-nil, zero value otherwise.

### GetIncludeDisabledOk

`func (o *StackAutoUpdatePolicyInput) GetIncludeDisabledOk() (*bool, bool)`

GetIncludeDisabledOk returns a tuple with the IncludeDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDisabled

`func (o *StackAutoUpdatePolicyInput) SetIncludeDisabled(v bool)`

SetIncludeDisabled sets IncludeDisabled field to given value.

### HasIncludeDisabled

`func (o *StackAutoUpdatePolicyInput) HasIncludeDisabled() bool`

HasIncludeDisabled returns a boolean if a field has been set.

### SetIncludeDisabledNil

`func (o *StackAutoUpdatePolicyInput) SetIncludeDisabledNil(b bool)`

 SetIncludeDisabledNil sets the value for IncludeDisabled to be an explicit nil

### UnsetIncludeDisabled
`func (o *StackAutoUpdatePolicyInput) UnsetIncludeDisabled()`

UnsetIncludeDisabled ensures that no value is present for IncludeDisabled, not even an explicit nil
### GetVersionPolicy

`func (o *StackAutoUpdatePolicyInput) GetVersionPolicy() StackAutoUpdateVersionPolicyInput`

GetVersionPolicy returns the VersionPolicy field if non-nil, zero value otherwise.

### GetVersionPolicyOk

`func (o *StackAutoUpdatePolicyInput) GetVersionPolicyOk() (*StackAutoUpdateVersionPolicyInput, bool)`

GetVersionPolicyOk returns a tuple with the VersionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionPolicy

`func (o *StackAutoUpdatePolicyInput) SetVersionPolicy(v StackAutoUpdateVersionPolicyInput)`

SetVersionPolicy sets VersionPolicy field to given value.

### HasVersionPolicy

`func (o *StackAutoUpdatePolicyInput) HasVersionPolicy() bool`

HasVersionPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


