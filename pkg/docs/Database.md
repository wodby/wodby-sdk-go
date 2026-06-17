# Database

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Title** | **string** |  | 
**Type** | **string** |  | 
**Kind** | **string** |  | 
**Status** | **string** |  | 
**Version** | **string** |  | 
**Region** | Pointer to **NullableString** |  | [optional] 
**Zone** | Pointer to **NullableString** |  | [optional] 
**IntegrationId** | Pointer to **NullableInt32** |  | [optional] 
**EnvId** | **int32** |  | 
**OrgId** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewDatabase

`func NewDatabase(id int32, name string, title string, type_ string, kind string, status string, version string, envId int32, orgId int32, createdAt time.Time, updatedAt time.Time, ) *Database`

NewDatabase instantiates a new Database object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseWithDefaults

`func NewDatabaseWithDefaults() *Database`

NewDatabaseWithDefaults instantiates a new Database object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Database) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Database) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Database) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Database) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Database) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Database) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *Database) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Database) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Database) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *Database) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Database) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Database) SetType(v string)`

SetType sets Type field to given value.


### GetKind

`func (o *Database) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Database) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Database) SetKind(v string)`

SetKind sets Kind field to given value.


### GetStatus

`func (o *Database) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Database) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Database) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetVersion

`func (o *Database) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Database) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Database) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetRegion

`func (o *Database) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Database) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Database) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Database) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *Database) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *Database) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetZone

`func (o *Database) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *Database) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *Database) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *Database) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *Database) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *Database) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil
### GetIntegrationId

`func (o *Database) GetIntegrationId() int32`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *Database) GetIntegrationIdOk() (*int32, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *Database) SetIntegrationId(v int32)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *Database) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### SetIntegrationIdNil

`func (o *Database) SetIntegrationIdNil(b bool)`

 SetIntegrationIdNil sets the value for IntegrationId to be an explicit nil

### UnsetIntegrationId
`func (o *Database) UnsetIntegrationId()`

UnsetIntegrationId ensures that no value is present for IntegrationId, not even an explicit nil
### GetEnvId

`func (o *Database) GetEnvId() int32`

GetEnvId returns the EnvId field if non-nil, zero value otherwise.

### GetEnvIdOk

`func (o *Database) GetEnvIdOk() (*int32, bool)`

GetEnvIdOk returns a tuple with the EnvId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvId

`func (o *Database) SetEnvId(v int32)`

SetEnvId sets EnvId field to given value.


### GetOrgId

`func (o *Database) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Database) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Database) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.


### GetCreatedAt

`func (o *Database) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Database) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Database) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Database) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Database) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Database) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


