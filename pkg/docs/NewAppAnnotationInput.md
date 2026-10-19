# NewAppAnnotationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppEnvironmentId** | **int32** |  | 
**AppServiceId** | Pointer to **NullableInt32** | Limit the annotation to one service&#39;s ingress. | [optional] 
**AppRouteId** | Pointer to **NullableInt32** | Limit the annotation to one route&#39;s ingress. | [optional] 
**Name** | **string** |  | 
**Value** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewNewAppAnnotationInput

`func NewNewAppAnnotationInput(appEnvironmentId int32, name string, value string, type_ string, ) *NewAppAnnotationInput`

NewNewAppAnnotationInput instantiates a new NewAppAnnotationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewAppAnnotationInputWithDefaults

`func NewNewAppAnnotationInputWithDefaults() *NewAppAnnotationInput`

NewNewAppAnnotationInputWithDefaults instantiates a new NewAppAnnotationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppEnvironmentId

`func (o *NewAppAnnotationInput) GetAppEnvironmentId() int32`

GetAppEnvironmentId returns the AppEnvironmentId field if non-nil, zero value otherwise.

### GetAppEnvironmentIdOk

`func (o *NewAppAnnotationInput) GetAppEnvironmentIdOk() (*int32, bool)`

GetAppEnvironmentIdOk returns a tuple with the AppEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEnvironmentId

`func (o *NewAppAnnotationInput) SetAppEnvironmentId(v int32)`

SetAppEnvironmentId sets AppEnvironmentId field to given value.


### GetAppServiceId

`func (o *NewAppAnnotationInput) GetAppServiceId() int32`

GetAppServiceId returns the AppServiceId field if non-nil, zero value otherwise.

### GetAppServiceIdOk

`func (o *NewAppAnnotationInput) GetAppServiceIdOk() (*int32, bool)`

GetAppServiceIdOk returns a tuple with the AppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceId

`func (o *NewAppAnnotationInput) SetAppServiceId(v int32)`

SetAppServiceId sets AppServiceId field to given value.

### HasAppServiceId

`func (o *NewAppAnnotationInput) HasAppServiceId() bool`

HasAppServiceId returns a boolean if a field has been set.

### SetAppServiceIdNil

`func (o *NewAppAnnotationInput) SetAppServiceIdNil(b bool)`

 SetAppServiceIdNil sets the value for AppServiceId to be an explicit nil

### UnsetAppServiceId
`func (o *NewAppAnnotationInput) UnsetAppServiceId()`

UnsetAppServiceId ensures that no value is present for AppServiceId, not even an explicit nil
### GetAppRouteId

`func (o *NewAppAnnotationInput) GetAppRouteId() int32`

GetAppRouteId returns the AppRouteId field if non-nil, zero value otherwise.

### GetAppRouteIdOk

`func (o *NewAppAnnotationInput) GetAppRouteIdOk() (*int32, bool)`

GetAppRouteIdOk returns a tuple with the AppRouteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRouteId

`func (o *NewAppAnnotationInput) SetAppRouteId(v int32)`

SetAppRouteId sets AppRouteId field to given value.

### HasAppRouteId

`func (o *NewAppAnnotationInput) HasAppRouteId() bool`

HasAppRouteId returns a boolean if a field has been set.

### SetAppRouteIdNil

`func (o *NewAppAnnotationInput) SetAppRouteIdNil(b bool)`

 SetAppRouteIdNil sets the value for AppRouteId to be an explicit nil

### UnsetAppRouteId
`func (o *NewAppAnnotationInput) UnsetAppRouteId()`

UnsetAppRouteId ensures that no value is present for AppRouteId, not even an explicit nil
### GetName

`func (o *NewAppAnnotationInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewAppAnnotationInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewAppAnnotationInput) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *NewAppAnnotationInput) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NewAppAnnotationInput) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NewAppAnnotationInput) SetValue(v string)`

SetValue sets Value field to given value.


### GetType

`func (o *NewAppAnnotationInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NewAppAnnotationInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NewAppAnnotationInput) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


