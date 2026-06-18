# AppServiceIntegrationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**IntegrationIds** | **[]int32** |  | 

## Methods

### NewAppServiceIntegrationInput

`func NewAppServiceIntegrationInput(name string, integrationIds []int32, ) *AppServiceIntegrationInput`

NewAppServiceIntegrationInput instantiates a new AppServiceIntegrationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceIntegrationInputWithDefaults

`func NewAppServiceIntegrationInputWithDefaults() *AppServiceIntegrationInput`

NewAppServiceIntegrationInputWithDefaults instantiates a new AppServiceIntegrationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AppServiceIntegrationInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppServiceIntegrationInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppServiceIntegrationInput) SetName(v string)`

SetName sets Name field to given value.


### GetIntegrationIds

`func (o *AppServiceIntegrationInput) GetIntegrationIds() []int32`

GetIntegrationIds returns the IntegrationIds field if non-nil, zero value otherwise.

### GetIntegrationIdsOk

`func (o *AppServiceIntegrationInput) GetIntegrationIdsOk() (*[]int32, bool)`

GetIntegrationIdsOk returns a tuple with the IntegrationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationIds

`func (o *AppServiceIntegrationInput) SetIntegrationIds(v []int32)`

SetIntegrationIds sets IntegrationIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


