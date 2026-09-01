# AppServiceBuildConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Title** | **string** |  | 
**Managed** | **bool** |  | 
**Main** | **bool** |  | 
**Image** | **string** |  | 
**Dockerfile** | Pointer to **NullableString** |  | [optional] 
**Dockerignore** | Pointer to **NullableString** |  | [optional] 
**CopySubdir** | **string** | Resolved subdirectory this build copies, applied under both the CI --from and --to paths. Empty means the whole context. | 
**Args** | Pointer to [**[]AppServiceBuildArg**](AppServiceBuildArg.md) |  | [optional] 

## Methods

### NewAppServiceBuildConfig

`func NewAppServiceBuildConfig(name string, title string, managed bool, main bool, image string, copySubdir string, ) *AppServiceBuildConfig`

NewAppServiceBuildConfig instantiates a new AppServiceBuildConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceBuildConfigWithDefaults

`func NewAppServiceBuildConfigWithDefaults() *AppServiceBuildConfig`

NewAppServiceBuildConfigWithDefaults instantiates a new AppServiceBuildConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AppServiceBuildConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppServiceBuildConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppServiceBuildConfig) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *AppServiceBuildConfig) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AppServiceBuildConfig) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AppServiceBuildConfig) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetManaged

`func (o *AppServiceBuildConfig) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *AppServiceBuildConfig) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *AppServiceBuildConfig) SetManaged(v bool)`

SetManaged sets Managed field to given value.


### GetMain

`func (o *AppServiceBuildConfig) GetMain() bool`

GetMain returns the Main field if non-nil, zero value otherwise.

### GetMainOk

`func (o *AppServiceBuildConfig) GetMainOk() (*bool, bool)`

GetMainOk returns a tuple with the Main field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMain

`func (o *AppServiceBuildConfig) SetMain(v bool)`

SetMain sets Main field to given value.


### GetImage

`func (o *AppServiceBuildConfig) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *AppServiceBuildConfig) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *AppServiceBuildConfig) SetImage(v string)`

SetImage sets Image field to given value.


### GetDockerfile

`func (o *AppServiceBuildConfig) GetDockerfile() string`

GetDockerfile returns the Dockerfile field if non-nil, zero value otherwise.

### GetDockerfileOk

`func (o *AppServiceBuildConfig) GetDockerfileOk() (*string, bool)`

GetDockerfileOk returns a tuple with the Dockerfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfile

`func (o *AppServiceBuildConfig) SetDockerfile(v string)`

SetDockerfile sets Dockerfile field to given value.

### HasDockerfile

`func (o *AppServiceBuildConfig) HasDockerfile() bool`

HasDockerfile returns a boolean if a field has been set.

### SetDockerfileNil

`func (o *AppServiceBuildConfig) SetDockerfileNil(b bool)`

 SetDockerfileNil sets the value for Dockerfile to be an explicit nil

### UnsetDockerfile
`func (o *AppServiceBuildConfig) UnsetDockerfile()`

UnsetDockerfile ensures that no value is present for Dockerfile, not even an explicit nil
### GetDockerignore

`func (o *AppServiceBuildConfig) GetDockerignore() string`

GetDockerignore returns the Dockerignore field if non-nil, zero value otherwise.

### GetDockerignoreOk

`func (o *AppServiceBuildConfig) GetDockerignoreOk() (*string, bool)`

GetDockerignoreOk returns a tuple with the Dockerignore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerignore

`func (o *AppServiceBuildConfig) SetDockerignore(v string)`

SetDockerignore sets Dockerignore field to given value.

### HasDockerignore

`func (o *AppServiceBuildConfig) HasDockerignore() bool`

HasDockerignore returns a boolean if a field has been set.

### SetDockerignoreNil

`func (o *AppServiceBuildConfig) SetDockerignoreNil(b bool)`

 SetDockerignoreNil sets the value for Dockerignore to be an explicit nil

### UnsetDockerignore
`func (o *AppServiceBuildConfig) UnsetDockerignore()`

UnsetDockerignore ensures that no value is present for Dockerignore, not even an explicit nil
### GetCopySubdir

`func (o *AppServiceBuildConfig) GetCopySubdir() string`

GetCopySubdir returns the CopySubdir field if non-nil, zero value otherwise.

### GetCopySubdirOk

`func (o *AppServiceBuildConfig) GetCopySubdirOk() (*string, bool)`

GetCopySubdirOk returns a tuple with the CopySubdir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopySubdir

`func (o *AppServiceBuildConfig) SetCopySubdir(v string)`

SetCopySubdir sets CopySubdir field to given value.


### GetArgs

`func (o *AppServiceBuildConfig) GetArgs() []AppServiceBuildArg`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *AppServiceBuildConfig) GetArgsOk() (*[]AppServiceBuildArg, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *AppServiceBuildConfig) SetArgs(v []AppServiceBuildArg)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *AppServiceBuildConfig) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### SetArgsNil

`func (o *AppServiceBuildConfig) SetArgsNil(b bool)`

 SetArgsNil sets the value for Args to be an explicit nil

### UnsetArgs
`func (o *AppServiceBuildConfig) UnsetArgs()`

UnsetArgs ensures that no value is present for Args, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


