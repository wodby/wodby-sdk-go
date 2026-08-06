# TaskTreeItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Task** | [**Task**](Task.md) |  | 
**ParentId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewTaskTreeItem

`func NewTaskTreeItem(task Task, ) *TaskTreeItem`

NewTaskTreeItem instantiates a new TaskTreeItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskTreeItemWithDefaults

`func NewTaskTreeItemWithDefaults() *TaskTreeItem`

NewTaskTreeItemWithDefaults instantiates a new TaskTreeItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTask

`func (o *TaskTreeItem) GetTask() Task`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *TaskTreeItem) GetTaskOk() (*Task, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *TaskTreeItem) SetTask(v Task)`

SetTask sets Task field to given value.


### GetParentId

`func (o *TaskTreeItem) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *TaskTreeItem) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *TaskTreeItem) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *TaskTreeItem) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *TaskTreeItem) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *TaskTreeItem) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


