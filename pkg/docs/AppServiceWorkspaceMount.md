# AppServiceWorkspaceMount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Path** | **string** | Absolute path in the container. | 
**SubPath** | **string** | Directory of the checkout that is mounted; empty for the whole repository. | 
**ReadOnly** | **bool** |  | 

## Methods

### NewAppServiceWorkspaceMount

`func NewAppServiceWorkspaceMount(name string, path string, subPath string, readOnly bool, ) *AppServiceWorkspaceMount`

NewAppServiceWorkspaceMount instantiates a new AppServiceWorkspaceMount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceWorkspaceMountWithDefaults

`func NewAppServiceWorkspaceMountWithDefaults() *AppServiceWorkspaceMount`

NewAppServiceWorkspaceMountWithDefaults instantiates a new AppServiceWorkspaceMount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AppServiceWorkspaceMount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppServiceWorkspaceMount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppServiceWorkspaceMount) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *AppServiceWorkspaceMount) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AppServiceWorkspaceMount) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AppServiceWorkspaceMount) SetPath(v string)`

SetPath sets Path field to given value.


### GetSubPath

`func (o *AppServiceWorkspaceMount) GetSubPath() string`

GetSubPath returns the SubPath field if non-nil, zero value otherwise.

### GetSubPathOk

`func (o *AppServiceWorkspaceMount) GetSubPathOk() (*string, bool)`

GetSubPathOk returns a tuple with the SubPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubPath

`func (o *AppServiceWorkspaceMount) SetSubPath(v string)`

SetSubPath sets SubPath field to given value.


### GetReadOnly

`func (o *AppServiceWorkspaceMount) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *AppServiceWorkspaceMount) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *AppServiceWorkspaceMount) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


