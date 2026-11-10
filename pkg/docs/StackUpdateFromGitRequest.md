# StackUpdateFromGitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitRef** | **string** |  | 
**GitRefType** | **string** |  | 
**ServiceUpdates** | Pointer to **NullableString** | What this update does with existing stack services whose stack.yml reference has no version, instead of the stack&#39;s gitServiceUpdates setting: latest, auto_update_settings or keep. | [optional] 

## Methods

### NewStackUpdateFromGitRequest

`func NewStackUpdateFromGitRequest(gitRef string, gitRefType string, ) *StackUpdateFromGitRequest`

NewStackUpdateFromGitRequest instantiates a new StackUpdateFromGitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackUpdateFromGitRequestWithDefaults

`func NewStackUpdateFromGitRequestWithDefaults() *StackUpdateFromGitRequest`

NewStackUpdateFromGitRequestWithDefaults instantiates a new StackUpdateFromGitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGitRef

`func (o *StackUpdateFromGitRequest) GetGitRef() string`

GetGitRef returns the GitRef field if non-nil, zero value otherwise.

### GetGitRefOk

`func (o *StackUpdateFromGitRequest) GetGitRefOk() (*string, bool)`

GetGitRefOk returns a tuple with the GitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRef

`func (o *StackUpdateFromGitRequest) SetGitRef(v string)`

SetGitRef sets GitRef field to given value.


### GetGitRefType

`func (o *StackUpdateFromGitRequest) GetGitRefType() string`

GetGitRefType returns the GitRefType field if non-nil, zero value otherwise.

### GetGitRefTypeOk

`func (o *StackUpdateFromGitRequest) GetGitRefTypeOk() (*string, bool)`

GetGitRefTypeOk returns a tuple with the GitRefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRefType

`func (o *StackUpdateFromGitRequest) SetGitRefType(v string)`

SetGitRefType sets GitRefType field to given value.


### GetServiceUpdates

`func (o *StackUpdateFromGitRequest) GetServiceUpdates() string`

GetServiceUpdates returns the ServiceUpdates field if non-nil, zero value otherwise.

### GetServiceUpdatesOk

`func (o *StackUpdateFromGitRequest) GetServiceUpdatesOk() (*string, bool)`

GetServiceUpdatesOk returns a tuple with the ServiceUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUpdates

`func (o *StackUpdateFromGitRequest) SetServiceUpdates(v string)`

SetServiceUpdates sets ServiceUpdates field to given value.

### HasServiceUpdates

`func (o *StackUpdateFromGitRequest) HasServiceUpdates() bool`

HasServiceUpdates returns a boolean if a field has been set.

### SetServiceUpdatesNil

`func (o *StackUpdateFromGitRequest) SetServiceUpdatesNil(b bool)`

 SetServiceUpdatesNil sets the value for ServiceUpdates to be an explicit nil

### UnsetServiceUpdates
`func (o *StackUpdateFromGitRequest) UnsetServiceUpdates()`

UnsetServiceUpdates ensures that no value is present for ServiceUpdates, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


