# Stack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Title** | **string** |  | 
**Icon** | **string** |  | 
**Status** | **string** |  | 
**Public** | **bool** |  | 
**RevId** | **int32** |  | 
**LatestRevNumber** | **int32** |  | 
**OrgId** | **int32** |  | 
**Settings** | Pointer to [**StackSettings**](StackSettings.md) |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewStack

`func NewStack(id int32, name string, title string, icon string, status string, public bool, revId int32, latestRevNumber int32, orgId int32, createdAt time.Time, updatedAt time.Time, ) *Stack`

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


