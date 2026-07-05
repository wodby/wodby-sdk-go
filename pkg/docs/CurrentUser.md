# CurrentUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Email** | **string** |  | 
**Name** | **string** |  | 
**Orgs** | Pointer to [**[]Org**](Org.md) |  | [optional] 
**Twofa** | **bool** |  | 
**DefaultOrg** | Pointer to [**NullableOrg**](Org.md) |  | [optional] 
**DefaultProjects** | Pointer to [**[]Project**](Project.md) |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewCurrentUser

`func NewCurrentUser(id int32, email string, name string, twofa bool, createdAt time.Time, updatedAt time.Time, ) *CurrentUser`

NewCurrentUser instantiates a new CurrentUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentUserWithDefaults

`func NewCurrentUserWithDefaults() *CurrentUser`

NewCurrentUserWithDefaults instantiates a new CurrentUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CurrentUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CurrentUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CurrentUser) SetId(v int32)`

SetId sets Id field to given value.


### GetEmail

`func (o *CurrentUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CurrentUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CurrentUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *CurrentUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CurrentUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CurrentUser) SetName(v string)`

SetName sets Name field to given value.


### GetOrgs

`func (o *CurrentUser) GetOrgs() []Org`

GetOrgs returns the Orgs field if non-nil, zero value otherwise.

### GetOrgsOk

`func (o *CurrentUser) GetOrgsOk() (*[]Org, bool)`

GetOrgsOk returns a tuple with the Orgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgs

`func (o *CurrentUser) SetOrgs(v []Org)`

SetOrgs sets Orgs field to given value.

### HasOrgs

`func (o *CurrentUser) HasOrgs() bool`

HasOrgs returns a boolean if a field has been set.

### GetTwofa

`func (o *CurrentUser) GetTwofa() bool`

GetTwofa returns the Twofa field if non-nil, zero value otherwise.

### GetTwofaOk

`func (o *CurrentUser) GetTwofaOk() (*bool, bool)`

GetTwofaOk returns a tuple with the Twofa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwofa

`func (o *CurrentUser) SetTwofa(v bool)`

SetTwofa sets Twofa field to given value.


### GetDefaultOrg

`func (o *CurrentUser) GetDefaultOrg() Org`

GetDefaultOrg returns the DefaultOrg field if non-nil, zero value otherwise.

### GetDefaultOrgOk

`func (o *CurrentUser) GetDefaultOrgOk() (*Org, bool)`

GetDefaultOrgOk returns a tuple with the DefaultOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOrg

`func (o *CurrentUser) SetDefaultOrg(v Org)`

SetDefaultOrg sets DefaultOrg field to given value.

### HasDefaultOrg

`func (o *CurrentUser) HasDefaultOrg() bool`

HasDefaultOrg returns a boolean if a field has been set.

### SetDefaultOrgNil

`func (o *CurrentUser) SetDefaultOrgNil(b bool)`

 SetDefaultOrgNil sets the value for DefaultOrg to be an explicit nil

### UnsetDefaultOrg
`func (o *CurrentUser) UnsetDefaultOrg()`

UnsetDefaultOrg ensures that no value is present for DefaultOrg, not even an explicit nil
### GetDefaultProjects

`func (o *CurrentUser) GetDefaultProjects() []Project`

GetDefaultProjects returns the DefaultProjects field if non-nil, zero value otherwise.

### GetDefaultProjectsOk

`func (o *CurrentUser) GetDefaultProjectsOk() (*[]Project, bool)`

GetDefaultProjectsOk returns a tuple with the DefaultProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultProjects

`func (o *CurrentUser) SetDefaultProjects(v []Project)`

SetDefaultProjects sets DefaultProjects field to given value.

### HasDefaultProjects

`func (o *CurrentUser) HasDefaultProjects() bool`

HasDefaultProjects returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CurrentUser) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CurrentUser) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CurrentUser) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CurrentUser) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CurrentUser) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CurrentUser) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


