# AppServicePod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Workload** | **string** |  | 
**KubeWorkloadName** | **string** |  | 
**Node** | Pointer to **NullableString** |  | [optional] 
**Phase** | **string** |  | 
**Ready** | **bool** |  | 
**Terminating** | **bool** |  | 
**Reason** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**CpuUsage** | Pointer to **NullableFloat32** | CPU cores in use. | [optional] 
**MemoryUsage** | Pointer to **NullableInt32** | Memory in use in bytes. | [optional] 
**CpuRequest** | Pointer to **NullableFloat32** |  | [optional] 
**MemoryRequest** | Pointer to **NullableInt32** |  | [optional] 
**CpuLimit** | Pointer to **NullableFloat32** |  | [optional] 
**MemoryLimit** | Pointer to **NullableInt32** |  | [optional] 
**Containers** | [**[]AppServicePodContainer**](AppServicePodContainer.md) |  | 

## Methods

### NewAppServicePod

`func NewAppServicePod(id string, name string, workload string, kubeWorkloadName string, phase string, ready bool, terminating bool, createdAt time.Time, containers []AppServicePodContainer, ) *AppServicePod`

NewAppServicePod instantiates a new AppServicePod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServicePodWithDefaults

`func NewAppServicePodWithDefaults() *AppServicePod`

NewAppServicePodWithDefaults instantiates a new AppServicePod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppServicePod) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppServicePod) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppServicePod) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AppServicePod) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppServicePod) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppServicePod) SetName(v string)`

SetName sets Name field to given value.


### GetWorkload

`func (o *AppServicePod) GetWorkload() string`

GetWorkload returns the Workload field if non-nil, zero value otherwise.

### GetWorkloadOk

`func (o *AppServicePod) GetWorkloadOk() (*string, bool)`

GetWorkloadOk returns a tuple with the Workload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkload

`func (o *AppServicePod) SetWorkload(v string)`

SetWorkload sets Workload field to given value.


### GetKubeWorkloadName

`func (o *AppServicePod) GetKubeWorkloadName() string`

GetKubeWorkloadName returns the KubeWorkloadName field if non-nil, zero value otherwise.

### GetKubeWorkloadNameOk

`func (o *AppServicePod) GetKubeWorkloadNameOk() (*string, bool)`

GetKubeWorkloadNameOk returns a tuple with the KubeWorkloadName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeWorkloadName

`func (o *AppServicePod) SetKubeWorkloadName(v string)`

SetKubeWorkloadName sets KubeWorkloadName field to given value.


### GetNode

`func (o *AppServicePod) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *AppServicePod) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *AppServicePod) SetNode(v string)`

SetNode sets Node field to given value.

### HasNode

`func (o *AppServicePod) HasNode() bool`

HasNode returns a boolean if a field has been set.

### SetNodeNil

`func (o *AppServicePod) SetNodeNil(b bool)`

 SetNodeNil sets the value for Node to be an explicit nil

### UnsetNode
`func (o *AppServicePod) UnsetNode()`

UnsetNode ensures that no value is present for Node, not even an explicit nil
### GetPhase

`func (o *AppServicePod) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *AppServicePod) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *AppServicePod) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetReady

`func (o *AppServicePod) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *AppServicePod) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *AppServicePod) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetTerminating

`func (o *AppServicePod) GetTerminating() bool`

GetTerminating returns the Terminating field if non-nil, zero value otherwise.

### GetTerminatingOk

`func (o *AppServicePod) GetTerminatingOk() (*bool, bool)`

GetTerminatingOk returns a tuple with the Terminating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminating

`func (o *AppServicePod) SetTerminating(v bool)`

SetTerminating sets Terminating field to given value.


### GetReason

`func (o *AppServicePod) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AppServicePod) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AppServicePod) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AppServicePod) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *AppServicePod) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *AppServicePod) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetMessage

`func (o *AppServicePod) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AppServicePod) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AppServicePod) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AppServicePod) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *AppServicePod) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *AppServicePod) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetCreatedAt

`func (o *AppServicePod) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppServicePod) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppServicePod) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCpuUsage

`func (o *AppServicePod) GetCpuUsage() float32`

GetCpuUsage returns the CpuUsage field if non-nil, zero value otherwise.

### GetCpuUsageOk

`func (o *AppServicePod) GetCpuUsageOk() (*float32, bool)`

GetCpuUsageOk returns a tuple with the CpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsage

`func (o *AppServicePod) SetCpuUsage(v float32)`

SetCpuUsage sets CpuUsage field to given value.

### HasCpuUsage

`func (o *AppServicePod) HasCpuUsage() bool`

HasCpuUsage returns a boolean if a field has been set.

### SetCpuUsageNil

`func (o *AppServicePod) SetCpuUsageNil(b bool)`

 SetCpuUsageNil sets the value for CpuUsage to be an explicit nil

### UnsetCpuUsage
`func (o *AppServicePod) UnsetCpuUsage()`

UnsetCpuUsage ensures that no value is present for CpuUsage, not even an explicit nil
### GetMemoryUsage

`func (o *AppServicePod) GetMemoryUsage() int32`

GetMemoryUsage returns the MemoryUsage field if non-nil, zero value otherwise.

### GetMemoryUsageOk

`func (o *AppServicePod) GetMemoryUsageOk() (*int32, bool)`

GetMemoryUsageOk returns a tuple with the MemoryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsage

`func (o *AppServicePod) SetMemoryUsage(v int32)`

SetMemoryUsage sets MemoryUsage field to given value.

### HasMemoryUsage

`func (o *AppServicePod) HasMemoryUsage() bool`

HasMemoryUsage returns a boolean if a field has been set.

### SetMemoryUsageNil

`func (o *AppServicePod) SetMemoryUsageNil(b bool)`

 SetMemoryUsageNil sets the value for MemoryUsage to be an explicit nil

### UnsetMemoryUsage
`func (o *AppServicePod) UnsetMemoryUsage()`

UnsetMemoryUsage ensures that no value is present for MemoryUsage, not even an explicit nil
### GetCpuRequest

`func (o *AppServicePod) GetCpuRequest() float32`

GetCpuRequest returns the CpuRequest field if non-nil, zero value otherwise.

### GetCpuRequestOk

`func (o *AppServicePod) GetCpuRequestOk() (*float32, bool)`

GetCpuRequestOk returns a tuple with the CpuRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuRequest

`func (o *AppServicePod) SetCpuRequest(v float32)`

SetCpuRequest sets CpuRequest field to given value.

### HasCpuRequest

`func (o *AppServicePod) HasCpuRequest() bool`

HasCpuRequest returns a boolean if a field has been set.

### SetCpuRequestNil

`func (o *AppServicePod) SetCpuRequestNil(b bool)`

 SetCpuRequestNil sets the value for CpuRequest to be an explicit nil

### UnsetCpuRequest
`func (o *AppServicePod) UnsetCpuRequest()`

UnsetCpuRequest ensures that no value is present for CpuRequest, not even an explicit nil
### GetMemoryRequest

`func (o *AppServicePod) GetMemoryRequest() int32`

GetMemoryRequest returns the MemoryRequest field if non-nil, zero value otherwise.

### GetMemoryRequestOk

`func (o *AppServicePod) GetMemoryRequestOk() (*int32, bool)`

GetMemoryRequestOk returns a tuple with the MemoryRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryRequest

`func (o *AppServicePod) SetMemoryRequest(v int32)`

SetMemoryRequest sets MemoryRequest field to given value.

### HasMemoryRequest

`func (o *AppServicePod) HasMemoryRequest() bool`

HasMemoryRequest returns a boolean if a field has been set.

### SetMemoryRequestNil

`func (o *AppServicePod) SetMemoryRequestNil(b bool)`

 SetMemoryRequestNil sets the value for MemoryRequest to be an explicit nil

### UnsetMemoryRequest
`func (o *AppServicePod) UnsetMemoryRequest()`

UnsetMemoryRequest ensures that no value is present for MemoryRequest, not even an explicit nil
### GetCpuLimit

`func (o *AppServicePod) GetCpuLimit() float32`

GetCpuLimit returns the CpuLimit field if non-nil, zero value otherwise.

### GetCpuLimitOk

`func (o *AppServicePod) GetCpuLimitOk() (*float32, bool)`

GetCpuLimitOk returns a tuple with the CpuLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLimit

`func (o *AppServicePod) SetCpuLimit(v float32)`

SetCpuLimit sets CpuLimit field to given value.

### HasCpuLimit

`func (o *AppServicePod) HasCpuLimit() bool`

HasCpuLimit returns a boolean if a field has been set.

### SetCpuLimitNil

`func (o *AppServicePod) SetCpuLimitNil(b bool)`

 SetCpuLimitNil sets the value for CpuLimit to be an explicit nil

### UnsetCpuLimit
`func (o *AppServicePod) UnsetCpuLimit()`

UnsetCpuLimit ensures that no value is present for CpuLimit, not even an explicit nil
### GetMemoryLimit

`func (o *AppServicePod) GetMemoryLimit() int32`

GetMemoryLimit returns the MemoryLimit field if non-nil, zero value otherwise.

### GetMemoryLimitOk

`func (o *AppServicePod) GetMemoryLimitOk() (*int32, bool)`

GetMemoryLimitOk returns a tuple with the MemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimit

`func (o *AppServicePod) SetMemoryLimit(v int32)`

SetMemoryLimit sets MemoryLimit field to given value.

### HasMemoryLimit

`func (o *AppServicePod) HasMemoryLimit() bool`

HasMemoryLimit returns a boolean if a field has been set.

### SetMemoryLimitNil

`func (o *AppServicePod) SetMemoryLimitNil(b bool)`

 SetMemoryLimitNil sets the value for MemoryLimit to be an explicit nil

### UnsetMemoryLimit
`func (o *AppServicePod) UnsetMemoryLimit()`

UnsetMemoryLimit ensures that no value is present for MemoryLimit, not even an explicit nil
### GetContainers

`func (o *AppServicePod) GetContainers() []AppServicePodContainer`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *AppServicePod) GetContainersOk() (*[]AppServicePodContainer, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *AppServicePod) SetContainers(v []AppServicePodContainer)`

SetContainers sets Containers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


