# UpdateOrgRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**RegistryIntegrationID** | Pointer to **NullableInt32** |  | [optional] 
**CiIntegrationID** | Pointer to **NullableInt32** |  | [optional] 

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


### GetRegistryIntegrationID

`func (o *UpdateOrgRequest) GetRegistryIntegrationID() int32`

GetRegistryIntegrationID returns the RegistryIntegrationID field if non-nil, zero value otherwise.

### GetRegistryIntegrationIDOk

`func (o *UpdateOrgRequest) GetRegistryIntegrationIDOk() (*int32, bool)`

GetRegistryIntegrationIDOk returns a tuple with the RegistryIntegrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryIntegrationID

`func (o *UpdateOrgRequest) SetRegistryIntegrationID(v int32)`

SetRegistryIntegrationID sets RegistryIntegrationID field to given value.

### HasRegistryIntegrationID

`func (o *UpdateOrgRequest) HasRegistryIntegrationID() bool`

HasRegistryIntegrationID returns a boolean if a field has been set.

### SetRegistryIntegrationIDNil

`func (o *UpdateOrgRequest) SetRegistryIntegrationIDNil(b bool)`

 SetRegistryIntegrationIDNil sets the value for RegistryIntegrationID to be an explicit nil

### UnsetRegistryIntegrationID
`func (o *UpdateOrgRequest) UnsetRegistryIntegrationID()`

UnsetRegistryIntegrationID ensures that no value is present for RegistryIntegrationID, not even an explicit nil
### GetCiIntegrationID

`func (o *UpdateOrgRequest) GetCiIntegrationID() int32`

GetCiIntegrationID returns the CiIntegrationID field if non-nil, zero value otherwise.

### GetCiIntegrationIDOk

`func (o *UpdateOrgRequest) GetCiIntegrationIDOk() (*int32, bool)`

GetCiIntegrationIDOk returns a tuple with the CiIntegrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiIntegrationID

`func (o *UpdateOrgRequest) SetCiIntegrationID(v int32)`

SetCiIntegrationID sets CiIntegrationID field to given value.

### HasCiIntegrationID

`func (o *UpdateOrgRequest) HasCiIntegrationID() bool`

HasCiIntegrationID returns a boolean if a field has been set.

### SetCiIntegrationIDNil

`func (o *UpdateOrgRequest) SetCiIntegrationIDNil(b bool)`

 SetCiIntegrationIDNil sets the value for CiIntegrationID to be an explicit nil

### UnsetCiIntegrationID
`func (o *UpdateOrgRequest) UnsetCiIntegrationID()`

UnsetCiIntegrationID ensures that no value is present for CiIntegrationID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


