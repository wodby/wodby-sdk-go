# StackOriginSyncChangelog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreviousVersion** | **string** |  | 
**Version** | **string** |  | 
**Entries** | [**[]StackServiceUpdateChangelogEntry**](StackServiceUpdateChangelogEntry.md) |  | 

## Methods

### NewStackOriginSyncChangelog

`func NewStackOriginSyncChangelog(previousVersion string, version string, entries []StackServiceUpdateChangelogEntry, ) *StackOriginSyncChangelog`

NewStackOriginSyncChangelog instantiates a new StackOriginSyncChangelog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackOriginSyncChangelogWithDefaults

`func NewStackOriginSyncChangelogWithDefaults() *StackOriginSyncChangelog`

NewStackOriginSyncChangelogWithDefaults instantiates a new StackOriginSyncChangelog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreviousVersion

`func (o *StackOriginSyncChangelog) GetPreviousVersion() string`

GetPreviousVersion returns the PreviousVersion field if non-nil, zero value otherwise.

### GetPreviousVersionOk

`func (o *StackOriginSyncChangelog) GetPreviousVersionOk() (*string, bool)`

GetPreviousVersionOk returns a tuple with the PreviousVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousVersion

`func (o *StackOriginSyncChangelog) SetPreviousVersion(v string)`

SetPreviousVersion sets PreviousVersion field to given value.


### GetVersion

`func (o *StackOriginSyncChangelog) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StackOriginSyncChangelog) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StackOriginSyncChangelog) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetEntries

`func (o *StackOriginSyncChangelog) GetEntries() []StackServiceUpdateChangelogEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *StackOriginSyncChangelog) GetEntriesOk() (*[]StackServiceUpdateChangelogEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *StackOriginSyncChangelog) SetEntries(v []StackServiceUpdateChangelogEntry)`

SetEntries sets Entries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


