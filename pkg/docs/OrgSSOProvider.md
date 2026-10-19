# OrgSSOProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**OrgId** | **int32** |  | 
**Name** | **string** |  | 
**Kind** | **string** | Provider kind, such as oidc, github or saml. | 
**IssuerUrl** | **string** |  | 
**ClientId** | **string** |  | 
**GithubOrg** | Pointer to **NullableString** |  | [optional] 
**SamlIdpMetadataUrl** | Pointer to **NullableString** |  | [optional] 
**SamlSpEntityId** | Pointer to **NullableString** | Entity ID to configure in the SAML identity provider. | [optional] 
**SamlSpMetadataUrl** | Pointer to **NullableString** |  | [optional] 
**SamlAcsUrl** | Pointer to **NullableString** | Assertion consumer service URL to configure in the SAML identity provider. | [optional] 
**Enabled** | **bool** |  | 
**JitEnabled** | **bool** | Create a membership on first sign-in for people from a verified domain. | 
**Domains** | [**[]OrgSSODomain**](OrgSSODomain.md) |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewOrgSSOProvider

`func NewOrgSSOProvider(id int32, orgId int32, name string, kind string, issuerUrl string, clientId string, enabled bool, jitEnabled bool, domains []OrgSSODomain, createdAt time.Time, updatedAt time.Time, ) *OrgSSOProvider`

NewOrgSSOProvider instantiates a new OrgSSOProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgSSOProviderWithDefaults

`func NewOrgSSOProviderWithDefaults() *OrgSSOProvider`

NewOrgSSOProviderWithDefaults instantiates a new OrgSSOProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrgSSOProvider) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgSSOProvider) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgSSOProvider) SetId(v int32)`

SetId sets Id field to given value.


### GetOrgId

`func (o *OrgSSOProvider) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *OrgSSOProvider) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *OrgSSOProvider) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.


### GetName

`func (o *OrgSSOProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrgSSOProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrgSSOProvider) SetName(v string)`

SetName sets Name field to given value.


### GetKind

`func (o *OrgSSOProvider) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *OrgSSOProvider) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *OrgSSOProvider) SetKind(v string)`

SetKind sets Kind field to given value.


### GetIssuerUrl

`func (o *OrgSSOProvider) GetIssuerUrl() string`

GetIssuerUrl returns the IssuerUrl field if non-nil, zero value otherwise.

### GetIssuerUrlOk

`func (o *OrgSSOProvider) GetIssuerUrlOk() (*string, bool)`

GetIssuerUrlOk returns a tuple with the IssuerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUrl

`func (o *OrgSSOProvider) SetIssuerUrl(v string)`

SetIssuerUrl sets IssuerUrl field to given value.


### GetClientId

`func (o *OrgSSOProvider) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OrgSSOProvider) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OrgSSOProvider) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetGithubOrg

`func (o *OrgSSOProvider) GetGithubOrg() string`

GetGithubOrg returns the GithubOrg field if non-nil, zero value otherwise.

### GetGithubOrgOk

`func (o *OrgSSOProvider) GetGithubOrgOk() (*string, bool)`

GetGithubOrgOk returns a tuple with the GithubOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubOrg

`func (o *OrgSSOProvider) SetGithubOrg(v string)`

SetGithubOrg sets GithubOrg field to given value.

### HasGithubOrg

`func (o *OrgSSOProvider) HasGithubOrg() bool`

HasGithubOrg returns a boolean if a field has been set.

### SetGithubOrgNil

`func (o *OrgSSOProvider) SetGithubOrgNil(b bool)`

 SetGithubOrgNil sets the value for GithubOrg to be an explicit nil

### UnsetGithubOrg
`func (o *OrgSSOProvider) UnsetGithubOrg()`

UnsetGithubOrg ensures that no value is present for GithubOrg, not even an explicit nil
### GetSamlIdpMetadataUrl

`func (o *OrgSSOProvider) GetSamlIdpMetadataUrl() string`

GetSamlIdpMetadataUrl returns the SamlIdpMetadataUrl field if non-nil, zero value otherwise.

### GetSamlIdpMetadataUrlOk

`func (o *OrgSSOProvider) GetSamlIdpMetadataUrlOk() (*string, bool)`

GetSamlIdpMetadataUrlOk returns a tuple with the SamlIdpMetadataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlIdpMetadataUrl

`func (o *OrgSSOProvider) SetSamlIdpMetadataUrl(v string)`

SetSamlIdpMetadataUrl sets SamlIdpMetadataUrl field to given value.

### HasSamlIdpMetadataUrl

`func (o *OrgSSOProvider) HasSamlIdpMetadataUrl() bool`

HasSamlIdpMetadataUrl returns a boolean if a field has been set.

### SetSamlIdpMetadataUrlNil

`func (o *OrgSSOProvider) SetSamlIdpMetadataUrlNil(b bool)`

 SetSamlIdpMetadataUrlNil sets the value for SamlIdpMetadataUrl to be an explicit nil

### UnsetSamlIdpMetadataUrl
`func (o *OrgSSOProvider) UnsetSamlIdpMetadataUrl()`

UnsetSamlIdpMetadataUrl ensures that no value is present for SamlIdpMetadataUrl, not even an explicit nil
### GetSamlSpEntityId

`func (o *OrgSSOProvider) GetSamlSpEntityId() string`

GetSamlSpEntityId returns the SamlSpEntityId field if non-nil, zero value otherwise.

### GetSamlSpEntityIdOk

`func (o *OrgSSOProvider) GetSamlSpEntityIdOk() (*string, bool)`

GetSamlSpEntityIdOk returns a tuple with the SamlSpEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlSpEntityId

`func (o *OrgSSOProvider) SetSamlSpEntityId(v string)`

SetSamlSpEntityId sets SamlSpEntityId field to given value.

### HasSamlSpEntityId

`func (o *OrgSSOProvider) HasSamlSpEntityId() bool`

HasSamlSpEntityId returns a boolean if a field has been set.

### SetSamlSpEntityIdNil

`func (o *OrgSSOProvider) SetSamlSpEntityIdNil(b bool)`

 SetSamlSpEntityIdNil sets the value for SamlSpEntityId to be an explicit nil

### UnsetSamlSpEntityId
`func (o *OrgSSOProvider) UnsetSamlSpEntityId()`

UnsetSamlSpEntityId ensures that no value is present for SamlSpEntityId, not even an explicit nil
### GetSamlSpMetadataUrl

`func (o *OrgSSOProvider) GetSamlSpMetadataUrl() string`

GetSamlSpMetadataUrl returns the SamlSpMetadataUrl field if non-nil, zero value otherwise.

### GetSamlSpMetadataUrlOk

`func (o *OrgSSOProvider) GetSamlSpMetadataUrlOk() (*string, bool)`

GetSamlSpMetadataUrlOk returns a tuple with the SamlSpMetadataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlSpMetadataUrl

`func (o *OrgSSOProvider) SetSamlSpMetadataUrl(v string)`

SetSamlSpMetadataUrl sets SamlSpMetadataUrl field to given value.

### HasSamlSpMetadataUrl

`func (o *OrgSSOProvider) HasSamlSpMetadataUrl() bool`

HasSamlSpMetadataUrl returns a boolean if a field has been set.

### SetSamlSpMetadataUrlNil

`func (o *OrgSSOProvider) SetSamlSpMetadataUrlNil(b bool)`

 SetSamlSpMetadataUrlNil sets the value for SamlSpMetadataUrl to be an explicit nil

### UnsetSamlSpMetadataUrl
`func (o *OrgSSOProvider) UnsetSamlSpMetadataUrl()`

UnsetSamlSpMetadataUrl ensures that no value is present for SamlSpMetadataUrl, not even an explicit nil
### GetSamlAcsUrl

`func (o *OrgSSOProvider) GetSamlAcsUrl() string`

GetSamlAcsUrl returns the SamlAcsUrl field if non-nil, zero value otherwise.

### GetSamlAcsUrlOk

`func (o *OrgSSOProvider) GetSamlAcsUrlOk() (*string, bool)`

GetSamlAcsUrlOk returns a tuple with the SamlAcsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAcsUrl

`func (o *OrgSSOProvider) SetSamlAcsUrl(v string)`

SetSamlAcsUrl sets SamlAcsUrl field to given value.

### HasSamlAcsUrl

`func (o *OrgSSOProvider) HasSamlAcsUrl() bool`

HasSamlAcsUrl returns a boolean if a field has been set.

### SetSamlAcsUrlNil

`func (o *OrgSSOProvider) SetSamlAcsUrlNil(b bool)`

 SetSamlAcsUrlNil sets the value for SamlAcsUrl to be an explicit nil

### UnsetSamlAcsUrl
`func (o *OrgSSOProvider) UnsetSamlAcsUrl()`

UnsetSamlAcsUrl ensures that no value is present for SamlAcsUrl, not even an explicit nil
### GetEnabled

`func (o *OrgSSOProvider) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OrgSSOProvider) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OrgSSOProvider) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetJitEnabled

`func (o *OrgSSOProvider) GetJitEnabled() bool`

GetJitEnabled returns the JitEnabled field if non-nil, zero value otherwise.

### GetJitEnabledOk

`func (o *OrgSSOProvider) GetJitEnabledOk() (*bool, bool)`

GetJitEnabledOk returns a tuple with the JitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitEnabled

`func (o *OrgSSOProvider) SetJitEnabled(v bool)`

SetJitEnabled sets JitEnabled field to given value.


### GetDomains

`func (o *OrgSSOProvider) GetDomains() []OrgSSODomain`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *OrgSSOProvider) GetDomainsOk() (*[]OrgSSODomain, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *OrgSSOProvider) SetDomains(v []OrgSSODomain)`

SetDomains sets Domains field to given value.


### GetCreatedAt

`func (o *OrgSSOProvider) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrgSSOProvider) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrgSSOProvider) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *OrgSSOProvider) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrgSSOProvider) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrgSSOProvider) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


