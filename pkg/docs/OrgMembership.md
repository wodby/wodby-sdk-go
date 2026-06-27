# OrgMembership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**UserId** | Pointer to **NullableInt32** |  | [optional] 
**User** | Pointer to [**NullableUser**](User.md) |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**OrgId** | **int32** |  | 
**Role** | **string** |  | 
**Status** | **string** |  | 
**JoinedAt** | Pointer to **NullableTime** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewOrgMembership

`func NewOrgMembership(id int32, orgId int32, role string, status string, createdAt time.Time, updatedAt time.Time, ) *OrgMembership`

NewOrgMembership instantiates a new OrgMembership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgMembershipWithDefaults

`func NewOrgMembershipWithDefaults() *OrgMembership`

NewOrgMembershipWithDefaults instantiates a new OrgMembership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrgMembership) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgMembership) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgMembership) SetId(v int32)`

SetId sets Id field to given value.


### GetUserId

`func (o *OrgMembership) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *OrgMembership) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *OrgMembership) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *OrgMembership) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *OrgMembership) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *OrgMembership) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetUser

`func (o *OrgMembership) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *OrgMembership) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *OrgMembership) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *OrgMembership) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *OrgMembership) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *OrgMembership) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetEmail

`func (o *OrgMembership) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrgMembership) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrgMembership) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrgMembership) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *OrgMembership) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *OrgMembership) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetOrgId

`func (o *OrgMembership) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *OrgMembership) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *OrgMembership) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.


### GetRole

`func (o *OrgMembership) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrgMembership) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrgMembership) SetRole(v string)`

SetRole sets Role field to given value.


### GetStatus

`func (o *OrgMembership) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrgMembership) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrgMembership) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetJoinedAt

`func (o *OrgMembership) GetJoinedAt() time.Time`

GetJoinedAt returns the JoinedAt field if non-nil, zero value otherwise.

### GetJoinedAtOk

`func (o *OrgMembership) GetJoinedAtOk() (*time.Time, bool)`

GetJoinedAtOk returns a tuple with the JoinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedAt

`func (o *OrgMembership) SetJoinedAt(v time.Time)`

SetJoinedAt sets JoinedAt field to given value.

### HasJoinedAt

`func (o *OrgMembership) HasJoinedAt() bool`

HasJoinedAt returns a boolean if a field has been set.

### SetJoinedAtNil

`func (o *OrgMembership) SetJoinedAtNil(b bool)`

 SetJoinedAtNil sets the value for JoinedAt to be an explicit nil

### UnsetJoinedAt
`func (o *OrgMembership) UnsetJoinedAt()`

UnsetJoinedAt ensures that no value is present for JoinedAt, not even an explicit nil
### GetCreatedAt

`func (o *OrgMembership) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrgMembership) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrgMembership) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *OrgMembership) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrgMembership) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrgMembership) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


