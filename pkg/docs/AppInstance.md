# AppInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Title** | **string** |  | 
**Status** | **string** |  | 
**MainDomain** | Pointer to **NullableString** |  | [optional] 
**AppId** | **int32** |  | 
**ClusterId** | **int32** |  | 
**EnvId** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewAppInstance

`func NewAppInstance(id int32, name string, title string, status string, appId int32, clusterId int32, envId int32, createdAt time.Time, updatedAt time.Time, ) *AppInstance`

NewAppInstance instantiates a new AppInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppInstanceWithDefaults

`func NewAppInstanceWithDefaults() *AppInstance`

NewAppInstanceWithDefaults instantiates a new AppInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppInstance) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppInstance) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppInstance) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *AppInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppInstance) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *AppInstance) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AppInstance) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AppInstance) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetStatus

`func (o *AppInstance) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppInstance) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppInstance) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMainDomain

`func (o *AppInstance) GetMainDomain() string`

GetMainDomain returns the MainDomain field if non-nil, zero value otherwise.

### GetMainDomainOk

`func (o *AppInstance) GetMainDomainOk() (*string, bool)`

GetMainDomainOk returns a tuple with the MainDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainDomain

`func (o *AppInstance) SetMainDomain(v string)`

SetMainDomain sets MainDomain field to given value.

### HasMainDomain

`func (o *AppInstance) HasMainDomain() bool`

HasMainDomain returns a boolean if a field has been set.

### SetMainDomainNil

`func (o *AppInstance) SetMainDomainNil(b bool)`

 SetMainDomainNil sets the value for MainDomain to be an explicit nil

### UnsetMainDomain
`func (o *AppInstance) UnsetMainDomain()`

UnsetMainDomain ensures that no value is present for MainDomain, not even an explicit nil
### GetAppId

`func (o *AppInstance) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *AppInstance) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *AppInstance) SetAppId(v int32)`

SetAppId sets AppId field to given value.


### GetClusterId

`func (o *AppInstance) GetClusterId() int32`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *AppInstance) GetClusterIdOk() (*int32, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *AppInstance) SetClusterId(v int32)`

SetClusterId sets ClusterId field to given value.


### GetEnvId

`func (o *AppInstance) GetEnvId() int32`

GetEnvId returns the EnvId field if non-nil, zero value otherwise.

### GetEnvIdOk

`func (o *AppInstance) GetEnvIdOk() (*int32, bool)`

GetEnvIdOk returns a tuple with the EnvId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvId

`func (o *AppInstance) SetEnvId(v int32)`

SetEnvId sets EnvId field to given value.


### GetCreatedAt

`func (o *AppInstance) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppInstance) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppInstance) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AppInstance) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AppInstance) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AppInstance) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


