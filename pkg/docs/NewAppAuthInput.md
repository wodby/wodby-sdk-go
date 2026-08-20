# NewAppAuthInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppInstanceId** | **int32** |  | 
**AppServiceIds** | Pointer to **[]int32** | App services to protect. Omit or pass an empty list to protect the whole app instance. | [optional] 
**AppServiceId** | Pointer to **NullableInt32** | Single-service scope. Ignored when appServiceIds is supplied. | [optional] 
**AppRouteId** | Pointer to **NullableInt32** | Route scope. The owning app service is derived from the route. | [optional] 
**Login** | **string** |  | 
**Password** | **string** |  | 
**Realm** | **string** |  | 

## Methods

### NewNewAppAuthInput

`func NewNewAppAuthInput(appInstanceId int32, login string, password string, realm string, ) *NewAppAuthInput`

NewNewAppAuthInput instantiates a new NewAppAuthInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewAppAuthInputWithDefaults

`func NewNewAppAuthInputWithDefaults() *NewAppAuthInput`

NewNewAppAuthInputWithDefaults instantiates a new NewAppAuthInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppInstanceId

`func (o *NewAppAuthInput) GetAppInstanceId() int32`

GetAppInstanceId returns the AppInstanceId field if non-nil, zero value otherwise.

### GetAppInstanceIdOk

`func (o *NewAppAuthInput) GetAppInstanceIdOk() (*int32, bool)`

GetAppInstanceIdOk returns a tuple with the AppInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceId

`func (o *NewAppAuthInput) SetAppInstanceId(v int32)`

SetAppInstanceId sets AppInstanceId field to given value.


### GetAppServiceIds

`func (o *NewAppAuthInput) GetAppServiceIds() []int32`

GetAppServiceIds returns the AppServiceIds field if non-nil, zero value otherwise.

### GetAppServiceIdsOk

`func (o *NewAppAuthInput) GetAppServiceIdsOk() (*[]int32, bool)`

GetAppServiceIdsOk returns a tuple with the AppServiceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceIds

`func (o *NewAppAuthInput) SetAppServiceIds(v []int32)`

SetAppServiceIds sets AppServiceIds field to given value.

### HasAppServiceIds

`func (o *NewAppAuthInput) HasAppServiceIds() bool`

HasAppServiceIds returns a boolean if a field has been set.

### SetAppServiceIdsNil

`func (o *NewAppAuthInput) SetAppServiceIdsNil(b bool)`

 SetAppServiceIdsNil sets the value for AppServiceIds to be an explicit nil

### UnsetAppServiceIds
`func (o *NewAppAuthInput) UnsetAppServiceIds()`

UnsetAppServiceIds ensures that no value is present for AppServiceIds, not even an explicit nil
### GetAppServiceId

`func (o *NewAppAuthInput) GetAppServiceId() int32`

GetAppServiceId returns the AppServiceId field if non-nil, zero value otherwise.

### GetAppServiceIdOk

`func (o *NewAppAuthInput) GetAppServiceIdOk() (*int32, bool)`

GetAppServiceIdOk returns a tuple with the AppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceId

`func (o *NewAppAuthInput) SetAppServiceId(v int32)`

SetAppServiceId sets AppServiceId field to given value.

### HasAppServiceId

`func (o *NewAppAuthInput) HasAppServiceId() bool`

HasAppServiceId returns a boolean if a field has been set.

### SetAppServiceIdNil

`func (o *NewAppAuthInput) SetAppServiceIdNil(b bool)`

 SetAppServiceIdNil sets the value for AppServiceId to be an explicit nil

### UnsetAppServiceId
`func (o *NewAppAuthInput) UnsetAppServiceId()`

UnsetAppServiceId ensures that no value is present for AppServiceId, not even an explicit nil
### GetAppRouteId

`func (o *NewAppAuthInput) GetAppRouteId() int32`

GetAppRouteId returns the AppRouteId field if non-nil, zero value otherwise.

### GetAppRouteIdOk

`func (o *NewAppAuthInput) GetAppRouteIdOk() (*int32, bool)`

GetAppRouteIdOk returns a tuple with the AppRouteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRouteId

`func (o *NewAppAuthInput) SetAppRouteId(v int32)`

SetAppRouteId sets AppRouteId field to given value.

### HasAppRouteId

`func (o *NewAppAuthInput) HasAppRouteId() bool`

HasAppRouteId returns a boolean if a field has been set.

### SetAppRouteIdNil

`func (o *NewAppAuthInput) SetAppRouteIdNil(b bool)`

 SetAppRouteIdNil sets the value for AppRouteId to be an explicit nil

### UnsetAppRouteId
`func (o *NewAppAuthInput) UnsetAppRouteId()`

UnsetAppRouteId ensures that no value is present for AppRouteId, not even an explicit nil
### GetLogin

`func (o *NewAppAuthInput) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *NewAppAuthInput) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *NewAppAuthInput) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetPassword

`func (o *NewAppAuthInput) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *NewAppAuthInput) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *NewAppAuthInput) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetRealm

`func (o *NewAppAuthInput) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *NewAppAuthInput) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *NewAppAuthInput) SetRealm(v string)`

SetRealm sets Realm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


