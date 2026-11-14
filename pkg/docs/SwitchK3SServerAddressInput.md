# SwitchK3SServerAddressInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | A private address from GET /cluster-k3s-private-addresses/{id}. | 
**Force** | Pointer to **bool** | Switch to an address the server&#39;s network configuration isn&#39;t known to keep after a reboot. | [optional] 

## Methods

### NewSwitchK3SServerAddressInput

`func NewSwitchK3SServerAddressInput(address string, ) *SwitchK3SServerAddressInput`

NewSwitchK3SServerAddressInput instantiates a new SwitchK3SServerAddressInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchK3SServerAddressInputWithDefaults

`func NewSwitchK3SServerAddressInputWithDefaults() *SwitchK3SServerAddressInput`

NewSwitchK3SServerAddressInputWithDefaults instantiates a new SwitchK3SServerAddressInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *SwitchK3SServerAddressInput) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SwitchK3SServerAddressInput) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SwitchK3SServerAddressInput) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetForce

`func (o *SwitchK3SServerAddressInput) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *SwitchK3SServerAddressInput) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *SwitchK3SServerAddressInput) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *SwitchK3SServerAddressInput) HasForce() bool`

HasForce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


