# Org

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Title** | **string** |  | 
**Domain** | **string** |  | 
**DefaultTimeZone** | **string** |  | 
**CiIntegrationId** | **int32** | Effective default CI integration ID. Zero selects the built-in Wodby CI service. | 
**RegistryIntegrationId** | **int32** | Effective default registry integration ID. Zero selects the built-in Wodby registry service. | 
**Capabilities** | Pointer to [**OrgCapabilities**](OrgCapabilities.md) |  | [optional] 
**Subscription** | Pointer to [**OrgSubscription**](OrgSubscription.md) |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewOrg

`func NewOrg(id int32, name string, title string, domain string, defaultTimeZone string, ciIntegrationId int32, registryIntegrationId int32, createdAt time.Time, updatedAt time.Time, ) *Org`

NewOrg instantiates a new Org object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgWithDefaults

`func NewOrgWithDefaults() *Org`

NewOrgWithDefaults instantiates a new Org object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Org) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Org) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Org) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Org) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Org) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Org) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *Org) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Org) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Org) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDomain

`func (o *Org) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Org) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Org) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetDefaultTimeZone

`func (o *Org) GetDefaultTimeZone() string`

GetDefaultTimeZone returns the DefaultTimeZone field if non-nil, zero value otherwise.

### GetDefaultTimeZoneOk

`func (o *Org) GetDefaultTimeZoneOk() (*string, bool)`

GetDefaultTimeZoneOk returns a tuple with the DefaultTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTimeZone

`func (o *Org) SetDefaultTimeZone(v string)`

SetDefaultTimeZone sets DefaultTimeZone field to given value.


### GetCiIntegrationId

`func (o *Org) GetCiIntegrationId() int32`

GetCiIntegrationId returns the CiIntegrationId field if non-nil, zero value otherwise.

### GetCiIntegrationIdOk

`func (o *Org) GetCiIntegrationIdOk() (*int32, bool)`

GetCiIntegrationIdOk returns a tuple with the CiIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiIntegrationId

`func (o *Org) SetCiIntegrationId(v int32)`

SetCiIntegrationId sets CiIntegrationId field to given value.


### GetRegistryIntegrationId

`func (o *Org) GetRegistryIntegrationId() int32`

GetRegistryIntegrationId returns the RegistryIntegrationId field if non-nil, zero value otherwise.

### GetRegistryIntegrationIdOk

`func (o *Org) GetRegistryIntegrationIdOk() (*int32, bool)`

GetRegistryIntegrationIdOk returns a tuple with the RegistryIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryIntegrationId

`func (o *Org) SetRegistryIntegrationId(v int32)`

SetRegistryIntegrationId sets RegistryIntegrationId field to given value.


### GetCapabilities

`func (o *Org) GetCapabilities() OrgCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *Org) GetCapabilitiesOk() (*OrgCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *Org) SetCapabilities(v OrgCapabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *Org) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetSubscription

`func (o *Org) GetSubscription() OrgSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *Org) GetSubscriptionOk() (*OrgSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *Org) SetSubscription(v OrgSubscription)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *Org) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Org) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Org) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Org) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Org) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Org) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Org) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


