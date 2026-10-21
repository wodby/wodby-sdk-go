# ServiceExcludedBaseVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseServiceRevisionId** | **int32** |  | 
**BaseServiceName** | **string** |  | 
**BaseServiceVersion** | **string** | The base service&#39;s newest version, which the inherited revision does not allow. | 
**BaseServiceRevNumber** | **int32** |  | 
**CurrentVersion** | **string** | The base service version this revision was built from. | 
**FromVersionConstraint** | Pointer to **NullableString** | The inherited manifest&#39;s fromVersionConstraint. Absent when the manifest pins an exact fromVersion. | [optional] 

## Methods

### NewServiceExcludedBaseVersion

`func NewServiceExcludedBaseVersion(baseServiceRevisionId int32, baseServiceName string, baseServiceVersion string, baseServiceRevNumber int32, currentVersion string, ) *ServiceExcludedBaseVersion`

NewServiceExcludedBaseVersion instantiates a new ServiceExcludedBaseVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceExcludedBaseVersionWithDefaults

`func NewServiceExcludedBaseVersionWithDefaults() *ServiceExcludedBaseVersion`

NewServiceExcludedBaseVersionWithDefaults instantiates a new ServiceExcludedBaseVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseServiceRevisionId

`func (o *ServiceExcludedBaseVersion) GetBaseServiceRevisionId() int32`

GetBaseServiceRevisionId returns the BaseServiceRevisionId field if non-nil, zero value otherwise.

### GetBaseServiceRevisionIdOk

`func (o *ServiceExcludedBaseVersion) GetBaseServiceRevisionIdOk() (*int32, bool)`

GetBaseServiceRevisionIdOk returns a tuple with the BaseServiceRevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseServiceRevisionId

`func (o *ServiceExcludedBaseVersion) SetBaseServiceRevisionId(v int32)`

SetBaseServiceRevisionId sets BaseServiceRevisionId field to given value.


### GetBaseServiceName

`func (o *ServiceExcludedBaseVersion) GetBaseServiceName() string`

GetBaseServiceName returns the BaseServiceName field if non-nil, zero value otherwise.

### GetBaseServiceNameOk

`func (o *ServiceExcludedBaseVersion) GetBaseServiceNameOk() (*string, bool)`

GetBaseServiceNameOk returns a tuple with the BaseServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseServiceName

`func (o *ServiceExcludedBaseVersion) SetBaseServiceName(v string)`

SetBaseServiceName sets BaseServiceName field to given value.


### GetBaseServiceVersion

`func (o *ServiceExcludedBaseVersion) GetBaseServiceVersion() string`

GetBaseServiceVersion returns the BaseServiceVersion field if non-nil, zero value otherwise.

### GetBaseServiceVersionOk

`func (o *ServiceExcludedBaseVersion) GetBaseServiceVersionOk() (*string, bool)`

GetBaseServiceVersionOk returns a tuple with the BaseServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseServiceVersion

`func (o *ServiceExcludedBaseVersion) SetBaseServiceVersion(v string)`

SetBaseServiceVersion sets BaseServiceVersion field to given value.


### GetBaseServiceRevNumber

`func (o *ServiceExcludedBaseVersion) GetBaseServiceRevNumber() int32`

GetBaseServiceRevNumber returns the BaseServiceRevNumber field if non-nil, zero value otherwise.

### GetBaseServiceRevNumberOk

`func (o *ServiceExcludedBaseVersion) GetBaseServiceRevNumberOk() (*int32, bool)`

GetBaseServiceRevNumberOk returns a tuple with the BaseServiceRevNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseServiceRevNumber

`func (o *ServiceExcludedBaseVersion) SetBaseServiceRevNumber(v int32)`

SetBaseServiceRevNumber sets BaseServiceRevNumber field to given value.


### GetCurrentVersion

`func (o *ServiceExcludedBaseVersion) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *ServiceExcludedBaseVersion) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *ServiceExcludedBaseVersion) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.


### GetFromVersionConstraint

`func (o *ServiceExcludedBaseVersion) GetFromVersionConstraint() string`

GetFromVersionConstraint returns the FromVersionConstraint field if non-nil, zero value otherwise.

### GetFromVersionConstraintOk

`func (o *ServiceExcludedBaseVersion) GetFromVersionConstraintOk() (*string, bool)`

GetFromVersionConstraintOk returns a tuple with the FromVersionConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromVersionConstraint

`func (o *ServiceExcludedBaseVersion) SetFromVersionConstraint(v string)`

SetFromVersionConstraint sets FromVersionConstraint field to given value.

### HasFromVersionConstraint

`func (o *ServiceExcludedBaseVersion) HasFromVersionConstraint() bool`

HasFromVersionConstraint returns a boolean if a field has been set.

### SetFromVersionConstraintNil

`func (o *ServiceExcludedBaseVersion) SetFromVersionConstraintNil(b bool)`

 SetFromVersionConstraintNil sets the value for FromVersionConstraint to be an explicit nil

### UnsetFromVersionConstraint
`func (o *ServiceExcludedBaseVersion) UnsetFromVersionConstraint()`

UnsetFromVersionConstraint ensures that no value is present for FromVersionConstraint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


