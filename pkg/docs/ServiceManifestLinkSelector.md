# ServiceManifestLinkSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The service type a target must have, in upper case, such as DB or DATASTORE. | 
**Option** | Pointer to **string** | The service version a target must run, when the selector asks for one. | [optional] 
**Labels** | Pointer to **[]string** | Labels a target service must all have. | [optional] 

## Methods

### NewServiceManifestLinkSelector

`func NewServiceManifestLinkSelector(type_ string, ) *ServiceManifestLinkSelector`

NewServiceManifestLinkSelector instantiates a new ServiceManifestLinkSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceManifestLinkSelectorWithDefaults

`func NewServiceManifestLinkSelectorWithDefaults() *ServiceManifestLinkSelector`

NewServiceManifestLinkSelectorWithDefaults instantiates a new ServiceManifestLinkSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ServiceManifestLinkSelector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceManifestLinkSelector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceManifestLinkSelector) SetType(v string)`

SetType sets Type field to given value.


### GetOption

`func (o *ServiceManifestLinkSelector) GetOption() string`

GetOption returns the Option field if non-nil, zero value otherwise.

### GetOptionOk

`func (o *ServiceManifestLinkSelector) GetOptionOk() (*string, bool)`

GetOptionOk returns a tuple with the Option field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption

`func (o *ServiceManifestLinkSelector) SetOption(v string)`

SetOption sets Option field to given value.

### HasOption

`func (o *ServiceManifestLinkSelector) HasOption() bool`

HasOption returns a boolean if a field has been set.

### GetLabels

`func (o *ServiceManifestLinkSelector) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ServiceManifestLinkSelector) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ServiceManifestLinkSelector) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ServiceManifestLinkSelector) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


