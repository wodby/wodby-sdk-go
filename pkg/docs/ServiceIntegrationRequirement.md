# ServiceIntegrationRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Title** | **string** |  | 
**Type** | **string** |  | 
**Labels** | Pointer to **[]string** |  | [optional] 
**Variables** | [**[]IntegrationVariableRequirement**](IntegrationVariableRequirement.md) |  | 
**Required** | **bool** |  | 
**Multiple** | **bool** |  | 

## Methods

### NewServiceIntegrationRequirement

`func NewServiceIntegrationRequirement(name string, title string, type_ string, variables []IntegrationVariableRequirement, required bool, multiple bool, ) *ServiceIntegrationRequirement`

NewServiceIntegrationRequirement instantiates a new ServiceIntegrationRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceIntegrationRequirementWithDefaults

`func NewServiceIntegrationRequirementWithDefaults() *ServiceIntegrationRequirement`

NewServiceIntegrationRequirementWithDefaults instantiates a new ServiceIntegrationRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServiceIntegrationRequirement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceIntegrationRequirement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceIntegrationRequirement) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *ServiceIntegrationRequirement) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ServiceIntegrationRequirement) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ServiceIntegrationRequirement) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *ServiceIntegrationRequirement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceIntegrationRequirement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceIntegrationRequirement) SetType(v string)`

SetType sets Type field to given value.


### GetLabels

`func (o *ServiceIntegrationRequirement) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ServiceIntegrationRequirement) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ServiceIntegrationRequirement) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ServiceIntegrationRequirement) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetVariables

`func (o *ServiceIntegrationRequirement) GetVariables() []IntegrationVariableRequirement`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ServiceIntegrationRequirement) GetVariablesOk() (*[]IntegrationVariableRequirement, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ServiceIntegrationRequirement) SetVariables(v []IntegrationVariableRequirement)`

SetVariables sets Variables field to given value.


### GetRequired

`func (o *ServiceIntegrationRequirement) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ServiceIntegrationRequirement) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ServiceIntegrationRequirement) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetMultiple

`func (o *ServiceIntegrationRequirement) GetMultiple() bool`

GetMultiple returns the Multiple field if non-nil, zero value otherwise.

### GetMultipleOk

`func (o *ServiceIntegrationRequirement) GetMultipleOk() (*bool, bool)`

GetMultipleOk returns a tuple with the Multiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiple

`func (o *ServiceIntegrationRequirement) SetMultiple(v bool)`

SetMultiple sets Multiple field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


