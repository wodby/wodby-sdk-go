# NewAppRouteInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppServiceId** | **int32** |  | 
**Disabled** | Pointer to **NullableBool** | Creates the custom domain disabled. Disabled domains do not require custom-domain feature access until enabled. | [optional] 
**Main** | **bool** |  | 
**Primary** | **bool** |  | 
**Port** | **int32** |  | 
**Host** | **string** |  | 
**Path** | Pointer to **NullableString** |  | [optional] 
**PathType** | Pointer to **NullableString** |  | [optional] 
**Action** | Pointer to **NullableString** | SERVE sends requests to the selected app service. BACKEND is accepted for backwards compatibility. | [optional] 
**RedirectScheme** | Pointer to **NullableString** |  | [optional] 
**RedirectHost** | Pointer to **NullableString** |  | [optional] 
**RedirectPath** | Pointer to **NullableString** |  | [optional] 
**RedirectStatusCode** | Pointer to **NullableInt32** |  | [optional] 
**Hsts** | Pointer to **NullableBool** | Enables HTTP Strict Transport Security for a serve route when TLS is active. | [optional] 
**Letsencrypt** | Pointer to **NullableBool** |  | [optional] 
**Tls** | Pointer to [**AppRouteTLSInput**](AppRouteTLSInput.md) |  | [optional] 

## Methods

### NewNewAppRouteInput

`func NewNewAppRouteInput(appServiceId int32, main bool, primary bool, port int32, host string, ) *NewAppRouteInput`

NewNewAppRouteInput instantiates a new NewAppRouteInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewAppRouteInputWithDefaults

`func NewNewAppRouteInputWithDefaults() *NewAppRouteInput`

NewNewAppRouteInputWithDefaults instantiates a new NewAppRouteInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppServiceId

`func (o *NewAppRouteInput) GetAppServiceId() int32`

GetAppServiceId returns the AppServiceId field if non-nil, zero value otherwise.

### GetAppServiceIdOk

`func (o *NewAppRouteInput) GetAppServiceIdOk() (*int32, bool)`

GetAppServiceIdOk returns a tuple with the AppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceId

`func (o *NewAppRouteInput) SetAppServiceId(v int32)`

SetAppServiceId sets AppServiceId field to given value.


### GetDisabled

`func (o *NewAppRouteInput) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *NewAppRouteInput) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *NewAppRouteInput) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *NewAppRouteInput) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### SetDisabledNil

`func (o *NewAppRouteInput) SetDisabledNil(b bool)`

 SetDisabledNil sets the value for Disabled to be an explicit nil

### UnsetDisabled
`func (o *NewAppRouteInput) UnsetDisabled()`

UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
### GetMain

`func (o *NewAppRouteInput) GetMain() bool`

GetMain returns the Main field if non-nil, zero value otherwise.

### GetMainOk

`func (o *NewAppRouteInput) GetMainOk() (*bool, bool)`

GetMainOk returns a tuple with the Main field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMain

`func (o *NewAppRouteInput) SetMain(v bool)`

SetMain sets Main field to given value.


### GetPrimary

`func (o *NewAppRouteInput) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *NewAppRouteInput) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *NewAppRouteInput) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.


### GetPort

`func (o *NewAppRouteInput) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *NewAppRouteInput) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *NewAppRouteInput) SetPort(v int32)`

SetPort sets Port field to given value.


### GetHost

`func (o *NewAppRouteInput) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *NewAppRouteInput) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *NewAppRouteInput) SetHost(v string)`

SetHost sets Host field to given value.


### GetPath

`func (o *NewAppRouteInput) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *NewAppRouteInput) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *NewAppRouteInput) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *NewAppRouteInput) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *NewAppRouteInput) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *NewAppRouteInput) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetPathType

`func (o *NewAppRouteInput) GetPathType() string`

GetPathType returns the PathType field if non-nil, zero value otherwise.

### GetPathTypeOk

`func (o *NewAppRouteInput) GetPathTypeOk() (*string, bool)`

GetPathTypeOk returns a tuple with the PathType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathType

`func (o *NewAppRouteInput) SetPathType(v string)`

SetPathType sets PathType field to given value.

### HasPathType

`func (o *NewAppRouteInput) HasPathType() bool`

HasPathType returns a boolean if a field has been set.

### SetPathTypeNil

`func (o *NewAppRouteInput) SetPathTypeNil(b bool)`

 SetPathTypeNil sets the value for PathType to be an explicit nil

### UnsetPathType
`func (o *NewAppRouteInput) UnsetPathType()`

UnsetPathType ensures that no value is present for PathType, not even an explicit nil
### GetAction

`func (o *NewAppRouteInput) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *NewAppRouteInput) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *NewAppRouteInput) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *NewAppRouteInput) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *NewAppRouteInput) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *NewAppRouteInput) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetRedirectScheme

`func (o *NewAppRouteInput) GetRedirectScheme() string`

GetRedirectScheme returns the RedirectScheme field if non-nil, zero value otherwise.

### GetRedirectSchemeOk

`func (o *NewAppRouteInput) GetRedirectSchemeOk() (*string, bool)`

GetRedirectSchemeOk returns a tuple with the RedirectScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectScheme

`func (o *NewAppRouteInput) SetRedirectScheme(v string)`

SetRedirectScheme sets RedirectScheme field to given value.

### HasRedirectScheme

`func (o *NewAppRouteInput) HasRedirectScheme() bool`

HasRedirectScheme returns a boolean if a field has been set.

### SetRedirectSchemeNil

`func (o *NewAppRouteInput) SetRedirectSchemeNil(b bool)`

 SetRedirectSchemeNil sets the value for RedirectScheme to be an explicit nil

### UnsetRedirectScheme
`func (o *NewAppRouteInput) UnsetRedirectScheme()`

UnsetRedirectScheme ensures that no value is present for RedirectScheme, not even an explicit nil
### GetRedirectHost

`func (o *NewAppRouteInput) GetRedirectHost() string`

GetRedirectHost returns the RedirectHost field if non-nil, zero value otherwise.

### GetRedirectHostOk

`func (o *NewAppRouteInput) GetRedirectHostOk() (*string, bool)`

GetRedirectHostOk returns a tuple with the RedirectHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectHost

`func (o *NewAppRouteInput) SetRedirectHost(v string)`

SetRedirectHost sets RedirectHost field to given value.

### HasRedirectHost

`func (o *NewAppRouteInput) HasRedirectHost() bool`

HasRedirectHost returns a boolean if a field has been set.

### SetRedirectHostNil

`func (o *NewAppRouteInput) SetRedirectHostNil(b bool)`

 SetRedirectHostNil sets the value for RedirectHost to be an explicit nil

### UnsetRedirectHost
`func (o *NewAppRouteInput) UnsetRedirectHost()`

UnsetRedirectHost ensures that no value is present for RedirectHost, not even an explicit nil
### GetRedirectPath

`func (o *NewAppRouteInput) GetRedirectPath() string`

GetRedirectPath returns the RedirectPath field if non-nil, zero value otherwise.

### GetRedirectPathOk

`func (o *NewAppRouteInput) GetRedirectPathOk() (*string, bool)`

GetRedirectPathOk returns a tuple with the RedirectPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectPath

`func (o *NewAppRouteInput) SetRedirectPath(v string)`

SetRedirectPath sets RedirectPath field to given value.

### HasRedirectPath

`func (o *NewAppRouteInput) HasRedirectPath() bool`

HasRedirectPath returns a boolean if a field has been set.

### SetRedirectPathNil

`func (o *NewAppRouteInput) SetRedirectPathNil(b bool)`

 SetRedirectPathNil sets the value for RedirectPath to be an explicit nil

### UnsetRedirectPath
`func (o *NewAppRouteInput) UnsetRedirectPath()`

UnsetRedirectPath ensures that no value is present for RedirectPath, not even an explicit nil
### GetRedirectStatusCode

`func (o *NewAppRouteInput) GetRedirectStatusCode() int32`

GetRedirectStatusCode returns the RedirectStatusCode field if non-nil, zero value otherwise.

### GetRedirectStatusCodeOk

`func (o *NewAppRouteInput) GetRedirectStatusCodeOk() (*int32, bool)`

GetRedirectStatusCodeOk returns a tuple with the RedirectStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectStatusCode

`func (o *NewAppRouteInput) SetRedirectStatusCode(v int32)`

SetRedirectStatusCode sets RedirectStatusCode field to given value.

### HasRedirectStatusCode

`func (o *NewAppRouteInput) HasRedirectStatusCode() bool`

HasRedirectStatusCode returns a boolean if a field has been set.

### SetRedirectStatusCodeNil

`func (o *NewAppRouteInput) SetRedirectStatusCodeNil(b bool)`

 SetRedirectStatusCodeNil sets the value for RedirectStatusCode to be an explicit nil

### UnsetRedirectStatusCode
`func (o *NewAppRouteInput) UnsetRedirectStatusCode()`

UnsetRedirectStatusCode ensures that no value is present for RedirectStatusCode, not even an explicit nil
### GetHsts

`func (o *NewAppRouteInput) GetHsts() bool`

GetHsts returns the Hsts field if non-nil, zero value otherwise.

### GetHstsOk

`func (o *NewAppRouteInput) GetHstsOk() (*bool, bool)`

GetHstsOk returns a tuple with the Hsts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsts

`func (o *NewAppRouteInput) SetHsts(v bool)`

SetHsts sets Hsts field to given value.

### HasHsts

`func (o *NewAppRouteInput) HasHsts() bool`

HasHsts returns a boolean if a field has been set.

### SetHstsNil

`func (o *NewAppRouteInput) SetHstsNil(b bool)`

 SetHstsNil sets the value for Hsts to be an explicit nil

### UnsetHsts
`func (o *NewAppRouteInput) UnsetHsts()`

UnsetHsts ensures that no value is present for Hsts, not even an explicit nil
### GetLetsencrypt

`func (o *NewAppRouteInput) GetLetsencrypt() bool`

GetLetsencrypt returns the Letsencrypt field if non-nil, zero value otherwise.

### GetLetsencryptOk

`func (o *NewAppRouteInput) GetLetsencryptOk() (*bool, bool)`

GetLetsencryptOk returns a tuple with the Letsencrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLetsencrypt

`func (o *NewAppRouteInput) SetLetsencrypt(v bool)`

SetLetsencrypt sets Letsencrypt field to given value.

### HasLetsencrypt

`func (o *NewAppRouteInput) HasLetsencrypt() bool`

HasLetsencrypt returns a boolean if a field has been set.

### SetLetsencryptNil

`func (o *NewAppRouteInput) SetLetsencryptNil(b bool)`

 SetLetsencryptNil sets the value for Letsencrypt to be an explicit nil

### UnsetLetsencrypt
`func (o *NewAppRouteInput) UnsetLetsencrypt()`

UnsetLetsencrypt ensures that no value is present for Letsencrypt, not even an explicit nil
### GetTls

`func (o *NewAppRouteInput) GetTls() AppRouteTLSInput`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *NewAppRouteInput) GetTlsOk() (*AppRouteTLSInput, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *NewAppRouteInput) SetTls(v AppRouteTLSInput)`

SetTls sets Tls field to given value.

### HasTls

`func (o *NewAppRouteInput) HasTls() bool`

HasTls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


