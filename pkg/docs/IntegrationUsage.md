# IntegrationUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Category** | **string** | current for what uses the integration now, historical for records such as builds and backups, cleanup for pending removal of resources it created. | 
**Kind** | **string** |  | 
**Detail** | Pointer to **NullableString** |  | [optional] 
**Target** | [**IntegrationUsageTarget**](IntegrationUsageTarget.md) |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewIntegrationUsage

`func NewIntegrationUsage(id string, category string, kind string, target IntegrationUsageTarget, createdAt time.Time, ) *IntegrationUsage`

NewIntegrationUsage instantiates a new IntegrationUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationUsageWithDefaults

`func NewIntegrationUsageWithDefaults() *IntegrationUsage`

NewIntegrationUsageWithDefaults instantiates a new IntegrationUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntegrationUsage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationUsage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationUsage) SetId(v string)`

SetId sets Id field to given value.


### GetCategory

`func (o *IntegrationUsage) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *IntegrationUsage) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *IntegrationUsage) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetKind

`func (o *IntegrationUsage) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IntegrationUsage) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IntegrationUsage) SetKind(v string)`

SetKind sets Kind field to given value.


### GetDetail

`func (o *IntegrationUsage) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *IntegrationUsage) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *IntegrationUsage) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *IntegrationUsage) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *IntegrationUsage) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *IntegrationUsage) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetTarget

`func (o *IntegrationUsage) GetTarget() IntegrationUsageTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *IntegrationUsage) GetTargetOk() (*IntegrationUsageTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *IntegrationUsage) SetTarget(v IntegrationUsageTarget)`

SetTarget sets Target field to given value.


### GetCreatedAt

`func (o *IntegrationUsage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IntegrationUsage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IntegrationUsage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


