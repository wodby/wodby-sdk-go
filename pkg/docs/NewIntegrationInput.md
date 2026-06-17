# NewIntegrationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgID** | **int32** |  | 
**ProviderID** | **int32** |  | 
**Name** | **string** |  | 
**Title** | **string** |  | 
**Kinds** | **[]string** |  | 
**Auth** | Pointer to **NullableString** |  | [optional] 
**ProjectID** | Pointer to **NullableInt32** |  | [optional] 
**FieldsInput** | Pointer to [**[]FieldInput**](FieldInput.md) |  | [optional] 
**Scope** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNewIntegrationInput

`func NewNewIntegrationInput(orgID int32, providerID int32, name string, title string, kinds []string, ) *NewIntegrationInput`

NewNewIntegrationInput instantiates a new NewIntegrationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewIntegrationInputWithDefaults

`func NewNewIntegrationInputWithDefaults() *NewIntegrationInput`

NewNewIntegrationInputWithDefaults instantiates a new NewIntegrationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgID

`func (o *NewIntegrationInput) GetOrgID() int32`

GetOrgID returns the OrgID field if non-nil, zero value otherwise.

### GetOrgIDOk

`func (o *NewIntegrationInput) GetOrgIDOk() (*int32, bool)`

GetOrgIDOk returns a tuple with the OrgID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgID

`func (o *NewIntegrationInput) SetOrgID(v int32)`

SetOrgID sets OrgID field to given value.


### GetProviderID

`func (o *NewIntegrationInput) GetProviderID() int32`

GetProviderID returns the ProviderID field if non-nil, zero value otherwise.

### GetProviderIDOk

`func (o *NewIntegrationInput) GetProviderIDOk() (*int32, bool)`

GetProviderIDOk returns a tuple with the ProviderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderID

`func (o *NewIntegrationInput) SetProviderID(v int32)`

SetProviderID sets ProviderID field to given value.


### GetName

`func (o *NewIntegrationInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewIntegrationInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewIntegrationInput) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *NewIntegrationInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewIntegrationInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NewIntegrationInput) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetKinds

`func (o *NewIntegrationInput) GetKinds() []string`

GetKinds returns the Kinds field if non-nil, zero value otherwise.

### GetKindsOk

`func (o *NewIntegrationInput) GetKindsOk() (*[]string, bool)`

GetKindsOk returns a tuple with the Kinds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKinds

`func (o *NewIntegrationInput) SetKinds(v []string)`

SetKinds sets Kinds field to given value.


### GetAuth

`func (o *NewIntegrationInput) GetAuth() string`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *NewIntegrationInput) GetAuthOk() (*string, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *NewIntegrationInput) SetAuth(v string)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *NewIntegrationInput) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### SetAuthNil

`func (o *NewIntegrationInput) SetAuthNil(b bool)`

 SetAuthNil sets the value for Auth to be an explicit nil

### UnsetAuth
`func (o *NewIntegrationInput) UnsetAuth()`

UnsetAuth ensures that no value is present for Auth, not even an explicit nil
### GetProjectID

`func (o *NewIntegrationInput) GetProjectID() int32`

GetProjectID returns the ProjectID field if non-nil, zero value otherwise.

### GetProjectIDOk

`func (o *NewIntegrationInput) GetProjectIDOk() (*int32, bool)`

GetProjectIDOk returns a tuple with the ProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectID

`func (o *NewIntegrationInput) SetProjectID(v int32)`

SetProjectID sets ProjectID field to given value.

### HasProjectID

`func (o *NewIntegrationInput) HasProjectID() bool`

HasProjectID returns a boolean if a field has been set.

### SetProjectIDNil

`func (o *NewIntegrationInput) SetProjectIDNil(b bool)`

 SetProjectIDNil sets the value for ProjectID to be an explicit nil

### UnsetProjectID
`func (o *NewIntegrationInput) UnsetProjectID()`

UnsetProjectID ensures that no value is present for ProjectID, not even an explicit nil
### GetFieldsInput

`func (o *NewIntegrationInput) GetFieldsInput() []FieldInput`

GetFieldsInput returns the FieldsInput field if non-nil, zero value otherwise.

### GetFieldsInputOk

`func (o *NewIntegrationInput) GetFieldsInputOk() (*[]FieldInput, bool)`

GetFieldsInputOk returns a tuple with the FieldsInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsInput

`func (o *NewIntegrationInput) SetFieldsInput(v []FieldInput)`

SetFieldsInput sets FieldsInput field to given value.

### HasFieldsInput

`func (o *NewIntegrationInput) HasFieldsInput() bool`

HasFieldsInput returns a boolean if a field has been set.

### GetScope

`func (o *NewIntegrationInput) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *NewIntegrationInput) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *NewIntegrationInput) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *NewIntegrationInput) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *NewIntegrationInput) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *NewIntegrationInput) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


