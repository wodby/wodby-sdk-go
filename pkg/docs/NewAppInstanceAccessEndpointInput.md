# NewAppInstanceAccessEndpointInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppServiceName** | **string** | Machine name of an enabled app service from the selected stack revision. | 
**AppPortName** | **string** | Machine name of a public HTTP port from that service manifest. | 
**Primary** | **bool** |  | 

## Methods

### NewNewAppInstanceAccessEndpointInput

`func NewNewAppInstanceAccessEndpointInput(appServiceName string, appPortName string, primary bool, ) *NewAppInstanceAccessEndpointInput`

NewNewAppInstanceAccessEndpointInput instantiates a new NewAppInstanceAccessEndpointInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewAppInstanceAccessEndpointInputWithDefaults

`func NewNewAppInstanceAccessEndpointInputWithDefaults() *NewAppInstanceAccessEndpointInput`

NewNewAppInstanceAccessEndpointInputWithDefaults instantiates a new NewAppInstanceAccessEndpointInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppServiceName

`func (o *NewAppInstanceAccessEndpointInput) GetAppServiceName() string`

GetAppServiceName returns the AppServiceName field if non-nil, zero value otherwise.

### GetAppServiceNameOk

`func (o *NewAppInstanceAccessEndpointInput) GetAppServiceNameOk() (*string, bool)`

GetAppServiceNameOk returns a tuple with the AppServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceName

`func (o *NewAppInstanceAccessEndpointInput) SetAppServiceName(v string)`

SetAppServiceName sets AppServiceName field to given value.


### GetAppPortName

`func (o *NewAppInstanceAccessEndpointInput) GetAppPortName() string`

GetAppPortName returns the AppPortName field if non-nil, zero value otherwise.

### GetAppPortNameOk

`func (o *NewAppInstanceAccessEndpointInput) GetAppPortNameOk() (*string, bool)`

GetAppPortNameOk returns a tuple with the AppPortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPortName

`func (o *NewAppInstanceAccessEndpointInput) SetAppPortName(v string)`

SetAppPortName sets AppPortName field to given value.


### GetPrimary

`func (o *NewAppInstanceAccessEndpointInput) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *NewAppInstanceAccessEndpointInput) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *NewAppInstanceAccessEndpointInput) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


