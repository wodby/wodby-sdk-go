# AppServiceLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**AppServiceId** | **int32** |  | 
**LinkedAppServiceId** | **int32** |  | 
**Name** | **string** |  | 
**CrossApp** | **bool** | The linked app service is in another app environment, of the same app or another one. | 
**TargetAppInstanceId** | Pointer to **NullableInt32** | The app environment of a linked app service in another environment; null otherwise. | [optional] 
**TargetAppInstanceName** | Pointer to **NullableString** |  | [optional] 
**TargetAppInstanceTitle** | Pointer to **NullableString** |  | [optional] 
**TargetAppId** | Pointer to **NullableInt32** | The app of a linked app service in another environment; null otherwise. | [optional] 
**TargetAppName** | Pointer to **NullableString** |  | [optional] 
**TargetAppTitle** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAppServiceLink

`func NewAppServiceLink(id int32, appServiceId int32, linkedAppServiceId int32, name string, crossApp bool, ) *AppServiceLink`

NewAppServiceLink instantiates a new AppServiceLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceLinkWithDefaults

`func NewAppServiceLinkWithDefaults() *AppServiceLink`

NewAppServiceLinkWithDefaults instantiates a new AppServiceLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppServiceLink) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppServiceLink) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppServiceLink) SetId(v int32)`

SetId sets Id field to given value.


### GetAppServiceId

`func (o *AppServiceLink) GetAppServiceId() int32`

GetAppServiceId returns the AppServiceId field if non-nil, zero value otherwise.

### GetAppServiceIdOk

`func (o *AppServiceLink) GetAppServiceIdOk() (*int32, bool)`

GetAppServiceIdOk returns a tuple with the AppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceId

`func (o *AppServiceLink) SetAppServiceId(v int32)`

SetAppServiceId sets AppServiceId field to given value.


### GetLinkedAppServiceId

`func (o *AppServiceLink) GetLinkedAppServiceId() int32`

GetLinkedAppServiceId returns the LinkedAppServiceId field if non-nil, zero value otherwise.

### GetLinkedAppServiceIdOk

`func (o *AppServiceLink) GetLinkedAppServiceIdOk() (*int32, bool)`

GetLinkedAppServiceIdOk returns a tuple with the LinkedAppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAppServiceId

`func (o *AppServiceLink) SetLinkedAppServiceId(v int32)`

SetLinkedAppServiceId sets LinkedAppServiceId field to given value.


### GetName

`func (o *AppServiceLink) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppServiceLink) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppServiceLink) SetName(v string)`

SetName sets Name field to given value.


### GetCrossApp

`func (o *AppServiceLink) GetCrossApp() bool`

GetCrossApp returns the CrossApp field if non-nil, zero value otherwise.

### GetCrossAppOk

`func (o *AppServiceLink) GetCrossAppOk() (*bool, bool)`

GetCrossAppOk returns a tuple with the CrossApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossApp

`func (o *AppServiceLink) SetCrossApp(v bool)`

SetCrossApp sets CrossApp field to given value.


### GetTargetAppInstanceId

`func (o *AppServiceLink) GetTargetAppInstanceId() int32`

GetTargetAppInstanceId returns the TargetAppInstanceId field if non-nil, zero value otherwise.

### GetTargetAppInstanceIdOk

`func (o *AppServiceLink) GetTargetAppInstanceIdOk() (*int32, bool)`

GetTargetAppInstanceIdOk returns a tuple with the TargetAppInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAppInstanceId

`func (o *AppServiceLink) SetTargetAppInstanceId(v int32)`

SetTargetAppInstanceId sets TargetAppInstanceId field to given value.

### HasTargetAppInstanceId

`func (o *AppServiceLink) HasTargetAppInstanceId() bool`

HasTargetAppInstanceId returns a boolean if a field has been set.

### SetTargetAppInstanceIdNil

`func (o *AppServiceLink) SetTargetAppInstanceIdNil(b bool)`

 SetTargetAppInstanceIdNil sets the value for TargetAppInstanceId to be an explicit nil

### UnsetTargetAppInstanceId
`func (o *AppServiceLink) UnsetTargetAppInstanceId()`

UnsetTargetAppInstanceId ensures that no value is present for TargetAppInstanceId, not even an explicit nil
### GetTargetAppInstanceName

`func (o *AppServiceLink) GetTargetAppInstanceName() string`

GetTargetAppInstanceName returns the TargetAppInstanceName field if non-nil, zero value otherwise.

### GetTargetAppInstanceNameOk

`func (o *AppServiceLink) GetTargetAppInstanceNameOk() (*string, bool)`

GetTargetAppInstanceNameOk returns a tuple with the TargetAppInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAppInstanceName

`func (o *AppServiceLink) SetTargetAppInstanceName(v string)`

SetTargetAppInstanceName sets TargetAppInstanceName field to given value.

### HasTargetAppInstanceName

`func (o *AppServiceLink) HasTargetAppInstanceName() bool`

HasTargetAppInstanceName returns a boolean if a field has been set.

### SetTargetAppInstanceNameNil

`func (o *AppServiceLink) SetTargetAppInstanceNameNil(b bool)`

 SetTargetAppInstanceNameNil sets the value for TargetAppInstanceName to be an explicit nil

### UnsetTargetAppInstanceName
`func (o *AppServiceLink) UnsetTargetAppInstanceName()`

UnsetTargetAppInstanceName ensures that no value is present for TargetAppInstanceName, not even an explicit nil
### GetTargetAppInstanceTitle

`func (o *AppServiceLink) GetTargetAppInstanceTitle() string`

GetTargetAppInstanceTitle returns the TargetAppInstanceTitle field if non-nil, zero value otherwise.

### GetTargetAppInstanceTitleOk

`func (o *AppServiceLink) GetTargetAppInstanceTitleOk() (*string, bool)`

GetTargetAppInstanceTitleOk returns a tuple with the TargetAppInstanceTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAppInstanceTitle

`func (o *AppServiceLink) SetTargetAppInstanceTitle(v string)`

SetTargetAppInstanceTitle sets TargetAppInstanceTitle field to given value.

### HasTargetAppInstanceTitle

`func (o *AppServiceLink) HasTargetAppInstanceTitle() bool`

HasTargetAppInstanceTitle returns a boolean if a field has been set.

### SetTargetAppInstanceTitleNil

`func (o *AppServiceLink) SetTargetAppInstanceTitleNil(b bool)`

 SetTargetAppInstanceTitleNil sets the value for TargetAppInstanceTitle to be an explicit nil

### UnsetTargetAppInstanceTitle
`func (o *AppServiceLink) UnsetTargetAppInstanceTitle()`

UnsetTargetAppInstanceTitle ensures that no value is present for TargetAppInstanceTitle, not even an explicit nil
### GetTargetAppId

`func (o *AppServiceLink) GetTargetAppId() int32`

GetTargetAppId returns the TargetAppId field if non-nil, zero value otherwise.

### GetTargetAppIdOk

`func (o *AppServiceLink) GetTargetAppIdOk() (*int32, bool)`

GetTargetAppIdOk returns a tuple with the TargetAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAppId

`func (o *AppServiceLink) SetTargetAppId(v int32)`

SetTargetAppId sets TargetAppId field to given value.

### HasTargetAppId

`func (o *AppServiceLink) HasTargetAppId() bool`

HasTargetAppId returns a boolean if a field has been set.

### SetTargetAppIdNil

`func (o *AppServiceLink) SetTargetAppIdNil(b bool)`

 SetTargetAppIdNil sets the value for TargetAppId to be an explicit nil

### UnsetTargetAppId
`func (o *AppServiceLink) UnsetTargetAppId()`

UnsetTargetAppId ensures that no value is present for TargetAppId, not even an explicit nil
### GetTargetAppName

`func (o *AppServiceLink) GetTargetAppName() string`

GetTargetAppName returns the TargetAppName field if non-nil, zero value otherwise.

### GetTargetAppNameOk

`func (o *AppServiceLink) GetTargetAppNameOk() (*string, bool)`

GetTargetAppNameOk returns a tuple with the TargetAppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAppName

`func (o *AppServiceLink) SetTargetAppName(v string)`

SetTargetAppName sets TargetAppName field to given value.

### HasTargetAppName

`func (o *AppServiceLink) HasTargetAppName() bool`

HasTargetAppName returns a boolean if a field has been set.

### SetTargetAppNameNil

`func (o *AppServiceLink) SetTargetAppNameNil(b bool)`

 SetTargetAppNameNil sets the value for TargetAppName to be an explicit nil

### UnsetTargetAppName
`func (o *AppServiceLink) UnsetTargetAppName()`

UnsetTargetAppName ensures that no value is present for TargetAppName, not even an explicit nil
### GetTargetAppTitle

`func (o *AppServiceLink) GetTargetAppTitle() string`

GetTargetAppTitle returns the TargetAppTitle field if non-nil, zero value otherwise.

### GetTargetAppTitleOk

`func (o *AppServiceLink) GetTargetAppTitleOk() (*string, bool)`

GetTargetAppTitleOk returns a tuple with the TargetAppTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAppTitle

`func (o *AppServiceLink) SetTargetAppTitle(v string)`

SetTargetAppTitle sets TargetAppTitle field to given value.

### HasTargetAppTitle

`func (o *AppServiceLink) HasTargetAppTitle() bool`

HasTargetAppTitle returns a boolean if a field has been set.

### SetTargetAppTitleNil

`func (o *AppServiceLink) SetTargetAppTitleNil(b bool)`

 SetTargetAppTitleNil sets the value for TargetAppTitle to be an explicit nil

### UnsetTargetAppTitle
`func (o *AppServiceLink) UnsetTargetAppTitle()`

UnsetTargetAppTitle ensures that no value is present for TargetAppTitle, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


