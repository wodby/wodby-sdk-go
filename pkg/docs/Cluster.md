# Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Title** | **string** |  | 
**Status** | **string** |  | 
**Serverless** | **bool** |  | 
**Version** | Pointer to **NullableString** |  | [optional] 
**Region** | Pointer to **NullableString** |  | [optional] 
**Zone** | Pointer to **NullableString** |  | [optional] 
**IntegrationId** | Pointer to **NullableInt32** |  | [optional] 
**OrgId** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewCluster

`func NewCluster(id int32, name string, title string, status string, serverless bool, orgId int32, createdAt time.Time, updatedAt time.Time, ) *Cluster`

NewCluster instantiates a new Cluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterWithDefaults

`func NewClusterWithDefaults() *Cluster`

NewClusterWithDefaults instantiates a new Cluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Cluster) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cluster) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cluster) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Cluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cluster) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *Cluster) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Cluster) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Cluster) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetStatus

`func (o *Cluster) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Cluster) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Cluster) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetServerless

`func (o *Cluster) GetServerless() bool`

GetServerless returns the Serverless field if non-nil, zero value otherwise.

### GetServerlessOk

`func (o *Cluster) GetServerlessOk() (*bool, bool)`

GetServerlessOk returns a tuple with the Serverless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerless

`func (o *Cluster) SetServerless(v bool)`

SetServerless sets Serverless field to given value.


### GetVersion

`func (o *Cluster) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Cluster) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Cluster) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Cluster) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *Cluster) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *Cluster) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetRegion

`func (o *Cluster) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Cluster) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Cluster) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Cluster) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *Cluster) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *Cluster) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetZone

`func (o *Cluster) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *Cluster) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *Cluster) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *Cluster) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *Cluster) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *Cluster) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil
### GetIntegrationId

`func (o *Cluster) GetIntegrationId() int32`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *Cluster) GetIntegrationIdOk() (*int32, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *Cluster) SetIntegrationId(v int32)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *Cluster) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### SetIntegrationIdNil

`func (o *Cluster) SetIntegrationIdNil(b bool)`

 SetIntegrationIdNil sets the value for IntegrationId to be an explicit nil

### UnsetIntegrationId
`func (o *Cluster) UnsetIntegrationId()`

UnsetIntegrationId ensures that no value is present for IntegrationId, not even an explicit nil
### GetOrgId

`func (o *Cluster) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Cluster) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Cluster) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.


### GetCreatedAt

`func (o *Cluster) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Cluster) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Cluster) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Cluster) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Cluster) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Cluster) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


