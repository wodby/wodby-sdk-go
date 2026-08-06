# UpdateOrgRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**DefaultTimeZone** | Pointer to **string** |  | [optional] 
**RegistryIntegrationId** | Pointer to **NullableInt32** |  | [optional] 
**CiIntegrationId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewUpdateOrgRequest

`func NewUpdateOrgRequest(title string, ) *UpdateOrgRequest`

NewUpdateOrgRequest instantiates a new UpdateOrgRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrgRequestWithDefaults

`func NewUpdateOrgRequestWithDefaults() *UpdateOrgRequest`

NewUpdateOrgRequestWithDefaults instantiates a new UpdateOrgRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *UpdateOrgRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateOrgRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateOrgRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDefaultTimeZone

`func (o *UpdateOrgRequest) GetDefaultTimeZone() string`

GetDefaultTimeZone returns the DefaultTimeZone field if non-nil, zero value otherwise.

### GetDefaultTimeZoneOk

`func (o *UpdateOrgRequest) GetDefaultTimeZoneOk() (*string, bool)`

GetDefaultTimeZoneOk returns a tuple with the DefaultTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTimeZone

`func (o *UpdateOrgRequest) SetDefaultTimeZone(v string)`

SetDefaultTimeZone sets DefaultTimeZone field to given value.

### HasDefaultTimeZone

`func (o *UpdateOrgRequest) HasDefaultTimeZone() bool`

HasDefaultTimeZone returns a boolean if a field has been set.

### GetRegistryIntegrationId

`func (o *UpdateOrgRequest) GetRegistryIntegrationId() int32`

GetRegistryIntegrationId returns the RegistryIntegrationId field if non-nil, zero value otherwise.

### GetRegistryIntegrationIdOk

`func (o *UpdateOrgRequest) GetRegistryIntegrationIdOk() (*int32, bool)`

GetRegistryIntegrationIdOk returns a tuple with the RegistryIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryIntegrationId

`func (o *UpdateOrgRequest) SetRegistryIntegrationId(v int32)`

SetRegistryIntegrationId sets RegistryIntegrationId field to given value.

### HasRegistryIntegrationId

`func (o *UpdateOrgRequest) HasRegistryIntegrationId() bool`

HasRegistryIntegrationId returns a boolean if a field has been set.

### SetRegistryIntegrationIdNil

`func (o *UpdateOrgRequest) SetRegistryIntegrationIdNil(b bool)`

 SetRegistryIntegrationIdNil sets the value for RegistryIntegrationId to be an explicit nil

### UnsetRegistryIntegrationId
`func (o *UpdateOrgRequest) UnsetRegistryIntegrationId()`

UnsetRegistryIntegrationId ensures that no value is present for RegistryIntegrationId, not even an explicit nil
### GetCiIntegrationId

`func (o *UpdateOrgRequest) GetCiIntegrationId() int32`

GetCiIntegrationId returns the CiIntegrationId field if non-nil, zero value otherwise.

### GetCiIntegrationIdOk

`func (o *UpdateOrgRequest) GetCiIntegrationIdOk() (*int32, bool)`

GetCiIntegrationIdOk returns a tuple with the CiIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiIntegrationId

`func (o *UpdateOrgRequest) SetCiIntegrationId(v int32)`

SetCiIntegrationId sets CiIntegrationId field to given value.

### HasCiIntegrationId

`func (o *UpdateOrgRequest) HasCiIntegrationId() bool`

HasCiIntegrationId returns a boolean if a field has been set.

### SetCiIntegrationIdNil

`func (o *UpdateOrgRequest) SetCiIntegrationIdNil(b bool)`

 SetCiIntegrationIdNil sets the value for CiIntegrationId to be an explicit nil

### UnsetCiIntegrationId
`func (o *UpdateOrgRequest) UnsetCiIntegrationId()`

UnsetCiIntegrationId ensures that no value is present for CiIntegrationId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


