# AppAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**AppInstanceId** | **int32** |  | 
**Scope** | [**AppAuthScope**](AppAuthScope.md) |  | 
**AppServiceIds** | **[]int32** | App services protected by this entry. Empty unless scope is SERVICE. | 
**AppServiceId** | Pointer to **NullableInt32** | Single protected app service. Null when the entry protects several services or the whole app environment. | [optional] 
**AppRouteId** | Pointer to **NullableInt32** |  | [optional] 
**Login** | **string** | Empty for exemptions. | 
**Realm** | **string** |  | 
**Source** | [**AppAuthSource**](AppAuthSource.md) |  | 
**Generated** | **bool** | True while the password is one Wodby generated. Setting a password by hand clears it. | 
**Exempt** | **bool** | Exemptions switch basic auth off for their services or route even when a wider entry covers them. They have no login or password. | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewAppAuth

`func NewAppAuth(id int32, appInstanceId int32, scope AppAuthScope, appServiceIds []int32, login string, realm string, source AppAuthSource, generated bool, exempt bool, createdAt time.Time, updatedAt time.Time, ) *AppAuth`

NewAppAuth instantiates a new AppAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppAuthWithDefaults

`func NewAppAuthWithDefaults() *AppAuth`

NewAppAuthWithDefaults instantiates a new AppAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppAuth) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppAuth) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppAuth) SetId(v int32)`

SetId sets Id field to given value.


### GetAppInstanceId

`func (o *AppAuth) GetAppInstanceId() int32`

GetAppInstanceId returns the AppInstanceId field if non-nil, zero value otherwise.

### GetAppInstanceIdOk

`func (o *AppAuth) GetAppInstanceIdOk() (*int32, bool)`

GetAppInstanceIdOk returns a tuple with the AppInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceId

`func (o *AppAuth) SetAppInstanceId(v int32)`

SetAppInstanceId sets AppInstanceId field to given value.


### GetScope

`func (o *AppAuth) GetScope() AppAuthScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AppAuth) GetScopeOk() (*AppAuthScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AppAuth) SetScope(v AppAuthScope)`

SetScope sets Scope field to given value.


### GetAppServiceIds

`func (o *AppAuth) GetAppServiceIds() []int32`

GetAppServiceIds returns the AppServiceIds field if non-nil, zero value otherwise.

### GetAppServiceIdsOk

`func (o *AppAuth) GetAppServiceIdsOk() (*[]int32, bool)`

GetAppServiceIdsOk returns a tuple with the AppServiceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceIds

`func (o *AppAuth) SetAppServiceIds(v []int32)`

SetAppServiceIds sets AppServiceIds field to given value.


### GetAppServiceId

`func (o *AppAuth) GetAppServiceId() int32`

GetAppServiceId returns the AppServiceId field if non-nil, zero value otherwise.

### GetAppServiceIdOk

`func (o *AppAuth) GetAppServiceIdOk() (*int32, bool)`

GetAppServiceIdOk returns a tuple with the AppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceId

`func (o *AppAuth) SetAppServiceId(v int32)`

SetAppServiceId sets AppServiceId field to given value.

### HasAppServiceId

`func (o *AppAuth) HasAppServiceId() bool`

HasAppServiceId returns a boolean if a field has been set.

### SetAppServiceIdNil

`func (o *AppAuth) SetAppServiceIdNil(b bool)`

 SetAppServiceIdNil sets the value for AppServiceId to be an explicit nil

### UnsetAppServiceId
`func (o *AppAuth) UnsetAppServiceId()`

UnsetAppServiceId ensures that no value is present for AppServiceId, not even an explicit nil
### GetAppRouteId

`func (o *AppAuth) GetAppRouteId() int32`

GetAppRouteId returns the AppRouteId field if non-nil, zero value otherwise.

### GetAppRouteIdOk

`func (o *AppAuth) GetAppRouteIdOk() (*int32, bool)`

GetAppRouteIdOk returns a tuple with the AppRouteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRouteId

`func (o *AppAuth) SetAppRouteId(v int32)`

SetAppRouteId sets AppRouteId field to given value.

### HasAppRouteId

`func (o *AppAuth) HasAppRouteId() bool`

HasAppRouteId returns a boolean if a field has been set.

### SetAppRouteIdNil

`func (o *AppAuth) SetAppRouteIdNil(b bool)`

 SetAppRouteIdNil sets the value for AppRouteId to be an explicit nil

### UnsetAppRouteId
`func (o *AppAuth) UnsetAppRouteId()`

UnsetAppRouteId ensures that no value is present for AppRouteId, not even an explicit nil
### GetLogin

`func (o *AppAuth) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *AppAuth) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *AppAuth) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetRealm

`func (o *AppAuth) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *AppAuth) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *AppAuth) SetRealm(v string)`

SetRealm sets Realm field to given value.


### GetSource

`func (o *AppAuth) GetSource() AppAuthSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AppAuth) GetSourceOk() (*AppAuthSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AppAuth) SetSource(v AppAuthSource)`

SetSource sets Source field to given value.


### GetGenerated

`func (o *AppAuth) GetGenerated() bool`

GetGenerated returns the Generated field if non-nil, zero value otherwise.

### GetGeneratedOk

`func (o *AppAuth) GetGeneratedOk() (*bool, bool)`

GetGeneratedOk returns a tuple with the Generated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerated

`func (o *AppAuth) SetGenerated(v bool)`

SetGenerated sets Generated field to given value.


### GetExempt

`func (o *AppAuth) GetExempt() bool`

GetExempt returns the Exempt field if non-nil, zero value otherwise.

### GetExemptOk

`func (o *AppAuth) GetExemptOk() (*bool, bool)`

GetExemptOk returns a tuple with the Exempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExempt

`func (o *AppAuth) SetExempt(v bool)`

SetExempt sets Exempt field to given value.


### GetCreatedAt

`func (o *AppAuth) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppAuth) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppAuth) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AppAuth) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AppAuth) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AppAuth) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


