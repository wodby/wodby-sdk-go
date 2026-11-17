# K3SServerNodesInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | Pointer to **[]string** | Node names from GET /cluster-k3s-nodes/{id}. Empty for the default nodes. | [optional] 

## Methods

### NewK3SServerNodesInput

`func NewK3SServerNodesInput() *K3SServerNodesInput`

NewK3SServerNodesInput instantiates a new K3SServerNodesInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK3SServerNodesInputWithDefaults

`func NewK3SServerNodesInputWithDefaults() *K3SServerNodesInput`

NewK3SServerNodesInputWithDefaults instantiates a new K3SServerNodesInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *K3SServerNodesInput) GetNodes() []string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *K3SServerNodesInput) GetNodesOk() (*[]string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *K3SServerNodesInput) SetNodes(v []string)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *K3SServerNodesInput) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


