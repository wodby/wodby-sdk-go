# IntegrationProviderRevisionUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** |  | 
**Reasons** | **[]string** |  | 
**RemovedFields** | **[]string** |  | 
**CanDropRemovedFields** | **bool** |  | 
**CurrentRevision** | [**ProviderRevision**](ProviderRevision.md) |  | 
**TargetRevision** | [**ProviderRevision**](ProviderRevision.md) |  | 

## Methods

### NewIntegrationProviderRevisionUpgrade

`func NewIntegrationProviderRevisionUpgrade(state string, reasons []string, removedFields []string, canDropRemovedFields bool, currentRevision ProviderRevision, targetRevision ProviderRevision, ) *IntegrationProviderRevisionUpgrade`

NewIntegrationProviderRevisionUpgrade instantiates a new IntegrationProviderRevisionUpgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationProviderRevisionUpgradeWithDefaults

`func NewIntegrationProviderRevisionUpgradeWithDefaults() *IntegrationProviderRevisionUpgrade`

NewIntegrationProviderRevisionUpgradeWithDefaults instantiates a new IntegrationProviderRevisionUpgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *IntegrationProviderRevisionUpgrade) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IntegrationProviderRevisionUpgrade) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IntegrationProviderRevisionUpgrade) SetState(v string)`

SetState sets State field to given value.


### GetReasons

`func (o *IntegrationProviderRevisionUpgrade) GetReasons() []string`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *IntegrationProviderRevisionUpgrade) GetReasonsOk() (*[]string, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *IntegrationProviderRevisionUpgrade) SetReasons(v []string)`

SetReasons sets Reasons field to given value.


### GetRemovedFields

`func (o *IntegrationProviderRevisionUpgrade) GetRemovedFields() []string`

GetRemovedFields returns the RemovedFields field if non-nil, zero value otherwise.

### GetRemovedFieldsOk

`func (o *IntegrationProviderRevisionUpgrade) GetRemovedFieldsOk() (*[]string, bool)`

GetRemovedFieldsOk returns a tuple with the RemovedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedFields

`func (o *IntegrationProviderRevisionUpgrade) SetRemovedFields(v []string)`

SetRemovedFields sets RemovedFields field to given value.


### GetCanDropRemovedFields

`func (o *IntegrationProviderRevisionUpgrade) GetCanDropRemovedFields() bool`

GetCanDropRemovedFields returns the CanDropRemovedFields field if non-nil, zero value otherwise.

### GetCanDropRemovedFieldsOk

`func (o *IntegrationProviderRevisionUpgrade) GetCanDropRemovedFieldsOk() (*bool, bool)`

GetCanDropRemovedFieldsOk returns a tuple with the CanDropRemovedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDropRemovedFields

`func (o *IntegrationProviderRevisionUpgrade) SetCanDropRemovedFields(v bool)`

SetCanDropRemovedFields sets CanDropRemovedFields field to given value.


### GetCurrentRevision

`func (o *IntegrationProviderRevisionUpgrade) GetCurrentRevision() ProviderRevision`

GetCurrentRevision returns the CurrentRevision field if non-nil, zero value otherwise.

### GetCurrentRevisionOk

`func (o *IntegrationProviderRevisionUpgrade) GetCurrentRevisionOk() (*ProviderRevision, bool)`

GetCurrentRevisionOk returns a tuple with the CurrentRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRevision

`func (o *IntegrationProviderRevisionUpgrade) SetCurrentRevision(v ProviderRevision)`

SetCurrentRevision sets CurrentRevision field to given value.


### GetTargetRevision

`func (o *IntegrationProviderRevisionUpgrade) GetTargetRevision() ProviderRevision`

GetTargetRevision returns the TargetRevision field if non-nil, zero value otherwise.

### GetTargetRevisionOk

`func (o *IntegrationProviderRevisionUpgrade) GetTargetRevisionOk() (*ProviderRevision, bool)`

GetTargetRevisionOk returns a tuple with the TargetRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRevision

`func (o *IntegrationProviderRevisionUpgrade) SetTargetRevision(v ProviderRevision)`

SetTargetRevision sets TargetRevision field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


