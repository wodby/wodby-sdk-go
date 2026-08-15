# StackRevisionLinkIssue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**StackServiceId** | **int32** |  | 
**StackServiceName** | **string** |  | 
**StackServiceTitle** | **string** |  | 
**LinkName** | Pointer to **NullableString** |  | [optional] 
**TargetStackServiceId** | Pointer to **NullableInt32** |  | [optional] 
**TargetStackServiceName** | Pointer to **NullableString** |  | [optional] 
**TargetStackServiceTitle** | Pointer to **NullableString** |  | [optional] 
**Message** | **string** |  | 

## Methods

### NewStackRevisionLinkIssue

`func NewStackRevisionLinkIssue(code string, stackServiceId int32, stackServiceName string, stackServiceTitle string, message string, ) *StackRevisionLinkIssue`

NewStackRevisionLinkIssue instantiates a new StackRevisionLinkIssue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackRevisionLinkIssueWithDefaults

`func NewStackRevisionLinkIssueWithDefaults() *StackRevisionLinkIssue`

NewStackRevisionLinkIssueWithDefaults instantiates a new StackRevisionLinkIssue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *StackRevisionLinkIssue) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *StackRevisionLinkIssue) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *StackRevisionLinkIssue) SetCode(v string)`

SetCode sets Code field to given value.


### GetStackServiceId

`func (o *StackRevisionLinkIssue) GetStackServiceId() int32`

GetStackServiceId returns the StackServiceId field if non-nil, zero value otherwise.

### GetStackServiceIdOk

`func (o *StackRevisionLinkIssue) GetStackServiceIdOk() (*int32, bool)`

GetStackServiceIdOk returns a tuple with the StackServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackServiceId

`func (o *StackRevisionLinkIssue) SetStackServiceId(v int32)`

SetStackServiceId sets StackServiceId field to given value.


### GetStackServiceName

`func (o *StackRevisionLinkIssue) GetStackServiceName() string`

GetStackServiceName returns the StackServiceName field if non-nil, zero value otherwise.

### GetStackServiceNameOk

`func (o *StackRevisionLinkIssue) GetStackServiceNameOk() (*string, bool)`

GetStackServiceNameOk returns a tuple with the StackServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackServiceName

`func (o *StackRevisionLinkIssue) SetStackServiceName(v string)`

SetStackServiceName sets StackServiceName field to given value.


### GetStackServiceTitle

`func (o *StackRevisionLinkIssue) GetStackServiceTitle() string`

GetStackServiceTitle returns the StackServiceTitle field if non-nil, zero value otherwise.

### GetStackServiceTitleOk

`func (o *StackRevisionLinkIssue) GetStackServiceTitleOk() (*string, bool)`

GetStackServiceTitleOk returns a tuple with the StackServiceTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackServiceTitle

`func (o *StackRevisionLinkIssue) SetStackServiceTitle(v string)`

SetStackServiceTitle sets StackServiceTitle field to given value.


### GetLinkName

`func (o *StackRevisionLinkIssue) GetLinkName() string`

GetLinkName returns the LinkName field if non-nil, zero value otherwise.

### GetLinkNameOk

`func (o *StackRevisionLinkIssue) GetLinkNameOk() (*string, bool)`

GetLinkNameOk returns a tuple with the LinkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkName

`func (o *StackRevisionLinkIssue) SetLinkName(v string)`

SetLinkName sets LinkName field to given value.

### HasLinkName

`func (o *StackRevisionLinkIssue) HasLinkName() bool`

HasLinkName returns a boolean if a field has been set.

### SetLinkNameNil

`func (o *StackRevisionLinkIssue) SetLinkNameNil(b bool)`

 SetLinkNameNil sets the value for LinkName to be an explicit nil

### UnsetLinkName
`func (o *StackRevisionLinkIssue) UnsetLinkName()`

UnsetLinkName ensures that no value is present for LinkName, not even an explicit nil
### GetTargetStackServiceId

`func (o *StackRevisionLinkIssue) GetTargetStackServiceId() int32`

GetTargetStackServiceId returns the TargetStackServiceId field if non-nil, zero value otherwise.

### GetTargetStackServiceIdOk

`func (o *StackRevisionLinkIssue) GetTargetStackServiceIdOk() (*int32, bool)`

GetTargetStackServiceIdOk returns a tuple with the TargetStackServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetStackServiceId

`func (o *StackRevisionLinkIssue) SetTargetStackServiceId(v int32)`

SetTargetStackServiceId sets TargetStackServiceId field to given value.

### HasTargetStackServiceId

`func (o *StackRevisionLinkIssue) HasTargetStackServiceId() bool`

HasTargetStackServiceId returns a boolean if a field has been set.

### SetTargetStackServiceIdNil

`func (o *StackRevisionLinkIssue) SetTargetStackServiceIdNil(b bool)`

 SetTargetStackServiceIdNil sets the value for TargetStackServiceId to be an explicit nil

### UnsetTargetStackServiceId
`func (o *StackRevisionLinkIssue) UnsetTargetStackServiceId()`

UnsetTargetStackServiceId ensures that no value is present for TargetStackServiceId, not even an explicit nil
### GetTargetStackServiceName

`func (o *StackRevisionLinkIssue) GetTargetStackServiceName() string`

GetTargetStackServiceName returns the TargetStackServiceName field if non-nil, zero value otherwise.

### GetTargetStackServiceNameOk

`func (o *StackRevisionLinkIssue) GetTargetStackServiceNameOk() (*string, bool)`

GetTargetStackServiceNameOk returns a tuple with the TargetStackServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetStackServiceName

`func (o *StackRevisionLinkIssue) SetTargetStackServiceName(v string)`

SetTargetStackServiceName sets TargetStackServiceName field to given value.

### HasTargetStackServiceName

`func (o *StackRevisionLinkIssue) HasTargetStackServiceName() bool`

HasTargetStackServiceName returns a boolean if a field has been set.

### SetTargetStackServiceNameNil

`func (o *StackRevisionLinkIssue) SetTargetStackServiceNameNil(b bool)`

 SetTargetStackServiceNameNil sets the value for TargetStackServiceName to be an explicit nil

### UnsetTargetStackServiceName
`func (o *StackRevisionLinkIssue) UnsetTargetStackServiceName()`

UnsetTargetStackServiceName ensures that no value is present for TargetStackServiceName, not even an explicit nil
### GetTargetStackServiceTitle

`func (o *StackRevisionLinkIssue) GetTargetStackServiceTitle() string`

GetTargetStackServiceTitle returns the TargetStackServiceTitle field if non-nil, zero value otherwise.

### GetTargetStackServiceTitleOk

`func (o *StackRevisionLinkIssue) GetTargetStackServiceTitleOk() (*string, bool)`

GetTargetStackServiceTitleOk returns a tuple with the TargetStackServiceTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetStackServiceTitle

`func (o *StackRevisionLinkIssue) SetTargetStackServiceTitle(v string)`

SetTargetStackServiceTitle sets TargetStackServiceTitle field to given value.

### HasTargetStackServiceTitle

`func (o *StackRevisionLinkIssue) HasTargetStackServiceTitle() bool`

HasTargetStackServiceTitle returns a boolean if a field has been set.

### SetTargetStackServiceTitleNil

`func (o *StackRevisionLinkIssue) SetTargetStackServiceTitleNil(b bool)`

 SetTargetStackServiceTitleNil sets the value for TargetStackServiceTitle to be an explicit nil

### UnsetTargetStackServiceTitle
`func (o *StackRevisionLinkIssue) UnsetTargetStackServiceTitle()`

UnsetTargetStackServiceTitle ensures that no value is present for TargetStackServiceTitle, not even an explicit nil
### GetMessage

`func (o *StackRevisionLinkIssue) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *StackRevisionLinkIssue) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *StackRevisionLinkIssue) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


