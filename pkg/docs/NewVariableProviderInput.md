# NewVariableProviderInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. | [optional] 
**ProjectId** | Pointer to **NullableInt32** |  | [optional] 
**Name** | **string** |  | 
**Title** | **string** |  | 
**Fields** | [**[]NewVariableProviderFieldInput**](NewVariableProviderFieldInput.md) |  | 

## Methods

### NewNewVariableProviderInput

`func NewNewVariableProviderInput(name string, title string, fields []NewVariableProviderFieldInput, ) *NewVariableProviderInput`

NewNewVariableProviderInput instantiates a new NewVariableProviderInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewVariableProviderInputWithDefaults

`func NewNewVariableProviderInputWithDefaults() *NewVariableProviderInput`

NewNewVariableProviderInputWithDefaults instantiates a new NewVariableProviderInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *NewVariableProviderInput) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *NewVariableProviderInput) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *NewVariableProviderInput) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *NewVariableProviderInput) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProjectId

`func (o *NewVariableProviderInput) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *NewVariableProviderInput) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *NewVariableProviderInput) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *NewVariableProviderInput) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *NewVariableProviderInput) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *NewVariableProviderInput) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetName

`func (o *NewVariableProviderInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewVariableProviderInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewVariableProviderInput) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *NewVariableProviderInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewVariableProviderInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NewVariableProviderInput) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetFields

`func (o *NewVariableProviderInput) GetFields() []NewVariableProviderFieldInput`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *NewVariableProviderInput) GetFieldsOk() (*[]NewVariableProviderFieldInput, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *NewVariableProviderInput) SetFields(v []NewVariableProviderFieldInput)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


