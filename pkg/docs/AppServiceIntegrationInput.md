# AppServiceIntegrationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**IntegrationIDs** | **[]int32** |  | 

## Methods

### NewAppServiceIntegrationInput

`func NewAppServiceIntegrationInput(name string, integrationIDs []int32, ) *AppServiceIntegrationInput`

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


### GetIntegrationIDs

`func (o *AppServiceIntegrationInput) GetIntegrationIDs() []int32`

GetIntegrationIDs returns the IntegrationIDs field if non-nil, zero value otherwise.

### GetIntegrationIDsOk

`func (o *AppServiceIntegrationInput) GetIntegrationIDsOk() (*[]int32, bool)`

GetIntegrationIDsOk returns a tuple with the IntegrationIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationIDs

`func (o *AppServiceIntegrationInput) SetIntegrationIDs(v []int32)`

SetIntegrationIDs sets IntegrationIDs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


