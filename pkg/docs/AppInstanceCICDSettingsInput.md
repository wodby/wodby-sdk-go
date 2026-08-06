# AppInstanceCICDSettingsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CiIntegrationId** | **int32** | CI integration ID. Set to zero to use the built-in Wodby CI service. | 
**RegistryIntegrationId** | **int32** | Registry integration ID. Set to zero to use the built-in Wodby registry service. | 

## Methods

### NewAppInstanceCICDSettingsInput

`func NewAppInstanceCICDSettingsInput(ciIntegrationId int32, registryIntegrationId int32, ) *AppInstanceCICDSettingsInput`

NewAppInstanceCICDSettingsInput instantiates a new AppInstanceCICDSettingsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppInstanceCICDSettingsInputWithDefaults

`func NewAppInstanceCICDSettingsInputWithDefaults() *AppInstanceCICDSettingsInput`

NewAppInstanceCICDSettingsInputWithDefaults instantiates a new AppInstanceCICDSettingsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCiIntegrationId

`func (o *AppInstanceCICDSettingsInput) GetCiIntegrationId() int32`

GetCiIntegrationId returns the CiIntegrationId field if non-nil, zero value otherwise.

### GetCiIntegrationIdOk

`func (o *AppInstanceCICDSettingsInput) GetCiIntegrationIdOk() (*int32, bool)`

GetCiIntegrationIdOk returns a tuple with the CiIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiIntegrationId

`func (o *AppInstanceCICDSettingsInput) SetCiIntegrationId(v int32)`

SetCiIntegrationId sets CiIntegrationId field to given value.


### GetRegistryIntegrationId

`func (o *AppInstanceCICDSettingsInput) GetRegistryIntegrationId() int32`

GetRegistryIntegrationId returns the RegistryIntegrationId field if non-nil, zero value otherwise.

### GetRegistryIntegrationIdOk

`func (o *AppInstanceCICDSettingsInput) GetRegistryIntegrationIdOk() (*int32, bool)`

GetRegistryIntegrationIdOk returns a tuple with the RegistryIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryIntegrationId

`func (o *AppInstanceCICDSettingsInput) SetRegistryIntegrationId(v int32)`

SetRegistryIntegrationId sets RegistryIntegrationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


