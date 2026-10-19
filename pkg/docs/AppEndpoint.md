# AppEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**AppServiceId** | **int32** |  | 
**AppEnvironmentId** | **int32** |  | 
**Primary** | **bool** |  | 
**RouteIds** | **[]int32** |  | 
**PortIds** | **[]int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewAppEndpoint

`func NewAppEndpoint(id int32, name string, appServiceId int32, appEnvironmentId int32, primary bool, routeIds []int32, portIds []int32, createdAt time.Time, updatedAt time.Time, ) *AppEndpoint`

NewAppEndpoint instantiates a new AppEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppEndpointWithDefaults

`func NewAppEndpointWithDefaults() *AppEndpoint`

NewAppEndpointWithDefaults instantiates a new AppEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppEndpoint) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppEndpoint) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppEndpoint) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *AppEndpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppEndpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppEndpoint) SetName(v string)`

SetName sets Name field to given value.


### GetAppServiceId

`func (o *AppEndpoint) GetAppServiceId() int32`

GetAppServiceId returns the AppServiceId field if non-nil, zero value otherwise.

### GetAppServiceIdOk

`func (o *AppEndpoint) GetAppServiceIdOk() (*int32, bool)`

GetAppServiceIdOk returns a tuple with the AppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceId

`func (o *AppEndpoint) SetAppServiceId(v int32)`

SetAppServiceId sets AppServiceId field to given value.


### GetAppEnvironmentId

`func (o *AppEndpoint) GetAppEnvironmentId() int32`

GetAppEnvironmentId returns the AppEnvironmentId field if non-nil, zero value otherwise.

### GetAppEnvironmentIdOk

`func (o *AppEndpoint) GetAppEnvironmentIdOk() (*int32, bool)`

GetAppEnvironmentIdOk returns a tuple with the AppEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEnvironmentId

`func (o *AppEndpoint) SetAppEnvironmentId(v int32)`

SetAppEnvironmentId sets AppEnvironmentId field to given value.


### GetPrimary

`func (o *AppEndpoint) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *AppEndpoint) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *AppEndpoint) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.


### GetRouteIds

`func (o *AppEndpoint) GetRouteIds() []int32`

GetRouteIds returns the RouteIds field if non-nil, zero value otherwise.

### GetRouteIdsOk

`func (o *AppEndpoint) GetRouteIdsOk() (*[]int32, bool)`

GetRouteIdsOk returns a tuple with the RouteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteIds

`func (o *AppEndpoint) SetRouteIds(v []int32)`

SetRouteIds sets RouteIds field to given value.


### GetPortIds

`func (o *AppEndpoint) GetPortIds() []int32`

GetPortIds returns the PortIds field if non-nil, zero value otherwise.

### GetPortIdsOk

`func (o *AppEndpoint) GetPortIdsOk() (*[]int32, bool)`

GetPortIdsOk returns a tuple with the PortIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortIds

`func (o *AppEndpoint) SetPortIds(v []int32)`

SetPortIds sets PortIds field to given value.


### GetCreatedAt

`func (o *AppEndpoint) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppEndpoint) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppEndpoint) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AppEndpoint) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AppEndpoint) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AppEndpoint) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


