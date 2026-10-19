# ResolvedServiceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Title** | **string** |  | 
**Filepath** | **string** |  | 
**ManifestVersion** | Pointer to **NullableString** |  | [optional] 
**DefaultSource** | **string** | Whether the default content comes from the service image or its repository. | 
**DefaultConfig** | Pointer to **NullableString** |  | [optional] 
**Status** | **string** |  | 
**LegacyFallback** | **bool** |  | 
**Image** | Pointer to **NullableString** |  | [optional] 
**ImageDigest** | Pointer to **NullableString** |  | [optional] 
**SourcePath** | Pointer to **NullableString** |  | [optional] 
**Platform** | **string** |  | 
**Error** | Pointer to **NullableString** |  | [optional] 
**FetchedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewResolvedServiceConfig

`func NewResolvedServiceConfig(name string, title string, filepath string, defaultSource string, status string, legacyFallback bool, platform string, ) *ResolvedServiceConfig`

NewResolvedServiceConfig instantiates a new ResolvedServiceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResolvedServiceConfigWithDefaults

`func NewResolvedServiceConfigWithDefaults() *ResolvedServiceConfig`

NewResolvedServiceConfigWithDefaults instantiates a new ResolvedServiceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResolvedServiceConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResolvedServiceConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResolvedServiceConfig) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *ResolvedServiceConfig) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ResolvedServiceConfig) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ResolvedServiceConfig) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetFilepath

`func (o *ResolvedServiceConfig) GetFilepath() string`

GetFilepath returns the Filepath field if non-nil, zero value otherwise.

### GetFilepathOk

`func (o *ResolvedServiceConfig) GetFilepathOk() (*string, bool)`

GetFilepathOk returns a tuple with the Filepath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilepath

`func (o *ResolvedServiceConfig) SetFilepath(v string)`

SetFilepath sets Filepath field to given value.


### GetManifestVersion

`func (o *ResolvedServiceConfig) GetManifestVersion() string`

GetManifestVersion returns the ManifestVersion field if non-nil, zero value otherwise.

### GetManifestVersionOk

`func (o *ResolvedServiceConfig) GetManifestVersionOk() (*string, bool)`

GetManifestVersionOk returns a tuple with the ManifestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestVersion

`func (o *ResolvedServiceConfig) SetManifestVersion(v string)`

SetManifestVersion sets ManifestVersion field to given value.

### HasManifestVersion

`func (o *ResolvedServiceConfig) HasManifestVersion() bool`

HasManifestVersion returns a boolean if a field has been set.

### SetManifestVersionNil

`func (o *ResolvedServiceConfig) SetManifestVersionNil(b bool)`

 SetManifestVersionNil sets the value for ManifestVersion to be an explicit nil

### UnsetManifestVersion
`func (o *ResolvedServiceConfig) UnsetManifestVersion()`

UnsetManifestVersion ensures that no value is present for ManifestVersion, not even an explicit nil
### GetDefaultSource

`func (o *ResolvedServiceConfig) GetDefaultSource() string`

GetDefaultSource returns the DefaultSource field if non-nil, zero value otherwise.

### GetDefaultSourceOk

`func (o *ResolvedServiceConfig) GetDefaultSourceOk() (*string, bool)`

GetDefaultSourceOk returns a tuple with the DefaultSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSource

`func (o *ResolvedServiceConfig) SetDefaultSource(v string)`

SetDefaultSource sets DefaultSource field to given value.


### GetDefaultConfig

`func (o *ResolvedServiceConfig) GetDefaultConfig() string`

GetDefaultConfig returns the DefaultConfig field if non-nil, zero value otherwise.

### GetDefaultConfigOk

`func (o *ResolvedServiceConfig) GetDefaultConfigOk() (*string, bool)`

GetDefaultConfigOk returns a tuple with the DefaultConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConfig

`func (o *ResolvedServiceConfig) SetDefaultConfig(v string)`

SetDefaultConfig sets DefaultConfig field to given value.

### HasDefaultConfig

`func (o *ResolvedServiceConfig) HasDefaultConfig() bool`

HasDefaultConfig returns a boolean if a field has been set.

### SetDefaultConfigNil

`func (o *ResolvedServiceConfig) SetDefaultConfigNil(b bool)`

 SetDefaultConfigNil sets the value for DefaultConfig to be an explicit nil

### UnsetDefaultConfig
`func (o *ResolvedServiceConfig) UnsetDefaultConfig()`

UnsetDefaultConfig ensures that no value is present for DefaultConfig, not even an explicit nil
### GetStatus

`func (o *ResolvedServiceConfig) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResolvedServiceConfig) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResolvedServiceConfig) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLegacyFallback

`func (o *ResolvedServiceConfig) GetLegacyFallback() bool`

GetLegacyFallback returns the LegacyFallback field if non-nil, zero value otherwise.

### GetLegacyFallbackOk

`func (o *ResolvedServiceConfig) GetLegacyFallbackOk() (*bool, bool)`

GetLegacyFallbackOk returns a tuple with the LegacyFallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyFallback

`func (o *ResolvedServiceConfig) SetLegacyFallback(v bool)`

SetLegacyFallback sets LegacyFallback field to given value.


### GetImage

`func (o *ResolvedServiceConfig) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ResolvedServiceConfig) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ResolvedServiceConfig) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *ResolvedServiceConfig) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *ResolvedServiceConfig) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *ResolvedServiceConfig) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetImageDigest

`func (o *ResolvedServiceConfig) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *ResolvedServiceConfig) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *ResolvedServiceConfig) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.

### HasImageDigest

`func (o *ResolvedServiceConfig) HasImageDigest() bool`

HasImageDigest returns a boolean if a field has been set.

### SetImageDigestNil

`func (o *ResolvedServiceConfig) SetImageDigestNil(b bool)`

 SetImageDigestNil sets the value for ImageDigest to be an explicit nil

### UnsetImageDigest
`func (o *ResolvedServiceConfig) UnsetImageDigest()`

UnsetImageDigest ensures that no value is present for ImageDigest, not even an explicit nil
### GetSourcePath

`func (o *ResolvedServiceConfig) GetSourcePath() string`

GetSourcePath returns the SourcePath field if non-nil, zero value otherwise.

### GetSourcePathOk

`func (o *ResolvedServiceConfig) GetSourcePathOk() (*string, bool)`

GetSourcePathOk returns a tuple with the SourcePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePath

`func (o *ResolvedServiceConfig) SetSourcePath(v string)`

SetSourcePath sets SourcePath field to given value.

### HasSourcePath

`func (o *ResolvedServiceConfig) HasSourcePath() bool`

HasSourcePath returns a boolean if a field has been set.

### SetSourcePathNil

`func (o *ResolvedServiceConfig) SetSourcePathNil(b bool)`

 SetSourcePathNil sets the value for SourcePath to be an explicit nil

### UnsetSourcePath
`func (o *ResolvedServiceConfig) UnsetSourcePath()`

UnsetSourcePath ensures that no value is present for SourcePath, not even an explicit nil
### GetPlatform

`func (o *ResolvedServiceConfig) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ResolvedServiceConfig) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ResolvedServiceConfig) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetError

`func (o *ResolvedServiceConfig) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ResolvedServiceConfig) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ResolvedServiceConfig) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ResolvedServiceConfig) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ResolvedServiceConfig) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ResolvedServiceConfig) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetFetchedAt

`func (o *ResolvedServiceConfig) GetFetchedAt() time.Time`

GetFetchedAt returns the FetchedAt field if non-nil, zero value otherwise.

### GetFetchedAtOk

`func (o *ResolvedServiceConfig) GetFetchedAtOk() (*time.Time, bool)`

GetFetchedAtOk returns a tuple with the FetchedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchedAt

`func (o *ResolvedServiceConfig) SetFetchedAt(v time.Time)`

SetFetchedAt sets FetchedAt field to given value.

### HasFetchedAt

`func (o *ResolvedServiceConfig) HasFetchedAt() bool`

HasFetchedAt returns a boolean if a field has been set.

### SetFetchedAtNil

`func (o *ResolvedServiceConfig) SetFetchedAtNil(b bool)`

 SetFetchedAtNil sets the value for FetchedAt to be an explicit nil

### UnsetFetchedAt
`func (o *ResolvedServiceConfig) UnsetFetchedAt()`

UnsetFetchedAt ensures that no value is present for FetchedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


