# NewIntegrationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **int32** |  | 
**ProviderId** | **int32** |  | 
**Name** | **string** |  | 
**Title** | **string** |  | 
**Kinds** | **[]string** |  | 
**Auth** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | Pointer to **NullableInt32** |  | [optional] 
**FieldsInput** | Pointer to [**[]FieldInput**](FieldInput.md) |  | [optional] 
**Scope** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNewIntegrationInput

`func NewNewIntegrationInput(orgId int32, providerId int32, name string, title string, kinds []string, ) *NewIntegrationInput`

NewNewIntegrationInput instantiates a new NewIntegrationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewIntegrationInputWithDefaults

`func NewNewIntegrationInputWithDefaults() *NewIntegrationInput`

NewNewIntegrationInputWithDefaults instantiates a new NewIntegrationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *NewIntegrationInput) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *NewIntegrationInput) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *NewIntegrationInput) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.


### GetProviderId

`func (o *NewIntegrationInput) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *NewIntegrationInput) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *NewIntegrationInput) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.


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
### GetProjectId

`func (o *NewIntegrationInput) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *NewIntegrationInput) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *NewIntegrationInput) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *NewIntegrationInput) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *NewIntegrationInput) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *NewIntegrationInput) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
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


