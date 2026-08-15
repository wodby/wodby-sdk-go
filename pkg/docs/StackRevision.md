# StackRevision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Title** | **string** |  | 
**Icon** | **string** |  | 
**Number** | **int32** |  | 
**Draft** | **bool** |  | 
**Version** | **string** |  | 
**StackId** | **int32** |  | 
**Manifest** | **string** |  | 
**LinkIssues** | [**[]StackRevisionLinkIssue**](StackRevisionLinkIssue.md) |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewStackRevision

`func NewStackRevision(id int32, name string, title string, icon string, number int32, draft bool, version string, stackId int32, manifest string, linkIssues []StackRevisionLinkIssue, createdAt time.Time, ) *StackRevision`

NewStackRevision instantiates a new StackRevision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackRevisionWithDefaults

`func NewStackRevisionWithDefaults() *StackRevision`

NewStackRevisionWithDefaults instantiates a new StackRevision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StackRevision) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StackRevision) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StackRevision) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *StackRevision) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StackRevision) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StackRevision) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *StackRevision) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StackRevision) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StackRevision) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetIcon

`func (o *StackRevision) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *StackRevision) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *StackRevision) SetIcon(v string)`

SetIcon sets Icon field to given value.


### GetNumber

`func (o *StackRevision) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *StackRevision) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *StackRevision) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetDraft

`func (o *StackRevision) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *StackRevision) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *StackRevision) SetDraft(v bool)`

SetDraft sets Draft field to given value.


### GetVersion

`func (o *StackRevision) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StackRevision) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StackRevision) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetStackId

`func (o *StackRevision) GetStackId() int32`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *StackRevision) GetStackIdOk() (*int32, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *StackRevision) SetStackId(v int32)`

SetStackId sets StackId field to given value.


### GetManifest

`func (o *StackRevision) GetManifest() string`

GetManifest returns the Manifest field if non-nil, zero value otherwise.

### GetManifestOk

`func (o *StackRevision) GetManifestOk() (*string, bool)`

GetManifestOk returns a tuple with the Manifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifest

`func (o *StackRevision) SetManifest(v string)`

SetManifest sets Manifest field to given value.


### GetLinkIssues

`func (o *StackRevision) GetLinkIssues() []StackRevisionLinkIssue`

GetLinkIssues returns the LinkIssues field if non-nil, zero value otherwise.

### GetLinkIssuesOk

`func (o *StackRevision) GetLinkIssuesOk() (*[]StackRevisionLinkIssue, bool)`

GetLinkIssuesOk returns a tuple with the LinkIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkIssues

`func (o *StackRevision) SetLinkIssues(v []StackRevisionLinkIssue)`

SetLinkIssues sets LinkIssues field to given value.


### GetCreatedAt

`func (o *StackRevision) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StackRevision) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StackRevision) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


