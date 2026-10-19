# BulkAppRouteOperationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppEnvironmentId** | **int32** |  | 
**RouteIds** | **[]int32** | Routes of the app environment to change. | 
**Operation** | **string** |  | 

## Methods

### NewBulkAppRouteOperationInput

`func NewBulkAppRouteOperationInput(appEnvironmentId int32, routeIds []int32, operation string, ) *BulkAppRouteOperationInput`

NewBulkAppRouteOperationInput instantiates a new BulkAppRouteOperationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkAppRouteOperationInputWithDefaults

`func NewBulkAppRouteOperationInputWithDefaults() *BulkAppRouteOperationInput`

NewBulkAppRouteOperationInputWithDefaults instantiates a new BulkAppRouteOperationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppEnvironmentId

`func (o *BulkAppRouteOperationInput) GetAppEnvironmentId() int32`

GetAppEnvironmentId returns the AppEnvironmentId field if non-nil, zero value otherwise.

### GetAppEnvironmentIdOk

`func (o *BulkAppRouteOperationInput) GetAppEnvironmentIdOk() (*int32, bool)`

GetAppEnvironmentIdOk returns a tuple with the AppEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEnvironmentId

`func (o *BulkAppRouteOperationInput) SetAppEnvironmentId(v int32)`

SetAppEnvironmentId sets AppEnvironmentId field to given value.


### GetRouteIds

`func (o *BulkAppRouteOperationInput) GetRouteIds() []int32`

GetRouteIds returns the RouteIds field if non-nil, zero value otherwise.

### GetRouteIdsOk

`func (o *BulkAppRouteOperationInput) GetRouteIdsOk() (*[]int32, bool)`

GetRouteIdsOk returns a tuple with the RouteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteIds

`func (o *BulkAppRouteOperationInput) SetRouteIds(v []int32)`

SetRouteIds sets RouteIds field to given value.


### GetOperation

`func (o *BulkAppRouteOperationInput) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *BulkAppRouteOperationInput) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *BulkAppRouteOperationInput) SetOperation(v string)`

SetOperation sets Operation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


