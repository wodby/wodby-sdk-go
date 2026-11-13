# RemoveK3SClusterNodeInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeName** | **string** |  | 
**Force** | Pointer to **bool** | Remove the node even though it stores persistent volumes, which loses their data. | [optional] 

## Methods

### NewRemoveK3SClusterNodeInput

`func NewRemoveK3SClusterNodeInput(nodeName string, ) *RemoveK3SClusterNodeInput`

NewRemoveK3SClusterNodeInput instantiates a new RemoveK3SClusterNodeInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveK3SClusterNodeInputWithDefaults

`func NewRemoveK3SClusterNodeInputWithDefaults() *RemoveK3SClusterNodeInput`

NewRemoveK3SClusterNodeInputWithDefaults instantiates a new RemoveK3SClusterNodeInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeName

`func (o *RemoveK3SClusterNodeInput) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *RemoveK3SClusterNodeInput) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *RemoveK3SClusterNodeInput) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.


### GetForce

`func (o *RemoveK3SClusterNodeInput) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *RemoveK3SClusterNodeInput) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *RemoveK3SClusterNodeInput) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *RemoveK3SClusterNodeInput) HasForce() bool`

HasForce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


