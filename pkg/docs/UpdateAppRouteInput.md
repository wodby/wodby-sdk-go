# UpdateAppRouteInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **NullableBool** |  | [optional] 
**Main** | Pointer to **NullableBool** |  | [optional] 
**Primary** | Pointer to **NullableBool** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**PathType** | Pointer to **NullableString** |  | [optional] 
**Action** | Pointer to **NullableString** |  | [optional] 
**RedirectScheme** | Pointer to **NullableString** |  | [optional] 
**RedirectHost** | Pointer to **NullableString** |  | [optional] 
**RedirectPath** | Pointer to **NullableString** |  | [optional] 
**RedirectStatusCode** | Pointer to **NullableInt32** |  | [optional] 
**Options** | Pointer to [**[]AppEndpointOptionInput**](AppEndpointOptionInput.md) |  | [optional] 

## Methods

### NewUpdateAppRouteInput

`func NewUpdateAppRouteInput() *UpdateAppRouteInput`

NewUpdateAppRouteInput instantiates a new UpdateAppRouteInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAppRouteInputWithDefaults

`func NewUpdateAppRouteInputWithDefaults() *UpdateAppRouteInput`

NewUpdateAppRouteInputWithDefaults instantiates a new UpdateAppRouteInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *UpdateAppRouteInput) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *UpdateAppRouteInput) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *UpdateAppRouteInput) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *UpdateAppRouteInput) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### SetDisabledNil

`func (o *UpdateAppRouteInput) SetDisabledNil(b bool)`

 SetDisabledNil sets the value for Disabled to be an explicit nil

### UnsetDisabled
`func (o *UpdateAppRouteInput) UnsetDisabled()`

UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
### GetMain

`func (o *UpdateAppRouteInput) GetMain() bool`

GetMain returns the Main field if non-nil, zero value otherwise.

### GetMainOk

`func (o *UpdateAppRouteInput) GetMainOk() (*bool, bool)`

GetMainOk returns a tuple with the Main field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMain

`func (o *UpdateAppRouteInput) SetMain(v bool)`

SetMain sets Main field to given value.

### HasMain

`func (o *UpdateAppRouteInput) HasMain() bool`

HasMain returns a boolean if a field has been set.

### SetMainNil

`func (o *UpdateAppRouteInput) SetMainNil(b bool)`

 SetMainNil sets the value for Main to be an explicit nil

### UnsetMain
`func (o *UpdateAppRouteInput) UnsetMain()`

UnsetMain ensures that no value is present for Main, not even an explicit nil
### GetPrimary

`func (o *UpdateAppRouteInput) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *UpdateAppRouteInput) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *UpdateAppRouteInput) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *UpdateAppRouteInput) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### SetPrimaryNil

`func (o *UpdateAppRouteInput) SetPrimaryNil(b bool)`

 SetPrimaryNil sets the value for Primary to be an explicit nil

### UnsetPrimary
`func (o *UpdateAppRouteInput) UnsetPrimary()`

UnsetPrimary ensures that no value is present for Primary, not even an explicit nil
### GetPath

`func (o *UpdateAppRouteInput) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *UpdateAppRouteInput) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *UpdateAppRouteInput) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *UpdateAppRouteInput) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *UpdateAppRouteInput) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *UpdateAppRouteInput) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetPathType

`func (o *UpdateAppRouteInput) GetPathType() string`

GetPathType returns the PathType field if non-nil, zero value otherwise.

### GetPathTypeOk

`func (o *UpdateAppRouteInput) GetPathTypeOk() (*string, bool)`

GetPathTypeOk returns a tuple with the PathType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathType

`func (o *UpdateAppRouteInput) SetPathType(v string)`

SetPathType sets PathType field to given value.

### HasPathType

`func (o *UpdateAppRouteInput) HasPathType() bool`

HasPathType returns a boolean if a field has been set.

### SetPathTypeNil

`func (o *UpdateAppRouteInput) SetPathTypeNil(b bool)`

 SetPathTypeNil sets the value for PathType to be an explicit nil

### UnsetPathType
`func (o *UpdateAppRouteInput) UnsetPathType()`

UnsetPathType ensures that no value is present for PathType, not even an explicit nil
### GetAction

`func (o *UpdateAppRouteInput) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UpdateAppRouteInput) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UpdateAppRouteInput) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *UpdateAppRouteInput) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *UpdateAppRouteInput) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *UpdateAppRouteInput) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetRedirectScheme

`func (o *UpdateAppRouteInput) GetRedirectScheme() string`

GetRedirectScheme returns the RedirectScheme field if non-nil, zero value otherwise.

### GetRedirectSchemeOk

`func (o *UpdateAppRouteInput) GetRedirectSchemeOk() (*string, bool)`

GetRedirectSchemeOk returns a tuple with the RedirectScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectScheme

`func (o *UpdateAppRouteInput) SetRedirectScheme(v string)`

SetRedirectScheme sets RedirectScheme field to given value.

### HasRedirectScheme

`func (o *UpdateAppRouteInput) HasRedirectScheme() bool`

HasRedirectScheme returns a boolean if a field has been set.

### SetRedirectSchemeNil

`func (o *UpdateAppRouteInput) SetRedirectSchemeNil(b bool)`

 SetRedirectSchemeNil sets the value for RedirectScheme to be an explicit nil

### UnsetRedirectScheme
`func (o *UpdateAppRouteInput) UnsetRedirectScheme()`

UnsetRedirectScheme ensures that no value is present for RedirectScheme, not even an explicit nil
### GetRedirectHost

`func (o *UpdateAppRouteInput) GetRedirectHost() string`

GetRedirectHost returns the RedirectHost field if non-nil, zero value otherwise.

### GetRedirectHostOk

`func (o *UpdateAppRouteInput) GetRedirectHostOk() (*string, bool)`

GetRedirectHostOk returns a tuple with the RedirectHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectHost

`func (o *UpdateAppRouteInput) SetRedirectHost(v string)`

SetRedirectHost sets RedirectHost field to given value.

### HasRedirectHost

`func (o *UpdateAppRouteInput) HasRedirectHost() bool`

HasRedirectHost returns a boolean if a field has been set.

### SetRedirectHostNil

`func (o *UpdateAppRouteInput) SetRedirectHostNil(b bool)`

 SetRedirectHostNil sets the value for RedirectHost to be an explicit nil

### UnsetRedirectHost
`func (o *UpdateAppRouteInput) UnsetRedirectHost()`

UnsetRedirectHost ensures that no value is present for RedirectHost, not even an explicit nil
### GetRedirectPath

`func (o *UpdateAppRouteInput) GetRedirectPath() string`

GetRedirectPath returns the RedirectPath field if non-nil, zero value otherwise.

### GetRedirectPathOk

`func (o *UpdateAppRouteInput) GetRedirectPathOk() (*string, bool)`

GetRedirectPathOk returns a tuple with the RedirectPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectPath

`func (o *UpdateAppRouteInput) SetRedirectPath(v string)`

SetRedirectPath sets RedirectPath field to given value.

### HasRedirectPath

`func (o *UpdateAppRouteInput) HasRedirectPath() bool`

HasRedirectPath returns a boolean if a field has been set.

### SetRedirectPathNil

`func (o *UpdateAppRouteInput) SetRedirectPathNil(b bool)`

 SetRedirectPathNil sets the value for RedirectPath to be an explicit nil

### UnsetRedirectPath
`func (o *UpdateAppRouteInput) UnsetRedirectPath()`

UnsetRedirectPath ensures that no value is present for RedirectPath, not even an explicit nil
### GetRedirectStatusCode

`func (o *UpdateAppRouteInput) GetRedirectStatusCode() int32`

GetRedirectStatusCode returns the RedirectStatusCode field if non-nil, zero value otherwise.

### GetRedirectStatusCodeOk

`func (o *UpdateAppRouteInput) GetRedirectStatusCodeOk() (*int32, bool)`

GetRedirectStatusCodeOk returns a tuple with the RedirectStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectStatusCode

`func (o *UpdateAppRouteInput) SetRedirectStatusCode(v int32)`

SetRedirectStatusCode sets RedirectStatusCode field to given value.

### HasRedirectStatusCode

`func (o *UpdateAppRouteInput) HasRedirectStatusCode() bool`

HasRedirectStatusCode returns a boolean if a field has been set.

### SetRedirectStatusCodeNil

`func (o *UpdateAppRouteInput) SetRedirectStatusCodeNil(b bool)`

 SetRedirectStatusCodeNil sets the value for RedirectStatusCode to be an explicit nil

### UnsetRedirectStatusCode
`func (o *UpdateAppRouteInput) UnsetRedirectStatusCode()`

UnsetRedirectStatusCode ensures that no value is present for RedirectStatusCode, not even an explicit nil
### GetOptions

`func (o *UpdateAppRouteInput) GetOptions() []AppEndpointOptionInput`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *UpdateAppRouteInput) GetOptionsOk() (*[]AppEndpointOptionInput, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *UpdateAppRouteInput) SetOptions(v []AppEndpointOptionInput)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *UpdateAppRouteInput) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


