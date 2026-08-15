# AppAccessCleanup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**AppAccessId** | **int32** |  | 
**AppInstanceId** | **int32** |  | 
**IntegrationId** | **int32** |  | 
**Provider** | **string** |  | 
**Status** | **string** |  | 
**Attempts** | **int32** |  | 
**LastError** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewAppAccessCleanup

`func NewAppAccessCleanup(id int32, appAccessId int32, appInstanceId int32, integrationId int32, provider string, status string, attempts int32, createdAt time.Time, updatedAt time.Time, ) *AppAccessCleanup`

NewAppAccessCleanup instantiates a new AppAccessCleanup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppAccessCleanupWithDefaults

`func NewAppAccessCleanupWithDefaults() *AppAccessCleanup`

NewAppAccessCleanupWithDefaults instantiates a new AppAccessCleanup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppAccessCleanup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppAccessCleanup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppAccessCleanup) SetId(v int32)`

SetId sets Id field to given value.


### GetAppAccessId

`func (o *AppAccessCleanup) GetAppAccessId() int32`

GetAppAccessId returns the AppAccessId field if non-nil, zero value otherwise.

### GetAppAccessIdOk

`func (o *AppAccessCleanup) GetAppAccessIdOk() (*int32, bool)`

GetAppAccessIdOk returns a tuple with the AppAccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppAccessId

`func (o *AppAccessCleanup) SetAppAccessId(v int32)`

SetAppAccessId sets AppAccessId field to given value.


### GetAppInstanceId

`func (o *AppAccessCleanup) GetAppInstanceId() int32`

GetAppInstanceId returns the AppInstanceId field if non-nil, zero value otherwise.

### GetAppInstanceIdOk

`func (o *AppAccessCleanup) GetAppInstanceIdOk() (*int32, bool)`

GetAppInstanceIdOk returns a tuple with the AppInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceId

`func (o *AppAccessCleanup) SetAppInstanceId(v int32)`

SetAppInstanceId sets AppInstanceId field to given value.


### GetIntegrationId

`func (o *AppAccessCleanup) GetIntegrationId() int32`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *AppAccessCleanup) GetIntegrationIdOk() (*int32, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *AppAccessCleanup) SetIntegrationId(v int32)`

SetIntegrationId sets IntegrationId field to given value.


### GetProvider

`func (o *AppAccessCleanup) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AppAccessCleanup) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AppAccessCleanup) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetStatus

`func (o *AppAccessCleanup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppAccessCleanup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppAccessCleanup) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAttempts

`func (o *AppAccessCleanup) GetAttempts() int32`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *AppAccessCleanup) GetAttemptsOk() (*int32, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *AppAccessCleanup) SetAttempts(v int32)`

SetAttempts sets Attempts field to given value.


### GetLastError

`func (o *AppAccessCleanup) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *AppAccessCleanup) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *AppAccessCleanup) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *AppAccessCleanup) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *AppAccessCleanup) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *AppAccessCleanup) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetCreatedAt

`func (o *AppAccessCleanup) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppAccessCleanup) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppAccessCleanup) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AppAccessCleanup) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AppAccessCleanup) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AppAccessCleanup) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


