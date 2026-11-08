# StackOriginSyncPreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreviousVersion** | **string** |  | 
**Version** | **string** |  | 
**ServiceUpdates** | [**[]StackOriginSyncServiceUpdate**](StackOriginSyncServiceUpdate.md) | Stack services the sync has to update to the service version the origin uses. The sync needs updateStackServices to update them. | 
**Error** | **NullableString** | Why the sync would fail. Null when it would go through once the listed stack services may be updated. | 
**Changes** | [**[]LogLine**](LogLine.md) | What the sync would change, as its task log will list it. Empty when it would fail. | 

## Methods

### NewStackOriginSyncPreview

`func NewStackOriginSyncPreview(previousVersion string, version string, serviceUpdates []StackOriginSyncServiceUpdate, error_ NullableString, changes []LogLine, ) *StackOriginSyncPreview`

NewStackOriginSyncPreview instantiates a new StackOriginSyncPreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackOriginSyncPreviewWithDefaults

`func NewStackOriginSyncPreviewWithDefaults() *StackOriginSyncPreview`

NewStackOriginSyncPreviewWithDefaults instantiates a new StackOriginSyncPreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreviousVersion

`func (o *StackOriginSyncPreview) GetPreviousVersion() string`

GetPreviousVersion returns the PreviousVersion field if non-nil, zero value otherwise.

### GetPreviousVersionOk

`func (o *StackOriginSyncPreview) GetPreviousVersionOk() (*string, bool)`

GetPreviousVersionOk returns a tuple with the PreviousVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousVersion

`func (o *StackOriginSyncPreview) SetPreviousVersion(v string)`

SetPreviousVersion sets PreviousVersion field to given value.


### GetVersion

`func (o *StackOriginSyncPreview) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StackOriginSyncPreview) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StackOriginSyncPreview) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetServiceUpdates

`func (o *StackOriginSyncPreview) GetServiceUpdates() []StackOriginSyncServiceUpdate`

GetServiceUpdates returns the ServiceUpdates field if non-nil, zero value otherwise.

### GetServiceUpdatesOk

`func (o *StackOriginSyncPreview) GetServiceUpdatesOk() (*[]StackOriginSyncServiceUpdate, bool)`

GetServiceUpdatesOk returns a tuple with the ServiceUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUpdates

`func (o *StackOriginSyncPreview) SetServiceUpdates(v []StackOriginSyncServiceUpdate)`

SetServiceUpdates sets ServiceUpdates field to given value.


### GetError

`func (o *StackOriginSyncPreview) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StackOriginSyncPreview) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StackOriginSyncPreview) SetError(v string)`

SetError sets Error field to given value.


### SetErrorNil

`func (o *StackOriginSyncPreview) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *StackOriginSyncPreview) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetChanges

`func (o *StackOriginSyncPreview) GetChanges() []LogLine`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *StackOriginSyncPreview) GetChangesOk() (*[]LogLine, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *StackOriginSyncPreview) SetChanges(v []LogLine)`

SetChanges sets Changes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


