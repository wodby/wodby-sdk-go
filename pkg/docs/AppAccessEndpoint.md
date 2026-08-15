# AppAccessEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**AppPortId** | **int32** |  | 
**Host** | **string** |  | 
**Primary** | **bool** |  | 
**Url** | **string** |  | 

## Methods

### NewAppAccessEndpoint

`func NewAppAccessEndpoint(id int32, appPortId int32, host string, primary bool, url string, ) *AppAccessEndpoint`

NewAppAccessEndpoint instantiates a new AppAccessEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppAccessEndpointWithDefaults

`func NewAppAccessEndpointWithDefaults() *AppAccessEndpoint`

NewAppAccessEndpointWithDefaults instantiates a new AppAccessEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppAccessEndpoint) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppAccessEndpoint) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppAccessEndpoint) SetId(v int32)`

SetId sets Id field to given value.


### GetAppPortId

`func (o *AppAccessEndpoint) GetAppPortId() int32`

GetAppPortId returns the AppPortId field if non-nil, zero value otherwise.

### GetAppPortIdOk

`func (o *AppAccessEndpoint) GetAppPortIdOk() (*int32, bool)`

GetAppPortIdOk returns a tuple with the AppPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPortId

`func (o *AppAccessEndpoint) SetAppPortId(v int32)`

SetAppPortId sets AppPortId field to given value.


### GetHost

`func (o *AppAccessEndpoint) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *AppAccessEndpoint) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *AppAccessEndpoint) SetHost(v string)`

SetHost sets Host field to given value.


### GetPrimary

`func (o *AppAccessEndpoint) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *AppAccessEndpoint) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *AppAccessEndpoint) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.


### GetUrl

`func (o *AppAccessEndpoint) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AppAccessEndpoint) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AppAccessEndpoint) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


