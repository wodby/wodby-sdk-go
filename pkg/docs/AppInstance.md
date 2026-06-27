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
**StackId** | **int32** |  | 
**StackRevId** | **int32** |  | 
**StackName** | **string** |  | 
**StackTitle** | **string** |  | 
**StackIcon** | **string** |  | 
**StackRevNumber** | **int32** |  | 
**StackVersion** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewAppInstance

`func NewAppInstance(id int32, name string, title string, status string, appId int32, clusterId int32, envId int32, stackId int32, stackRevId int32, stackName string, stackTitle string, stackIcon string, stackRevNumber int32, stackVersion string, createdAt time.Time, updatedAt time.Time, ) *AppInstance`

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


### GetStackId

`func (o *AppInstance) GetStackId() int32`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *AppInstance) GetStackIdOk() (*int32, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *AppInstance) SetStackId(v int32)`

SetStackId sets StackId field to given value.


### GetStackRevId

`func (o *AppInstance) GetStackRevId() int32`

GetStackRevId returns the StackRevId field if non-nil, zero value otherwise.

### GetStackRevIdOk

`func (o *AppInstance) GetStackRevIdOk() (*int32, bool)`

GetStackRevIdOk returns a tuple with the StackRevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackRevId

`func (o *AppInstance) SetStackRevId(v int32)`

SetStackRevId sets StackRevId field to given value.


### GetStackName

`func (o *AppInstance) GetStackName() string`

GetStackName returns the StackName field if non-nil, zero value otherwise.

### GetStackNameOk

`func (o *AppInstance) GetStackNameOk() (*string, bool)`

GetStackNameOk returns a tuple with the StackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackName

`func (o *AppInstance) SetStackName(v string)`

SetStackName sets StackName field to given value.


### GetStackTitle

`func (o *AppInstance) GetStackTitle() string`

GetStackTitle returns the StackTitle field if non-nil, zero value otherwise.

### GetStackTitleOk

`func (o *AppInstance) GetStackTitleOk() (*string, bool)`

GetStackTitleOk returns a tuple with the StackTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackTitle

`func (o *AppInstance) SetStackTitle(v string)`

SetStackTitle sets StackTitle field to given value.


### GetStackIcon

`func (o *AppInstance) GetStackIcon() string`

GetStackIcon returns the StackIcon field if non-nil, zero value otherwise.

### GetStackIconOk

`func (o *AppInstance) GetStackIconOk() (*string, bool)`

GetStackIconOk returns a tuple with the StackIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackIcon

`func (o *AppInstance) SetStackIcon(v string)`

SetStackIcon sets StackIcon field to given value.


### GetStackRevNumber

`func (o *AppInstance) GetStackRevNumber() int32`

GetStackRevNumber returns the StackRevNumber field if non-nil, zero value otherwise.

### GetStackRevNumberOk

`func (o *AppInstance) GetStackRevNumberOk() (*int32, bool)`

GetStackRevNumberOk returns a tuple with the StackRevNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackRevNumber

`func (o *AppInstance) SetStackRevNumber(v int32)`

SetStackRevNumber sets StackRevNumber field to given value.


### GetStackVersion

`func (o *AppInstance) GetStackVersion() string`

GetStackVersion returns the StackVersion field if non-nil, zero value otherwise.

### GetStackVersionOk

`func (o *AppInstance) GetStackVersionOk() (*string, bool)`

GetStackVersionOk returns a tuple with the StackVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackVersion

`func (o *AppInstance) SetStackVersion(v string)`

SetStackVersion sets StackVersion field to given value.


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


