# OrgSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildImageCleanup** | Pointer to [**BuildImageCleanupSettings**](BuildImageCleanupSettings.md) | Default build image retention for the organization&#39;s app environments. | [optional] 

## Methods

### NewOrgSettings

`func NewOrgSettings() *OrgSettings`

NewOrgSettings instantiates a new OrgSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgSettingsWithDefaults

`func NewOrgSettingsWithDefaults() *OrgSettings`

NewOrgSettingsWithDefaults instantiates a new OrgSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildImageCleanup

`func (o *OrgSettings) GetBuildImageCleanup() BuildImageCleanupSettings`

GetBuildImageCleanup returns the BuildImageCleanup field if non-nil, zero value otherwise.

### GetBuildImageCleanupOk

`func (o *OrgSettings) GetBuildImageCleanupOk() (*BuildImageCleanupSettings, bool)`

GetBuildImageCleanupOk returns a tuple with the BuildImageCleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildImageCleanup

`func (o *OrgSettings) SetBuildImageCleanup(v BuildImageCleanupSettings)`

SetBuildImageCleanup sets BuildImageCleanup field to given value.

### HasBuildImageCleanup

`func (o *OrgSettings) HasBuildImageCleanup() bool`

HasBuildImageCleanup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


