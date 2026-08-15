# Cert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Title** | **string** |  | 
**Custom** | **bool** |  | 
**Issuer** | **string** | Human-readable certificate authority name parsed from uploaded certificates, or the managed issuer identifier. | 
**Domain** | **string** |  | 
**DnsNames** | **[]string** |  | 
**RouteIds** | **[]int32** |  | 
**Fingerprint** | Pointer to **NullableString** |  | [optional] 
**KeyType** | **string** |  | 
**KeyLength** | **int32** |  | 
**Status** | **string** |  | 
**AppInstanceId** | Pointer to **NullableInt32** |  | [optional] 
**AppServiceId** | Pointer to **NullableInt32** |  | [optional] 
**DatabaseId** | Pointer to **NullableInt32** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**IssuedAt** | Pointer to **NullableTime** |  | [optional] 
**RenewsAt** | Pointer to **NullableTime** |  | [optional] 
**ExpiresAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewCert

`func NewCert(id int32, title string, custom bool, issuer string, domain string, dnsNames []string, routeIds []int32, keyType string, keyLength int32, status string, createdAt time.Time, updatedAt time.Time, ) *Cert`

NewCert instantiates a new Cert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertWithDefaults

`func NewCertWithDefaults() *Cert`

NewCertWithDefaults instantiates a new Cert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Cert) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cert) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cert) SetId(v int32)`

SetId sets Id field to given value.


### GetTitle

`func (o *Cert) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Cert) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Cert) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetCustom

`func (o *Cert) GetCustom() bool`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *Cert) GetCustomOk() (*bool, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *Cert) SetCustom(v bool)`

SetCustom sets Custom field to given value.


### GetIssuer

`func (o *Cert) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *Cert) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *Cert) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetDomain

`func (o *Cert) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Cert) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Cert) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetDnsNames

`func (o *Cert) GetDnsNames() []string`

GetDnsNames returns the DnsNames field if non-nil, zero value otherwise.

### GetDnsNamesOk

`func (o *Cert) GetDnsNamesOk() (*[]string, bool)`

GetDnsNamesOk returns a tuple with the DnsNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNames

`func (o *Cert) SetDnsNames(v []string)`

SetDnsNames sets DnsNames field to given value.


### GetRouteIds

`func (o *Cert) GetRouteIds() []int32`

GetRouteIds returns the RouteIds field if non-nil, zero value otherwise.

### GetRouteIdsOk

`func (o *Cert) GetRouteIdsOk() (*[]int32, bool)`

GetRouteIdsOk returns a tuple with the RouteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteIds

`func (o *Cert) SetRouteIds(v []int32)`

SetRouteIds sets RouteIds field to given value.


### GetFingerprint

`func (o *Cert) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *Cert) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *Cert) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *Cert) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### SetFingerprintNil

`func (o *Cert) SetFingerprintNil(b bool)`

 SetFingerprintNil sets the value for Fingerprint to be an explicit nil

### UnsetFingerprint
`func (o *Cert) UnsetFingerprint()`

UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil
### GetKeyType

`func (o *Cert) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *Cert) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *Cert) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.


### GetKeyLength

`func (o *Cert) GetKeyLength() int32`

GetKeyLength returns the KeyLength field if non-nil, zero value otherwise.

### GetKeyLengthOk

`func (o *Cert) GetKeyLengthOk() (*int32, bool)`

GetKeyLengthOk returns a tuple with the KeyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyLength

`func (o *Cert) SetKeyLength(v int32)`

SetKeyLength sets KeyLength field to given value.


### GetStatus

`func (o *Cert) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Cert) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Cert) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAppInstanceId

`func (o *Cert) GetAppInstanceId() int32`

GetAppInstanceId returns the AppInstanceId field if non-nil, zero value otherwise.

### GetAppInstanceIdOk

`func (o *Cert) GetAppInstanceIdOk() (*int32, bool)`

GetAppInstanceIdOk returns a tuple with the AppInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceId

`func (o *Cert) SetAppInstanceId(v int32)`

SetAppInstanceId sets AppInstanceId field to given value.

### HasAppInstanceId

`func (o *Cert) HasAppInstanceId() bool`

HasAppInstanceId returns a boolean if a field has been set.

### SetAppInstanceIdNil

`func (o *Cert) SetAppInstanceIdNil(b bool)`

 SetAppInstanceIdNil sets the value for AppInstanceId to be an explicit nil

### UnsetAppInstanceId
`func (o *Cert) UnsetAppInstanceId()`

UnsetAppInstanceId ensures that no value is present for AppInstanceId, not even an explicit nil
### GetAppServiceId

`func (o *Cert) GetAppServiceId() int32`

GetAppServiceId returns the AppServiceId field if non-nil, zero value otherwise.

### GetAppServiceIdOk

`func (o *Cert) GetAppServiceIdOk() (*int32, bool)`

GetAppServiceIdOk returns a tuple with the AppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceId

`func (o *Cert) SetAppServiceId(v int32)`

SetAppServiceId sets AppServiceId field to given value.

### HasAppServiceId

`func (o *Cert) HasAppServiceId() bool`

HasAppServiceId returns a boolean if a field has been set.

### SetAppServiceIdNil

`func (o *Cert) SetAppServiceIdNil(b bool)`

 SetAppServiceIdNil sets the value for AppServiceId to be an explicit nil

### UnsetAppServiceId
`func (o *Cert) UnsetAppServiceId()`

UnsetAppServiceId ensures that no value is present for AppServiceId, not even an explicit nil
### GetDatabaseId

`func (o *Cert) GetDatabaseId() int32`

GetDatabaseId returns the DatabaseId field if non-nil, zero value otherwise.

### GetDatabaseIdOk

`func (o *Cert) GetDatabaseIdOk() (*int32, bool)`

GetDatabaseIdOk returns a tuple with the DatabaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseId

`func (o *Cert) SetDatabaseId(v int32)`

SetDatabaseId sets DatabaseId field to given value.

### HasDatabaseId

`func (o *Cert) HasDatabaseId() bool`

HasDatabaseId returns a boolean if a field has been set.

### SetDatabaseIdNil

`func (o *Cert) SetDatabaseIdNil(b bool)`

 SetDatabaseIdNil sets the value for DatabaseId to be an explicit nil

### UnsetDatabaseId
`func (o *Cert) UnsetDatabaseId()`

UnsetDatabaseId ensures that no value is present for DatabaseId, not even an explicit nil
### GetCreatedAt

`func (o *Cert) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Cert) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Cert) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Cert) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Cert) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Cert) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIssuedAt

`func (o *Cert) GetIssuedAt() time.Time`

GetIssuedAt returns the IssuedAt field if non-nil, zero value otherwise.

### GetIssuedAtOk

`func (o *Cert) GetIssuedAtOk() (*time.Time, bool)`

GetIssuedAtOk returns a tuple with the IssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAt

`func (o *Cert) SetIssuedAt(v time.Time)`

SetIssuedAt sets IssuedAt field to given value.

### HasIssuedAt

`func (o *Cert) HasIssuedAt() bool`

HasIssuedAt returns a boolean if a field has been set.

### SetIssuedAtNil

`func (o *Cert) SetIssuedAtNil(b bool)`

 SetIssuedAtNil sets the value for IssuedAt to be an explicit nil

### UnsetIssuedAt
`func (o *Cert) UnsetIssuedAt()`

UnsetIssuedAt ensures that no value is present for IssuedAt, not even an explicit nil
### GetRenewsAt

`func (o *Cert) GetRenewsAt() time.Time`

GetRenewsAt returns the RenewsAt field if non-nil, zero value otherwise.

### GetRenewsAtOk

`func (o *Cert) GetRenewsAtOk() (*time.Time, bool)`

GetRenewsAtOk returns a tuple with the RenewsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewsAt

`func (o *Cert) SetRenewsAt(v time.Time)`

SetRenewsAt sets RenewsAt field to given value.

### HasRenewsAt

`func (o *Cert) HasRenewsAt() bool`

HasRenewsAt returns a boolean if a field has been set.

### SetRenewsAtNil

`func (o *Cert) SetRenewsAtNil(b bool)`

 SetRenewsAtNil sets the value for RenewsAt to be an explicit nil

### UnsetRenewsAt
`func (o *Cert) UnsetRenewsAt()`

UnsetRenewsAt ensures that no value is present for RenewsAt, not even an explicit nil
### GetExpiresAt

`func (o *Cert) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Cert) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Cert) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *Cert) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *Cert) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *Cert) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


