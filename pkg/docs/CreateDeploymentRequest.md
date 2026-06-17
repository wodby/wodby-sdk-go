# CreateDeploymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Services** | [**[]AppServiceDeploymentRequest**](AppServiceDeploymentRequest.md) |  | 
**SkipRollback** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewCreateDeploymentRequest

`func NewCreateDeploymentRequest(services []AppServiceDeploymentRequest, ) *CreateDeploymentRequest`

NewCreateDeploymentRequest instantiates a new CreateDeploymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDeploymentRequestWithDefaults

`func NewCreateDeploymentRequestWithDefaults() *CreateDeploymentRequest`

NewCreateDeploymentRequestWithDefaults instantiates a new CreateDeploymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServices

`func (o *CreateDeploymentRequest) GetServices() []AppServiceDeploymentRequest`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *CreateDeploymentRequest) GetServicesOk() (*[]AppServiceDeploymentRequest, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *CreateDeploymentRequest) SetServices(v []AppServiceDeploymentRequest)`

SetServices sets Services field to given value.


### GetSkipRollback

`func (o *CreateDeploymentRequest) GetSkipRollback() bool`

GetSkipRollback returns the SkipRollback field if non-nil, zero value otherwise.

### GetSkipRollbackOk

`func (o *CreateDeploymentRequest) GetSkipRollbackOk() (*bool, bool)`

GetSkipRollbackOk returns a tuple with the SkipRollback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipRollback

`func (o *CreateDeploymentRequest) SetSkipRollback(v bool)`

SetSkipRollback sets SkipRollback field to given value.

### HasSkipRollback

`func (o *CreateDeploymentRequest) HasSkipRollback() bool`

HasSkipRollback returns a boolean if a field has been set.

### SetSkipRollbackNil

`func (o *CreateDeploymentRequest) SetSkipRollbackNil(b bool)`

 SetSkipRollbackNil sets the value for SkipRollback to be an explicit nil

### UnsetSkipRollback
`func (o *CreateDeploymentRequest) UnsetSkipRollback()`

UnsetSkipRollback ensures that no value is present for SkipRollback, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


