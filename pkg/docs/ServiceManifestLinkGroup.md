# ServiceManifestLinkGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Title** | Pointer to **NullableString** |  | [optional] 
**Links** | **[]string** | Names of the links in the group. | 
**Min** | **int32** |  | 

## Methods

### NewServiceManifestLinkGroup

`func NewServiceManifestLinkGroup(name string, links []string, min int32, ) *ServiceManifestLinkGroup`

NewServiceManifestLinkGroup instantiates a new ServiceManifestLinkGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceManifestLinkGroupWithDefaults

`func NewServiceManifestLinkGroupWithDefaults() *ServiceManifestLinkGroup`

NewServiceManifestLinkGroupWithDefaults instantiates a new ServiceManifestLinkGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServiceManifestLinkGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceManifestLinkGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceManifestLinkGroup) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *ServiceManifestLinkGroup) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ServiceManifestLinkGroup) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ServiceManifestLinkGroup) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ServiceManifestLinkGroup) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ServiceManifestLinkGroup) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ServiceManifestLinkGroup) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetLinks

`func (o *ServiceManifestLinkGroup) GetLinks() []string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServiceManifestLinkGroup) GetLinksOk() (*[]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServiceManifestLinkGroup) SetLinks(v []string)`

SetLinks sets Links field to given value.


### GetMin

`func (o *ServiceManifestLinkGroup) GetMin() int32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *ServiceManifestLinkGroup) GetMinOk() (*int32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *ServiceManifestLinkGroup) SetMin(v int32)`

SetMin sets Min field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


