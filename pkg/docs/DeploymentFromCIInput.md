# DeploymentFromCIInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppBuildId** | **int32** |  | 
**Services** | [**[]ServiceDeploymentInput**](ServiceDeploymentInput.md) |  | 
**SkipPostDeployment** | **bool** |  | 

## Methods

### NewDeploymentFromCIInput

`func NewDeploymentFromCIInput(appBuildId int32, services []ServiceDeploymentInput, skipPostDeployment bool, ) *DeploymentFromCIInput`

NewDeploymentFromCIInput instantiates a new DeploymentFromCIInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentFromCIInputWithDefaults

`func NewDeploymentFromCIInputWithDefaults() *DeploymentFromCIInput`

NewDeploymentFromCIInputWithDefaults instantiates a new DeploymentFromCIInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppBuildId

`func (o *DeploymentFromCIInput) GetAppBuildId() int32`

GetAppBuildId returns the AppBuildId field if non-nil, zero value otherwise.

### GetAppBuildIdOk

`func (o *DeploymentFromCIInput) GetAppBuildIdOk() (*int32, bool)`

GetAppBuildIdOk returns a tuple with the AppBuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppBuildId

`func (o *DeploymentFromCIInput) SetAppBuildId(v int32)`

SetAppBuildId sets AppBuildId field to given value.


### GetServices

`func (o *DeploymentFromCIInput) GetServices() []ServiceDeploymentInput`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *DeploymentFromCIInput) GetServicesOk() (*[]ServiceDeploymentInput, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *DeploymentFromCIInput) SetServices(v []ServiceDeploymentInput)`

SetServices sets Services field to given value.


### GetSkipPostDeployment

`func (o *DeploymentFromCIInput) GetSkipPostDeployment() bool`

GetSkipPostDeployment returns the SkipPostDeployment field if non-nil, zero value otherwise.

### GetSkipPostDeploymentOk

`func (o *DeploymentFromCIInput) GetSkipPostDeploymentOk() (*bool, bool)`

GetSkipPostDeploymentOk returns a tuple with the SkipPostDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipPostDeployment

`func (o *DeploymentFromCIInput) SetSkipPostDeployment(v bool)`

SetSkipPostDeployment sets SkipPostDeployment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


