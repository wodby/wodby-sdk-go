# AppRouteSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**AppInstanceId** | **int32** |  | 
**RouteId** | **int32** |  | 
**Default** | **bool** |  | 
**Name** | [**AppRouteSettingName**](AppRouteSettingName.md) |  | 
**Value** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewAppRouteSetting

`func NewAppRouteSetting(id int32, appInstanceId int32, routeId int32, default_ bool, name AppRouteSettingName, value string, createdAt time.Time, updatedAt time.Time, ) *AppRouteSetting`

NewAppRouteSetting instantiates a new AppRouteSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRouteSettingWithDefaults

`func NewAppRouteSettingWithDefaults() *AppRouteSetting`

NewAppRouteSettingWithDefaults instantiates a new AppRouteSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppRouteSetting) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppRouteSetting) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppRouteSetting) SetId(v int32)`

SetId sets Id field to given value.


### GetAppInstanceId

`func (o *AppRouteSetting) GetAppInstanceId() int32`

GetAppInstanceId returns the AppInstanceId field if non-nil, zero value otherwise.

### GetAppInstanceIdOk

`func (o *AppRouteSetting) GetAppInstanceIdOk() (*int32, bool)`

GetAppInstanceIdOk returns a tuple with the AppInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceId

`func (o *AppRouteSetting) SetAppInstanceId(v int32)`

SetAppInstanceId sets AppInstanceId field to given value.


### GetRouteId

`func (o *AppRouteSetting) GetRouteId() int32`

GetRouteId returns the RouteId field if non-nil, zero value otherwise.

### GetRouteIdOk

`func (o *AppRouteSetting) GetRouteIdOk() (*int32, bool)`

GetRouteIdOk returns a tuple with the RouteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteId

`func (o *AppRouteSetting) SetRouteId(v int32)`

SetRouteId sets RouteId field to given value.


### GetDefault

`func (o *AppRouteSetting) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *AppRouteSetting) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *AppRouteSetting) SetDefault(v bool)`

SetDefault sets Default field to given value.


### GetName

`func (o *AppRouteSetting) GetName() AppRouteSettingName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppRouteSetting) GetNameOk() (*AppRouteSettingName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppRouteSetting) SetName(v AppRouteSettingName)`

SetName sets Name field to given value.


### GetValue

`func (o *AppRouteSetting) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AppRouteSetting) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AppRouteSetting) SetValue(v string)`

SetValue sets Value field to given value.


### GetCreatedAt

`func (o *AppRouteSetting) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppRouteSetting) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppRouteSetting) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AppRouteSetting) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AppRouteSetting) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AppRouteSetting) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


