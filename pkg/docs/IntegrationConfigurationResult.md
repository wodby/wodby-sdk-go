# IntegrationConfigurationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Integration** | [**Integration**](Integration.md) |  | 
**TaskId** | Pointer to **NullableInt32** |  | [optional] 
**Warnings** | **[]string** |  | 

## Methods

### NewIntegrationConfigurationResult

`func NewIntegrationConfigurationResult(integration Integration, warnings []string, ) *IntegrationConfigurationResult`

NewIntegrationConfigurationResult instantiates a new IntegrationConfigurationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationConfigurationResultWithDefaults

`func NewIntegrationConfigurationResultWithDefaults() *IntegrationConfigurationResult`

NewIntegrationConfigurationResultWithDefaults instantiates a new IntegrationConfigurationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegration

`func (o *IntegrationConfigurationResult) GetIntegration() Integration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *IntegrationConfigurationResult) GetIntegrationOk() (*Integration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *IntegrationConfigurationResult) SetIntegration(v Integration)`

SetIntegration sets Integration field to given value.


### GetTaskId

`func (o *IntegrationConfigurationResult) GetTaskId() int32`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *IntegrationConfigurationResult) GetTaskIdOk() (*int32, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *IntegrationConfigurationResult) SetTaskId(v int32)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *IntegrationConfigurationResult) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskIdNil

`func (o *IntegrationConfigurationResult) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *IntegrationConfigurationResult) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
### GetWarnings

`func (o *IntegrationConfigurationResult) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *IntegrationConfigurationResult) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *IntegrationConfigurationResult) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


