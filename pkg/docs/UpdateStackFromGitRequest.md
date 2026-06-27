# UpdateStackFromGitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitRef** | **string** |  | 
**GitRefType** | **string** |  | 

## Methods

### NewUpdateStackFromGitRequest

`func NewUpdateStackFromGitRequest(gitRef string, gitRefType string, ) *UpdateStackFromGitRequest`

NewUpdateStackFromGitRequest instantiates a new UpdateStackFromGitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStackFromGitRequestWithDefaults

`func NewUpdateStackFromGitRequestWithDefaults() *UpdateStackFromGitRequest`

NewUpdateStackFromGitRequestWithDefaults instantiates a new UpdateStackFromGitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGitRef

`func (o *UpdateStackFromGitRequest) GetGitRef() string`

GetGitRef returns the GitRef field if non-nil, zero value otherwise.

### GetGitRefOk

`func (o *UpdateStackFromGitRequest) GetGitRefOk() (*string, bool)`

GetGitRefOk returns a tuple with the GitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRef

`func (o *UpdateStackFromGitRequest) SetGitRef(v string)`

SetGitRef sets GitRef field to given value.


### GetGitRefType

`func (o *UpdateStackFromGitRequest) GetGitRefType() string`

GetGitRefType returns the GitRefType field if non-nil, zero value otherwise.

### GetGitRefTypeOk

`func (o *UpdateStackFromGitRequest) GetGitRefTypeOk() (*string, bool)`

GetGitRefTypeOk returns a tuple with the GitRefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRefType

`func (o *UpdateStackFromGitRequest) SetGitRefType(v string)`

SetGitRefType sets GitRefType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


