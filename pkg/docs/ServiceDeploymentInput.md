# ServiceDeploymentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Image** | **string** |  | 
**UnmanagedImage** | Pointer to **NullableBool** | Set by the CI build when the image was produced from a Dockerfile that does not derive from the service image. | [optional] 
**DockerfilePath** | Pointer to **NullableString** | Repository path of an author-provided Dockerfile, reported by the CI build. | [optional] 
**DockerfileHash** | Pointer to **NullableString** | SHA-256 of the Dockerfile that produced the image, reported by the CI build. | [optional] 

## Methods

### NewServiceDeploymentInput

`func NewServiceDeploymentInput(name string, image string, ) *ServiceDeploymentInput`

NewServiceDeploymentInput instantiates a new ServiceDeploymentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDeploymentInputWithDefaults

`func NewServiceDeploymentInputWithDefaults() *ServiceDeploymentInput`

NewServiceDeploymentInputWithDefaults instantiates a new ServiceDeploymentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServiceDeploymentInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceDeploymentInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceDeploymentInput) SetName(v string)`

SetName sets Name field to given value.


### GetImage

`func (o *ServiceDeploymentInput) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ServiceDeploymentInput) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ServiceDeploymentInput) SetImage(v string)`

SetImage sets Image field to given value.


### GetUnmanagedImage

`func (o *ServiceDeploymentInput) GetUnmanagedImage() bool`

GetUnmanagedImage returns the UnmanagedImage field if non-nil, zero value otherwise.

### GetUnmanagedImageOk

`func (o *ServiceDeploymentInput) GetUnmanagedImageOk() (*bool, bool)`

GetUnmanagedImageOk returns a tuple with the UnmanagedImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmanagedImage

`func (o *ServiceDeploymentInput) SetUnmanagedImage(v bool)`

SetUnmanagedImage sets UnmanagedImage field to given value.

### HasUnmanagedImage

`func (o *ServiceDeploymentInput) HasUnmanagedImage() bool`

HasUnmanagedImage returns a boolean if a field has been set.

### SetUnmanagedImageNil

`func (o *ServiceDeploymentInput) SetUnmanagedImageNil(b bool)`

 SetUnmanagedImageNil sets the value for UnmanagedImage to be an explicit nil

### UnsetUnmanagedImage
`func (o *ServiceDeploymentInput) UnsetUnmanagedImage()`

UnsetUnmanagedImage ensures that no value is present for UnmanagedImage, not even an explicit nil
### GetDockerfilePath

`func (o *ServiceDeploymentInput) GetDockerfilePath() string`

GetDockerfilePath returns the DockerfilePath field if non-nil, zero value otherwise.

### GetDockerfilePathOk

`func (o *ServiceDeploymentInput) GetDockerfilePathOk() (*string, bool)`

GetDockerfilePathOk returns a tuple with the DockerfilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfilePath

`func (o *ServiceDeploymentInput) SetDockerfilePath(v string)`

SetDockerfilePath sets DockerfilePath field to given value.

### HasDockerfilePath

`func (o *ServiceDeploymentInput) HasDockerfilePath() bool`

HasDockerfilePath returns a boolean if a field has been set.

### SetDockerfilePathNil

`func (o *ServiceDeploymentInput) SetDockerfilePathNil(b bool)`

 SetDockerfilePathNil sets the value for DockerfilePath to be an explicit nil

### UnsetDockerfilePath
`func (o *ServiceDeploymentInput) UnsetDockerfilePath()`

UnsetDockerfilePath ensures that no value is present for DockerfilePath, not even an explicit nil
### GetDockerfileHash

`func (o *ServiceDeploymentInput) GetDockerfileHash() string`

GetDockerfileHash returns the DockerfileHash field if non-nil, zero value otherwise.

### GetDockerfileHashOk

`func (o *ServiceDeploymentInput) GetDockerfileHashOk() (*string, bool)`

GetDockerfileHashOk returns a tuple with the DockerfileHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfileHash

`func (o *ServiceDeploymentInput) SetDockerfileHash(v string)`

SetDockerfileHash sets DockerfileHash field to given value.

### HasDockerfileHash

`func (o *ServiceDeploymentInput) HasDockerfileHash() bool`

HasDockerfileHash returns a boolean if a field has been set.

### SetDockerfileHashNil

`func (o *ServiceDeploymentInput) SetDockerfileHashNil(b bool)`

 SetDockerfileHashNil sets the value for DockerfileHash to be an explicit nil

### UnsetDockerfileHash
`func (o *ServiceDeploymentInput) UnsetDockerfileHash()`

UnsetDockerfileHash ensures that no value is present for DockerfileHash, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


