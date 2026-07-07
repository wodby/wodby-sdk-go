# ManifestValidationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Valid** | **bool** |  | 
**Error** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to **map[string]interface{}** |  | [optional] 
**Manifest** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewManifestValidationResponse

`func NewManifestValidationResponse(valid bool, ) *ManifestValidationResponse`

NewManifestValidationResponse instantiates a new ManifestValidationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManifestValidationResponseWithDefaults

`func NewManifestValidationResponseWithDefaults() *ManifestValidationResponse`

NewManifestValidationResponseWithDefaults instantiates a new ManifestValidationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValid

`func (o *ManifestValidationResponse) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *ManifestValidationResponse) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *ManifestValidationResponse) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetError

`func (o *ManifestValidationResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ManifestValidationResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ManifestValidationResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ManifestValidationResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetResource

`func (o *ManifestValidationResponse) GetResource() map[string]interface{}`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ManifestValidationResponse) GetResourceOk() (*map[string]interface{}, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ManifestValidationResponse) SetResource(v map[string]interface{})`

SetResource sets Resource field to given value.

### HasResource

`func (o *ManifestValidationResponse) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetManifest

`func (o *ManifestValidationResponse) GetManifest() map[string]interface{}`

GetManifest returns the Manifest field if non-nil, zero value otherwise.

### GetManifestOk

`func (o *ManifestValidationResponse) GetManifestOk() (*map[string]interface{}, bool)`

GetManifestOk returns a tuple with the Manifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifest

`func (o *ManifestValidationResponse) SetManifest(v map[string]interface{})`

SetManifest sets Manifest field to given value.

### HasManifest

`func (o *ManifestValidationResponse) HasManifest() bool`

HasManifest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


