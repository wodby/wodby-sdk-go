# K3SClusterNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Role** | **string** | server runs the K3s control plane; agent is a worker node joined with the join command. | 
**Ready** | **bool** |  | 
**Schedulable** | **bool** | Whether the node accepts new pods. A node being removed or upgraded doesn&#39;t. | 
**Version** | **string** | The node&#39;s K3s release, such as v1.36.6+k3s1. | 
**InternalIp** | Pointer to **string** | The private address the node uses for cluster traffic. | [optional] 
**ExternalIp** | Pointer to **string** |  | [optional] 
**LocalVolumes** | **[]string** | Persistent volumes stored on the node, as namespace/name of their claims. Their data stays on that server, so removing the node loses it. | 
**Capacity** | [**NodeCapacity**](NodeCapacity.md) |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewK3SClusterNode

`func NewK3SClusterNode(name string, role string, ready bool, schedulable bool, version string, localVolumes []string, capacity NodeCapacity, createdAt time.Time, ) *K3SClusterNode`

NewK3SClusterNode instantiates a new K3SClusterNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK3SClusterNodeWithDefaults

`func NewK3SClusterNodeWithDefaults() *K3SClusterNode`

NewK3SClusterNodeWithDefaults instantiates a new K3SClusterNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *K3SClusterNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *K3SClusterNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *K3SClusterNode) SetName(v string)`

SetName sets Name field to given value.


### GetRole

`func (o *K3SClusterNode) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *K3SClusterNode) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *K3SClusterNode) SetRole(v string)`

SetRole sets Role field to given value.


### GetReady

`func (o *K3SClusterNode) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *K3SClusterNode) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *K3SClusterNode) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetSchedulable

`func (o *K3SClusterNode) GetSchedulable() bool`

GetSchedulable returns the Schedulable field if non-nil, zero value otherwise.

### GetSchedulableOk

`func (o *K3SClusterNode) GetSchedulableOk() (*bool, bool)`

GetSchedulableOk returns a tuple with the Schedulable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulable

`func (o *K3SClusterNode) SetSchedulable(v bool)`

SetSchedulable sets Schedulable field to given value.


### GetVersion

`func (o *K3SClusterNode) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *K3SClusterNode) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *K3SClusterNode) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetInternalIp

`func (o *K3SClusterNode) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *K3SClusterNode) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *K3SClusterNode) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *K3SClusterNode) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### GetExternalIp

`func (o *K3SClusterNode) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *K3SClusterNode) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *K3SClusterNode) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *K3SClusterNode) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### GetLocalVolumes

`func (o *K3SClusterNode) GetLocalVolumes() []string`

GetLocalVolumes returns the LocalVolumes field if non-nil, zero value otherwise.

### GetLocalVolumesOk

`func (o *K3SClusterNode) GetLocalVolumesOk() (*[]string, bool)`

GetLocalVolumesOk returns a tuple with the LocalVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalVolumes

`func (o *K3SClusterNode) SetLocalVolumes(v []string)`

SetLocalVolumes sets LocalVolumes field to given value.


### GetCapacity

`func (o *K3SClusterNode) GetCapacity() NodeCapacity`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *K3SClusterNode) GetCapacityOk() (*NodeCapacity, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *K3SClusterNode) SetCapacity(v NodeCapacity)`

SetCapacity sets Capacity field to given value.


### GetCreatedAt

`func (o *K3SClusterNode) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *K3SClusterNode) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *K3SClusterNode) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


