# AppRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Host** | **string** |  | 
**Path** | **string** |  | 
**PathType** | **string** |  | 
**Action** | **string** |  | 
**RedirectScheme** | Pointer to **NullableString** |  | [optional] 
**RedirectHost** | Pointer to **NullableString** |  | [optional] 
**RedirectPath** | Pointer to **NullableString** |  | [optional] 
**RedirectStatusCode** | Pointer to **NullableInt32** |  | [optional] 
**Status** | **string** |  | 
**Disabled** | **bool** |  | 
**Main** | **bool** |  | 
**Primary** | **bool** |  | 
**Private** | **bool** |  | 
**AppInstanceId** | **int32** |  | 
**AppServiceId** | **int32** |  | 
**PortId** | **int32** |  | 
**Cert** | Pointer to [**NullableCert**](Cert.md) |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**LastSyncedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewAppRoute

`func NewAppRoute(id int32, host string, path string, pathType string, action string, status string, disabled bool, main bool, primary bool, private bool, appInstanceId int32, appServiceId int32, portId int32, createdAt time.Time, updatedAt time.Time, ) *AppRoute`

NewAppRoute instantiates a new AppRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRouteWithDefaults

`func NewAppRouteWithDefaults() *AppRoute`

NewAppRouteWithDefaults instantiates a new AppRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppRoute) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppRoute) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppRoute) SetId(v int32)`

SetId sets Id field to given value.


### GetHost

`func (o *AppRoute) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *AppRoute) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *AppRoute) SetHost(v string)`

SetHost sets Host field to given value.


### GetPath

`func (o *AppRoute) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AppRoute) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AppRoute) SetPath(v string)`

SetPath sets Path field to given value.


### GetPathType

`func (o *AppRoute) GetPathType() string`

GetPathType returns the PathType field if non-nil, zero value otherwise.

### GetPathTypeOk

`func (o *AppRoute) GetPathTypeOk() (*string, bool)`

GetPathTypeOk returns a tuple with the PathType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathType

`func (o *AppRoute) SetPathType(v string)`

SetPathType sets PathType field to given value.


### GetAction

`func (o *AppRoute) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AppRoute) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AppRoute) SetAction(v string)`

SetAction sets Action field to given value.


### GetRedirectScheme

`func (o *AppRoute) GetRedirectScheme() string`

GetRedirectScheme returns the RedirectScheme field if non-nil, zero value otherwise.

### GetRedirectSchemeOk

`func (o *AppRoute) GetRedirectSchemeOk() (*string, bool)`

GetRedirectSchemeOk returns a tuple with the RedirectScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectScheme

`func (o *AppRoute) SetRedirectScheme(v string)`

SetRedirectScheme sets RedirectScheme field to given value.

### HasRedirectScheme

`func (o *AppRoute) HasRedirectScheme() bool`

HasRedirectScheme returns a boolean if a field has been set.

### SetRedirectSchemeNil

`func (o *AppRoute) SetRedirectSchemeNil(b bool)`

 SetRedirectSchemeNil sets the value for RedirectScheme to be an explicit nil

### UnsetRedirectScheme
`func (o *AppRoute) UnsetRedirectScheme()`

UnsetRedirectScheme ensures that no value is present for RedirectScheme, not even an explicit nil
### GetRedirectHost

`func (o *AppRoute) GetRedirectHost() string`

GetRedirectHost returns the RedirectHost field if non-nil, zero value otherwise.

### GetRedirectHostOk

`func (o *AppRoute) GetRedirectHostOk() (*string, bool)`

GetRedirectHostOk returns a tuple with the RedirectHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectHost

`func (o *AppRoute) SetRedirectHost(v string)`

SetRedirectHost sets RedirectHost field to given value.

### HasRedirectHost

`func (o *AppRoute) HasRedirectHost() bool`

HasRedirectHost returns a boolean if a field has been set.

### SetRedirectHostNil

`func (o *AppRoute) SetRedirectHostNil(b bool)`

 SetRedirectHostNil sets the value for RedirectHost to be an explicit nil

### UnsetRedirectHost
`func (o *AppRoute) UnsetRedirectHost()`

UnsetRedirectHost ensures that no value is present for RedirectHost, not even an explicit nil
### GetRedirectPath

`func (o *AppRoute) GetRedirectPath() string`

GetRedirectPath returns the RedirectPath field if non-nil, zero value otherwise.

### GetRedirectPathOk

`func (o *AppRoute) GetRedirectPathOk() (*string, bool)`

GetRedirectPathOk returns a tuple with the RedirectPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectPath

`func (o *AppRoute) SetRedirectPath(v string)`

SetRedirectPath sets RedirectPath field to given value.

### HasRedirectPath

`func (o *AppRoute) HasRedirectPath() bool`

HasRedirectPath returns a boolean if a field has been set.

### SetRedirectPathNil

`func (o *AppRoute) SetRedirectPathNil(b bool)`

 SetRedirectPathNil sets the value for RedirectPath to be an explicit nil

### UnsetRedirectPath
`func (o *AppRoute) UnsetRedirectPath()`

UnsetRedirectPath ensures that no value is present for RedirectPath, not even an explicit nil
### GetRedirectStatusCode

`func (o *AppRoute) GetRedirectStatusCode() int32`

GetRedirectStatusCode returns the RedirectStatusCode field if non-nil, zero value otherwise.

### GetRedirectStatusCodeOk

`func (o *AppRoute) GetRedirectStatusCodeOk() (*int32, bool)`

GetRedirectStatusCodeOk returns a tuple with the RedirectStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectStatusCode

`func (o *AppRoute) SetRedirectStatusCode(v int32)`

SetRedirectStatusCode sets RedirectStatusCode field to given value.

### HasRedirectStatusCode

`func (o *AppRoute) HasRedirectStatusCode() bool`

HasRedirectStatusCode returns a boolean if a field has been set.

### SetRedirectStatusCodeNil

`func (o *AppRoute) SetRedirectStatusCodeNil(b bool)`

 SetRedirectStatusCodeNil sets the value for RedirectStatusCode to be an explicit nil

### UnsetRedirectStatusCode
`func (o *AppRoute) UnsetRedirectStatusCode()`

UnsetRedirectStatusCode ensures that no value is present for RedirectStatusCode, not even an explicit nil
### GetStatus

`func (o *AppRoute) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppRoute) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppRoute) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDisabled

`func (o *AppRoute) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *AppRoute) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *AppRoute) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetMain

`func (o *AppRoute) GetMain() bool`

GetMain returns the Main field if non-nil, zero value otherwise.

### GetMainOk

`func (o *AppRoute) GetMainOk() (*bool, bool)`

GetMainOk returns a tuple with the Main field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMain

`func (o *AppRoute) SetMain(v bool)`

SetMain sets Main field to given value.


### GetPrimary

`func (o *AppRoute) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *AppRoute) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *AppRoute) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.


### GetPrivate

`func (o *AppRoute) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *AppRoute) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *AppRoute) SetPrivate(v bool)`

SetPrivate sets Private field to given value.


### GetAppInstanceId

`func (o *AppRoute) GetAppInstanceId() int32`

GetAppInstanceId returns the AppInstanceId field if non-nil, zero value otherwise.

### GetAppInstanceIdOk

`func (o *AppRoute) GetAppInstanceIdOk() (*int32, bool)`

GetAppInstanceIdOk returns a tuple with the AppInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceId

`func (o *AppRoute) SetAppInstanceId(v int32)`

SetAppInstanceId sets AppInstanceId field to given value.


### GetAppServiceId

`func (o *AppRoute) GetAppServiceId() int32`

GetAppServiceId returns the AppServiceId field if non-nil, zero value otherwise.

### GetAppServiceIdOk

`func (o *AppRoute) GetAppServiceIdOk() (*int32, bool)`

GetAppServiceIdOk returns a tuple with the AppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceId

`func (o *AppRoute) SetAppServiceId(v int32)`

SetAppServiceId sets AppServiceId field to given value.


### GetPortId

`func (o *AppRoute) GetPortId() int32`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *AppRoute) GetPortIdOk() (*int32, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *AppRoute) SetPortId(v int32)`

SetPortId sets PortId field to given value.


### GetCert

`func (o *AppRoute) GetCert() Cert`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *AppRoute) GetCertOk() (*Cert, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCert

`func (o *AppRoute) SetCert(v Cert)`

SetCert sets Cert field to given value.

### HasCert

`func (o *AppRoute) HasCert() bool`

HasCert returns a boolean if a field has been set.

### SetCertNil

`func (o *AppRoute) SetCertNil(b bool)`

 SetCertNil sets the value for Cert to be an explicit nil

### UnsetCert
`func (o *AppRoute) UnsetCert()`

UnsetCert ensures that no value is present for Cert, not even an explicit nil
### GetCreatedAt

`func (o *AppRoute) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppRoute) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppRoute) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AppRoute) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AppRoute) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AppRoute) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLastSyncedAt

`func (o *AppRoute) GetLastSyncedAt() time.Time`

GetLastSyncedAt returns the LastSyncedAt field if non-nil, zero value otherwise.

### GetLastSyncedAtOk

`func (o *AppRoute) GetLastSyncedAtOk() (*time.Time, bool)`

GetLastSyncedAtOk returns a tuple with the LastSyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncedAt

`func (o *AppRoute) SetLastSyncedAt(v time.Time)`

SetLastSyncedAt sets LastSyncedAt field to given value.

### HasLastSyncedAt

`func (o *AppRoute) HasLastSyncedAt() bool`

HasLastSyncedAt returns a boolean if a field has been set.

### SetLastSyncedAtNil

`func (o *AppRoute) SetLastSyncedAtNil(b bool)`

 SetLastSyncedAtNil sets the value for LastSyncedAt to be an explicit nil

### UnsetLastSyncedAt
`func (o *AppRoute) UnsetLastSyncedAt()`

UnsetLastSyncedAt ensures that no value is present for LastSyncedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


