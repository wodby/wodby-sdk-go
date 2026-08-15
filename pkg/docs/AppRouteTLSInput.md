# AppRouteTLSInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** |  | 
**CertId** | Pointer to **NullableInt32** | Required only when mode is CUSTOM. | [optional] 

## Methods

### NewAppRouteTLSInput

`func NewAppRouteTLSInput(mode string, ) *AppRouteTLSInput`

NewAppRouteTLSInput instantiates a new AppRouteTLSInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRouteTLSInputWithDefaults

`func NewAppRouteTLSInputWithDefaults() *AppRouteTLSInput`

NewAppRouteTLSInputWithDefaults instantiates a new AppRouteTLSInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *AppRouteTLSInput) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *AppRouteTLSInput) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *AppRouteTLSInput) SetMode(v string)`

SetMode sets Mode field to given value.


### GetCertId

`func (o *AppRouteTLSInput) GetCertId() int32`

GetCertId returns the CertId field if non-nil, zero value otherwise.

### GetCertIdOk

`func (o *AppRouteTLSInput) GetCertIdOk() (*int32, bool)`

GetCertIdOk returns a tuple with the CertId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertId

`func (o *AppRouteTLSInput) SetCertId(v int32)`

SetCertId sets CertId field to given value.

### HasCertId

`func (o *AppRouteTLSInput) HasCertId() bool`

HasCertId returns a boolean if a field has been set.

### SetCertIdNil

`func (o *AppRouteTLSInput) SetCertIdNil(b bool)`

 SetCertIdNil sets the value for CertId to be an explicit nil

### UnsetCertId
`func (o *AppRouteTLSInput) UnsetCertId()`

UnsetCertId ensures that no value is present for CertId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


