# StackOriginSyncServiceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StackServiceId** | **int32** |  | 
**Name** | **string** |  | 
**Title** | **string** |  | 
**ServiceTitle** | **string** |  | 
**PreviousVersion** | **string** |  | 
**Version** | **string** |  | 
**Reason** | **string** | The origin&#39;s configuration that needs the newer version, such as option 3. | 
**Refusal** | **NullableString** | Why the sync can&#39;t update this stack service, such as a pin. Null when it can. | 

## Methods

### NewStackOriginSyncServiceUpdate

`func NewStackOriginSyncServiceUpdate(stackServiceId int32, name string, title string, serviceTitle string, previousVersion string, version string, reason string, refusal NullableString, ) *StackOriginSyncServiceUpdate`

NewStackOriginSyncServiceUpdate instantiates a new StackOriginSyncServiceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackOriginSyncServiceUpdateWithDefaults

`func NewStackOriginSyncServiceUpdateWithDefaults() *StackOriginSyncServiceUpdate`

NewStackOriginSyncServiceUpdateWithDefaults instantiates a new StackOriginSyncServiceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStackServiceId

`func (o *StackOriginSyncServiceUpdate) GetStackServiceId() int32`

GetStackServiceId returns the StackServiceId field if non-nil, zero value otherwise.

### GetStackServiceIdOk

`func (o *StackOriginSyncServiceUpdate) GetStackServiceIdOk() (*int32, bool)`

GetStackServiceIdOk returns a tuple with the StackServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackServiceId

`func (o *StackOriginSyncServiceUpdate) SetStackServiceId(v int32)`

SetStackServiceId sets StackServiceId field to given value.


### GetName

`func (o *StackOriginSyncServiceUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StackOriginSyncServiceUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StackOriginSyncServiceUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *StackOriginSyncServiceUpdate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StackOriginSyncServiceUpdate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StackOriginSyncServiceUpdate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetServiceTitle

`func (o *StackOriginSyncServiceUpdate) GetServiceTitle() string`

GetServiceTitle returns the ServiceTitle field if non-nil, zero value otherwise.

### GetServiceTitleOk

`func (o *StackOriginSyncServiceUpdate) GetServiceTitleOk() (*string, bool)`

GetServiceTitleOk returns a tuple with the ServiceTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTitle

`func (o *StackOriginSyncServiceUpdate) SetServiceTitle(v string)`

SetServiceTitle sets ServiceTitle field to given value.


### GetPreviousVersion

`func (o *StackOriginSyncServiceUpdate) GetPreviousVersion() string`

GetPreviousVersion returns the PreviousVersion field if non-nil, zero value otherwise.

### GetPreviousVersionOk

`func (o *StackOriginSyncServiceUpdate) GetPreviousVersionOk() (*string, bool)`

GetPreviousVersionOk returns a tuple with the PreviousVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousVersion

`func (o *StackOriginSyncServiceUpdate) SetPreviousVersion(v string)`

SetPreviousVersion sets PreviousVersion field to given value.


### GetVersion

`func (o *StackOriginSyncServiceUpdate) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StackOriginSyncServiceUpdate) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StackOriginSyncServiceUpdate) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetReason

`func (o *StackOriginSyncServiceUpdate) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *StackOriginSyncServiceUpdate) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *StackOriginSyncServiceUpdate) SetReason(v string)`

SetReason sets Reason field to given value.


### GetRefusal

`func (o *StackOriginSyncServiceUpdate) GetRefusal() string`

GetRefusal returns the Refusal field if non-nil, zero value otherwise.

### GetRefusalOk

`func (o *StackOriginSyncServiceUpdate) GetRefusalOk() (*string, bool)`

GetRefusalOk returns a tuple with the Refusal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusal

`func (o *StackOriginSyncServiceUpdate) SetRefusal(v string)`

SetRefusal sets Refusal field to given value.


### SetRefusalNil

`func (o *StackOriginSyncServiceUpdate) SetRefusalNil(b bool)`

 SetRefusalNil sets the value for Refusal to be an explicit nil

### UnsetRefusal
`func (o *StackOriginSyncServiceUpdate) UnsetRefusal()`

UnsetRefusal ensures that no value is present for Refusal, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


