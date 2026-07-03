# StackAutoUpdatePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | **string** |  | 
**VersionPolicy** | Pointer to [**StackAutoUpdateVersionPolicy**](StackAutoUpdateVersionPolicy.md) |  | [optional] 

## Methods

### NewStackAutoUpdatePolicy

`func NewStackAutoUpdatePolicy(scope string, ) *StackAutoUpdatePolicy`

NewStackAutoUpdatePolicy instantiates a new StackAutoUpdatePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackAutoUpdatePolicyWithDefaults

`func NewStackAutoUpdatePolicyWithDefaults() *StackAutoUpdatePolicy`

NewStackAutoUpdatePolicyWithDefaults instantiates a new StackAutoUpdatePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *StackAutoUpdatePolicy) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *StackAutoUpdatePolicy) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *StackAutoUpdatePolicy) SetScope(v string)`

SetScope sets Scope field to given value.


### GetVersionPolicy

`func (o *StackAutoUpdatePolicy) GetVersionPolicy() StackAutoUpdateVersionPolicy`

GetVersionPolicy returns the VersionPolicy field if non-nil, zero value otherwise.

### GetVersionPolicyOk

`func (o *StackAutoUpdatePolicy) GetVersionPolicyOk() (*StackAutoUpdateVersionPolicy, bool)`

GetVersionPolicyOk returns a tuple with the VersionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionPolicy

`func (o *StackAutoUpdatePolicy) SetVersionPolicy(v StackAutoUpdateVersionPolicy)`

SetVersionPolicy sets VersionPolicy field to given value.

### HasVersionPolicy

`func (o *StackAutoUpdatePolicy) HasVersionPolicy() bool`

HasVersionPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


