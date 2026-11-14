# K3SPrivateAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**NetworkInterface** | **string** | The network interface the address is on, such as enp0s31f6.4000 for a VLAN. | 
**Mtu** | **int32** | The interface&#39;s MTU, or 0 when the server doesn&#39;t report it. | 
**Persistent** | **bool** | Whether the server&#39;s network configuration brings the address back after a reboot. k3s doesn&#39;t start when its address is gone, so switching to an address that isn&#39;t persistent needs force. | 
**PersistentSource** | Pointer to **string** | What keeps the address after a reboot, such as systemd-networkd or NetworkManager. | [optional] 

## Methods

### NewK3SPrivateAddress

`func NewK3SPrivateAddress(address string, networkInterface string, mtu int32, persistent bool, ) *K3SPrivateAddress`

NewK3SPrivateAddress instantiates a new K3SPrivateAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK3SPrivateAddressWithDefaults

`func NewK3SPrivateAddressWithDefaults() *K3SPrivateAddress`

NewK3SPrivateAddressWithDefaults instantiates a new K3SPrivateAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *K3SPrivateAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *K3SPrivateAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *K3SPrivateAddress) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetNetworkInterface

`func (o *K3SPrivateAddress) GetNetworkInterface() string`

GetNetworkInterface returns the NetworkInterface field if non-nil, zero value otherwise.

### GetNetworkInterfaceOk

`func (o *K3SPrivateAddress) GetNetworkInterfaceOk() (*string, bool)`

GetNetworkInterfaceOk returns a tuple with the NetworkInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterface

`func (o *K3SPrivateAddress) SetNetworkInterface(v string)`

SetNetworkInterface sets NetworkInterface field to given value.


### GetMtu

`func (o *K3SPrivateAddress) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *K3SPrivateAddress) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *K3SPrivateAddress) SetMtu(v int32)`

SetMtu sets Mtu field to given value.


### GetPersistent

`func (o *K3SPrivateAddress) GetPersistent() bool`

GetPersistent returns the Persistent field if non-nil, zero value otherwise.

### GetPersistentOk

`func (o *K3SPrivateAddress) GetPersistentOk() (*bool, bool)`

GetPersistentOk returns a tuple with the Persistent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistent

`func (o *K3SPrivateAddress) SetPersistent(v bool)`

SetPersistent sets Persistent field to given value.


### GetPersistentSource

`func (o *K3SPrivateAddress) GetPersistentSource() string`

GetPersistentSource returns the PersistentSource field if non-nil, zero value otherwise.

### GetPersistentSourceOk

`func (o *K3SPrivateAddress) GetPersistentSourceOk() (*string, bool)`

GetPersistentSourceOk returns a tuple with the PersistentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentSource

`func (o *K3SPrivateAddress) SetPersistentSource(v string)`

SetPersistentSource sets PersistentSource field to given value.

### HasPersistentSource

`func (o *K3SPrivateAddress) HasPersistentSource() bool`

HasPersistentSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


