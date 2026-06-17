# BuildSourceInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildSourceType** | **string** |  | 
**Template** | Pointer to **NullableString** |  | [optional] 
**NewRepoName** | Pointer to **NullableString** |  | [optional] 
**IntegrationID** | Pointer to **NullableInt32** |  | [optional] 
**RemoteGitRepoID** | Pointer to **NullableString** |  | [optional] 
**GitRef** | Pointer to **NullableString** |  | [optional] 
**GitRefType** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBuildSourceInput

`func NewBuildSourceInput(buildSourceType string, ) *BuildSourceInput`

NewBuildSourceInput instantiates a new BuildSourceInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildSourceInputWithDefaults

`func NewBuildSourceInputWithDefaults() *BuildSourceInput`

NewBuildSourceInputWithDefaults instantiates a new BuildSourceInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildSourceType

`func (o *BuildSourceInput) GetBuildSourceType() string`

GetBuildSourceType returns the BuildSourceType field if non-nil, zero value otherwise.

### GetBuildSourceTypeOk

`func (o *BuildSourceInput) GetBuildSourceTypeOk() (*string, bool)`

GetBuildSourceTypeOk returns a tuple with the BuildSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildSourceType

`func (o *BuildSourceInput) SetBuildSourceType(v string)`

SetBuildSourceType sets BuildSourceType field to given value.


### GetTemplate

`func (o *BuildSourceInput) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *BuildSourceInput) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *BuildSourceInput) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *BuildSourceInput) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### SetTemplateNil

`func (o *BuildSourceInput) SetTemplateNil(b bool)`

 SetTemplateNil sets the value for Template to be an explicit nil

### UnsetTemplate
`func (o *BuildSourceInput) UnsetTemplate()`

UnsetTemplate ensures that no value is present for Template, not even an explicit nil
### GetNewRepoName

`func (o *BuildSourceInput) GetNewRepoName() string`

GetNewRepoName returns the NewRepoName field if non-nil, zero value otherwise.

### GetNewRepoNameOk

`func (o *BuildSourceInput) GetNewRepoNameOk() (*string, bool)`

GetNewRepoNameOk returns a tuple with the NewRepoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewRepoName

`func (o *BuildSourceInput) SetNewRepoName(v string)`

SetNewRepoName sets NewRepoName field to given value.

### HasNewRepoName

`func (o *BuildSourceInput) HasNewRepoName() bool`

HasNewRepoName returns a boolean if a field has been set.

### SetNewRepoNameNil

`func (o *BuildSourceInput) SetNewRepoNameNil(b bool)`

 SetNewRepoNameNil sets the value for NewRepoName to be an explicit nil

### UnsetNewRepoName
`func (o *BuildSourceInput) UnsetNewRepoName()`

UnsetNewRepoName ensures that no value is present for NewRepoName, not even an explicit nil
### GetIntegrationID

`func (o *BuildSourceInput) GetIntegrationID() int32`

GetIntegrationID returns the IntegrationID field if non-nil, zero value otherwise.

### GetIntegrationIDOk

`func (o *BuildSourceInput) GetIntegrationIDOk() (*int32, bool)`

GetIntegrationIDOk returns a tuple with the IntegrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationID

`func (o *BuildSourceInput) SetIntegrationID(v int32)`

SetIntegrationID sets IntegrationID field to given value.

### HasIntegrationID

`func (o *BuildSourceInput) HasIntegrationID() bool`

HasIntegrationID returns a boolean if a field has been set.

### SetIntegrationIDNil

`func (o *BuildSourceInput) SetIntegrationIDNil(b bool)`

 SetIntegrationIDNil sets the value for IntegrationID to be an explicit nil

### UnsetIntegrationID
`func (o *BuildSourceInput) UnsetIntegrationID()`

UnsetIntegrationID ensures that no value is present for IntegrationID, not even an explicit nil
### GetRemoteGitRepoID

`func (o *BuildSourceInput) GetRemoteGitRepoID() string`

GetRemoteGitRepoID returns the RemoteGitRepoID field if non-nil, zero value otherwise.

### GetRemoteGitRepoIDOk

`func (o *BuildSourceInput) GetRemoteGitRepoIDOk() (*string, bool)`

GetRemoteGitRepoIDOk returns a tuple with the RemoteGitRepoID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteGitRepoID

`func (o *BuildSourceInput) SetRemoteGitRepoID(v string)`

SetRemoteGitRepoID sets RemoteGitRepoID field to given value.

### HasRemoteGitRepoID

`func (o *BuildSourceInput) HasRemoteGitRepoID() bool`

HasRemoteGitRepoID returns a boolean if a field has been set.

### SetRemoteGitRepoIDNil

`func (o *BuildSourceInput) SetRemoteGitRepoIDNil(b bool)`

 SetRemoteGitRepoIDNil sets the value for RemoteGitRepoID to be an explicit nil

### UnsetRemoteGitRepoID
`func (o *BuildSourceInput) UnsetRemoteGitRepoID()`

UnsetRemoteGitRepoID ensures that no value is present for RemoteGitRepoID, not even an explicit nil
### GetGitRef

`func (o *BuildSourceInput) GetGitRef() string`

GetGitRef returns the GitRef field if non-nil, zero value otherwise.

### GetGitRefOk

`func (o *BuildSourceInput) GetGitRefOk() (*string, bool)`

GetGitRefOk returns a tuple with the GitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRef

`func (o *BuildSourceInput) SetGitRef(v string)`

SetGitRef sets GitRef field to given value.

### HasGitRef

`func (o *BuildSourceInput) HasGitRef() bool`

HasGitRef returns a boolean if a field has been set.

### SetGitRefNil

`func (o *BuildSourceInput) SetGitRefNil(b bool)`

 SetGitRefNil sets the value for GitRef to be an explicit nil

### UnsetGitRef
`func (o *BuildSourceInput) UnsetGitRef()`

UnsetGitRef ensures that no value is present for GitRef, not even an explicit nil
### GetGitRefType

`func (o *BuildSourceInput) GetGitRefType() string`

GetGitRefType returns the GitRefType field if non-nil, zero value otherwise.

### GetGitRefTypeOk

`func (o *BuildSourceInput) GetGitRefTypeOk() (*string, bool)`

GetGitRefTypeOk returns a tuple with the GitRefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRefType

`func (o *BuildSourceInput) SetGitRefType(v string)`

SetGitRefType sets GitRefType field to given value.

### HasGitRefType

`func (o *BuildSourceInput) HasGitRefType() bool`

HasGitRefType returns a boolean if a field has been set.

### SetGitRefTypeNil

`func (o *BuildSourceInput) SetGitRefTypeNil(b bool)`

 SetGitRefTypeNil sets the value for GitRefType to be an explicit nil

### UnsetGitRefType
`func (o *BuildSourceInput) UnsetGitRefType()`

UnsetGitRefType ensures that no value is present for GitRefType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


