# AppEnvironmentPreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **int32** |  | 
**Title** | **string** |  | 
**Url** | **string** |  | 
**HeadRef** | **string** |  | 
**HeadSha** | **string** |  | 
**BaseRef** | **string** |  | 
**Author** | **string** |  | 
**Draft** | **bool** |  | 
**BaseAppEnvironmentId** | Pointer to **NullableInt32** |  | [optional] 
**LastActivityAt** | **time.Time** |  | 
**ClosedAt** | Pointer to **NullableTime** |  | [optional] 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewAppEnvironmentPreview

`func NewAppEnvironmentPreview(number int32, title string, url string, headRef string, headSha string, baseRef string, author string, draft bool, lastActivityAt time.Time, createdAt time.Time, ) *AppEnvironmentPreview`

NewAppEnvironmentPreview instantiates a new AppEnvironmentPreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppEnvironmentPreviewWithDefaults

`func NewAppEnvironmentPreviewWithDefaults() *AppEnvironmentPreview`

NewAppEnvironmentPreviewWithDefaults instantiates a new AppEnvironmentPreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *AppEnvironmentPreview) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *AppEnvironmentPreview) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *AppEnvironmentPreview) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetTitle

`func (o *AppEnvironmentPreview) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AppEnvironmentPreview) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AppEnvironmentPreview) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUrl

`func (o *AppEnvironmentPreview) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AppEnvironmentPreview) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AppEnvironmentPreview) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHeadRef

`func (o *AppEnvironmentPreview) GetHeadRef() string`

GetHeadRef returns the HeadRef field if non-nil, zero value otherwise.

### GetHeadRefOk

`func (o *AppEnvironmentPreview) GetHeadRefOk() (*string, bool)`

GetHeadRefOk returns a tuple with the HeadRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadRef

`func (o *AppEnvironmentPreview) SetHeadRef(v string)`

SetHeadRef sets HeadRef field to given value.


### GetHeadSha

`func (o *AppEnvironmentPreview) GetHeadSha() string`

GetHeadSha returns the HeadSha field if non-nil, zero value otherwise.

### GetHeadShaOk

`func (o *AppEnvironmentPreview) GetHeadShaOk() (*string, bool)`

GetHeadShaOk returns a tuple with the HeadSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadSha

`func (o *AppEnvironmentPreview) SetHeadSha(v string)`

SetHeadSha sets HeadSha field to given value.


### GetBaseRef

`func (o *AppEnvironmentPreview) GetBaseRef() string`

GetBaseRef returns the BaseRef field if non-nil, zero value otherwise.

### GetBaseRefOk

`func (o *AppEnvironmentPreview) GetBaseRefOk() (*string, bool)`

GetBaseRefOk returns a tuple with the BaseRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseRef

`func (o *AppEnvironmentPreview) SetBaseRef(v string)`

SetBaseRef sets BaseRef field to given value.


### GetAuthor

`func (o *AppEnvironmentPreview) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *AppEnvironmentPreview) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *AppEnvironmentPreview) SetAuthor(v string)`

SetAuthor sets Author field to given value.


### GetDraft

`func (o *AppEnvironmentPreview) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *AppEnvironmentPreview) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *AppEnvironmentPreview) SetDraft(v bool)`

SetDraft sets Draft field to given value.


### GetBaseAppEnvironmentId

`func (o *AppEnvironmentPreview) GetBaseAppEnvironmentId() int32`

GetBaseAppEnvironmentId returns the BaseAppEnvironmentId field if non-nil, zero value otherwise.

### GetBaseAppEnvironmentIdOk

`func (o *AppEnvironmentPreview) GetBaseAppEnvironmentIdOk() (*int32, bool)`

GetBaseAppEnvironmentIdOk returns a tuple with the BaseAppEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAppEnvironmentId

`func (o *AppEnvironmentPreview) SetBaseAppEnvironmentId(v int32)`

SetBaseAppEnvironmentId sets BaseAppEnvironmentId field to given value.

### HasBaseAppEnvironmentId

`func (o *AppEnvironmentPreview) HasBaseAppEnvironmentId() bool`

HasBaseAppEnvironmentId returns a boolean if a field has been set.

### SetBaseAppEnvironmentIdNil

`func (o *AppEnvironmentPreview) SetBaseAppEnvironmentIdNil(b bool)`

 SetBaseAppEnvironmentIdNil sets the value for BaseAppEnvironmentId to be an explicit nil

### UnsetBaseAppEnvironmentId
`func (o *AppEnvironmentPreview) UnsetBaseAppEnvironmentId()`

UnsetBaseAppEnvironmentId ensures that no value is present for BaseAppEnvironmentId, not even an explicit nil
### GetLastActivityAt

`func (o *AppEnvironmentPreview) GetLastActivityAt() time.Time`

GetLastActivityAt returns the LastActivityAt field if non-nil, zero value otherwise.

### GetLastActivityAtOk

`func (o *AppEnvironmentPreview) GetLastActivityAtOk() (*time.Time, bool)`

GetLastActivityAtOk returns a tuple with the LastActivityAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityAt

`func (o *AppEnvironmentPreview) SetLastActivityAt(v time.Time)`

SetLastActivityAt sets LastActivityAt field to given value.


### GetClosedAt

`func (o *AppEnvironmentPreview) GetClosedAt() time.Time`

GetClosedAt returns the ClosedAt field if non-nil, zero value otherwise.

### GetClosedAtOk

`func (o *AppEnvironmentPreview) GetClosedAtOk() (*time.Time, bool)`

GetClosedAtOk returns a tuple with the ClosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedAt

`func (o *AppEnvironmentPreview) SetClosedAt(v time.Time)`

SetClosedAt sets ClosedAt field to given value.

### HasClosedAt

`func (o *AppEnvironmentPreview) HasClosedAt() bool`

HasClosedAt returns a boolean if a field has been set.

### SetClosedAtNil

`func (o *AppEnvironmentPreview) SetClosedAtNil(b bool)`

 SetClosedAtNil sets the value for ClosedAt to be an explicit nil

### UnsetClosedAt
`func (o *AppEnvironmentPreview) UnsetClosedAt()`

UnsetClosedAt ensures that no value is present for ClosedAt, not even an explicit nil
### GetCreatedAt

`func (o *AppEnvironmentPreview) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppEnvironmentPreview) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppEnvironmentPreview) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


