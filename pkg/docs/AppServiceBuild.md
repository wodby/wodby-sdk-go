# AppServiceBuild

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Status** | **string** |  | 
**Image** | **string** |  | 
**ImageDeleted** | **bool** |  | 
**UnmanagedImage** | **bool** | True when the image was built from a Dockerfile that does not derive from the service image, so it no longer tracks service image updates. | 
**DockerfilePath** | **string** | Repository path of an author-provided Dockerfile. Empty when the build used a service-provided or generated Dockerfile. | 
**DockerfileHash** | **string** | SHA-256 of the Dockerfile that produced the image. Empty when the build did not report it. | 
**Size** | **int32** |  | 
**AppServiceId** | **int32** |  | 
**PreviouslyDeployed** | **bool** |  | 
**CurrentlyDeployed** | **bool** |  | 
**CurrentBuildNumber** | Pointer to **NullableInt32** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewAppServiceBuild

`func NewAppServiceBuild(id int32, status string, image string, imageDeleted bool, unmanagedImage bool, dockerfilePath string, dockerfileHash string, size int32, appServiceId int32, previouslyDeployed bool, currentlyDeployed bool, createdAt time.Time, updatedAt time.Time, ) *AppServiceBuild`

NewAppServiceBuild instantiates a new AppServiceBuild object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceBuildWithDefaults

`func NewAppServiceBuildWithDefaults() *AppServiceBuild`

NewAppServiceBuildWithDefaults instantiates a new AppServiceBuild object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppServiceBuild) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppServiceBuild) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppServiceBuild) SetId(v int32)`

SetId sets Id field to given value.


### GetStatus

`func (o *AppServiceBuild) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppServiceBuild) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppServiceBuild) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetImage

`func (o *AppServiceBuild) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *AppServiceBuild) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *AppServiceBuild) SetImage(v string)`

SetImage sets Image field to given value.


### GetImageDeleted

`func (o *AppServiceBuild) GetImageDeleted() bool`

GetImageDeleted returns the ImageDeleted field if non-nil, zero value otherwise.

### GetImageDeletedOk

`func (o *AppServiceBuild) GetImageDeletedOk() (*bool, bool)`

GetImageDeletedOk returns a tuple with the ImageDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDeleted

`func (o *AppServiceBuild) SetImageDeleted(v bool)`

SetImageDeleted sets ImageDeleted field to given value.


### GetUnmanagedImage

`func (o *AppServiceBuild) GetUnmanagedImage() bool`

GetUnmanagedImage returns the UnmanagedImage field if non-nil, zero value otherwise.

### GetUnmanagedImageOk

`func (o *AppServiceBuild) GetUnmanagedImageOk() (*bool, bool)`

GetUnmanagedImageOk returns a tuple with the UnmanagedImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmanagedImage

`func (o *AppServiceBuild) SetUnmanagedImage(v bool)`

SetUnmanagedImage sets UnmanagedImage field to given value.


### GetDockerfilePath

`func (o *AppServiceBuild) GetDockerfilePath() string`

GetDockerfilePath returns the DockerfilePath field if non-nil, zero value otherwise.

### GetDockerfilePathOk

`func (o *AppServiceBuild) GetDockerfilePathOk() (*string, bool)`

GetDockerfilePathOk returns a tuple with the DockerfilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfilePath

`func (o *AppServiceBuild) SetDockerfilePath(v string)`

SetDockerfilePath sets DockerfilePath field to given value.


### GetDockerfileHash

`func (o *AppServiceBuild) GetDockerfileHash() string`

GetDockerfileHash returns the DockerfileHash field if non-nil, zero value otherwise.

### GetDockerfileHashOk

`func (o *AppServiceBuild) GetDockerfileHashOk() (*string, bool)`

GetDockerfileHashOk returns a tuple with the DockerfileHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfileHash

`func (o *AppServiceBuild) SetDockerfileHash(v string)`

SetDockerfileHash sets DockerfileHash field to given value.


### GetSize

`func (o *AppServiceBuild) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AppServiceBuild) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AppServiceBuild) SetSize(v int32)`

SetSize sets Size field to given value.


### GetAppServiceId

`func (o *AppServiceBuild) GetAppServiceId() int32`

GetAppServiceId returns the AppServiceId field if non-nil, zero value otherwise.

### GetAppServiceIdOk

`func (o *AppServiceBuild) GetAppServiceIdOk() (*int32, bool)`

GetAppServiceIdOk returns a tuple with the AppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceId

`func (o *AppServiceBuild) SetAppServiceId(v int32)`

SetAppServiceId sets AppServiceId field to given value.


### GetPreviouslyDeployed

`func (o *AppServiceBuild) GetPreviouslyDeployed() bool`

GetPreviouslyDeployed returns the PreviouslyDeployed field if non-nil, zero value otherwise.

### GetPreviouslyDeployedOk

`func (o *AppServiceBuild) GetPreviouslyDeployedOk() (*bool, bool)`

GetPreviouslyDeployedOk returns a tuple with the PreviouslyDeployed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviouslyDeployed

`func (o *AppServiceBuild) SetPreviouslyDeployed(v bool)`

SetPreviouslyDeployed sets PreviouslyDeployed field to given value.


### GetCurrentlyDeployed

`func (o *AppServiceBuild) GetCurrentlyDeployed() bool`

GetCurrentlyDeployed returns the CurrentlyDeployed field if non-nil, zero value otherwise.

### GetCurrentlyDeployedOk

`func (o *AppServiceBuild) GetCurrentlyDeployedOk() (*bool, bool)`

GetCurrentlyDeployedOk returns a tuple with the CurrentlyDeployed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentlyDeployed

`func (o *AppServiceBuild) SetCurrentlyDeployed(v bool)`

SetCurrentlyDeployed sets CurrentlyDeployed field to given value.


### GetCurrentBuildNumber

`func (o *AppServiceBuild) GetCurrentBuildNumber() int32`

GetCurrentBuildNumber returns the CurrentBuildNumber field if non-nil, zero value otherwise.

### GetCurrentBuildNumberOk

`func (o *AppServiceBuild) GetCurrentBuildNumberOk() (*int32, bool)`

GetCurrentBuildNumberOk returns a tuple with the CurrentBuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBuildNumber

`func (o *AppServiceBuild) SetCurrentBuildNumber(v int32)`

SetCurrentBuildNumber sets CurrentBuildNumber field to given value.

### HasCurrentBuildNumber

`func (o *AppServiceBuild) HasCurrentBuildNumber() bool`

HasCurrentBuildNumber returns a boolean if a field has been set.

### SetCurrentBuildNumberNil

`func (o *AppServiceBuild) SetCurrentBuildNumberNil(b bool)`

 SetCurrentBuildNumberNil sets the value for CurrentBuildNumber to be an explicit nil

### UnsetCurrentBuildNumber
`func (o *AppServiceBuild) UnsetCurrentBuildNumber()`

UnsetCurrentBuildNumber ensures that no value is present for CurrentBuildNumber, not even an explicit nil
### GetCreatedAt

`func (o *AppServiceBuild) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppServiceBuild) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppServiceBuild) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AppServiceBuild) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AppServiceBuild) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AppServiceBuild) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


