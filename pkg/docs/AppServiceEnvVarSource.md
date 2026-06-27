# AppServiceEnvVarSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromService** | **bool** |  | 
**FromStack** | **bool** |  | 
**FromWodby** | **bool** |  | 
**Setting** | Pointer to **NullableString** |  | [optional] 
**Link** | Pointer to **NullableString** |  | [optional] 
**Integration** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAppServiceEnvVarSource

`func NewAppServiceEnvVarSource(fromService bool, fromStack bool, fromWodby bool, ) *AppServiceEnvVarSource`

NewAppServiceEnvVarSource instantiates a new AppServiceEnvVarSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceEnvVarSourceWithDefaults

`func NewAppServiceEnvVarSourceWithDefaults() *AppServiceEnvVarSource`

NewAppServiceEnvVarSourceWithDefaults instantiates a new AppServiceEnvVarSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromService

`func (o *AppServiceEnvVarSource) GetFromService() bool`

GetFromService returns the FromService field if non-nil, zero value otherwise.

### GetFromServiceOk

`func (o *AppServiceEnvVarSource) GetFromServiceOk() (*bool, bool)`

GetFromServiceOk returns a tuple with the FromService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromService

`func (o *AppServiceEnvVarSource) SetFromService(v bool)`

SetFromService sets FromService field to given value.


### GetFromStack

`func (o *AppServiceEnvVarSource) GetFromStack() bool`

GetFromStack returns the FromStack field if non-nil, zero value otherwise.

### GetFromStackOk

`func (o *AppServiceEnvVarSource) GetFromStackOk() (*bool, bool)`

GetFromStackOk returns a tuple with the FromStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromStack

`func (o *AppServiceEnvVarSource) SetFromStack(v bool)`

SetFromStack sets FromStack field to given value.


### GetFromWodby

`func (o *AppServiceEnvVarSource) GetFromWodby() bool`

GetFromWodby returns the FromWodby field if non-nil, zero value otherwise.

### GetFromWodbyOk

`func (o *AppServiceEnvVarSource) GetFromWodbyOk() (*bool, bool)`

GetFromWodbyOk returns a tuple with the FromWodby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromWodby

`func (o *AppServiceEnvVarSource) SetFromWodby(v bool)`

SetFromWodby sets FromWodby field to given value.


### GetSetting

`func (o *AppServiceEnvVarSource) GetSetting() string`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *AppServiceEnvVarSource) GetSettingOk() (*string, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *AppServiceEnvVarSource) SetSetting(v string)`

SetSetting sets Setting field to given value.

### HasSetting

`func (o *AppServiceEnvVarSource) HasSetting() bool`

HasSetting returns a boolean if a field has been set.

### SetSettingNil

`func (o *AppServiceEnvVarSource) SetSettingNil(b bool)`

 SetSettingNil sets the value for Setting to be an explicit nil

### UnsetSetting
`func (o *AppServiceEnvVarSource) UnsetSetting()`

UnsetSetting ensures that no value is present for Setting, not even an explicit nil
### GetLink

`func (o *AppServiceEnvVarSource) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *AppServiceEnvVarSource) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *AppServiceEnvVarSource) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *AppServiceEnvVarSource) HasLink() bool`

HasLink returns a boolean if a field has been set.

### SetLinkNil

`func (o *AppServiceEnvVarSource) SetLinkNil(b bool)`

 SetLinkNil sets the value for Link to be an explicit nil

### UnsetLink
`func (o *AppServiceEnvVarSource) UnsetLink()`

UnsetLink ensures that no value is present for Link, not even an explicit nil
### GetIntegration

`func (o *AppServiceEnvVarSource) GetIntegration() string`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *AppServiceEnvVarSource) GetIntegrationOk() (*string, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *AppServiceEnvVarSource) SetIntegration(v string)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *AppServiceEnvVarSource) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### SetIntegrationNil

`func (o *AppServiceEnvVarSource) SetIntegrationNil(b bool)`

 SetIntegrationNil sets the value for Integration to be an explicit nil

### UnsetIntegration
`func (o *AppServiceEnvVarSource) UnsetIntegration()`

UnsetIntegration ensures that no value is present for Integration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


