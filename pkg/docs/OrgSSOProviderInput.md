# OrgSSOProviderInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **int32** | Used when creating a provider. Optional for API-key requests; defaults to the API key&#39;s organization. | [optional] 
**Kind** | Pointer to **NullableString** | Provider kind when creating, such as oidc, github or saml. It cannot be changed. | [optional] 
**Name** | **string** |  | 
**IssuerUrl** | Pointer to **NullableString** |  | [optional] 
**ClientId** | Pointer to **NullableString** |  | [optional] 
**ClientSecret** | Pointer to **NullableString** | OIDC or GitHub client secret. Leave it out on update to keep the stored secret. | [optional] 
**GithubOrg** | Pointer to **NullableString** |  | [optional] 
**SamlIdpMetadataUrl** | Pointer to **NullableString** |  | [optional] 
**SamlIdpMetadataXml** | Pointer to **NullableString** |  | [optional] 
**Enabled** | **bool** |  | 
**JitEnabled** | **bool** |  | 
**Domains** | **[]string** |  | 

## Methods

### NewOrgSSOProviderInput

`func NewOrgSSOProviderInput(name string, enabled bool, jitEnabled bool, domains []string, ) *OrgSSOProviderInput`

NewOrgSSOProviderInput instantiates a new OrgSSOProviderInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgSSOProviderInputWithDefaults

`func NewOrgSSOProviderInputWithDefaults() *OrgSSOProviderInput`

NewOrgSSOProviderInputWithDefaults instantiates a new OrgSSOProviderInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *OrgSSOProviderInput) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *OrgSSOProviderInput) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *OrgSSOProviderInput) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *OrgSSOProviderInput) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetKind

`func (o *OrgSSOProviderInput) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *OrgSSOProviderInput) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *OrgSSOProviderInput) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *OrgSSOProviderInput) HasKind() bool`

HasKind returns a boolean if a field has been set.

### SetKindNil

`func (o *OrgSSOProviderInput) SetKindNil(b bool)`

 SetKindNil sets the value for Kind to be an explicit nil

### UnsetKind
`func (o *OrgSSOProviderInput) UnsetKind()`

UnsetKind ensures that no value is present for Kind, not even an explicit nil
### GetName

`func (o *OrgSSOProviderInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrgSSOProviderInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrgSSOProviderInput) SetName(v string)`

SetName sets Name field to given value.


### GetIssuerUrl

`func (o *OrgSSOProviderInput) GetIssuerUrl() string`

GetIssuerUrl returns the IssuerUrl field if non-nil, zero value otherwise.

### GetIssuerUrlOk

`func (o *OrgSSOProviderInput) GetIssuerUrlOk() (*string, bool)`

GetIssuerUrlOk returns a tuple with the IssuerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUrl

`func (o *OrgSSOProviderInput) SetIssuerUrl(v string)`

SetIssuerUrl sets IssuerUrl field to given value.

### HasIssuerUrl

`func (o *OrgSSOProviderInput) HasIssuerUrl() bool`

HasIssuerUrl returns a boolean if a field has been set.

### SetIssuerUrlNil

`func (o *OrgSSOProviderInput) SetIssuerUrlNil(b bool)`

 SetIssuerUrlNil sets the value for IssuerUrl to be an explicit nil

### UnsetIssuerUrl
`func (o *OrgSSOProviderInput) UnsetIssuerUrl()`

UnsetIssuerUrl ensures that no value is present for IssuerUrl, not even an explicit nil
### GetClientId

`func (o *OrgSSOProviderInput) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OrgSSOProviderInput) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OrgSSOProviderInput) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *OrgSSOProviderInput) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *OrgSSOProviderInput) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *OrgSSOProviderInput) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetClientSecret

`func (o *OrgSSOProviderInput) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OrgSSOProviderInput) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OrgSSOProviderInput) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *OrgSSOProviderInput) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### SetClientSecretNil

`func (o *OrgSSOProviderInput) SetClientSecretNil(b bool)`

 SetClientSecretNil sets the value for ClientSecret to be an explicit nil

### UnsetClientSecret
`func (o *OrgSSOProviderInput) UnsetClientSecret()`

UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
### GetGithubOrg

`func (o *OrgSSOProviderInput) GetGithubOrg() string`

GetGithubOrg returns the GithubOrg field if non-nil, zero value otherwise.

### GetGithubOrgOk

`func (o *OrgSSOProviderInput) GetGithubOrgOk() (*string, bool)`

GetGithubOrgOk returns a tuple with the GithubOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubOrg

`func (o *OrgSSOProviderInput) SetGithubOrg(v string)`

SetGithubOrg sets GithubOrg field to given value.

### HasGithubOrg

`func (o *OrgSSOProviderInput) HasGithubOrg() bool`

HasGithubOrg returns a boolean if a field has been set.

### SetGithubOrgNil

`func (o *OrgSSOProviderInput) SetGithubOrgNil(b bool)`

 SetGithubOrgNil sets the value for GithubOrg to be an explicit nil

### UnsetGithubOrg
`func (o *OrgSSOProviderInput) UnsetGithubOrg()`

UnsetGithubOrg ensures that no value is present for GithubOrg, not even an explicit nil
### GetSamlIdpMetadataUrl

`func (o *OrgSSOProviderInput) GetSamlIdpMetadataUrl() string`

GetSamlIdpMetadataUrl returns the SamlIdpMetadataUrl field if non-nil, zero value otherwise.

### GetSamlIdpMetadataUrlOk

`func (o *OrgSSOProviderInput) GetSamlIdpMetadataUrlOk() (*string, bool)`

GetSamlIdpMetadataUrlOk returns a tuple with the SamlIdpMetadataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlIdpMetadataUrl

`func (o *OrgSSOProviderInput) SetSamlIdpMetadataUrl(v string)`

SetSamlIdpMetadataUrl sets SamlIdpMetadataUrl field to given value.

### HasSamlIdpMetadataUrl

`func (o *OrgSSOProviderInput) HasSamlIdpMetadataUrl() bool`

HasSamlIdpMetadataUrl returns a boolean if a field has been set.

### SetSamlIdpMetadataUrlNil

`func (o *OrgSSOProviderInput) SetSamlIdpMetadataUrlNil(b bool)`

 SetSamlIdpMetadataUrlNil sets the value for SamlIdpMetadataUrl to be an explicit nil

### UnsetSamlIdpMetadataUrl
`func (o *OrgSSOProviderInput) UnsetSamlIdpMetadataUrl()`

UnsetSamlIdpMetadataUrl ensures that no value is present for SamlIdpMetadataUrl, not even an explicit nil
### GetSamlIdpMetadataXml

`func (o *OrgSSOProviderInput) GetSamlIdpMetadataXml() string`

GetSamlIdpMetadataXml returns the SamlIdpMetadataXml field if non-nil, zero value otherwise.

### GetSamlIdpMetadataXmlOk

`func (o *OrgSSOProviderInput) GetSamlIdpMetadataXmlOk() (*string, bool)`

GetSamlIdpMetadataXmlOk returns a tuple with the SamlIdpMetadataXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlIdpMetadataXml

`func (o *OrgSSOProviderInput) SetSamlIdpMetadataXml(v string)`

SetSamlIdpMetadataXml sets SamlIdpMetadataXml field to given value.

### HasSamlIdpMetadataXml

`func (o *OrgSSOProviderInput) HasSamlIdpMetadataXml() bool`

HasSamlIdpMetadataXml returns a boolean if a field has been set.

### SetSamlIdpMetadataXmlNil

`func (o *OrgSSOProviderInput) SetSamlIdpMetadataXmlNil(b bool)`

 SetSamlIdpMetadataXmlNil sets the value for SamlIdpMetadataXml to be an explicit nil

### UnsetSamlIdpMetadataXml
`func (o *OrgSSOProviderInput) UnsetSamlIdpMetadataXml()`

UnsetSamlIdpMetadataXml ensures that no value is present for SamlIdpMetadataXml, not even an explicit nil
### GetEnabled

`func (o *OrgSSOProviderInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OrgSSOProviderInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OrgSSOProviderInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetJitEnabled

`func (o *OrgSSOProviderInput) GetJitEnabled() bool`

GetJitEnabled returns the JitEnabled field if non-nil, zero value otherwise.

### GetJitEnabledOk

`func (o *OrgSSOProviderInput) GetJitEnabledOk() (*bool, bool)`

GetJitEnabledOk returns a tuple with the JitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitEnabled

`func (o *OrgSSOProviderInput) SetJitEnabled(v bool)`

SetJitEnabled sets JitEnabled field to given value.


### GetDomains

`func (o *OrgSSOProviderInput) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *OrgSSOProviderInput) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *OrgSSOProviderInput) SetDomains(v []string)`

SetDomains sets Domains field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


