# AppPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Protocol** | **string** |  | 
**Number** | **int32** |  | 
**PublicPort** | Pointer to **NullableInt32** |  | [optional] 
**Private** | **bool** |  | 
**AppEndpointId** | **int32** |  | 
**AppInstanceId** | **int32** |  | 
**AppServiceId** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewAppPort

`func NewAppPort(id int32, name string, protocol string, number int32, private bool, appEndpointId int32, appInstanceId int32, appServiceId int32, createdAt time.Time, updatedAt time.Time, ) *AppPort`

NewAppPort instantiates a new AppPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPortWithDefaults

`func NewAppPortWithDefaults() *AppPort`

NewAppPortWithDefaults instantiates a new AppPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppPort) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppPort) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppPort) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *AppPort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppPort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppPort) SetName(v string)`

SetName sets Name field to given value.


### GetProtocol

`func (o *AppPort) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *AppPort) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *AppPort) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetNumber

`func (o *AppPort) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *AppPort) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *AppPort) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetPublicPort

`func (o *AppPort) GetPublicPort() int32`

GetPublicPort returns the PublicPort field if non-nil, zero value otherwise.

### GetPublicPortOk

`func (o *AppPort) GetPublicPortOk() (*int32, bool)`

GetPublicPortOk returns a tuple with the PublicPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicPort

`func (o *AppPort) SetPublicPort(v int32)`

SetPublicPort sets PublicPort field to given value.

### HasPublicPort

`func (o *AppPort) HasPublicPort() bool`

HasPublicPort returns a boolean if a field has been set.

### SetPublicPortNil

`func (o *AppPort) SetPublicPortNil(b bool)`

 SetPublicPortNil sets the value for PublicPort to be an explicit nil

### UnsetPublicPort
`func (o *AppPort) UnsetPublicPort()`

UnsetPublicPort ensures that no value is present for PublicPort, not even an explicit nil
### GetPrivate

`func (o *AppPort) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *AppPort) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *AppPort) SetPrivate(v bool)`

SetPrivate sets Private field to given value.


### GetAppEndpointId

`func (o *AppPort) GetAppEndpointId() int32`

GetAppEndpointId returns the AppEndpointId field if non-nil, zero value otherwise.

### GetAppEndpointIdOk

`func (o *AppPort) GetAppEndpointIdOk() (*int32, bool)`

GetAppEndpointIdOk returns a tuple with the AppEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEndpointId

`func (o *AppPort) SetAppEndpointId(v int32)`

SetAppEndpointId sets AppEndpointId field to given value.


### GetAppInstanceId

`func (o *AppPort) GetAppInstanceId() int32`

GetAppInstanceId returns the AppInstanceId field if non-nil, zero value otherwise.

### GetAppInstanceIdOk

`func (o *AppPort) GetAppInstanceIdOk() (*int32, bool)`

GetAppInstanceIdOk returns a tuple with the AppInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceId

`func (o *AppPort) SetAppInstanceId(v int32)`

SetAppInstanceId sets AppInstanceId field to given value.


### GetAppServiceId

`func (o *AppPort) GetAppServiceId() int32`

GetAppServiceId returns the AppServiceId field if non-nil, zero value otherwise.

### GetAppServiceIdOk

`func (o *AppPort) GetAppServiceIdOk() (*int32, bool)`

GetAppServiceIdOk returns a tuple with the AppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceId

`func (o *AppPort) SetAppServiceId(v int32)`

SetAppServiceId sets AppServiceId field to given value.


### GetCreatedAt

`func (o *AppPort) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppPort) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppPort) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AppPort) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AppPort) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AppPort) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


