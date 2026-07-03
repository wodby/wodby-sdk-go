# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Title** | **string** |  | 
**Type** | **string** |  | 
**Status** | **string** |  | 
**External** | **bool** |  | 
**Public** | **bool** |  | 
**RevId** | **int32** |  | 
**DraftRevId** | Pointer to **NullableInt32** |  | [optional] 
**LatestRevNumber** | **int32** |  | 
**GitRepoId** | Pointer to **NullableInt32** |  | [optional] 
**GitRepoRemoteId** | Pointer to **NullableString** |  | [optional] 
**GitRepoRef** | Pointer to **NullableString** |  | [optional] 
**GitRepoRefType** | Pointer to **NullableString** |  | [optional] 
**OriginStackRevId** | Pointer to **NullableInt32** |  | [optional] 
**OriginStackRevStackId** | Pointer to **NullableInt32** |  | [optional] 
**OriginStackRevName** | Pointer to **NullableString** |  | [optional] 
**OriginStackRevNumber** | Pointer to **NullableInt32** |  | [optional] 
**OriginStackRevVersion** | Pointer to **NullableString** |  | [optional] 
**OriginStackRevCreatedAt** | Pointer to **NullableTime** |  | [optional] 
**OrgId** | **int32** |  | 
**Settings** | Pointer to [**ServiceSettings**](ServiceSettings.md) |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewService

`func NewService(id int32, name string, title string, type_ string, status string, external bool, public bool, revId int32, latestRevNumber int32, orgId int32, createdAt time.Time, updatedAt time.Time, ) *Service`

NewService instantiates a new Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWithDefaults

`func NewServiceWithDefaults() *Service`

NewServiceWithDefaults instantiates a new Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Service) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Service) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Service) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Service) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Service) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Service) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *Service) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Service) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Service) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *Service) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Service) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Service) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *Service) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Service) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Service) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetExternal

`func (o *Service) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *Service) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *Service) SetExternal(v bool)`

SetExternal sets External field to given value.


### GetPublic

`func (o *Service) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *Service) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *Service) SetPublic(v bool)`

SetPublic sets Public field to given value.


### GetRevId

`func (o *Service) GetRevId() int32`

GetRevId returns the RevId field if non-nil, zero value otherwise.

### GetRevIdOk

`func (o *Service) GetRevIdOk() (*int32, bool)`

GetRevIdOk returns a tuple with the RevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevId

`func (o *Service) SetRevId(v int32)`

SetRevId sets RevId field to given value.


### GetDraftRevId

`func (o *Service) GetDraftRevId() int32`

GetDraftRevId returns the DraftRevId field if non-nil, zero value otherwise.

### GetDraftRevIdOk

`func (o *Service) GetDraftRevIdOk() (*int32, bool)`

GetDraftRevIdOk returns a tuple with the DraftRevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftRevId

`func (o *Service) SetDraftRevId(v int32)`

SetDraftRevId sets DraftRevId field to given value.

### HasDraftRevId

`func (o *Service) HasDraftRevId() bool`

HasDraftRevId returns a boolean if a field has been set.

### SetDraftRevIdNil

`func (o *Service) SetDraftRevIdNil(b bool)`

 SetDraftRevIdNil sets the value for DraftRevId to be an explicit nil

### UnsetDraftRevId
`func (o *Service) UnsetDraftRevId()`

UnsetDraftRevId ensures that no value is present for DraftRevId, not even an explicit nil
### GetLatestRevNumber

`func (o *Service) GetLatestRevNumber() int32`

GetLatestRevNumber returns the LatestRevNumber field if non-nil, zero value otherwise.

### GetLatestRevNumberOk

`func (o *Service) GetLatestRevNumberOk() (*int32, bool)`

GetLatestRevNumberOk returns a tuple with the LatestRevNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRevNumber

`func (o *Service) SetLatestRevNumber(v int32)`

SetLatestRevNumber sets LatestRevNumber field to given value.


### GetGitRepoId

`func (o *Service) GetGitRepoId() int32`

GetGitRepoId returns the GitRepoId field if non-nil, zero value otherwise.

### GetGitRepoIdOk

`func (o *Service) GetGitRepoIdOk() (*int32, bool)`

GetGitRepoIdOk returns a tuple with the GitRepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepoId

`func (o *Service) SetGitRepoId(v int32)`

SetGitRepoId sets GitRepoId field to given value.

### HasGitRepoId

`func (o *Service) HasGitRepoId() bool`

HasGitRepoId returns a boolean if a field has been set.

### SetGitRepoIdNil

`func (o *Service) SetGitRepoIdNil(b bool)`

 SetGitRepoIdNil sets the value for GitRepoId to be an explicit nil

### UnsetGitRepoId
`func (o *Service) UnsetGitRepoId()`

UnsetGitRepoId ensures that no value is present for GitRepoId, not even an explicit nil
### GetGitRepoRemoteId

`func (o *Service) GetGitRepoRemoteId() string`

GetGitRepoRemoteId returns the GitRepoRemoteId field if non-nil, zero value otherwise.

### GetGitRepoRemoteIdOk

`func (o *Service) GetGitRepoRemoteIdOk() (*string, bool)`

GetGitRepoRemoteIdOk returns a tuple with the GitRepoRemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepoRemoteId

`func (o *Service) SetGitRepoRemoteId(v string)`

SetGitRepoRemoteId sets GitRepoRemoteId field to given value.

### HasGitRepoRemoteId

`func (o *Service) HasGitRepoRemoteId() bool`

HasGitRepoRemoteId returns a boolean if a field has been set.

### SetGitRepoRemoteIdNil

`func (o *Service) SetGitRepoRemoteIdNil(b bool)`

 SetGitRepoRemoteIdNil sets the value for GitRepoRemoteId to be an explicit nil

### UnsetGitRepoRemoteId
`func (o *Service) UnsetGitRepoRemoteId()`

UnsetGitRepoRemoteId ensures that no value is present for GitRepoRemoteId, not even an explicit nil
### GetGitRepoRef

`func (o *Service) GetGitRepoRef() string`

GetGitRepoRef returns the GitRepoRef field if non-nil, zero value otherwise.

### GetGitRepoRefOk

`func (o *Service) GetGitRepoRefOk() (*string, bool)`

GetGitRepoRefOk returns a tuple with the GitRepoRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepoRef

`func (o *Service) SetGitRepoRef(v string)`

SetGitRepoRef sets GitRepoRef field to given value.

### HasGitRepoRef

`func (o *Service) HasGitRepoRef() bool`

HasGitRepoRef returns a boolean if a field has been set.

### SetGitRepoRefNil

`func (o *Service) SetGitRepoRefNil(b bool)`

 SetGitRepoRefNil sets the value for GitRepoRef to be an explicit nil

### UnsetGitRepoRef
`func (o *Service) UnsetGitRepoRef()`

UnsetGitRepoRef ensures that no value is present for GitRepoRef, not even an explicit nil
### GetGitRepoRefType

`func (o *Service) GetGitRepoRefType() string`

GetGitRepoRefType returns the GitRepoRefType field if non-nil, zero value otherwise.

### GetGitRepoRefTypeOk

`func (o *Service) GetGitRepoRefTypeOk() (*string, bool)`

GetGitRepoRefTypeOk returns a tuple with the GitRepoRefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepoRefType

`func (o *Service) SetGitRepoRefType(v string)`

SetGitRepoRefType sets GitRepoRefType field to given value.

### HasGitRepoRefType

`func (o *Service) HasGitRepoRefType() bool`

HasGitRepoRefType returns a boolean if a field has been set.

### SetGitRepoRefTypeNil

`func (o *Service) SetGitRepoRefTypeNil(b bool)`

 SetGitRepoRefTypeNil sets the value for GitRepoRefType to be an explicit nil

### UnsetGitRepoRefType
`func (o *Service) UnsetGitRepoRefType()`

UnsetGitRepoRefType ensures that no value is present for GitRepoRefType, not even an explicit nil
### GetOriginStackRevId

`func (o *Service) GetOriginStackRevId() int32`

GetOriginStackRevId returns the OriginStackRevId field if non-nil, zero value otherwise.

### GetOriginStackRevIdOk

`func (o *Service) GetOriginStackRevIdOk() (*int32, bool)`

GetOriginStackRevIdOk returns a tuple with the OriginStackRevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStackRevId

`func (o *Service) SetOriginStackRevId(v int32)`

SetOriginStackRevId sets OriginStackRevId field to given value.

### HasOriginStackRevId

`func (o *Service) HasOriginStackRevId() bool`

HasOriginStackRevId returns a boolean if a field has been set.

### SetOriginStackRevIdNil

`func (o *Service) SetOriginStackRevIdNil(b bool)`

 SetOriginStackRevIdNil sets the value for OriginStackRevId to be an explicit nil

### UnsetOriginStackRevId
`func (o *Service) UnsetOriginStackRevId()`

UnsetOriginStackRevId ensures that no value is present for OriginStackRevId, not even an explicit nil
### GetOriginStackRevStackId

`func (o *Service) GetOriginStackRevStackId() int32`

GetOriginStackRevStackId returns the OriginStackRevStackId field if non-nil, zero value otherwise.

### GetOriginStackRevStackIdOk

`func (o *Service) GetOriginStackRevStackIdOk() (*int32, bool)`

GetOriginStackRevStackIdOk returns a tuple with the OriginStackRevStackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStackRevStackId

`func (o *Service) SetOriginStackRevStackId(v int32)`

SetOriginStackRevStackId sets OriginStackRevStackId field to given value.

### HasOriginStackRevStackId

`func (o *Service) HasOriginStackRevStackId() bool`

HasOriginStackRevStackId returns a boolean if a field has been set.

### SetOriginStackRevStackIdNil

`func (o *Service) SetOriginStackRevStackIdNil(b bool)`

 SetOriginStackRevStackIdNil sets the value for OriginStackRevStackId to be an explicit nil

### UnsetOriginStackRevStackId
`func (o *Service) UnsetOriginStackRevStackId()`

UnsetOriginStackRevStackId ensures that no value is present for OriginStackRevStackId, not even an explicit nil
### GetOriginStackRevName

`func (o *Service) GetOriginStackRevName() string`

GetOriginStackRevName returns the OriginStackRevName field if non-nil, zero value otherwise.

### GetOriginStackRevNameOk

`func (o *Service) GetOriginStackRevNameOk() (*string, bool)`

GetOriginStackRevNameOk returns a tuple with the OriginStackRevName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStackRevName

`func (o *Service) SetOriginStackRevName(v string)`

SetOriginStackRevName sets OriginStackRevName field to given value.

### HasOriginStackRevName

`func (o *Service) HasOriginStackRevName() bool`

HasOriginStackRevName returns a boolean if a field has been set.

### SetOriginStackRevNameNil

`func (o *Service) SetOriginStackRevNameNil(b bool)`

 SetOriginStackRevNameNil sets the value for OriginStackRevName to be an explicit nil

### UnsetOriginStackRevName
`func (o *Service) UnsetOriginStackRevName()`

UnsetOriginStackRevName ensures that no value is present for OriginStackRevName, not even an explicit nil
### GetOriginStackRevNumber

`func (o *Service) GetOriginStackRevNumber() int32`

GetOriginStackRevNumber returns the OriginStackRevNumber field if non-nil, zero value otherwise.

### GetOriginStackRevNumberOk

`func (o *Service) GetOriginStackRevNumberOk() (*int32, bool)`

GetOriginStackRevNumberOk returns a tuple with the OriginStackRevNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStackRevNumber

`func (o *Service) SetOriginStackRevNumber(v int32)`

SetOriginStackRevNumber sets OriginStackRevNumber field to given value.

### HasOriginStackRevNumber

`func (o *Service) HasOriginStackRevNumber() bool`

HasOriginStackRevNumber returns a boolean if a field has been set.

### SetOriginStackRevNumberNil

`func (o *Service) SetOriginStackRevNumberNil(b bool)`

 SetOriginStackRevNumberNil sets the value for OriginStackRevNumber to be an explicit nil

### UnsetOriginStackRevNumber
`func (o *Service) UnsetOriginStackRevNumber()`

UnsetOriginStackRevNumber ensures that no value is present for OriginStackRevNumber, not even an explicit nil
### GetOriginStackRevVersion

`func (o *Service) GetOriginStackRevVersion() string`

GetOriginStackRevVersion returns the OriginStackRevVersion field if non-nil, zero value otherwise.

### GetOriginStackRevVersionOk

`func (o *Service) GetOriginStackRevVersionOk() (*string, bool)`

GetOriginStackRevVersionOk returns a tuple with the OriginStackRevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStackRevVersion

`func (o *Service) SetOriginStackRevVersion(v string)`

SetOriginStackRevVersion sets OriginStackRevVersion field to given value.

### HasOriginStackRevVersion

`func (o *Service) HasOriginStackRevVersion() bool`

HasOriginStackRevVersion returns a boolean if a field has been set.

### SetOriginStackRevVersionNil

`func (o *Service) SetOriginStackRevVersionNil(b bool)`

 SetOriginStackRevVersionNil sets the value for OriginStackRevVersion to be an explicit nil

### UnsetOriginStackRevVersion
`func (o *Service) UnsetOriginStackRevVersion()`

UnsetOriginStackRevVersion ensures that no value is present for OriginStackRevVersion, not even an explicit nil
### GetOriginStackRevCreatedAt

`func (o *Service) GetOriginStackRevCreatedAt() time.Time`

GetOriginStackRevCreatedAt returns the OriginStackRevCreatedAt field if non-nil, zero value otherwise.

### GetOriginStackRevCreatedAtOk

`func (o *Service) GetOriginStackRevCreatedAtOk() (*time.Time, bool)`

GetOriginStackRevCreatedAtOk returns a tuple with the OriginStackRevCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStackRevCreatedAt

`func (o *Service) SetOriginStackRevCreatedAt(v time.Time)`

SetOriginStackRevCreatedAt sets OriginStackRevCreatedAt field to given value.

### HasOriginStackRevCreatedAt

`func (o *Service) HasOriginStackRevCreatedAt() bool`

HasOriginStackRevCreatedAt returns a boolean if a field has been set.

### SetOriginStackRevCreatedAtNil

`func (o *Service) SetOriginStackRevCreatedAtNil(b bool)`

 SetOriginStackRevCreatedAtNil sets the value for OriginStackRevCreatedAt to be an explicit nil

### UnsetOriginStackRevCreatedAt
`func (o *Service) UnsetOriginStackRevCreatedAt()`

UnsetOriginStackRevCreatedAt ensures that no value is present for OriginStackRevCreatedAt, not even an explicit nil
### GetOrgId

`func (o *Service) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Service) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Service) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.


### GetSettings

`func (o *Service) GetSettings() ServiceSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *Service) GetSettingsOk() (*ServiceSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *Service) SetSettings(v ServiceSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *Service) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Service) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Service) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Service) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Service) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Service) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Service) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


