# ServiceManifestLinkValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Title** | Pointer to **NullableString** |  | [optional] 
**Secret** | **bool** |  | 
**Optional** | **bool** |  | 

## Methods

### NewServiceManifestLinkValue

`func NewServiceManifestLinkValue(name string, secret bool, optional bool, ) *ServiceManifestLinkValue`

NewServiceManifestLinkValue instantiates a new ServiceManifestLinkValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceManifestLinkValueWithDefaults

`func NewServiceManifestLinkValueWithDefaults() *ServiceManifestLinkValue`

NewServiceManifestLinkValueWithDefaults instantiates a new ServiceManifestLinkValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServiceManifestLinkValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceManifestLinkValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceManifestLinkValue) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *ServiceManifestLinkValue) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ServiceManifestLinkValue) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ServiceManifestLinkValue) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ServiceManifestLinkValue) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ServiceManifestLinkValue) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ServiceManifestLinkValue) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetSecret

`func (o *ServiceManifestLinkValue) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ServiceManifestLinkValue) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ServiceManifestLinkValue) SetSecret(v bool)`

SetSecret sets Secret field to given value.


### GetOptional

`func (o *ServiceManifestLinkValue) GetOptional() bool`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *ServiceManifestLinkValue) GetOptionalOk() (*bool, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *ServiceManifestLinkValue) SetOptional(v bool)`

SetOptional sets Optional field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


