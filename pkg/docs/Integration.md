# Integration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Title** | **string** |  | 
**Status** | **string** |  | 
**Scope** | Pointer to **NullableString** |  | [optional] 
**Auth** | Pointer to **NullableString** |  | [optional] 
**ProviderRevId** | **int32** |  | 
**OrgId** | **int32** |  | 
**PrimaryEnvId** | Pointer to **NullableInt32** |  | [optional] 
**EnvScope** | **string** |  | 
**AllowedEnvIds** | **[]int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewIntegration

`func NewIntegration(id int32, title string, status string, providerRevId int32, orgId int32, envScope string, allowedEnvIds []int32, createdAt time.Time, updatedAt time.Time, ) *Integration`

NewIntegration instantiates a new Integration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationWithDefaults

`func NewIntegrationWithDefaults() *Integration`

NewIntegrationWithDefaults instantiates a new Integration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Integration) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Integration) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Integration) SetId(v int32)`

SetId sets Id field to given value.


### GetTitle

`func (o *Integration) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Integration) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Integration) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetStatus

`func (o *Integration) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Integration) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Integration) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetScope

`func (o *Integration) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Integration) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Integration) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *Integration) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *Integration) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *Integration) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetAuth

`func (o *Integration) GetAuth() string`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *Integration) GetAuthOk() (*string, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *Integration) SetAuth(v string)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *Integration) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### SetAuthNil

`func (o *Integration) SetAuthNil(b bool)`

 SetAuthNil sets the value for Auth to be an explicit nil

### UnsetAuth
`func (o *Integration) UnsetAuth()`

UnsetAuth ensures that no value is present for Auth, not even an explicit nil
### GetProviderRevId

`func (o *Integration) GetProviderRevId() int32`

GetProviderRevId returns the ProviderRevId field if non-nil, zero value otherwise.

### GetProviderRevIdOk

`func (o *Integration) GetProviderRevIdOk() (*int32, bool)`

GetProviderRevIdOk returns a tuple with the ProviderRevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderRevId

`func (o *Integration) SetProviderRevId(v int32)`

SetProviderRevId sets ProviderRevId field to given value.


### GetOrgId

`func (o *Integration) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Integration) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Integration) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.


### GetPrimaryEnvId

`func (o *Integration) GetPrimaryEnvId() int32`

GetPrimaryEnvId returns the PrimaryEnvId field if non-nil, zero value otherwise.

### GetPrimaryEnvIdOk

`func (o *Integration) GetPrimaryEnvIdOk() (*int32, bool)`

GetPrimaryEnvIdOk returns a tuple with the PrimaryEnvId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryEnvId

`func (o *Integration) SetPrimaryEnvId(v int32)`

SetPrimaryEnvId sets PrimaryEnvId field to given value.

### HasPrimaryEnvId

`func (o *Integration) HasPrimaryEnvId() bool`

HasPrimaryEnvId returns a boolean if a field has been set.

### SetPrimaryEnvIdNil

`func (o *Integration) SetPrimaryEnvIdNil(b bool)`

 SetPrimaryEnvIdNil sets the value for PrimaryEnvId to be an explicit nil

### UnsetPrimaryEnvId
`func (o *Integration) UnsetPrimaryEnvId()`

UnsetPrimaryEnvId ensures that no value is present for PrimaryEnvId, not even an explicit nil
### GetEnvScope

`func (o *Integration) GetEnvScope() string`

GetEnvScope returns the EnvScope field if non-nil, zero value otherwise.

### GetEnvScopeOk

`func (o *Integration) GetEnvScopeOk() (*string, bool)`

GetEnvScopeOk returns a tuple with the EnvScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvScope

`func (o *Integration) SetEnvScope(v string)`

SetEnvScope sets EnvScope field to given value.


### GetAllowedEnvIds

`func (o *Integration) GetAllowedEnvIds() []int32`

GetAllowedEnvIds returns the AllowedEnvIds field if non-nil, zero value otherwise.

### GetAllowedEnvIdsOk

`func (o *Integration) GetAllowedEnvIdsOk() (*[]int32, bool)`

GetAllowedEnvIdsOk returns a tuple with the AllowedEnvIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedEnvIds

`func (o *Integration) SetAllowedEnvIds(v []int32)`

SetAllowedEnvIds sets AllowedEnvIds field to given value.


### GetCreatedAt

`func (o *Integration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Integration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Integration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Integration) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Integration) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Integration) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


