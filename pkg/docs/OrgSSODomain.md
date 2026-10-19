# OrgSSODomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Domain** | **string** |  | 
**Verified** | **bool** |  | 
**VerificationRecordName** | **string** | DNS record name to create for verification. | 
**VerificationRecordType** | **string** |  | 
**VerificationRecordValue** | **string** |  | 
**VerifiedAt** | Pointer to **NullableTime** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewOrgSSODomain

`func NewOrgSSODomain(id int32, domain string, verified bool, verificationRecordName string, verificationRecordType string, verificationRecordValue string, createdAt time.Time, updatedAt time.Time, ) *OrgSSODomain`

NewOrgSSODomain instantiates a new OrgSSODomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgSSODomainWithDefaults

`func NewOrgSSODomainWithDefaults() *OrgSSODomain`

NewOrgSSODomainWithDefaults instantiates a new OrgSSODomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrgSSODomain) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgSSODomain) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgSSODomain) SetId(v int32)`

SetId sets Id field to given value.


### GetDomain

`func (o *OrgSSODomain) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *OrgSSODomain) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *OrgSSODomain) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetVerified

`func (o *OrgSSODomain) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *OrgSSODomain) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *OrgSSODomain) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### GetVerificationRecordName

`func (o *OrgSSODomain) GetVerificationRecordName() string`

GetVerificationRecordName returns the VerificationRecordName field if non-nil, zero value otherwise.

### GetVerificationRecordNameOk

`func (o *OrgSSODomain) GetVerificationRecordNameOk() (*string, bool)`

GetVerificationRecordNameOk returns a tuple with the VerificationRecordName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationRecordName

`func (o *OrgSSODomain) SetVerificationRecordName(v string)`

SetVerificationRecordName sets VerificationRecordName field to given value.


### GetVerificationRecordType

`func (o *OrgSSODomain) GetVerificationRecordType() string`

GetVerificationRecordType returns the VerificationRecordType field if non-nil, zero value otherwise.

### GetVerificationRecordTypeOk

`func (o *OrgSSODomain) GetVerificationRecordTypeOk() (*string, bool)`

GetVerificationRecordTypeOk returns a tuple with the VerificationRecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationRecordType

`func (o *OrgSSODomain) SetVerificationRecordType(v string)`

SetVerificationRecordType sets VerificationRecordType field to given value.


### GetVerificationRecordValue

`func (o *OrgSSODomain) GetVerificationRecordValue() string`

GetVerificationRecordValue returns the VerificationRecordValue field if non-nil, zero value otherwise.

### GetVerificationRecordValueOk

`func (o *OrgSSODomain) GetVerificationRecordValueOk() (*string, bool)`

GetVerificationRecordValueOk returns a tuple with the VerificationRecordValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationRecordValue

`func (o *OrgSSODomain) SetVerificationRecordValue(v string)`

SetVerificationRecordValue sets VerificationRecordValue field to given value.


### GetVerifiedAt

`func (o *OrgSSODomain) GetVerifiedAt() time.Time`

GetVerifiedAt returns the VerifiedAt field if non-nil, zero value otherwise.

### GetVerifiedAtOk

`func (o *OrgSSODomain) GetVerifiedAtOk() (*time.Time, bool)`

GetVerifiedAtOk returns a tuple with the VerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAt

`func (o *OrgSSODomain) SetVerifiedAt(v time.Time)`

SetVerifiedAt sets VerifiedAt field to given value.

### HasVerifiedAt

`func (o *OrgSSODomain) HasVerifiedAt() bool`

HasVerifiedAt returns a boolean if a field has been set.

### SetVerifiedAtNil

`func (o *OrgSSODomain) SetVerifiedAtNil(b bool)`

 SetVerifiedAtNil sets the value for VerifiedAt to be an explicit nil

### UnsetVerifiedAt
`func (o *OrgSSODomain) UnsetVerifiedAt()`

UnsetVerifiedAt ensures that no value is present for VerifiedAt, not even an explicit nil
### GetCreatedAt

`func (o *OrgSSODomain) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrgSSODomain) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrgSSODomain) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *OrgSSODomain) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrgSSODomain) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrgSSODomain) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


