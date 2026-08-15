# ServiceRevisionChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Title** | **string** |  | 
**Icon** | **string** |  | 
**Kind** | **string** |  | 
**PreviousVersion** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**PreviousRevNumber** | Pointer to **NullableInt32** |  | [optional] 
**RevNumber** | Pointer to **NullableInt32** |  | [optional] 
**Entries** | [**[]StackServiceUpdateChangelogEntry**](StackServiceUpdateChangelogEntry.md) |  | 

## Methods

### NewServiceRevisionChange

`func NewServiceRevisionChange(name string, title string, icon string, kind string, entries []StackServiceUpdateChangelogEntry, ) *ServiceRevisionChange`

NewServiceRevisionChange instantiates a new ServiceRevisionChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceRevisionChangeWithDefaults

`func NewServiceRevisionChangeWithDefaults() *ServiceRevisionChange`

NewServiceRevisionChangeWithDefaults instantiates a new ServiceRevisionChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServiceRevisionChange) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceRevisionChange) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceRevisionChange) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *ServiceRevisionChange) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ServiceRevisionChange) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ServiceRevisionChange) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetIcon

`func (o *ServiceRevisionChange) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *ServiceRevisionChange) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *ServiceRevisionChange) SetIcon(v string)`

SetIcon sets Icon field to given value.


### GetKind

`func (o *ServiceRevisionChange) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ServiceRevisionChange) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ServiceRevisionChange) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPreviousVersion

`func (o *ServiceRevisionChange) GetPreviousVersion() string`

GetPreviousVersion returns the PreviousVersion field if non-nil, zero value otherwise.

### GetPreviousVersionOk

`func (o *ServiceRevisionChange) GetPreviousVersionOk() (*string, bool)`

GetPreviousVersionOk returns a tuple with the PreviousVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousVersion

`func (o *ServiceRevisionChange) SetPreviousVersion(v string)`

SetPreviousVersion sets PreviousVersion field to given value.

### HasPreviousVersion

`func (o *ServiceRevisionChange) HasPreviousVersion() bool`

HasPreviousVersion returns a boolean if a field has been set.

### SetPreviousVersionNil

`func (o *ServiceRevisionChange) SetPreviousVersionNil(b bool)`

 SetPreviousVersionNil sets the value for PreviousVersion to be an explicit nil

### UnsetPreviousVersion
`func (o *ServiceRevisionChange) UnsetPreviousVersion()`

UnsetPreviousVersion ensures that no value is present for PreviousVersion, not even an explicit nil
### GetVersion

`func (o *ServiceRevisionChange) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ServiceRevisionChange) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ServiceRevisionChange) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ServiceRevisionChange) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ServiceRevisionChange) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ServiceRevisionChange) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetPreviousRevNumber

`func (o *ServiceRevisionChange) GetPreviousRevNumber() int32`

GetPreviousRevNumber returns the PreviousRevNumber field if non-nil, zero value otherwise.

### GetPreviousRevNumberOk

`func (o *ServiceRevisionChange) GetPreviousRevNumberOk() (*int32, bool)`

GetPreviousRevNumberOk returns a tuple with the PreviousRevNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousRevNumber

`func (o *ServiceRevisionChange) SetPreviousRevNumber(v int32)`

SetPreviousRevNumber sets PreviousRevNumber field to given value.

### HasPreviousRevNumber

`func (o *ServiceRevisionChange) HasPreviousRevNumber() bool`

HasPreviousRevNumber returns a boolean if a field has been set.

### SetPreviousRevNumberNil

`func (o *ServiceRevisionChange) SetPreviousRevNumberNil(b bool)`

 SetPreviousRevNumberNil sets the value for PreviousRevNumber to be an explicit nil

### UnsetPreviousRevNumber
`func (o *ServiceRevisionChange) UnsetPreviousRevNumber()`

UnsetPreviousRevNumber ensures that no value is present for PreviousRevNumber, not even an explicit nil
### GetRevNumber

`func (o *ServiceRevisionChange) GetRevNumber() int32`

GetRevNumber returns the RevNumber field if non-nil, zero value otherwise.

### GetRevNumberOk

`func (o *ServiceRevisionChange) GetRevNumberOk() (*int32, bool)`

GetRevNumberOk returns a tuple with the RevNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevNumber

`func (o *ServiceRevisionChange) SetRevNumber(v int32)`

SetRevNumber sets RevNumber field to given value.

### HasRevNumber

`func (o *ServiceRevisionChange) HasRevNumber() bool`

HasRevNumber returns a boolean if a field has been set.

### SetRevNumberNil

`func (o *ServiceRevisionChange) SetRevNumberNil(b bool)`

 SetRevNumberNil sets the value for RevNumber to be an explicit nil

### UnsetRevNumber
`func (o *ServiceRevisionChange) UnsetRevNumber()`

UnsetRevNumber ensures that no value is present for RevNumber, not even an explicit nil
### GetEntries

`func (o *ServiceRevisionChange) GetEntries() []StackServiceUpdateChangelogEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *ServiceRevisionChange) GetEntriesOk() (*[]StackServiceUpdateChangelogEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *ServiceRevisionChange) SetEntries(v []StackServiceUpdateChangelogEntry)`

SetEntries sets Entries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


