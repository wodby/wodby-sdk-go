# ProviderRevision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Title** | **string** |  | 
**Number** | **int32** |  | 
**Version** | **string** |  | 
**ProviderId** | **int32** |  | 
**Manifest** | Pointer to **map[string]interface{}** |  | [optional] 
**PermissionAudit** | **bool** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewProviderRevision

`func NewProviderRevision(id int32, name string, title string, number int32, version string, providerId int32, permissionAudit bool, createdAt time.Time, ) *ProviderRevision`

NewProviderRevision instantiates a new ProviderRevision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderRevisionWithDefaults

`func NewProviderRevisionWithDefaults() *ProviderRevision`

NewProviderRevisionWithDefaults instantiates a new ProviderRevision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProviderRevision) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProviderRevision) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProviderRevision) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ProviderRevision) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderRevision) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderRevision) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *ProviderRevision) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProviderRevision) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProviderRevision) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetNumber

`func (o *ProviderRevision) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *ProviderRevision) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *ProviderRevision) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetVersion

`func (o *ProviderRevision) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ProviderRevision) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ProviderRevision) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetProviderId

`func (o *ProviderRevision) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *ProviderRevision) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *ProviderRevision) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.


### GetManifest

`func (o *ProviderRevision) GetManifest() map[string]interface{}`

GetManifest returns the Manifest field if non-nil, zero value otherwise.

### GetManifestOk

`func (o *ProviderRevision) GetManifestOk() (*map[string]interface{}, bool)`

GetManifestOk returns a tuple with the Manifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifest

`func (o *ProviderRevision) SetManifest(v map[string]interface{})`

SetManifest sets Manifest field to given value.

### HasManifest

`func (o *ProviderRevision) HasManifest() bool`

HasManifest returns a boolean if a field has been set.

### GetPermissionAudit

`func (o *ProviderRevision) GetPermissionAudit() bool`

GetPermissionAudit returns the PermissionAudit field if non-nil, zero value otherwise.

### GetPermissionAuditOk

`func (o *ProviderRevision) GetPermissionAuditOk() (*bool, bool)`

GetPermissionAuditOk returns a tuple with the PermissionAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionAudit

`func (o *ProviderRevision) SetPermissionAudit(v bool)`

SetPermissionAudit sets PermissionAudit field to given value.


### GetCreatedAt

`func (o *ProviderRevision) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProviderRevision) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProviderRevision) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


