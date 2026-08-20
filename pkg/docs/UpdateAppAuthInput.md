# UpdateAppAuthInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppServiceIds** | Pointer to **[]int32** | App services to protect. Omit every scope field to preserve the current scope, or pass an empty list to protect the whole app instance. | [optional] 
**AppServiceId** | Pointer to **NullableInt32** | Single-service scope. Ignored when appServiceIds is supplied. | [optional] 
**AppRouteId** | Pointer to **NullableInt32** | Moves the entry to route scope. The owning app service is derived from the route. | [optional] 
**Login** | **string** |  | 
**Password** | Pointer to **NullableString** | Replaces the existing secret when supplied; omit to keep the current password. | [optional] 
**Realm** | **string** |  | 

## Methods

### NewUpdateAppAuthInput

`func NewUpdateAppAuthInput(login string, realm string, ) *UpdateAppAuthInput`

NewUpdateAppAuthInput instantiates a new UpdateAppAuthInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAppAuthInputWithDefaults

`func NewUpdateAppAuthInputWithDefaults() *UpdateAppAuthInput`

NewUpdateAppAuthInputWithDefaults instantiates a new UpdateAppAuthInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppServiceIds

`func (o *UpdateAppAuthInput) GetAppServiceIds() []int32`

GetAppServiceIds returns the AppServiceIds field if non-nil, zero value otherwise.

### GetAppServiceIdsOk

`func (o *UpdateAppAuthInput) GetAppServiceIdsOk() (*[]int32, bool)`

GetAppServiceIdsOk returns a tuple with the AppServiceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceIds

`func (o *UpdateAppAuthInput) SetAppServiceIds(v []int32)`

SetAppServiceIds sets AppServiceIds field to given value.

### HasAppServiceIds

`func (o *UpdateAppAuthInput) HasAppServiceIds() bool`

HasAppServiceIds returns a boolean if a field has been set.

### SetAppServiceIdsNil

`func (o *UpdateAppAuthInput) SetAppServiceIdsNil(b bool)`

 SetAppServiceIdsNil sets the value for AppServiceIds to be an explicit nil

### UnsetAppServiceIds
`func (o *UpdateAppAuthInput) UnsetAppServiceIds()`

UnsetAppServiceIds ensures that no value is present for AppServiceIds, not even an explicit nil
### GetAppServiceId

`func (o *UpdateAppAuthInput) GetAppServiceId() int32`

GetAppServiceId returns the AppServiceId field if non-nil, zero value otherwise.

### GetAppServiceIdOk

`func (o *UpdateAppAuthInput) GetAppServiceIdOk() (*int32, bool)`

GetAppServiceIdOk returns a tuple with the AppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceId

`func (o *UpdateAppAuthInput) SetAppServiceId(v int32)`

SetAppServiceId sets AppServiceId field to given value.

### HasAppServiceId

`func (o *UpdateAppAuthInput) HasAppServiceId() bool`

HasAppServiceId returns a boolean if a field has been set.

### SetAppServiceIdNil

`func (o *UpdateAppAuthInput) SetAppServiceIdNil(b bool)`

 SetAppServiceIdNil sets the value for AppServiceId to be an explicit nil

### UnsetAppServiceId
`func (o *UpdateAppAuthInput) UnsetAppServiceId()`

UnsetAppServiceId ensures that no value is present for AppServiceId, not even an explicit nil
### GetAppRouteId

`func (o *UpdateAppAuthInput) GetAppRouteId() int32`

GetAppRouteId returns the AppRouteId field if non-nil, zero value otherwise.

### GetAppRouteIdOk

`func (o *UpdateAppAuthInput) GetAppRouteIdOk() (*int32, bool)`

GetAppRouteIdOk returns a tuple with the AppRouteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRouteId

`func (o *UpdateAppAuthInput) SetAppRouteId(v int32)`

SetAppRouteId sets AppRouteId field to given value.

### HasAppRouteId

`func (o *UpdateAppAuthInput) HasAppRouteId() bool`

HasAppRouteId returns a boolean if a field has been set.

### SetAppRouteIdNil

`func (o *UpdateAppAuthInput) SetAppRouteIdNil(b bool)`

 SetAppRouteIdNil sets the value for AppRouteId to be an explicit nil

### UnsetAppRouteId
`func (o *UpdateAppAuthInput) UnsetAppRouteId()`

UnsetAppRouteId ensures that no value is present for AppRouteId, not even an explicit nil
### GetLogin

`func (o *UpdateAppAuthInput) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *UpdateAppAuthInput) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *UpdateAppAuthInput) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetPassword

`func (o *UpdateAppAuthInput) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateAppAuthInput) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateAppAuthInput) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateAppAuthInput) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *UpdateAppAuthInput) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *UpdateAppAuthInput) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetRealm

`func (o *UpdateAppAuthInput) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *UpdateAppAuthInput) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *UpdateAppAuthInput) SetRealm(v string)`

SetRealm sets Realm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


