# UpdateStackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique stack machine name within the organization. | 
**Title** | **string** | Human-readable stack title. | 

## Methods

### NewUpdateStackRequest

`func NewUpdateStackRequest(name string, title string, ) *UpdateStackRequest`

NewUpdateStackRequest instantiates a new UpdateStackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStackRequestWithDefaults

`func NewUpdateStackRequestWithDefaults() *UpdateStackRequest`

NewUpdateStackRequestWithDefaults instantiates a new UpdateStackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateStackRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateStackRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateStackRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *UpdateStackRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateStackRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateStackRequest) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


