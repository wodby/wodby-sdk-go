# AppBuildsCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]AppBuild**](AppBuild.md) |  | 
**TaskId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewAppBuildsCreateResponse

`func NewAppBuildsCreateResponse(items []AppBuild, ) *AppBuildsCreateResponse`

NewAppBuildsCreateResponse instantiates a new AppBuildsCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppBuildsCreateResponseWithDefaults

`func NewAppBuildsCreateResponseWithDefaults() *AppBuildsCreateResponse`

NewAppBuildsCreateResponseWithDefaults instantiates a new AppBuildsCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *AppBuildsCreateResponse) GetItems() []AppBuild`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AppBuildsCreateResponse) GetItemsOk() (*[]AppBuild, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AppBuildsCreateResponse) SetItems(v []AppBuild)`

SetItems sets Items field to given value.


### GetTaskId

`func (o *AppBuildsCreateResponse) GetTaskId() int32`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *AppBuildsCreateResponse) GetTaskIdOk() (*int32, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *AppBuildsCreateResponse) SetTaskId(v int32)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *AppBuildsCreateResponse) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskIdNil

`func (o *AppBuildsCreateResponse) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *AppBuildsCreateResponse) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


