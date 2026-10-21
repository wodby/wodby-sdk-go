# UpdateStackServiceRevisionsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StackServiceIds** | Pointer to **[]int32** | Root stack services to advance, by their IDs in the published revision or the draft. Their derivatives advance with them. Omit or leave empty to advance every outdated service. | [optional] 

## Methods

### NewUpdateStackServiceRevisionsInput

`func NewUpdateStackServiceRevisionsInput() *UpdateStackServiceRevisionsInput`

NewUpdateStackServiceRevisionsInput instantiates a new UpdateStackServiceRevisionsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStackServiceRevisionsInputWithDefaults

`func NewUpdateStackServiceRevisionsInputWithDefaults() *UpdateStackServiceRevisionsInput`

NewUpdateStackServiceRevisionsInputWithDefaults instantiates a new UpdateStackServiceRevisionsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStackServiceIds

`func (o *UpdateStackServiceRevisionsInput) GetStackServiceIds() []int32`

GetStackServiceIds returns the StackServiceIds field if non-nil, zero value otherwise.

### GetStackServiceIdsOk

`func (o *UpdateStackServiceRevisionsInput) GetStackServiceIdsOk() (*[]int32, bool)`

GetStackServiceIdsOk returns a tuple with the StackServiceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackServiceIds

`func (o *UpdateStackServiceRevisionsInput) SetStackServiceIds(v []int32)`

SetStackServiceIds sets StackServiceIds field to given value.

### HasStackServiceIds

`func (o *UpdateStackServiceRevisionsInput) HasStackServiceIds() bool`

HasStackServiceIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


