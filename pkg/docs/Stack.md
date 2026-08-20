# Stack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Title** | **string** |  | 
**Icon** | **string** |  | 
**Status** | **string** |  | 
**Outdated** | **bool** |  | 
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
**Settings** | Pointer to [**StackSettings**](StackSettings.md) |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewStack

`func NewStack(id int32, name string, title string, icon string, status string, outdated bool, public bool, revId int32, latestRevNumber int32, orgId int32, createdAt time.Time, updatedAt time.Time, ) *Stack`

NewStack instantiates a new Stack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackWithDefaults

`func NewStackWithDefaults() *Stack`

NewStackWithDefaults instantiates a new Stack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Stack) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Stack) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Stack) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Stack) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Stack) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Stack) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *Stack) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Stack) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Stack) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetIcon

`func (o *Stack) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *Stack) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *Stack) SetIcon(v string)`

SetIcon sets Icon field to given value.


### GetStatus

`func (o *Stack) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Stack) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Stack) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetOutdated

`func (o *Stack) GetOutdated() bool`

GetOutdated returns the Outdated field if non-nil, zero value otherwise.

### GetOutdatedOk

`func (o *Stack) GetOutdatedOk() (*bool, bool)`

GetOutdatedOk returns a tuple with the Outdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutdated

`func (o *Stack) SetOutdated(v bool)`

SetOutdated sets Outdated field to given value.


### GetPublic

`func (o *Stack) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *Stack) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *Stack) SetPublic(v bool)`

SetPublic sets Public field to given value.


### GetRevId

`func (o *Stack) GetRevId() int32`

GetRevId returns the RevId field if non-nil, zero value otherwise.

### GetRevIdOk

`func (o *Stack) GetRevIdOk() (*int32, bool)`

GetRevIdOk returns a tuple with the RevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevId

`func (o *Stack) SetRevId(v int32)`

SetRevId sets RevId field to given value.


### GetDraftRevId

`func (o *Stack) GetDraftRevId() int32`

GetDraftRevId returns the DraftRevId field if non-nil, zero value otherwise.

### GetDraftRevIdOk

`func (o *Stack) GetDraftRevIdOk() (*int32, bool)`

GetDraftRevIdOk returns a tuple with the DraftRevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftRevId

`func (o *Stack) SetDraftRevId(v int32)`

SetDraftRevId sets DraftRevId field to given value.

### HasDraftRevId

`func (o *Stack) HasDraftRevId() bool`

HasDraftRevId returns a boolean if a field has been set.

### SetDraftRevIdNil

`func (o *Stack) SetDraftRevIdNil(b bool)`

 SetDraftRevIdNil sets the value for DraftRevId to be an explicit nil

### UnsetDraftRevId
`func (o *Stack) UnsetDraftRevId()`

UnsetDraftRevId ensures that no value is present for DraftRevId, not even an explicit nil
### GetLatestRevNumber

`func (o *Stack) GetLatestRevNumber() int32`

GetLatestRevNumber returns the LatestRevNumber field if non-nil, zero value otherwise.

### GetLatestRevNumberOk

`func (o *Stack) GetLatestRevNumberOk() (*int32, bool)`

GetLatestRevNumberOk returns a tuple with the LatestRevNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRevNumber

`func (o *Stack) SetLatestRevNumber(v int32)`

SetLatestRevNumber sets LatestRevNumber field to given value.


### GetGitRepoId

`func (o *Stack) GetGitRepoId() int32`

GetGitRepoId returns the GitRepoId field if non-nil, zero value otherwise.

### GetGitRepoIdOk

`func (o *Stack) GetGitRepoIdOk() (*int32, bool)`

GetGitRepoIdOk returns a tuple with the GitRepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepoId

`func (o *Stack) SetGitRepoId(v int32)`

SetGitRepoId sets GitRepoId field to given value.

### HasGitRepoId

`func (o *Stack) HasGitRepoId() bool`

HasGitRepoId returns a boolean if a field has been set.

### SetGitRepoIdNil

`func (o *Stack) SetGitRepoIdNil(b bool)`

 SetGitRepoIdNil sets the value for GitRepoId to be an explicit nil

### UnsetGitRepoId
`func (o *Stack) UnsetGitRepoId()`

UnsetGitRepoId ensures that no value is present for GitRepoId, not even an explicit nil
### GetGitRepoRemoteId

`func (o *Stack) GetGitRepoRemoteId() string`

GetGitRepoRemoteId returns the GitRepoRemoteId field if non-nil, zero value otherwise.

### GetGitRepoRemoteIdOk

`func (o *Stack) GetGitRepoRemoteIdOk() (*string, bool)`

GetGitRepoRemoteIdOk returns a tuple with the GitRepoRemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepoRemoteId

`func (o *Stack) SetGitRepoRemoteId(v string)`

SetGitRepoRemoteId sets GitRepoRemoteId field to given value.

### HasGitRepoRemoteId

`func (o *Stack) HasGitRepoRemoteId() bool`

HasGitRepoRemoteId returns a boolean if a field has been set.

### SetGitRepoRemoteIdNil

`func (o *Stack) SetGitRepoRemoteIdNil(b bool)`

 SetGitRepoRemoteIdNil sets the value for GitRepoRemoteId to be an explicit nil

### UnsetGitRepoRemoteId
`func (o *Stack) UnsetGitRepoRemoteId()`

UnsetGitRepoRemoteId ensures that no value is present for GitRepoRemoteId, not even an explicit nil
### GetGitRepoRef

`func (o *Stack) GetGitRepoRef() string`

GetGitRepoRef returns the GitRepoRef field if non-nil, zero value otherwise.

### GetGitRepoRefOk

`func (o *Stack) GetGitRepoRefOk() (*string, bool)`

GetGitRepoRefOk returns a tuple with the GitRepoRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepoRef

`func (o *Stack) SetGitRepoRef(v string)`

SetGitRepoRef sets GitRepoRef field to given value.

### HasGitRepoRef

`func (o *Stack) HasGitRepoRef() bool`

HasGitRepoRef returns a boolean if a field has been set.

### SetGitRepoRefNil

`func (o *Stack) SetGitRepoRefNil(b bool)`

 SetGitRepoRefNil sets the value for GitRepoRef to be an explicit nil

### UnsetGitRepoRef
`func (o *Stack) UnsetGitRepoRef()`

UnsetGitRepoRef ensures that no value is present for GitRepoRef, not even an explicit nil
### GetGitRepoRefType

`func (o *Stack) GetGitRepoRefType() string`

GetGitRepoRefType returns the GitRepoRefType field if non-nil, zero value otherwise.

### GetGitRepoRefTypeOk

`func (o *Stack) GetGitRepoRefTypeOk() (*string, bool)`

GetGitRepoRefTypeOk returns a tuple with the GitRepoRefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepoRefType

`func (o *Stack) SetGitRepoRefType(v string)`

SetGitRepoRefType sets GitRepoRefType field to given value.

### HasGitRepoRefType

`func (o *Stack) HasGitRepoRefType() bool`

HasGitRepoRefType returns a boolean if a field has been set.

### SetGitRepoRefTypeNil

`func (o *Stack) SetGitRepoRefTypeNil(b bool)`

 SetGitRepoRefTypeNil sets the value for GitRepoRefType to be an explicit nil

### UnsetGitRepoRefType
`func (o *Stack) UnsetGitRepoRefType()`

UnsetGitRepoRefType ensures that no value is present for GitRepoRefType, not even an explicit nil
### GetOriginStackRevId

`func (o *Stack) GetOriginStackRevId() int32`

GetOriginStackRevId returns the OriginStackRevId field if non-nil, zero value otherwise.

### GetOriginStackRevIdOk

`func (o *Stack) GetOriginStackRevIdOk() (*int32, bool)`

GetOriginStackRevIdOk returns a tuple with the OriginStackRevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStackRevId

`func (o *Stack) SetOriginStackRevId(v int32)`

SetOriginStackRevId sets OriginStackRevId field to given value.

### HasOriginStackRevId

`func (o *Stack) HasOriginStackRevId() bool`

HasOriginStackRevId returns a boolean if a field has been set.

### SetOriginStackRevIdNil

`func (o *Stack) SetOriginStackRevIdNil(b bool)`

 SetOriginStackRevIdNil sets the value for OriginStackRevId to be an explicit nil

### UnsetOriginStackRevId
`func (o *Stack) UnsetOriginStackRevId()`

UnsetOriginStackRevId ensures that no value is present for OriginStackRevId, not even an explicit nil
### GetOriginStackRevStackId

`func (o *Stack) GetOriginStackRevStackId() int32`

GetOriginStackRevStackId returns the OriginStackRevStackId field if non-nil, zero value otherwise.

### GetOriginStackRevStackIdOk

`func (o *Stack) GetOriginStackRevStackIdOk() (*int32, bool)`

GetOriginStackRevStackIdOk returns a tuple with the OriginStackRevStackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStackRevStackId

`func (o *Stack) SetOriginStackRevStackId(v int32)`

SetOriginStackRevStackId sets OriginStackRevStackId field to given value.

### HasOriginStackRevStackId

`func (o *Stack) HasOriginStackRevStackId() bool`

HasOriginStackRevStackId returns a boolean if a field has been set.

### SetOriginStackRevStackIdNil

`func (o *Stack) SetOriginStackRevStackIdNil(b bool)`

 SetOriginStackRevStackIdNil sets the value for OriginStackRevStackId to be an explicit nil

### UnsetOriginStackRevStackId
`func (o *Stack) UnsetOriginStackRevStackId()`

UnsetOriginStackRevStackId ensures that no value is present for OriginStackRevStackId, not even an explicit nil
### GetOriginStackRevName

`func (o *Stack) GetOriginStackRevName() string`

GetOriginStackRevName returns the OriginStackRevName field if non-nil, zero value otherwise.

### GetOriginStackRevNameOk

`func (o *Stack) GetOriginStackRevNameOk() (*string, bool)`

GetOriginStackRevNameOk returns a tuple with the OriginStackRevName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStackRevName

`func (o *Stack) SetOriginStackRevName(v string)`

SetOriginStackRevName sets OriginStackRevName field to given value.

### HasOriginStackRevName

`func (o *Stack) HasOriginStackRevName() bool`

HasOriginStackRevName returns a boolean if a field has been set.

### SetOriginStackRevNameNil

`func (o *Stack) SetOriginStackRevNameNil(b bool)`

 SetOriginStackRevNameNil sets the value for OriginStackRevName to be an explicit nil

### UnsetOriginStackRevName
`func (o *Stack) UnsetOriginStackRevName()`

UnsetOriginStackRevName ensures that no value is present for OriginStackRevName, not even an explicit nil
### GetOriginStackRevNumber

`func (o *Stack) GetOriginStackRevNumber() int32`

GetOriginStackRevNumber returns the OriginStackRevNumber field if non-nil, zero value otherwise.

### GetOriginStackRevNumberOk

`func (o *Stack) GetOriginStackRevNumberOk() (*int32, bool)`

GetOriginStackRevNumberOk returns a tuple with the OriginStackRevNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStackRevNumber

`func (o *Stack) SetOriginStackRevNumber(v int32)`

SetOriginStackRevNumber sets OriginStackRevNumber field to given value.

### HasOriginStackRevNumber

`func (o *Stack) HasOriginStackRevNumber() bool`

HasOriginStackRevNumber returns a boolean if a field has been set.

### SetOriginStackRevNumberNil

`func (o *Stack) SetOriginStackRevNumberNil(b bool)`

 SetOriginStackRevNumberNil sets the value for OriginStackRevNumber to be an explicit nil

### UnsetOriginStackRevNumber
`func (o *Stack) UnsetOriginStackRevNumber()`

UnsetOriginStackRevNumber ensures that no value is present for OriginStackRevNumber, not even an explicit nil
### GetOriginStackRevVersion

`func (o *Stack) GetOriginStackRevVersion() string`

GetOriginStackRevVersion returns the OriginStackRevVersion field if non-nil, zero value otherwise.

### GetOriginStackRevVersionOk

`func (o *Stack) GetOriginStackRevVersionOk() (*string, bool)`

GetOriginStackRevVersionOk returns a tuple with the OriginStackRevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStackRevVersion

`func (o *Stack) SetOriginStackRevVersion(v string)`

SetOriginStackRevVersion sets OriginStackRevVersion field to given value.

### HasOriginStackRevVersion

`func (o *Stack) HasOriginStackRevVersion() bool`

HasOriginStackRevVersion returns a boolean if a field has been set.

### SetOriginStackRevVersionNil

`func (o *Stack) SetOriginStackRevVersionNil(b bool)`

 SetOriginStackRevVersionNil sets the value for OriginStackRevVersion to be an explicit nil

### UnsetOriginStackRevVersion
`func (o *Stack) UnsetOriginStackRevVersion()`

UnsetOriginStackRevVersion ensures that no value is present for OriginStackRevVersion, not even an explicit nil
### GetOriginStackRevCreatedAt

`func (o *Stack) GetOriginStackRevCreatedAt() time.Time`

GetOriginStackRevCreatedAt returns the OriginStackRevCreatedAt field if non-nil, zero value otherwise.

### GetOriginStackRevCreatedAtOk

`func (o *Stack) GetOriginStackRevCreatedAtOk() (*time.Time, bool)`

GetOriginStackRevCreatedAtOk returns a tuple with the OriginStackRevCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStackRevCreatedAt

`func (o *Stack) SetOriginStackRevCreatedAt(v time.Time)`

SetOriginStackRevCreatedAt sets OriginStackRevCreatedAt field to given value.

### HasOriginStackRevCreatedAt

`func (o *Stack) HasOriginStackRevCreatedAt() bool`

HasOriginStackRevCreatedAt returns a boolean if a field has been set.

### SetOriginStackRevCreatedAtNil

`func (o *Stack) SetOriginStackRevCreatedAtNil(b bool)`

 SetOriginStackRevCreatedAtNil sets the value for OriginStackRevCreatedAt to be an explicit nil

### UnsetOriginStackRevCreatedAt
`func (o *Stack) UnsetOriginStackRevCreatedAt()`

UnsetOriginStackRevCreatedAt ensures that no value is present for OriginStackRevCreatedAt, not even an explicit nil
### GetOrgId

`func (o *Stack) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Stack) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Stack) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.


### GetSettings

`func (o *Stack) GetSettings() StackSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *Stack) GetSettingsOk() (*StackSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *Stack) SetSettings(v StackSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *Stack) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Stack) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Stack) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Stack) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Stack) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Stack) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Stack) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


