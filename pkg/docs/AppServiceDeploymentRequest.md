# AppServiceDeploymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppServiceId** | **int32** |  | 
**AppServiceBuildId** | Pointer to **NullableInt32** |  | [optional] 
**SkipPostDeployment** | Pointer to **NullableBool** |  | [optional] 
**Force** | **bool** |  | 

## Methods

### NewAppServiceDeploymentRequest

`func NewAppServiceDeploymentRequest(appServiceId int32, force bool, ) *AppServiceDeploymentRequest`

NewAppServiceDeploymentRequest instantiates a new AppServiceDeploymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceDeploymentRequestWithDefaults

`func NewAppServiceDeploymentRequestWithDefaults() *AppServiceDeploymentRequest`

NewAppServiceDeploymentRequestWithDefaults instantiates a new AppServiceDeploymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppServiceId

`func (o *AppServiceDeploymentRequest) GetAppServiceId() int32`

GetAppServiceId returns the AppServiceId field if non-nil, zero value otherwise.

### GetAppServiceIdOk

`func (o *AppServiceDeploymentRequest) GetAppServiceIdOk() (*int32, bool)`

GetAppServiceIdOk returns a tuple with the AppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceId

`func (o *AppServiceDeploymentRequest) SetAppServiceId(v int32)`

SetAppServiceId sets AppServiceId field to given value.


### GetAppServiceBuildId

`func (o *AppServiceDeploymentRequest) GetAppServiceBuildId() int32`

GetAppServiceBuildId returns the AppServiceBuildId field if non-nil, zero value otherwise.

### GetAppServiceBuildIdOk

`func (o *AppServiceDeploymentRequest) GetAppServiceBuildIdOk() (*int32, bool)`

GetAppServiceBuildIdOk returns a tuple with the AppServiceBuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceBuildId

`func (o *AppServiceDeploymentRequest) SetAppServiceBuildId(v int32)`

SetAppServiceBuildId sets AppServiceBuildId field to given value.

### HasAppServiceBuildId

`func (o *AppServiceDeploymentRequest) HasAppServiceBuildId() bool`

HasAppServiceBuildId returns a boolean if a field has been set.

### SetAppServiceBuildIdNil

`func (o *AppServiceDeploymentRequest) SetAppServiceBuildIdNil(b bool)`

 SetAppServiceBuildIdNil sets the value for AppServiceBuildId to be an explicit nil

### UnsetAppServiceBuildId
`func (o *AppServiceDeploymentRequest) UnsetAppServiceBuildId()`

UnsetAppServiceBuildId ensures that no value is present for AppServiceBuildId, not even an explicit nil
### GetSkipPostDeployment

`func (o *AppServiceDeploymentRequest) GetSkipPostDeployment() bool`

GetSkipPostDeployment returns the SkipPostDeployment field if non-nil, zero value otherwise.

### GetSkipPostDeploymentOk

`func (o *AppServiceDeploymentRequest) GetSkipPostDeploymentOk() (*bool, bool)`

GetSkipPostDeploymentOk returns a tuple with the SkipPostDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipPostDeployment

`func (o *AppServiceDeploymentRequest) SetSkipPostDeployment(v bool)`

SetSkipPostDeployment sets SkipPostDeployment field to given value.

### HasSkipPostDeployment

`func (o *AppServiceDeploymentRequest) HasSkipPostDeployment() bool`

HasSkipPostDeployment returns a boolean if a field has been set.

### SetSkipPostDeploymentNil

`func (o *AppServiceDeploymentRequest) SetSkipPostDeploymentNil(b bool)`

 SetSkipPostDeploymentNil sets the value for SkipPostDeployment to be an explicit nil

### UnsetSkipPostDeployment
`func (o *AppServiceDeploymentRequest) UnsetSkipPostDeployment()`

UnsetSkipPostDeployment ensures that no value is present for SkipPostDeployment, not even an explicit nil
### GetForce

`func (o *AppServiceDeploymentRequest) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *AppServiceDeploymentRequest) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *AppServiceDeploymentRequest) SetForce(v bool)`

SetForce sets Force field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


