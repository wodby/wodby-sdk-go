# TasksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]Task**](Task.md) |  | 
**TreeItems** | Pointer to [**[]TaskTreeItem**](TaskTreeItem.md) | Bounded current-page roots and authorized descendants for tree view, linked by parentId. | [optional] 
**TreeTruncated** | **bool** | True when treeItems omitted visible descendants after reaching the 250-item response limit. Always false for flat view. | 
**CanIncludeSystem** | **bool** | True when the current user may request operator-only system tasks. | 
**TotalCount** | **int32** |  | 
**NextPage** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewTasksResponse

`func NewTasksResponse(items []Task, treeTruncated bool, canIncludeSystem bool, totalCount int32, ) *TasksResponse`

NewTasksResponse instantiates a new TasksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTasksResponseWithDefaults

`func NewTasksResponseWithDefaults() *TasksResponse`

NewTasksResponseWithDefaults instantiates a new TasksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *TasksResponse) GetItems() []Task`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TasksResponse) GetItemsOk() (*[]Task, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TasksResponse) SetItems(v []Task)`

SetItems sets Items field to given value.


### GetTreeItems

`func (o *TasksResponse) GetTreeItems() []TaskTreeItem`

GetTreeItems returns the TreeItems field if non-nil, zero value otherwise.

### GetTreeItemsOk

`func (o *TasksResponse) GetTreeItemsOk() (*[]TaskTreeItem, bool)`

GetTreeItemsOk returns a tuple with the TreeItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreeItems

`func (o *TasksResponse) SetTreeItems(v []TaskTreeItem)`

SetTreeItems sets TreeItems field to given value.

### HasTreeItems

`func (o *TasksResponse) HasTreeItems() bool`

HasTreeItems returns a boolean if a field has been set.

### SetTreeItemsNil

`func (o *TasksResponse) SetTreeItemsNil(b bool)`

 SetTreeItemsNil sets the value for TreeItems to be an explicit nil

### UnsetTreeItems
`func (o *TasksResponse) UnsetTreeItems()`

UnsetTreeItems ensures that no value is present for TreeItems, not even an explicit nil
### GetTreeTruncated

`func (o *TasksResponse) GetTreeTruncated() bool`

GetTreeTruncated returns the TreeTruncated field if non-nil, zero value otherwise.

### GetTreeTruncatedOk

`func (o *TasksResponse) GetTreeTruncatedOk() (*bool, bool)`

GetTreeTruncatedOk returns a tuple with the TreeTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreeTruncated

`func (o *TasksResponse) SetTreeTruncated(v bool)`

SetTreeTruncated sets TreeTruncated field to given value.


### GetCanIncludeSystem

`func (o *TasksResponse) GetCanIncludeSystem() bool`

GetCanIncludeSystem returns the CanIncludeSystem field if non-nil, zero value otherwise.

### GetCanIncludeSystemOk

`func (o *TasksResponse) GetCanIncludeSystemOk() (*bool, bool)`

GetCanIncludeSystemOk returns a tuple with the CanIncludeSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanIncludeSystem

`func (o *TasksResponse) SetCanIncludeSystem(v bool)`

SetCanIncludeSystem sets CanIncludeSystem field to given value.


### GetTotalCount

`func (o *TasksResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *TasksResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *TasksResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetNextPage

`func (o *TasksResponse) GetNextPage() int32`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *TasksResponse) GetNextPageOk() (*int32, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *TasksResponse) SetNextPage(v int32)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *TasksResponse) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### SetNextPageNil

`func (o *TasksResponse) SetNextPageNil(b bool)`

 SetNextPageNil sets the value for NextPage to be an explicit nil

### UnsetNextPage
`func (o *TasksResponse) UnsetNextPage()`

UnsetNextPage ensures that no value is present for NextPage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


